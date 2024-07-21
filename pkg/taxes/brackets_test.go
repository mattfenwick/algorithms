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
			gomega.Expect(b.GetTax(45)).To(gomega.BeEquivalentTo(&BracketTax{0, 0}))
		})
		It("handles partially full bracket", func() {
			gomega.Expect(b.GetTax(7500)).To(gomega.BeEquivalentTo(&BracketTax{2500, 250}))
		})
		It("handles completely full bracket", func() {
			gomega.Expect(b.GetTax(10000)).To(gomega.BeEquivalentTo(&BracketTax{5000, 500}))
			gomega.Expect(b.GetTax(10001)).To(gomega.BeEquivalentTo(&BracketTax{5000, 500}))
		})
		It("handles over full bracket", func() {
			gomega.Expect(b.GetTax(100_000)).To(gomega.BeEquivalentTo(&BracketTax{5000, 500}))
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
			gomega.Expect(b2.GetTax(45)).To(gomega.BeEquivalentTo(&BracketTax{0, 0}))
			gomega.Expect(b2.GetTax(5000)).To(gomega.BeEquivalentTo(&BracketTax{0, 0}))
			gomega.Expect(b2.GetTax(5001)).To(gomega.BeEquivalentTo(&BracketTax{1, 0}))
			gomega.Expect(b2.GetTax(7500)).To(gomega.BeEquivalentTo(&BracketTax{2500, 250}))
			gomega.Expect(b2.GetTax(10_000)).To(gomega.BeEquivalentTo(&BracketTax{5000, 500}))
			gomega.Expect(b2.GetTax(100_000)).To(gomega.BeEquivalentTo(&BracketTax{95000, 9500}))
		})

		ltcgBrackets := NewStatusBrackets([]*RawBracket{
			{0, 94050},
			{15, 583750},
			{20, math.MaxInt},
		}).GetBrackets()
		It("handles LTCG taxes -- 100K/50K", func() {
			gomega.Expect(ltcgBrackets[0].GetLongTermCapitalGainsTax(50_000, 100_000)).To(gomega.BeEquivalentTo(&BracketTax{44_050, 0}))
			gomega.Expect(ltcgBrackets[1].GetLongTermCapitalGainsTax(50_000, 100_000)).To(gomega.BeEquivalentTo(&BracketTax{5950, 892}))
			gomega.Expect(ltcgBrackets[2].GetLongTermCapitalGainsTax(50_000, 100_000)).To(gomega.BeEquivalentTo(&BracketTax{0, 0}))
		})
		It("handles LTCG taxes -- 100K/0K", func() {
			gomega.Expect(ltcgBrackets[0].GetLongTermCapitalGainsTax(0, 100_000)).To(gomega.BeEquivalentTo(&BracketTax{94050, 0}))
			gomega.Expect(ltcgBrackets[1].GetLongTermCapitalGainsTax(0, 100_000)).To(gomega.BeEquivalentTo(&BracketTax{5950, 892}))
			gomega.Expect(ltcgBrackets[2].GetLongTermCapitalGainsTax(0, 100_000)).To(gomega.BeEquivalentTo(&BracketTax{0, 0}))
		})
		It("handles LTCG taxes -- 800K/0K", func() {
			gomega.Expect(ltcgBrackets[0].GetLongTermCapitalGainsTax(0, 800_000)).To(gomega.BeEquivalentTo(&BracketTax{94050, 0}))
			gomega.Expect(ltcgBrackets[1].GetLongTermCapitalGainsTax(0, 800_000)).To(gomega.BeEquivalentTo(&BracketTax{489700, 73455}))
			gomega.Expect(ltcgBrackets[2].GetLongTermCapitalGainsTax(0, 800_000)).To(gomega.BeEquivalentTo(&BracketTax{216250, 43250}))
		})
		It("handles LTCG taxes -- 800K/400K", func() {
			gomega.Expect(ltcgBrackets[0].GetLongTermCapitalGainsTax(400_000, 800_000)).To(gomega.BeEquivalentTo(&BracketTax{0, 0}))
			gomega.Expect(ltcgBrackets[1].GetLongTermCapitalGainsTax(400_000, 800_000)).To(gomega.BeEquivalentTo(&BracketTax{183750, 27562}))
			gomega.Expect(ltcgBrackets[2].GetLongTermCapitalGainsTax(400_000, 800_000)).To(gomega.BeEquivalentTo(&BracketTax{216250, 43250}))
		})
		It("handles LTCG taxes -- 800K/795K", func() {
			gomega.Expect(ltcgBrackets[0].GetLongTermCapitalGainsTax(795_000, 800_000)).To(gomega.BeEquivalentTo(&BracketTax{0, 0}))
			gomega.Expect(ltcgBrackets[1].GetLongTermCapitalGainsTax(795_000, 800_000)).To(gomega.BeEquivalentTo(&BracketTax{0, 0}))
			gomega.Expect(ltcgBrackets[2].GetLongTermCapitalGainsTax(795_000, 800_000)).To(gomega.BeEquivalentTo(&BracketTax{5000, 1000}))
		})
	})
}
