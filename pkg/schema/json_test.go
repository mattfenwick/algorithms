package schema

import (
	"encoding/json"
	"strings"
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
	Expected any
}

// var _ = Describe("Type parsing", func() {
// 	cases := []*Case{
// 		{"grab bag", `{"a":1,"b":true,"c":[1,1.2,"d",{}],"e":null}`, "", &Type{Dict: map[string]*Type{
// 			"a": {Number: true},
// 			"b": {Bool: true},
// 			"c": {Array: Union([]*Type{{Number: true}, {String: true}, {Dict: map[string]*Type{}}})},
// 			"e": {Null: true},
// 		}}},
// 		{"array of arrays -- unions", `[[null, true], [null, true], [3, false]]`, "", &Type{Array: Union([]*Type{
// 			{Array: Union([]*Type{{Bool: true}, {Null: true}})},
// 			{Array: Union([]*Type{{Bool: true}, {Number: true}})},
// 		})}},
// 	}

// 	for _, c := range cases {
// 		It(c.Name, func() {
// 			var out any
// 			utils.Die0(json.Unmarshal([]byte(c.Input), &out))
// 			res := GetType(out)
// 			Expect(res).To(BeEquivalentTo(c.Expected))
// 			Expect(res.PrettyPrint()).To(BeEquivalentTo("?"))
// 		})
// 	}
// })

var _ = Describe("Traverse", func() {
	cases := []*Case{
		{"grab bag", `{"a":1,"b":true,"c":[1,1.2,"d",{}],"e":null}`, "",
			`: object (1)
["a"]: number (1)
["b"]: bool (1)
["c"]: array (1)
["c"][*]: number (2)
["c"][*]: object (1)
["c"][*]: string (1)
["e"]: null (1)`,
		},
		{"array of arrays", `[[null, true], [null, true], [3, false], "abc"]`, "",
			`: array (1)
[*]: array (3)
[*]: string (1)
[*][*]: bool (3)
[*][*]: null (2)
[*][*]: number (1)`,
		},
		{"objects in array", `{"a": [{"b": 1, "c": null, "e": 2}, {"b": [], "d": false, "e": 3}]}`, "",
			`: object (1)
["a"]: array (1)
["a"][*]: object (2)
["a"][*]["b"]: array (1)
["a"][*]["b"]: number (1)
["a"][*]["c"]: null (1)
["a"][*]["d"]: bool (1)
["a"][*]["e"]: number (2)`,
		},
		{"case-insensitive sort", `{"a":1,"b":2,"c":3,"A":4,"B":5,"C":6}`, "",
			`: object (1)
["A"]: number (1)
["a"]: number (1)
["B"]: number (1)
["b"]: number (1)
["C"]: number (1)
["c"]: number (1)`,
		},
	}

	for _, c := range cases {
		It(c.Name, func() {
			var out any
			utils.Die0(json.Unmarshal([]byte(c.Input), &out))
			t := NewTraverser()
			t.Add(out)
			res := strings.Join(t.Lines(), "\n")
			Expect(res).To(BeEquivalentTo(c.Expected))
		})
	}
})
