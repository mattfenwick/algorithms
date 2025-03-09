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
		Expect(ParseString("3")).To(BeEquivalentTo(NewNumNode("3")))
	})
	It("handles binop", func() {
		Expect(ParseString("3 + 4")).To(BeEquivalentTo(
			NewOpNode("+",
				NewNumNode("3"),
				NewNumNode("4")),
		))
	})
	It("handles increasing precedence", func() {
		Expect(ParseString("3 + 4 * 5")).To(BeEquivalentTo(
			NewOpNode("+",
				NewNumNode("3"),
				NewOpNode("*",
					NewNumNode("4"),
					NewNumNode("5"))),
		))
	})
	It("handles decreasing precedence", func() {
		Expect(ParseString("3 * 4 + 5")).To(BeEquivalentTo(
			NewOpNode("+",
				NewOpNode("*",
					NewNumNode("3"),
					NewNumNode("4")),
				NewNumNode("5")),
		))
	})
	It("handles increasing and decreasing precedence (part 1)", func() {
		Expect(ParseString("3 + 4 * 5 ^ 6 * 7")).To(BeEquivalentTo(
			NewOpNode("+",
				NewNumNode("3"),
				NewOpNode("*",
					NewOpNode("*",
						NewNumNode("4"),
						NewOpNode("^",
							NewNumNode("5"),
							NewNumNode("6"))),
					NewNumNode("7"))),
		))
	})
	It("handles increasing and decreasing precedence (part 2)", func() {
		Expect(ParseString("3 + 4 * 5 ^ 6 + 7")).To(BeEquivalentTo(
			NewOpNode("+",
				NewOpNode("+",
					NewNumNode("3"),
					NewOpNode("*",
						NewNumNode("4"),
						NewOpNode("^",
							NewNumNode("5"),
							NewNumNode("6")))),
				NewNumNode("7")),
		))
	})
	It("handles left-associativity", func() {
		Expect(ParseString("3 + 4 + 5")).To(BeEquivalentTo(
			NewOpNode("+",
				NewOpNode("+",
					NewNumNode("3"),
					NewNumNode("4")),
				NewNumNode("5")),
		))
	})
	It("handles right-associativity", func() {
		Expect(ParseString("3 ** 4 ** 5")).To(BeEquivalentTo(
			NewOpNode("**",
				NewNumNode("3"),
				NewOpNode("**",
					NewNumNode("4"),
					NewNumNode("5"))),
		))
	})
})
