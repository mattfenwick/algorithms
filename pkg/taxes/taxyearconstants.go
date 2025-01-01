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
// - the medicare threshold limit is per-person according to https://www.irs.gov/taxtopics/tc751 -- but TODO, am I misunderstanding that?
type TaxYearConstants struct {
	SocialSecurityLimit    int64
	SocialSecurityRate     *TaxRate
	MedicareBaseRate       *TaxRate
	MedicareAdditionalRate *TaxRate
	NetInvestmentTaxRate   *TaxRate
	ByStatus               map[FilingStatus]*TaxStatusConstants
}

var (
	TaxYear2023 = &TaxYearConstants{
		SocialSecurityLimit:    160_200,
		SocialSecurityRate:     &TaxRate{Rate: 62, Divisor: 1000},   // .062  = 6.2%
		MedicareBaseRate:       &TaxRate{Rate: 145, Divisor: 10000}, // .0145 = 1.45%
		MedicareAdditionalRate: &TaxRate{Rate: 235, Divisor: 10000}, // .0235 = 2.35%
		NetInvestmentTaxRate:   &TaxRate{Rate: 38, Divisor: 1000},   // .038  = 3.8%
		ByStatus: map[FilingStatus]*TaxStatusConstants{
			FilingStatusSingle: {
				OrdinaryIncomeBrackets: NewStatusBrackets([]*RawBracket{
					{Rate10Percent, 11_000},
					{Rate12Percent, 44_725},
					{Rate22Percent, 95_375},
					{Rate24Percent, 182_100},
					{Rate32Percent, 231_250},
					{Rate35Percent, 578_125},
					{Rate37Percent, math.MaxInt64},
				}),
				LTCGIncomeBrackets: NewStatusBrackets([]*RawBracket{
					{Rate_0Percent, 44_625},
					{Rate15Percent, 492_300},
					{Rate20Percent, math.MaxInt64},
				}),
				StandardDeduction:           13_850,
				MedicareAdditionalThreshold: 200_000,
			},
			FilingStatusMarriedJointly: {
				OrdinaryIncomeBrackets: NewStatusBrackets([]*RawBracket{
					{Rate10Percent, 22_000},
					{Rate12Percent, 89_450},
					{Rate22Percent, 190_750},
					{Rate24Percent, 364_200},
					{Rate32Percent, 462_500},
					{Rate35Percent, 693_750},
					{Rate37Percent, math.MaxInt64},
				}),
				LTCGIncomeBrackets: NewStatusBrackets([]*RawBracket{
					{Rate_0Percent, 89_250},
					{Rate15Percent, 553_850},
					{Rate20Percent, math.MaxInt64},
				}),
				StandardDeduction:           27_700,
				MedicareAdditionalThreshold: 250_000,
			},
			FilingStatusMarriedSeparately: {
				OrdinaryIncomeBrackets: NewStatusBrackets([]*RawBracket{
					{Rate10Percent, 11_000},
					{Rate12Percent, 44_725},
					{Rate22Percent, 95_375},
					{Rate24Percent, 182_100},
					{Rate32Percent, 231_250},
					{Rate35Percent, 346_875},
					{Rate37Percent, math.MaxInt64},
				}),
				LTCGIncomeBrackets: NewStatusBrackets([]*RawBracket{
					{Rate_0Percent, 44_625},
					{Rate15Percent, 276_900},
					{Rate20Percent, math.MaxInt64},
				}),
				StandardDeduction:           13_850,
				MedicareAdditionalThreshold: 125_000,
			},
			FilingStatusHeadOfHouseHold: {
				OrdinaryIncomeBrackets: NewStatusBrackets([]*RawBracket{
					{Rate10Percent, 15_700},
					{Rate12Percent, 59_63100850},
					{Rate22Percent, 95_350},
					{Rate24Percent, 182_100},
					{Rate32Percent, 231_250},
					{Rate35Percent, 578_100},
					{Rate37Percent, math.MaxInt64},
				}),
				LTCGIncomeBrackets: NewStatusBrackets([]*RawBracket{
					{Rate_0Percent, 59_750},
					{Rate15Percent, 523_050},
					{Rate20Percent, math.MaxInt64},
				}),
				StandardDeduction:           20_800,
				MedicareAdditionalThreshold: 200_000,
			},
		},
	}

	TaxYear2024 = &TaxYearConstants{
		SocialSecurityLimit:    168_600,
		SocialSecurityRate:     &TaxRate{Rate: 62, Divisor: 1000},   // .062  = 6.2%
		MedicareBaseRate:       &TaxRate{Rate: 145, Divisor: 10000}, // .0145 = 1.45%
		MedicareAdditionalRate: &TaxRate{Rate: 235, Divisor: 10000}, // .0235 = 2.35%
		NetInvestmentTaxRate:   &TaxRate{Rate: 38, Divisor: 1000},   // .038  = 3.8%
		ByStatus: map[FilingStatus]*TaxStatusConstants{
			FilingStatusSingle: {
				OrdinaryIncomeBrackets: NewStatusBrackets([]*RawBracket{
					{Rate10Percent, 11600},
					{Rate12Percent, 47150},
					{Rate22Percent, 100525},
					{Rate24Percent, 191950},
					{Rate32Percent, 243725},
					{Rate35Percent, 609350},
					{Rate37Percent, math.MaxInt64},
				}),
				LTCGIncomeBrackets: NewStatusBrackets([]*RawBracket{
					{Rate_0Percent, 47025},
					{Rate15Percent, 518900},
					{Rate20Percent, math.MaxInt64},
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
					{Rate37Percent, math.MaxInt64},
				}),
				LTCGIncomeBrackets: NewStatusBrackets([]*RawBracket{
					{Rate_0Percent, 94050},
					{Rate15Percent, 583750},
					{Rate20Percent, math.MaxInt64},
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
					{Rate37Percent, math.MaxInt64},
				}),
				LTCGIncomeBrackets: NewStatusBrackets([]*RawBracket{
					{Rate_0Percent, 47025},
					{Rate15Percent, 291850},
					{Rate20Percent, math.MaxInt64},
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
					{Rate37Percent, math.MaxInt64},
				}),
				LTCGIncomeBrackets: NewStatusBrackets([]*RawBracket{
					{Rate_0Percent, 63000},
					{Rate15Percent, 551350},
					{Rate20Percent, math.MaxInt64},
				}),
				StandardDeduction:           21900,
				MedicareAdditionalThreshold: 200_000,
			},
		},
	}

	TaxYear2025 = &TaxYearConstants{
		SocialSecurityLimit:    176_100,
		SocialSecurityRate:     &TaxRate{Rate: 62, Divisor: 1000},   // .062  = 6.2%
		MedicareBaseRate:       &TaxRate{Rate: 145, Divisor: 10000}, // .0145 = 1.45%
		MedicareAdditionalRate: &TaxRate{Rate: 235, Divisor: 10000}, // .0235 = 2.35%
		NetInvestmentTaxRate:   &TaxRate{Rate: 38, Divisor: 1000},   // .038  = 3.8%
		ByStatus: map[FilingStatus]*TaxStatusConstants{
			FilingStatusSingle: {
				OrdinaryIncomeBrackets: NewStatusBrackets([]*RawBracket{
					{Rate10Percent, 11925},
					{Rate12Percent, 48475},
					{Rate22Percent, 103350},
					{Rate24Percent, 197300},
					{Rate32Percent, 250525},
					{Rate35Percent, 626350},
					{Rate37Percent, math.MaxInt64},
				}),
				LTCGIncomeBrackets: NewStatusBrackets([]*RawBracket{
					{Rate_0Percent, 48350},
					{Rate15Percent, 533400},
					{Rate20Percent, math.MaxInt64},
				}),
				StandardDeduction:           15000,
				MedicareAdditionalThreshold: 200_000,
			},
			FilingStatusMarriedJointly: {
				OrdinaryIncomeBrackets: NewStatusBrackets([]*RawBracket{
					{Rate10Percent, 23850},
					{Rate12Percent, 96950},
					{Rate22Percent, 206700},
					{Rate24Percent, 394600},
					{Rate32Percent, 501050},
					{Rate35Percent, 751600},
					{Rate37Percent, math.MaxInt64},
				}),
				LTCGIncomeBrackets: NewStatusBrackets([]*RawBracket{
					{Rate_0Percent, 96700},
					{Rate15Percent, 600_050},
					{Rate20Percent, math.MaxInt64},
				}),
				StandardDeduction:           30000,
				MedicareAdditionalThreshold: 250_000,
			},
			FilingStatusMarriedSeparately: {
				OrdinaryIncomeBrackets: NewStatusBrackets([]*RawBracket{
					{Rate10Percent, 11925},
					{Rate12Percent, 48475},
					{Rate22Percent, 103350},
					{Rate24Percent, 197300},
					{Rate32Percent, 250525},
					{Rate35Percent, 375800},
					{Rate37Percent, math.MaxInt64},
				}),
				LTCGIncomeBrackets: NewStatusBrackets([]*RawBracket{
					{Rate_0Percent, 48350},
					{Rate15Percent, 300_000},
					{Rate20Percent, math.MaxInt64},
				}),
				StandardDeduction:           15000,
				MedicareAdditionalThreshold: 125_000,
			},
			FilingStatusHeadOfHouseHold: {
				OrdinaryIncomeBrackets: NewStatusBrackets([]*RawBracket{
					{Rate10Percent, 17000},
					{Rate12Percent, 64850},
					{Rate22Percent, 103350},
					{Rate24Percent, 197300},
					{Rate32Percent, 250500},
					{Rate35Percent, 626350},
					{Rate37Percent, math.MaxInt64},
				}),
				LTCGIncomeBrackets: NewStatusBrackets([]*RawBracket{
					{Rate_0Percent, 64750},
					{Rate15Percent, 566700},
					{Rate20Percent, math.MaxInt64},
				}),
				StandardDeduction:           22500,
				MedicareAdditionalThreshold: 200_000,
			},
		},
	}

	TaxYears = map[int]*TaxYearConstants{
		2023: TaxYear2023,
		2024: TaxYear2024,
		2025: TaxYear2025,
	}
)
