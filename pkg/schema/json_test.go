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
		{"grab bag", `{"a":1,"b":true,"c":[1,1.2,"d",{}],"e":null}`, "", &Type{Dict: map[string]*Type{
			"a": {Number: true},
			"b": {Bool: true},
			"c": {Array: Union([]*Type{{Number: true}, {String: true}, {Dict: map[string]*Type{}}})},
			"e": {Null: true},
		}}},
		{"array of arrays -- unions", `[[null, true], [null, true], [3, false]]`, "", &Type{Array: Union([]*Type{
			{Array: Union([]*Type{{Bool: true}, {Null: true}})},
			{Array: Union([]*Type{{Bool: true}, {Number: true}})},
		})}},
	}

	for _, c := range cases {
		It(c.Name, func() {
			var out any
			utils.Die0(json.Unmarshal([]byte(c.Input), &out))
			res := GetType(out)
			Expect(res).To(BeEquivalentTo(c.Expected))
			Expect(res.PrettyPrint()).To(BeEquivalentTo("?"))
		})
	}
})
