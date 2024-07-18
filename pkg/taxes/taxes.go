package taxes

import (
	"math"

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

type Bracket struct {
	Rate int
	Max  int
}

type TaxStatusConstants struct {
	OrdinaryIncomeBrackets []*Bracket
	LTCGIncomeBrackets     []*Bracket
	StandardDeduction      int
}

type TaxYearConstants struct {
	ByStatus map[FilingStatus]*TaxStatusConstants
}

var (
	TaxYear2024 = &TaxYearConstants{
		ByStatus: map[FilingStatus]*TaxStatusConstants{
			FilingStatusSingle: &TaxStatusConstants{
				OrdinaryIncomeBrackets: []*Bracket{
					{10, 11600},
					{12, 47150},
					{22, 100525},
					{24, 191950},
					{32, 243725},
					{35, 609350},
					{37, math.MaxInt}, // TODO actually there's no max here
				},
				LTCGIncomeBrackets: []*Bracket{
					{0, 47025},
					{15, 518900},
					{20, math.MaxInt},
				},
				StandardDeduction: 14600,
			},
			FilingStatusMarriedJointly: &TaxStatusConstants{
				OrdinaryIncomeBrackets: []*Bracket{
					{10, 23200},
					{12, 94300},
					{22, 201050},
					{24, 383900},
					{32, 487450},
					{35, 731200},
					{37, math.MaxInt}, // TODO actually there's no max here
				},
				LTCGIncomeBrackets: []*Bracket{
					{0, 94050},
					{15, 583750},
					{20, math.MaxInt},
				},
				StandardDeduction: 29200,
			},
			FilingStatusMarriedSeparately: &TaxStatusConstants{
				OrdinaryIncomeBrackets: []*Bracket{
					{10, 11600},
					{12, 47150},
					{22, 100525},
					{24, 191950},
					{32, 243725},
					{35, 365600},
					{37, math.MaxInt}, // TODO actually there's no max here
				},
				LTCGIncomeBrackets: []*Bracket{
					{0, 47025},
					{15, 291850},
					{20, math.MaxInt},
				},
				StandardDeduction: 14600,
			},
			FilingStatusHeadOfHouseHold: &TaxStatusConstants{
				OrdinaryIncomeBrackets: []*Bracket{
					{10, 16550},
					{12, 63100},
					{22, 100500},
					{24, 191950},
					{32, 243700},
					{35, 609350},
					{37, math.MaxInt}, // TODO actually there's no max here
				},
				LTCGIncomeBrackets: []*Bracket{
					{0, 63000},
					{15, 551350},
					{20, math.MaxInt},
				},
				StandardDeduction: 21900,
			},
		},
	}
)

// TODO standard deductions

// TODO capital gains brackets

// TODO social security

// TODO medicare

// TODO investment income
