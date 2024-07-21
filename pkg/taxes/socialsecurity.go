package taxes

import "github.com/mattfenwick/collections/pkg/builtin"

// some social security resources:
//
//	https://www.ssa.gov/oact/cola/cbb.html#:~:text=For%20earnings%20in%202024%2C%20this,for%20employees%20and%20employers%2C%20each.
//	https://www.irs.gov/taxtopics/tc751
func EstimateSocialSecurityTax(income *Income) []*SocialSecurityTax {
	var out []*SocialSecurityTax

	yearConstants := TaxYears[income.Year]

	for _, i := range income.IncomeSources {
		if i.IncomeType == IncomeTypeWage {
			taxableAmount := builtin.Min(i.Amount, yearConstants.SocialSecurityLimit)
			tax := int(float32(taxableAmount) * yearConstants.SocialSecurityRate)
			out = append(out, &SocialSecurityTax{
				Description:   i.Description,
				WageIncome:    i.Amount,
				TaxableAmount: taxableAmount,
				Tax:           tax,
			})
		}
	}
	return out
}

type SocialSecurityTax struct {
	Description   string
	WageIncome    int
	TaxableAmount int
	Tax           int
}
