package pratt

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestPratt(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Pratt Suite")
}

func O(s string, args ...Node) Node {
	return Op(s, len(args), args...)
}

type Case struct {
	Name  string
	Input string
	Error string
	Node  Node
}

var _ = Describe("Operator parsing", func() {
	cases := []*Case{
		{"handles number", "3", "",
			Num("3")},
		{"handles binop", "3 + 4", "",
			O("+",
				Num("3"),
				Num("4")),
		},
		{"handles increasing precedence", "3 + 4 * 5", "",
			O("+",
				Num("3"),
				O("*",
					Num("4"),
					Num("5")))},
		{"handles decreasing precedence", "3 * 4 + 5", "",
			O("+",
				O("*",
					Num("3"),
					Num("4")),
				Num("5"))},
		{"handles increasing and decreasing precedence (part 1)", "3 + 4 * 5 ^ 6 * 7", "",
			O("+",
				Num("3"),
				O("*",
					O("*",
						Num("4"),
						O("^",
							Num("5"),
							Num("6"))),
					Num("7")))},
		{"handles increasing and decreasing precedence (part 2)", "3 + 4 * 5 ^ 6 + 7", "",
			O("+",
				O("+",
					Num("3"),
					O("*",
						Num("4"),
						O("^",
							Num("5"),
							Num("6")))),
				Num("7"))},
		{"handles left-associativity", "3 + 4 + 5", "",
			O("+",
				O("+",
					Num("3"),
					Num("4")),
				Num("5"))},
		{"handles right-associativity", "3 ** 4 ** 5", "",
			O("**",
				Num("3"),
				O("**",
					Num("4"),
					Num("5")))},
		{"handles simple prefix", "! 3", "",
			O("!",
				Num("3"))},
		{"handles stacked prefix", "! <- ~ 3", "",
			O("!",
				O("<-",
					O("~",
						Num("3"))))},
		{"distinguishes prefix and binary op of same name", "- 3 - - 4", "",
			O("-",
				O("-", Num("3")),
				O("-", Num("4")))},
		{"handles precedence: lower prefix vs higher binary", "pre-low 3 - 4", "",
			O("pre-low",
				O("-",
					Num("3"),
					Num("4")))},
		{"handles precedence: higher prefix vs lower binary", "! 3 - 4", "",
			O("-",
				O("!", Num("3")),
				Num("4"))},
		{"handles prefix vs. binary left-associativity", "pre-mid 3 bin-mid-left 4", "",
			O("bin-mid-left",
				O("pre-mid", Num("3")),
				Num("4"))},
		{"refuses to handle associativity mismatch of prefix vs. binary", "pre-mid 3 bin-mid-right 4",
			"unable to handle same precedence but different associativity: bin-mid-right vs pre-mid",
			nil},
	}

	for _, c := range cases {
		It(c.Name, func() {
			node, err := ParseString(c.Input)
			if c.Error == "" {
				Expect(err).To(BeNil())
				Expect(node).To(BeEquivalentTo(c.Node))
			} else {
				Expect(err).To(MatchError(c.Error))
				Expect(node).To(BeNil())
			}
		})
	}
})
