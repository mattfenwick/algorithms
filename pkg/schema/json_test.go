package schema

import (
	"encoding/json"
	"testing"

	"github.com/mattfenwick/algorithms/pkg/utils"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestSchema(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Schema Suite")
}

type Case struct {
	Name     string
	Input    string
	Error    string
	Expected interface{}
}

var _ = Describe("Type parsing", func() {
	cases := []*Case{
		{"first, crazy", `{"a": 1, "b": true, "c": [1,1.2,"d",{}]}`, "", &Type{Dict: map[string]*Type{
			"a": {Float64: ptr(1.0)},
			"b": {Bool: ptr(true)},
			"c": {Array: &Type{Union: []*Type{
				{Float64: ptr(1.0)},
				{Float64: ptr(1.2)},
				{String: ptr("d")},
				{Dict: map[string]*Type{}},
			}}},
		}}},
	}

	for _, c := range cases {
		It(c.Name, func() {
			var out interface{}
			utils.Die0(json.Unmarshal([]byte(c.Input), &out))
			res := GetType(out)
			Expect(res).To(BeEquivalentTo(c.Expected))
		})
	}
})
