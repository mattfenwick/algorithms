package schema

import (
	"fmt"
	"strings"

	"github.com/mattfenwick/algorithms/pkg/utils"
	"github.com/mattfenwick/collections/pkg/dict"
	"github.com/mattfenwick/collections/pkg/json"
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

func Union(ts []*Type) *Type {
	types := map[string]*Type{}
	for _, t := range ts {
		if v, ok := types[t.Key()]; ok {
			logrus.Infof("overwriting %+v with %+v for key %s, should be same though", v, t, t.Key())
		}
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
