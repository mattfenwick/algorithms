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
			NewOpNode("+", []Node{NewNumNode("3"), NewNumNode("4")}),
		))
	})
	It("handles increasing precedence", func() {
		Expect(ParseString("3 + 4 * 5")).To(BeEquivalentTo(
			NewOpNode("+", []Node{
				NewNumNode("3"),
				NewOpNode("*", []Node{
					NewNumNode("4"),
					NewNumNode("5")})}),
		))
	})
	It("handles decreasing precedence", func() {
		Expect(ParseString("3 * 4 + 5")).To(BeEquivalentTo(
			NewOpNode("+", []Node{
				NewOpNode("*", []Node{
					NewNumNode("3"),
					NewNumNode("4")}),
				NewNumNode("5")}),
		))
	})
})
