package taxes

import (
	"math"

	"github.com/mattfenwick/collections/pkg/slice"
	"github.com/pkg/errors"
)

type FilingStatus string

const (
	FilingStatusSingle            FilingStatus = "FilingStatusSingle"
	FilingStatusMarriedJointly    FilingStatus = "FilingStatusMarriedJointly"
	FilingStatusMarriedSeparately FilingStatus = "FilingStatusMarriedSeparately"
	FilingStatusHeadOfHouseHold   FilingStatus = "FilingStatusHeadOfHouseHold"
)

func (f FilingStatus) ToString() string {
	switch f {
	case FilingStatusSingle:
		return "single"
	case FilingStatusMarriedJointly:
		return "married/jointly"
	case FilingStatusMarriedSeparately:
		return "married/separately"
	case FilingStatusHeadOfHouseHold:
		return "head of household"
	default:
		panic(errors.Errorf("invalid status: %s", f))
	}
}

type TaxStatusConstants struct {
	OrdinaryIncomeBrackets *StatusBrackets
	LTCGIncomeBrackets     *StatusBrackets
	StandardDeduction      int
	SocialSecurityLimit    int
	// InvestmentTaxLimit todo name ?
	// MedicareThreshold todo name ?
}

type TaxYearConstants struct {
	ByStatus map[FilingStatus]*TaxStatusConstants
}

var (
	TaxYear2024 = &TaxYearConstants{
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
				StandardDeduction: 14600,
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
				StandardDeduction: 29200,
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
				StandardDeduction: 14600,
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
				StandardDeduction: 21900,
			},
		},
	}

	TaxYears = map[int]*TaxYearConstants{
		2024: TaxYear2024,
	}
)

type IncomeType string

const (
	IncomeTypeW2        IncomeType = "IncomeTypeW2"
	IncomeTypeShortTerm IncomeType = "IncomeTypeShortTerm"
	IncomeTypeLongTerm  IncomeType = "IncomeTypeLongTerm"
)

type IncomeSource struct {
	Description string
	IncomeType  IncomeType
	Amount      int
}

type Income struct {
	Year          int
	Status        FilingStatus
	IncomeSources []*IncomeSource
	Deduction     int
}

type TaxResult struct {
	Income   *Income
	Brackets []*BracketResult
}

type BracketResult struct {
	Income  int
	Tax     int
	Bracket *Bracket
}

type TaxEstimate struct {
	Income               *Income
	TaxYearConstants     *TaxYearConstants
	StatusConstants      *TaxStatusConstants
	TotalIncome          int
	IncomeAfterDeduction int
}

func EstimateTaxes(income *Income) *TaxEstimate {
	yearConstants, ok := TaxYears[income.Year]
	if !ok {
		panic(errors.Errorf("no tax constant info for year %d", income.Year))
	}
	statusConstants := yearConstants.ByStatus[income.Status]
	// TODO decision -- include deduction as first bracket?

	totalIncome := slice.Sum(slice.Map(func(i *IncomeSource) int { return i.Amount }, income.IncomeSources))

	incomeAfterDeduction := totalIncome - income.Deduction

	// brackets := statusConstants.OrdinaryIncomeBrackets
	// remainingIncome := incomeAfterDeduction
	// for remainingIncome > 0 {
	// 	taxed := builtin.Min()
	// }

	// TODO LTCG
	// TODO Social Security
	// TODO medicare
	// TODO investment

	return &TaxEstimate{
		Income:               income,
		TaxYearConstants:     yearConstants,
		StatusConstants:      statusConstants,
		TotalIncome:          totalIncome,
		IncomeAfterDeduction: incomeAfterDeduction,
	}
}

// TODO social security

// TODO medicare

// TODO investment income
