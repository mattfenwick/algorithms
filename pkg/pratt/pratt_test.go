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

var _ = Describe("Parse -- binary precedence", func() {
	It("handles number", func() {
		Expect(ParseString("3")).To(BeEquivalentTo(Num("3")))
	})
	It("handles binop", func() {
		Expect(ParseString("3 + 4")).To(BeEquivalentTo(
			Op("+",
				Num("3"),
				Num("4")),
		))
	})
	It("handles increasing precedence", func() {
		Expect(ParseString("3 + 4 * 5")).To(BeEquivalentTo(
			Op("+",
				Num("3"),
				Op("*",
					Num("4"),
					Num("5"))),
		))
	})
	It("handles decreasing precedence", func() {
		Expect(ParseString("3 * 4 + 5")).To(BeEquivalentTo(
			Op("+",
				Op("*",
					Num("3"),
					Num("4")),
				Num("5")),
		))
	})
	It("handles increasing and decreasing precedence (part 1)", func() {
		Expect(ParseString("3 + 4 * 5 ^ 6 * 7")).To(BeEquivalentTo(
			Op("+",
				Num("3"),
				Op("*",
					Op("*",
						Num("4"),
						Op("^",
							Num("5"),
							Num("6"))),
					Num("7"))),
		))
	})
	It("handles increasing and decreasing precedence (part 2)", func() {
		Expect(ParseString("3 + 4 * 5 ^ 6 + 7")).To(BeEquivalentTo(
			Op("+",
				Op("+",
					Num("3"),
					Op("*",
						Num("4"),
						Op("^",
							Num("5"),
							Num("6")))),
				Num("7")),
		))
	})
	It("handles left-associativity", func() {
		Expect(ParseString("3 + 4 + 5")).To(BeEquivalentTo(
			Op("+",
				Op("+",
					Num("3"),
					Num("4")),
				Num("5")),
		))
	})
	It("handles right-associativity", func() {
		Expect(ParseString("3 ** 4 ** 5")).To(BeEquivalentTo(
			Op("**",
				Num("3"),
				Op("**",
					Num("4"),
					Num("5"))),
		))
	})
})
