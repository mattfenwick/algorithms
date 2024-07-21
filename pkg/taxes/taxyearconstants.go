package taxes

import "math"

type TaxStatusConstants struct {
	OrdinaryIncomeBrackets      *StatusBrackets
	LTCGIncomeBrackets          *StatusBrackets
	StandardDeduction           int64
	MedicareAdditionalThreshold int64
}

// TaxYearConstants organizes key constants for each tax year
//
// Some notes:
// - half of social security is typically paid by the employer
// - the social security limit is per-person
// - half of medicare is typically paid by the employer
// - the medicare threshold limit is per-person according to https://www.irs.gov/taxtopics/tc751
type TaxYearConstants struct {
	SocialSecurityLimit    int64
	SocialSecurityRate     *TaxRate
	MedicareBaseRate       *TaxRate
	MedicareAdditionalRate *TaxRate
	NetInvestmentTaxRate   *TaxRate
	ByStatus               map[FilingStatus]*TaxStatusConstants
}

var (
	TaxYear2024 = &TaxYearConstants{
		SocialSecurityLimit:    168_600,
		SocialSecurityRate:     &TaxRate{Rate: 62, Divisor: 1000},   // .062 = 6.2%
		MedicareBaseRate:       &TaxRate{Rate: 145, Divisor: 10000}, // .0145 = 1.45%
		MedicareAdditionalRate: &TaxRate{Rate: 9, Divisor: 1000},    // .009 = .9%
		NetInvestmentTaxRate:   &TaxRate{Rate: 38, Divisor: 1000},   // .038 = 3.8%
		ByStatus: map[FilingStatus]*TaxStatusConstants{
			FilingStatusSingle: {
				OrdinaryIncomeBrackets: NewStatusBrackets([]*RawBracket{
					{Rate10Percent, 11600},
					{Rate12Percent, 47150},
					{Rate22Percent, 100525},
					{Rate24Percent, 191950},
					{Rate32Percent, 243725},
					{Rate35Percent, 609350},
					{Rate37Percent, math.MaxInt},
				}),
				LTCGIncomeBrackets: NewStatusBrackets([]*RawBracket{
					{Rate_0Percent, 47025},
					{Rate15Percent, 518900},
					{Rate20Percent, math.MaxInt},
				}),
				StandardDeduction:           14600,
				MedicareAdditionalThreshold: 200_000,
			},
			FilingStatusMarriedJointly: {
				OrdinaryIncomeBrackets: NewStatusBrackets([]*RawBracket{
					{Rate10Percent, 23200},
					{Rate12Percent, 94300},
					{Rate22Percent, 201050},
					{Rate24Percent, 383900},
					{Rate32Percent, 487450},
					{Rate35Percent, 731200},
					{Rate37Percent, math.MaxInt},
				}),
				LTCGIncomeBrackets: NewStatusBrackets([]*RawBracket{
					{Rate_0Percent, 94050},
					{Rate15Percent, 583750},
					{Rate20Percent, math.MaxInt},
				}),
				StandardDeduction:           29200,
				MedicareAdditionalThreshold: 250_000,
			},
			FilingStatusMarriedSeparately: {
				OrdinaryIncomeBrackets: NewStatusBrackets([]*RawBracket{
					{Rate10Percent, 11600},
					{Rate12Percent, 47150},
					{Rate22Percent, 100525},
					{Rate24Percent, 191950},
					{Rate32Percent, 243725},
					{Rate35Percent, 365600},
					{Rate37Percent, math.MaxInt},
				}),
				LTCGIncomeBrackets: NewStatusBrackets([]*RawBracket{
					{Rate_0Percent, 47025},
					{Rate15Percent, 291850},
					{Rate20Percent, math.MaxInt},
				}),
				StandardDeduction:           14600,
				MedicareAdditionalThreshold: 125_000,
			},
			FilingStatusHeadOfHouseHold: {
				OrdinaryIncomeBrackets: NewStatusBrackets([]*RawBracket{
					{Rate10Percent, 16550},
					{Rate12Percent, 63100},
					{Rate22Percent, 100500},
					{Rate24Percent, 191950},
					{Rate32Percent, 243700},
					{Rate35Percent, 609350},
					{Rate37Percent, math.MaxInt},
				}),
				LTCGIncomeBrackets: NewStatusBrackets([]*RawBracket{
					{Rate_0Percent, 63000},
					{Rate15Percent, 551350},
					{Rate20Percent, math.MaxInt},
				}),
				StandardDeduction:           21900,
				MedicareAdditionalThreshold: 200_000,
			},
		},
	}

	TaxYears = map[int]*TaxYearConstants{
		2024: TaxYear2024,
	}
)
