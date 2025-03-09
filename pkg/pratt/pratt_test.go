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

var _ = Describe("Parse -- binary precedence", func() {
	It("handles number", func() {
		Expect(ParseString("3")).To(BeEquivalentTo(Num("3")))
	})
	It("handles binop", func() {
		Expect(ParseString("3 + 4")).To(BeEquivalentTo(
			O("+",
				Num("3"),
				Num("4")),
		))
	})
	It("handles increasing precedence", func() {
		Expect(ParseString("3 + 4 * 5")).To(BeEquivalentTo(
			O("+",
				Num("3"),
				O("*",
					Num("4"),
					Num("5"))),
		))
	})
	It("handles decreasing precedence", func() {
		Expect(ParseString("3 * 4 + 5")).To(BeEquivalentTo(
			O("+",
				O("*",
					Num("3"),
					Num("4")),
				Num("5")),
		))
	})
	It("handles increasing and decreasing precedence (part 1)", func() {
		Expect(ParseString("3 + 4 * 5 ^ 6 * 7")).To(BeEquivalentTo(
			O("+",
				Num("3"),
				O("*",
					O("*",
						Num("4"),
						O("^",
							Num("5"),
							Num("6"))),
					Num("7"))),
		))
	})
	It("handles increasing and decreasing precedence (part 2)", func() {
		Expect(ParseString("3 + 4 * 5 ^ 6 + 7")).To(BeEquivalentTo(
			O("+",
				O("+",
					Num("3"),
					O("*",
						Num("4"),
						O("^",
							Num("5"),
							Num("6")))),
				Num("7")),
		))
	})
	It("handles left-associativity", func() {
		Expect(ParseString("3 + 4 + 5")).To(BeEquivalentTo(
			O("+",
				O("+",
					Num("3"),
					Num("4")),
				Num("5")),
		))
	})
	It("handles right-associativity", func() {
		Expect(ParseString("3 ** 4 ** 5")).To(BeEquivalentTo(
			O("**",
				Num("3"),
				O("**",
					Num("4"),
					Num("5"))),
		))
	})
	It("handles simple prefix", func() {
		Expect(ParseString("! 3")).To(BeEquivalentTo(
			O("!",
				Num("3")),
		))
	})
	It("handles stacked prefix", func() {
		Expect(ParseString("! <- ~ 3")).To(BeEquivalentTo(
			O("!",
				O("<-",
					O("~",
						Num("3")))),
		))
	})
	It("distinguishes prefix and binary op of same name", func() {
		Expect(ParseString("- 3 - - 4")).To(BeEquivalentTo(
			O("-",
				O("-", Num("3")),
				O("-", Num("4"))),
		))
	})
	It("handles precedence: lower prefix vs higher binary", func() {
		Expect(ParseString("pre-low 3 - 4")).To(BeEquivalentTo(
			O("pre-low",
				O("-",
					Num("3"),
					Num("4"))),
		))
	})
	It("handles precedence: higher prefix vs lower binary", func() {
		Expect(ParseString("! 3 - 4")).To(BeEquivalentTo(
			O("-",
				O("!", Num("3")),
				Num("4")),
		))
	})
	It("handles prefix vs. binary left-associativity", func() {
		Expect(ParseString("pre-mid 3 bin-mid-left 4")).To(BeEquivalentTo(
			O("bin-mid-left",
				O("pre-mid", Num("3")),
				Num("4")),
		))
	})
	It("refuses to handle associativity mismatch of prefix vs. binary", func() {
		node, err := ParseString("pre-mid 3 bin-mid-right 4")
		Expect(node).To(BeNil())
		Expect(err).To(MatchError("unable to handle same precedence but different associativity: bin-mid-right vs pre-mid"))
	})
})
