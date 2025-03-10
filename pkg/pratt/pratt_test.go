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
		{"number", "3", "", "3"},
		{"binop", "3 + 4", "", "( 3 + 4 )"},
		{"increasing precedence", "3 + 4 * 5", "", "( 3 + ( 4 * 5 ) )"},
		{"decreasing precedence", "3 * 4 + 5", "", "( ( 3 * 4 ) + 5 )"},
		{"increasing and decreasing precedence (part 1)",
			"3 + 4 * 5 ^ 6 * 7",
			"",
			"( 3 + ( ( 4 * ( 5 ^ 6 ) ) * 7 ) )"},
		{"increasing and decreasing precedence (part 2)",
			"3 + 4 * 5 ^ 6 + 7",
			"",
			"( ( 3 + ( 4 * ( 5 ^ 6 ) ) ) + 7 )"},
		{"left-associativity", "3 + 4 + 5", "",
			"( ( 3 + 4 ) + 5 )"},
		{"right-associativity", "3 ** 4 ** 5", "",
			"( 3 ** ( 4 ** 5 ) )"},

		// prefix
		{"simple prefix", "! 3", "", "( ! 3 )"},
		{"stacked prefix", "! <- ~ 3", "", "( ! ( <- ( ~ 3 ) ) )"},
		{"distinguishes prefix and binary op of same name", "- 3 - - 4", "",
			"( ( - 3 ) - ( - 4 ) )"},
		{"precedence: lower prefix vs higher binary", "pre-low 3 - 4", "",
			"( pre-low ( 3 - 4 ) )"},
		{"precedence: higher prefix vs lower binary", "! 3 - 4", "",
			"( ( ! 3 ) - 4 )"},
		{"precedence: lower to higher prefix vs medium binary", "pre-low ! 3 - 4", "",
			"( pre-low ( ( ! 3 ) - 4 ) )"},
		{"precedence: higher to lower prefix vs medium binary", "! pre-low 3 - 4", "",
			"( ! ( pre-low ( 3 - 4 ) ) )"},
		{"prefix vs. binary left-associativity (part 1)", "pre-mid 3 bin-mid-left 4", "",
			"( ( pre-mid 3 ) bin-mid-left 4 )"},
		{"prefix vs. binary left-associativity (part 2)",
			"pre-mid pre-mid pre-mid 3 bin-mid-left 4",
			"",
			"( ( pre-mid ( pre-mid ( pre-mid 3 ) ) ) bin-mid-left 4 )"},
		{"refuses to handle associativity mismatch of prefix vs. binary", "pre-mid 3 bin-mid-right 4",
			"unable to handle same precedence but different associativity: bin-mid-right vs pre-mid",
			""},

		// postfix
		{"simple postfix", "3 ;", "", "( 3 ; )"},
		{"lower postfix vs. higher binary", "3 + 4 post-low", "", "( ( 3 + 4 ) post-low )"},
		{"higher postfix vs. lower binary", "3 + 4 ;", "", "( 3 + ( 4 ; ) )"},
		{"higher prefix vs. lower postfix", "! 3 post-low", "", "( ( ! 3 ) post-low )"},
		{"lower prefix vs. higher postfix", "pre-low 3 ;", "", "( pre-low ( 3 ; ) )"},
		{"postfix vs. binary right-associativity (part 1)",
			"3 bin-mid-right 4 post-mid",
			"",
			"( 3 bin-mid-right ( 4 post-mid ) )"},
		{"postfix vs. binary right-associativity (part 2)",
			"3 bin-mid-right 4 post-mid post-mid post-mid",
			"",
			"( 3 bin-mid-right ( ( ( 4 post-mid ) post-mid ) post-mid ) )"},
		{"refuses to handle associativity mismatch of binary vs. postfix", "3 bin-mid-left 4 post-mid",
			"unable to handle same precedence but different associativity: post-mid vs bin-mid-left",
			""},

		// ternary
		{"simple ternary", "3 ? 4 : 5", "", "( 3 ? 4 : 5 )"},
		{"right-associative ternary", "3 ? 4 : 5 ? 6 : 7", "", "( 3 ? 4 : ( 5 ? 6 : 7 ) )"},
		{"left-associative ternary", "3 <l 4 <lb 5 <l 6 <lb 7", "", "( ( 3 <l 4 <lb 5 ) <l 6 <lb 7 )"},
		{"ternary precedence (part 1)", "3 <l 4 <lb 5 <l1 6 <l1b 7", "", "( 3 <l 4 <lb ( 5 <l1 6 <l1b 7 ) )"},
		{"ternary precedence (part 2)", "3 <l1 4 <l1b 5 <l 6 <lb 7", "", "( ( 3 <l1 4 <l1b 5 ) <l 6 <lb 7 )"},
		{"prefix/ternary precedence (part 1)", "! 3 ? 4 : 5", "", "( ( ! 3 ) ? 4 : 5 )"},
		{"prefix/ternary precedence (part 2)", "pre-low 3 <l1 4 <l1b 5", "", "( pre-low ( 3 <l1 4 <l1b 5 ) )"},
		{"ternary/postfix precedence (part 1)", "3 ? 4 : 5 ;", "", "( 3 ? 4 : ( 5 ; ) )"},
		{"ternary/postfix precedence (part 2)", "3 <l1 4 <l1b 5 post-low", "", "( ( 3 <l1 4 <l1b 5 ) post-low )"},
		{"binary/ternary precedence (part 1)", "2 + 3 ? 4 : 5", "", "( ( 2 + 3 ) ? 4 : 5 )"},
		{"binary/ternary precedence (part 2)", "2 <-0 3 <l1 4 <l1b 5", "", "( 2 <-0 ( 3 <l1 4 <l1b 5 ) )"},
		{"ternary/binary precedence (part 1)", "3 ? 4 : 5 + 6", "", "( 3 ? 4 : ( 5 + 6 ) )"},
		{"ternary/binary precedence (part 2)", "3 <l1 4 <l1b 5 <-0 6", "", "( ( 3 <l1 4 <l1b 5 ) <-0 6 )"},
	}

	for _, c := range cases {
		It(c.Name, func() {
			node, err := ParseString(c.Input)
			if c.Error == "" {
				Expect(err).To(BeNil())
				parensString := Parens(node)
				Expect(parensString).To(BeEquivalentTo(c.Parenthesized))
			} else {
				Expect(err).To(MatchError(c.Error))
				Expect(node).To(BeNil())
			}
		})
	}
})
