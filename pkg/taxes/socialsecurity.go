package taxes

import "github.com/mattfenwick/collections/pkg/builtin"

func EstimateSocialSecurityTax(income *Income) []*SocialSecurityTax {
	var out []*SocialSecurityTax

	yearConstants := TaxYears[income.Year]

	for _, i := range income.IncomeSources {
		if i.IncomeType == IncomeTypeWage {
			taxableAmount := builtin.Min(i.Amount, yearConstants.SocialSecurityLimit)
			tax := int(float32(taxableAmount) * yearConstants.SocialSecurityRate)
			out = append(out, &SocialSecurityTax{
				WageIncome:    i.Amount,
				TaxableAmount: taxableAmount,
				Tax:           tax,
			})
		}
	}
	return out
}

type SocialSecurityTax struct {
	WageIncome    int
	TaxableAmount int
	Tax           int
}
