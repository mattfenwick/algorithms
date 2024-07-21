package taxes

import (
	. "github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

func RunMedicareTests() {
	Describe("Medicare", func() {
		It("sub limit -- wages", func() {
			i := &Income{
				Year:   2024,
				Status: FilingStatusSingle,
				IncomeSources: []*IncomeSource{
					{Description: "job1", IncomeType: IncomeTypeWage, Amount: 50_000},
				},
				Deduction: TaxYear2024.ByStatus[FilingStatusSingle].StandardDeduction,
			}
			o := &MedicareTax{
				BaseWageIncome: 50_000,
				BaseWageTax:    725,
			}
			gomega.Expect(EstimateMedicareTax(i)).To(gomega.BeEquivalentTo(o))
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
			o := &MedicareTax{
				BaseWageIncome:       200_000,
				BaseWageTax:          2900,
				AdditionalWageIncome: 50_000,
				AdditionalWageTax:    1175,
			}
			gomega.Expect(EstimateMedicareTax(i)).To(gomega.BeEquivalentTo(o))
		})
		It("sub limit -- investment", func() {
			i := &Income{
				Year:   2024,
				Status: FilingStatusSingle,
				IncomeSources: []*IncomeSource{
					{Description: "inv1", IncomeType: IncomeTypeLongTerm, Amount: 50_000},
				},
				Deduction: TaxYear2024.ByStatus[FilingStatusSingle].StandardDeduction,
			}
			o := &MedicareTax{}
			gomega.Expect(EstimateMedicareTax(i)).To(gomega.BeEquivalentTo(o))
		})
		It("over limit -- investment", func() {
			i := &Income{
				Year:   2024,
				Status: FilingStatusSingle,
				IncomeSources: []*IncomeSource{
					{Description: "inv1", IncomeType: IncomeTypeLongTerm, Amount: 250_000},
				},
				Deduction: TaxYear2024.ByStatus[FilingStatusSingle].StandardDeduction,
			}
			o := &MedicareTax{
				NiitIncome: 50_000,
				NiitTax:    1900,
			}
			gomega.Expect(EstimateMedicareTax(i)).To(gomega.BeEquivalentTo(o))
		})
		It("sub limit -- wages and investment", func() {
			i := &Income{
				Year:   2024,
				Status: FilingStatusSingle,
				IncomeSources: []*IncomeSource{
					{Description: "job1", IncomeType: IncomeTypeWage, Amount: 50_000},
					{Description: "inv1", IncomeType: IncomeTypeLongTerm, Amount: 50_000},
				},
				Deduction: TaxYear2024.ByStatus[FilingStatusSingle].StandardDeduction,
			}
			o := &MedicareTax{
				BaseWageIncome: 50_000,
				BaseWageTax:    725,
			}
			gomega.Expect(EstimateMedicareTax(i)).To(gomega.BeEquivalentTo(o))
		})
		It("over limit -- wages and investment", func() {
			i := &Income{
				Year:   2024,
				Status: FilingStatusSingle,
				IncomeSources: []*IncomeSource{
					{Description: "job1", IncomeType: IncomeTypeWage, Amount: 250_000},
					{Description: "inv1", IncomeType: IncomeTypeLongTerm, Amount: 250_000},
				},
				Deduction: TaxYear2024.ByStatus[FilingStatusSingle].StandardDeduction,
			}
			o := &MedicareTax{
				BaseWageIncome:       200_000,
				BaseWageTax:          2900,
				AdditionalWageIncome: 50_000,
				AdditionalWageTax:    1175,
				NiitIncome:           250_000,
				NiitTax:              9500,
			}
			gomega.Expect(EstimateMedicareTax(i)).To(gomega.BeEquivalentTo(o))
		})
	})
}
