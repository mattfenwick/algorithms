package taxes

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
	baseWageIncome, additionalWageIncome, niitIncome := income.MedicareIncome()
	yearConstants, _ := income.GetTaxConstants()

	marginalRate := Rate_0Percent
	if niitIncome > 0 {
		marginalRate = yearConstants.NetInvestmentTaxRate
	} else if additionalWageIncome > 0 {
		marginalRate = yearConstants.MedicareAdditionalRate
	} else if baseWageIncome > 0 {
		marginalRate = yearConstants.MedicareBaseRate
	}
	return &MedicareTax{
		BaseWageIncome:       baseWageIncome,
		BaseWageTax:          yearConstants.MedicareBaseRate.GetTax(baseWageIncome),
		AdditionalWageIncome: additionalWageIncome,
		AdditionalWageTax:    yearConstants.MedicareAdditionalRate.GetTax(additionalWageIncome),
		NiitIncome:           niitIncome,
		NiitTax:              yearConstants.NetInvestmentTaxRate.GetTax(niitIncome),
		MarginalRate:         marginalRate,
	}
}

type MedicareTax struct {
	BaseWageIncome       int64
	BaseWageTax          int64
	AdditionalWageIncome int64
	AdditionalWageTax    int64
	NiitIncome           int64
	NiitTax              int64
	MarginalRate         *TaxRate
}
