package taxes

import "math"

type TaxStatusConstants struct {
	OrdinaryIncomeBrackets *StatusBrackets
	LTCGIncomeBrackets     *StatusBrackets
	StandardDeduction      int
	NetInvestmentTaxLimit  int
	// MedicareThreshold todo name ?
}

// TaxYearConstants organizes key constants for each tax year
//
// Some notes:
// - half of social security is typically paid by the employer
// - the social security limit is per-person
// - half of medicare is typically paid by the employer
// - the medicare threshold limit is per-person according to https://www.irs.gov/taxtopics/tc751
type TaxYearConstants struct {
	SocialSecurityLimit         int
	SocialSecurityRate          float32
	MedicareBaseRate            float32
	MedicareAdditionalThreshold int
	MedicareAdditionalRate      float32
	NetInvestmentTaxRate        float32
	ByStatus                    map[FilingStatus]*TaxStatusConstants
}

var (
	TaxYear2024 = &TaxYearConstants{
		SocialSecurityLimit:         168_600,
		SocialSecurityRate:          0.062,
		MedicareBaseRate:            0.0145,
		MedicareAdditionalThreshold: 200_000,
		MedicareAdditionalRate:      0.009,
		NetInvestmentTaxRate:        0.038,
		ByStatus: map[FilingStatus]*TaxStatusConstants{
			FilingStatusSingle: {
				OrdinaryIncomeBrackets: NewStatusBrackets([]*RawBracket{
					{10, 11600},
					{12, 47150},
					{22, 100525},
					{24, 191950},
					{32, 243725},
					{35, 609350},
					{37, math.MaxInt},
				}),
				LTCGIncomeBrackets: NewStatusBrackets([]*RawBracket{
					{0, 47025},
					{15, 518900},
					{20, math.MaxInt},
				}),
				StandardDeduction:     14600,
				NetInvestmentTaxLimit: 200_000,
			},
			FilingStatusMarriedJointly: {
				OrdinaryIncomeBrackets: NewStatusBrackets([]*RawBracket{
					{10, 23200},
					{12, 94300},
					{22, 201050},
					{24, 383900},
					{32, 487450},
					{35, 731200},
					{37, math.MaxInt},
				}),
				LTCGIncomeBrackets: NewStatusBrackets([]*RawBracket{
					{0, 94050},
					{15, 583750},
					{20, math.MaxInt},
				}),
				StandardDeduction:     29200,
				NetInvestmentTaxLimit: 250_000,
			},
			FilingStatusMarriedSeparately: {
				OrdinaryIncomeBrackets: NewStatusBrackets([]*RawBracket{
					{10, 11600},
					{12, 47150},
					{22, 100525},
					{24, 191950},
					{32, 243725},
					{35, 365600},
					{37, math.MaxInt},
				}),
				LTCGIncomeBrackets: NewStatusBrackets([]*RawBracket{
					{0, 47025},
					{15, 291850},
					{20, math.MaxInt},
				}),
				StandardDeduction:     14600,
				NetInvestmentTaxLimit: 125_000,
			},
			FilingStatusHeadOfHouseHold: {
				OrdinaryIncomeBrackets: NewStatusBrackets([]*RawBracket{
					{10, 16550},
					{12, 63100},
					{22, 100500},
					{24, 191950},
					{32, 243700},
					{35, 609350},
					{37, math.MaxInt},
				}),
				LTCGIncomeBrackets: NewStatusBrackets([]*RawBracket{
					{0, 63000},
					{15, 551350},
					{20, math.MaxInt},
				}),
				StandardDeduction:     21900,
				NetInvestmentTaxLimit: 200_000,
			},
		},
	}

	TaxYears = map[int]*TaxYearConstants{
		2024: TaxYear2024,
	}
)
