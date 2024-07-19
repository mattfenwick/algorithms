package taxes

import (
	"math"

	. "github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

func RunBracketTests() {
	Describe("GetTax", func() {
		tenK := 10000
		b := Bracket{
			RawBracket: &RawBracket{
				Rate: 10,
				Max:  10000,
			},
			Start: 5000,
			End:   &tenK,
		}
		It("handles empty bracket", func() {
			gomega.Expect(b.GetTax(45)).To(gomega.BeEquivalentTo(0))
		})
		It("handles partially full bracket", func() {
			gomega.Expect(b.GetTax(7500)).To(gomega.BeEquivalentTo(250))
		})
		It("handles completely full bracket", func() {
			gomega.Expect(b.GetTax(10000)).To(gomega.BeEquivalentTo(500))
			gomega.Expect(b.GetTax(10001)).To(gomega.BeEquivalentTo(500))
		})
		It("handles over full bracket", func() {
			gomega.Expect(b.GetTax(100_000)).To(gomega.BeEquivalentTo(500))
		})

		b2 := Bracket{
			RawBracket: &RawBracket{
				Rate: 10,
				Max:  math.MaxInt,
			},
			Start: 5000,
			End:   nil,
		}
		It("handles bracket without max", func() {
			gomega.Expect(b2.GetTax(45)).To(gomega.BeEquivalentTo(0))
			gomega.Expect(b2.GetTax(5000)).To(gomega.BeEquivalentTo(0))
			gomega.Expect(b2.GetTax(5001)).To(gomega.BeEquivalentTo(0))
			gomega.Expect(b2.GetTax(7500)).To(gomega.BeEquivalentTo(250))
			gomega.Expect(b2.GetTax(10_000)).To(gomega.BeEquivalentTo(500))
			gomega.Expect(b2.GetTax(100_000)).To(gomega.BeEquivalentTo(9500))
		})
	})
}
