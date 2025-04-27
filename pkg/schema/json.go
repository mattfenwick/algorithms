package schema

import (
	"fmt"
	"strings"

	"github.com/mattfenwick/algorithms/pkg/utils"
	"github.com/mattfenwick/collections/pkg/dict"
	"github.com/mattfenwick/collections/pkg/json"
	"github.com/mattfenwick/collections/pkg/set"
	"github.com/mattfenwick/collections/pkg/slice"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type Type struct {
	Null   bool
	Bool   bool
	Number bool
	String bool
	Array  *Type
	Dict   map[string]*Type // key is from the raw json
	Union  map[string]*Type // key is calculated from the corresponding Type -- so sort of like a set
	// internal use
	key string
}

func (t *Type) Key() string {
	if t.key == "" {
		if t.Null {
			t.key = "null"
		} else if t.Bool {
			t.key = "bool"
		} else if t.Number {
			t.key = "number"
		} else if t.String {
			t.key = "string"
		} else if t.Array != nil {
			t.key = fmt.Sprintf("[%s]", t.Array.Key())
		} else if t.Dict != nil {
			t.key = string(utils.Die(json.MarshalWithOptions(dict.Map(func(t *Type) string { return t.Key() }, t.Dict), &json.MarshalOptions{
				Indent: false,
				Sort:   true,
			})))
		} else if t.Union != nil {
			t.key = strings.Join(slice.Sort(dict.Keys(t.Union)), " | ")
		} else {
			utils.Die0(errors.Errorf("invalid Type %+v", t))
		}
	}
	return t.key
}

func (t *Type) PrettyPrint() string {
	return strings.Join(
		slice.Map(func(l *line) string {
			return strings.Repeat("  ", l.indent) + l.content
		},
			t.prettyPrintHelper(0)),
		"\n")
}

type line struct {
	indent  int
	content string
}

func (t *Type) prettyPrintHelper(indent int) []*line {
	if t.Null {
		return []*line{{indent, "null"}}
	} else if t.Bool {
		return []*line{{indent, "bool"}}
	} else if t.Number {
		return []*line{{indent, "number"}}
	} else if t.String {
		return []*line{{indent, "string"}}
	} else if t.Array != nil {
		ls := t.Array.prettyPrintHelper(indent + 1)
		return append(append([]*line{{indent, "["}}, ls...), &line{indent, "]"})
	} else if t.Dict != nil {
		keys := slice.Sort(dict.Keys(t.Dict))
		ls := []*line{{indent, "{"}}
		for _, k := range keys {
			ls = append(ls, &line{indent + 1, k + ":"})
			ls = append(ls, t.Dict[k].prettyPrintHelper(indent+2)...)
		}
		ls = append(ls, &line{indent, "}"})
		return ls
	} else if t.Union != nil {
		ls := []*line{{indent, "("}}
		for _, k := range slice.Sort(dict.Keys(t.Union)) {
			ls = append(ls, t.Union[k].prettyPrintHelper(indent+1)...)
		}
		ls = append(ls, &line{indent, ")"})
		return ls
	} else {
		panic(errors.Errorf("invalid Type %+v", t))
	}
}

func Union(ts []*Type) *Type {
	types := map[string]*Type{}
	for _, t := range ts {
		if v, ok := types[t.Key()]; ok {
			logrus.Infof("overwriting %+v with %+v for key %s, should be same though", v, t, t.Key())
		}
		// TODO 3 ways to look at {}. 1: arbitrary dict, kvs don't matter. 2: ks matter, vs don't. 3: ks and vs both matter. which one to use? should this choice be configurable?
		// TODO do pairwise union of objects?
		// TODO -- if it's a union, pull out every member and add into `types`
		types[t.Key()] = t
	}
	return &Type{Union: types}
}

func GetType(o any) *Type {
	switch val := o.(type) {
	case bool:
		return &Type{Bool: true}
	case int, float64:
		return &Type{Number: true}
	case string:
		return &Type{String: true}
	case nil:
		return &Type{Null: true}
	case []any:
		return &Type{Array: Union(slice.Map(GetType, val))}
	case map[string]any:
		// if (compactType) // TODO
		return &Type{Dict: dict.Map(GetType, val)}
	default:
		panic(errors.Errorf("invalid type: %+v, %T", o, o))
	}
}

type Index struct {
	Object *string
	Array  *int
}

func copySlice[A any](as []A) []A {
	out := make([]A, len(as))
	copy(out, as)
	return out
}

func ObjectIndex(k string) *Index {
	return &Index{Object: &k}
}

func ArrayIndex(i int) *Index {
	return &Index{Array: &i}
}

type Path []*Index

func (p Path) Key() string {
	return strings.Join(slice.Map(func(i *Index) string {
		if i.Array != nil {
			return "[*]"
		}
		return fmt.Sprintf(`["%s"]`, strings.ReplaceAll(*i.Object, `"`, `\"`))
	}, p), "")
}

type Listener func(Path, any)

func Traverse(o any) string {
	paths := TraversePaths(o)
	var lines []string
	for _, path := range slice.Sort(dict.Keys(paths)) {
		for _, t := range slice.Sort(paths[path].ToSlice()) {
			lines = append(lines, path+": "+t)
		}
	}
	return strings.Join(lines, "\n")
}

func TraversePaths(o any) map[string]*set.Set[string] {
	paths := map[string]*set.Set[string]{}
	f := func(path Path, o any) {
		if _, ok := paths[path.Key()]; !ok {
			paths[path.Key()] = set.Empty[string]()
		}
		var typeName string
		switch o.(type) {
		case nil:
			typeName = "null"
		case bool:
			typeName = "bool"
		case string:
			typeName = "string"
		case int, float64:
			typeName = "number"
		case []any:
			typeName = "array"
		case map[string]any:
			typeName = "object"
		default:
			panic(errors.Errorf("invalid type: %T", o))
		}
		paths[path.Key()].Add(typeName)
	}
	TraverseHelp(nil, o, f)
	return paths
}

func TraverseHelp(path Path, o any, f Listener) {
	f(copySlice(path), o)
	switch val := o.(type) {
	case bool, int, float64, string, nil:
		// no need to recur for these types
	case []any:
		for i, v := range val {
			newPath := append(copySlice(path), ArrayIndex(i))
			TraverseHelp(newPath, v, f)
		}
	case map[string]any:
		for _, k := range slice.Sort(dict.Keys(val)) {
			newPath := append(copySlice(path), ObjectIndex(k))
			TraverseHelp(newPath, val[k], f)
		}
	default:
		panic(errors.Errorf("invalid type: %+v, %T", o, o))
	}

}
