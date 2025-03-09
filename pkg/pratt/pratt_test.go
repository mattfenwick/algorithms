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

type Case struct {
	Name          string
	Input         string
	Error         string
	Parenthesized string
}

var _ = Describe("Operator parsing", func() {
	cases := []*Case{
		{"handles number", "3", "", "3"},
		{"handles binop", "3 + 4", "", "( 3 + 4 )"},
		{"handles increasing precedence", "3 + 4 * 5", "", "( 3 + ( 4 * 5 ) )"},
		{"handles decreasing precedence", "3 * 4 + 5", "", "( ( 3 * 4 ) + 5 )"},
		{"handles increasing and decreasing precedence (part 1)",
			"3 + 4 * 5 ^ 6 * 7",
			"",
			"( 3 + ( ( 4 * ( 5 ^ 6 ) ) * 7 ) )"},
		{"handles increasing and decreasing precedence (part 2)",
			"3 + 4 * 5 ^ 6 + 7",
			"",
			"( ( 3 + ( 4 * ( 5 ^ 6 ) ) ) + 7 )"},
		{"handles left-associativity", "3 + 4 + 5", "",
			"( ( 3 + 4 ) + 5 )"},
		{"handles right-associativity", "3 ** 4 ** 5", "",
			"( 3 ** ( 4 ** 5 ) )"},
		{"handles simple prefix", "! 3", "", "( ! 3 )"},
		{"handles stacked prefix", "! <- ~ 3", "", "( ! ( <- ( ~ 3 ) ) )"},
		{"distinguishes prefix and binary op of same name", "- 3 - - 4", "",
			"( ( - 3 ) - ( - 4 ) )"},
		{"handles precedence: lower prefix vs higher binary", "pre-low 3 - 4", "",
			"( pre-low ( 3 - 4 ) )"},
		{"handles precedence: higher prefix vs lower binary", "! 3 - 4", "",
			"( ( ! 3 ) - 4 )"},
		{"handles precedence: lower to higher prefix vs medium binary", "pre-low ! 3 - 4", "",
			"( pre-low ( ( ! 3 ) - 4 ) )"},
		{"handles precedence: higher to lower prefix vs medium binary", "! pre-low 3 - 4", "",
			"( ! ( pre-low ( 3 - 4 ) ) )"},
		{"handles prefix vs. binary left-associativity", "pre-mid 3 bin-mid-left 4", "",
			"( ( pre-mid 3 ) bin-mid-left 4 )"},
		{"refuses to handle associativity mismatch of prefix vs. binary", "pre-mid 3 bin-mid-right 4",
			"unable to handle same precedence but different associativity: bin-mid-right vs pre-mid",
			""},
	}

	for _, c := range cases {
		It(c.Name, func() {
			node, err := ParseString(c.Input)
			if c.Error == "" {
				parensString := Parens(node)
				Expect(err).To(BeNil())
				Expect(parensString).To(BeEquivalentTo(c.Parenthesized))
			} else {
				Expect(err).To(MatchError(c.Error))
				Expect(node).To(BeNil())
			}
		})
	}
})
