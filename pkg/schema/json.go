package schema

import (
	"github.com/mattfenwick/collections/pkg/dict"
	"github.com/mattfenwick/collections/pkg/slice"
	"github.com/pkg/errors"
)

type Type struct {
	Null    bool
	Bool    *bool
	Int     *int
	Float64 *float64
	String  *string
	Array   *Type
	Dict    map[string]*Type
	Union   []*Type // TODO change this to a set or something where membership test is constant time
}

func ptr[A any](a A) *A {
	return &a
}

func CompactType(ts []*Type) *Type {
	return &Type{Union: ts}
}

func GetType(o any) *Type {
	switch val := o.(type) {
	case bool:
		return &Type{Bool: ptr(val)}
	case int:
		return &Type{Int: ptr(val)}
	case float64:
		return &Type{Float64: ptr(val)}
	case string:
		return &Type{String: ptr(val)}
	case nil:
		return &Type{Null: true}
	case []any:
		return &Type{Array: CompactType(slice.Map(GetType, val))}
	case map[string]any:
		// if (compactType) // TODO
		return &Type{Dict: dict.Map(GetType, val)}
	default:
		panic(errors.Errorf("invalid type: %+v, %T", o, o))
	}
}
