package taxes

import (
	"math"

	. "github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

func RunBracketTests() {
	Describe("GetTax -- ordinary income", func() {
		tenK := int64(10000)
		b := Bracket{
			RawBracket: &RawBracket{
				Rate: Rate10Percent,
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
				Rate: Rate10Percent,
				Max:  math.MaxInt64,
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
	})

	Describe("GetLongTermCapitalGainsTax (total / ordinary)", func() {
		ltcgBrackets := NewStatusBrackets([]*RawBracket{
			{Rate_0Percent, 94050},
			{Rate15Percent, 583750},
			{Rate20Percent, math.MaxInt64},
		}).GetBrackets()
		It("0K/0K", func() {
			gomega.Expect(ltcgBrackets[0].GetLongTermCapitalGainsTax(0, 0)).To(gomega.BeEquivalentTo(&BracketTax{0, 0}))
			gomega.Expect(ltcgBrackets[1].GetLongTermCapitalGainsTax(0, 0)).To(gomega.BeEquivalentTo(&BracketTax{0, 0}))
			gomega.Expect(ltcgBrackets[2].GetLongTermCapitalGainsTax(0, 0)).To(gomega.BeEquivalentTo(&BracketTax{0, 0}))
		})
		It("10K/0K", func() {
			gomega.Expect(ltcgBrackets[0].GetLongTermCapitalGainsTax(10_000, 0)).To(gomega.BeEquivalentTo(&BracketTax{10_000, 0}))
			gomega.Expect(ltcgBrackets[1].GetLongTermCapitalGainsTax(10_000, 0)).To(gomega.BeEquivalentTo(&BracketTax{0, 0}))
			gomega.Expect(ltcgBrackets[2].GetLongTermCapitalGainsTax(10_000, 0)).To(gomega.BeEquivalentTo(&BracketTax{0, 0}))
		})
		It("10K/10K", func() {
			gomega.Expect(ltcgBrackets[0].GetLongTermCapitalGainsTax(10_000, 10_000)).To(gomega.BeEquivalentTo(&BracketTax{0, 0}))
			gomega.Expect(ltcgBrackets[1].GetLongTermCapitalGainsTax(10_000, 10_000)).To(gomega.BeEquivalentTo(&BracketTax{0, 0}))
			gomega.Expect(ltcgBrackets[2].GetLongTermCapitalGainsTax(10_000, 10_000)).To(gomega.BeEquivalentTo(&BracketTax{0, 0}))
		})
		It("10K/5K", func() {
			gomega.Expect(ltcgBrackets[0].GetLongTermCapitalGainsTax(10_000, 5_000)).To(gomega.BeEquivalentTo(&BracketTax{5_000, 0}))
			gomega.Expect(ltcgBrackets[1].GetLongTermCapitalGainsTax(10_000, 5_000)).To(gomega.BeEquivalentTo(&BracketTax{0, 0}))
			gomega.Expect(ltcgBrackets[2].GetLongTermCapitalGainsTax(10_000, 5_000)).To(gomega.BeEquivalentTo(&BracketTax{0, 0}))
		})
		It("100K/50K", func() {
			gomega.Expect(ltcgBrackets[0].GetLongTermCapitalGainsTax(100_000, 50_000)).To(gomega.BeEquivalentTo(&BracketTax{44_050, 0}))
			gomega.Expect(ltcgBrackets[1].GetLongTermCapitalGainsTax(100_000, 50_000)).To(gomega.BeEquivalentTo(&BracketTax{5950, 892}))
			gomega.Expect(ltcgBrackets[2].GetLongTermCapitalGainsTax(100_000, 50_000)).To(gomega.BeEquivalentTo(&BracketTax{0, 0}))
		})
		It("100K/0K", func() {
			gomega.Expect(ltcgBrackets[0].GetLongTermCapitalGainsTax(100_000, 0)).To(gomega.BeEquivalentTo(&BracketTax{94050, 0}))
			gomega.Expect(ltcgBrackets[1].GetLongTermCapitalGainsTax(100_000, 0)).To(gomega.BeEquivalentTo(&BracketTax{5950, 892}))
			gomega.Expect(ltcgBrackets[2].GetLongTermCapitalGainsTax(100_000, 0)).To(gomega.BeEquivalentTo(&BracketTax{0, 0}))
		})
		It("800K/0K", func() {
			gomega.Expect(ltcgBrackets[0].GetLongTermCapitalGainsTax(800_000, 0)).To(gomega.BeEquivalentTo(&BracketTax{94050, 0}))
			gomega.Expect(ltcgBrackets[1].GetLongTermCapitalGainsTax(800_000, 0)).To(gomega.BeEquivalentTo(&BracketTax{489700, 73455}))
			gomega.Expect(ltcgBrackets[2].GetLongTermCapitalGainsTax(800_000, 0)).To(gomega.BeEquivalentTo(&BracketTax{216250, 43250}))
		})
		It("800K/400K", func() {
			gomega.Expect(ltcgBrackets[0].GetLongTermCapitalGainsTax(800_000, 400_000)).To(gomega.BeEquivalentTo(&BracketTax{0, 0}))
			gomega.Expect(ltcgBrackets[1].GetLongTermCapitalGainsTax(800_000, 400_000)).To(gomega.BeEquivalentTo(&BracketTax{183750, 27562}))
			gomega.Expect(ltcgBrackets[2].GetLongTermCapitalGainsTax(800_000, 400_000)).To(gomega.BeEquivalentTo(&BracketTax{216250, 43250}))
		})
		It("800K/795K", func() {
			gomega.Expect(ltcgBrackets[0].GetLongTermCapitalGainsTax(800_000, 795_000)).To(gomega.BeEquivalentTo(&BracketTax{0, 0}))
			gomega.Expect(ltcgBrackets[1].GetLongTermCapitalGainsTax(800_000, 795_000)).To(gomega.BeEquivalentTo(&BracketTax{0, 0}))
			gomega.Expect(ltcgBrackets[2].GetLongTermCapitalGainsTax(800_000, 795_000)).To(gomega.BeEquivalentTo(&BracketTax{5000, 1000}))
		})
	})
}
