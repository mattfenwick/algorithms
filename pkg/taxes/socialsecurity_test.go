package taxes

import (
	. "github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

func RunSocialSecurityTests() {
	Describe("Social security", func() {
		It("sub limit -- wages", func() {
			i := &Income{
				Year:   2024,
				Status: FilingStatusSingle,
				IncomeSources: []*IncomeSource{
					{Description: "job1", IncomeType: IncomeTypeWage, Amount: 50_000},
				},
				Deduction: TaxYear2024.ByStatus[FilingStatusSingle].StandardDeduction,
			}
			o := []*SocialSecurityTax{
				{WageIncome: 50_000, TaxableAmount: 50_000, Tax: 3100},
			}
			gomega.Expect(EstimateSocialSecurityTax(i)).To(gomega.BeEquivalentTo(o))
		})
		It("over limit -- wages", func() {
			i := &Income{
				Year:   2024,
				Status: FilingStatusSingle,
				IncomeSources: []*IncomeSource{
					{Description: "job1", IncomeType: IncomeTypeWage, Amount: 250_000},
				},
				Deduction: TaxYear2024.ByStatus[FilingStatusSingle].StandardDeduction,
			}
			o := []*SocialSecurityTax{
				{WageIncome: 250_000, TaxableAmount: 168_600, Tax: 10453},
			}
			gomega.Expect(EstimateSocialSecurityTax(i)).To(gomega.BeEquivalentTo(o))
		})
		It("ignored -- investment", func() {
			i := &Income{
				Year:   2024,
				Status: FilingStatusSingle,
				IncomeSources: []*IncomeSource{
					{Description: "inv1", IncomeType: IncomeTypeLongTerm, Amount: 50_000},
				},
				Deduction: TaxYear2024.ByStatus[FilingStatusSingle].StandardDeduction,
			}
			gomega.Expect(EstimateSocialSecurityTax(i)).To(gomega.BeEmpty())
		})
		It("separate jobs have tax imposed equally", func() {
			i := &Income{
				Year:   2024,
				Status: FilingStatusMarriedJointly,
				IncomeSources: []*IncomeSource{
					{Description: "job1", IncomeType: IncomeTypeWage, Amount: 50_000},
					{Description: "job2", IncomeType: IncomeTypeWage, Amount: 250_000},
				},
				Deduction: TaxYear2024.ByStatus[FilingStatusMarriedJointly].StandardDeduction,
			}
			o := []*SocialSecurityTax{
				{WageIncome: 50_000, TaxableAmount: 50_000, Tax: 3100},
				{WageIncome: 250_000, TaxableAmount: 168_600, Tax: 10453},
			}
			gomega.Expect(EstimateSocialSecurityTax(i)).To(gomega.BeEquivalentTo(o))
		})
	})
}
