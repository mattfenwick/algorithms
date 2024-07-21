package taxes

import "github.com/mattfenwick/collections/pkg/builtin"

// EstimateMedicareTax takes into account:
//
//   - filing status
//
//   - for wages: the base medicare rate, plus the threshold above which an additional rate is imposed
//
//   - for capital gains: the threshold above which NIIT is imposed
//
//     As far as I can tell, the way this works is:
//
//   - wage income is taxed "first", with two brackets: 1.45%, and 2.35%
//
//   - investment income is taxed "second", with two brackets: 0%, and 3.8%
//     "second" meaning: start with the wage income from 0, then add the investment income on top of that
//
//   - no deductions are applicable
//
//   - income is taxed either as wage or as investment, but never as both
//
//     Some resources:
//
//   - https://www.irs.gov/taxtopics/tc559
func EstimateMedicareTax(income *Income) *MedicareTax {
	wageIncome, capitalGains := income.WageIncome(), income.InvestmentIncome()

	yearConstants := TaxYears[income.Year]
	statusConstants := yearConstants.ByStatus[income.Status]

	baseWageIncome := builtin.Min(statusConstants.MedicareAdditionalThreshold, wageIncome)
	baseWageTax := yearConstants.MedicareBaseRate.GetTax(int64(baseWageIncome))

	// NIIT income -- investments go "on top"
	// so start by figuring out how much income is over the threshold
	niitIncome := builtin.Max(0, wageIncome+capitalGains-statusConstants.MedicareAdditionalThreshold)
	// then figure out how much over-the-threshold is investment income
	niitIncome = builtin.Min(niitIncome, capitalGains)
	niitTax := yearConstants.NetInvestmentTaxRate.GetTax(niitIncome)

	additionalWageIncome := builtin.Max(0, wageIncome-statusConstants.MedicareAdditionalThreshold)
	additionalWageTax := yearConstants.MedicareBaseRate.GetTax(additionalWageIncome) + yearConstants.MedicareAdditionalRate.GetTax(additionalWageIncome)

	return &MedicareTax{
		BaseWageIncome:       baseWageIncome,
		BaseWageTax:          baseWageTax,
		AdditionalWageIncome: additionalWageIncome,
		AdditionalWageTax:    additionalWageTax,
		NiitIncome:           niitIncome,
		NiitTax:              niitTax,
	}
}

type MedicareTax struct {
	BaseWageIncome       int64
	BaseWageTax          int64
	AdditionalWageIncome int64
	AdditionalWageTax    int64
	NiitIncome           int64
	NiitTax              int64
}
