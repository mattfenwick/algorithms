package taxes

import (
	"github.com/mattfenwick/collections/pkg/builtin"
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

type TaxEstimateBracket struct {
	BracketTax *BracketTax
	Bracket    *Bracket
}

type SocialSecurityTax struct {
	Income int
	Tax    int
}

type MedicareTax struct {
	BaseIncome           int
	BaseTax              int
	IncomeAboveThreshold int
	TaxAboveThreshold    int
}

type TaxEstimate struct {
	Income                   *Income
	TaxYearConstants         *TaxYearConstants
	StatusConstants          *TaxStatusConstants
	OrdinaryTaxes            []*TaxEstimateBracket
	NiitIncome               int
	NiitTax                  int
	AdditionalMedicareIncome int
	AdditionalMedicareTax    int
	TotalIncome              int
	IncomeAfterDeduction     int
}

func EstimateTaxes(income *Income) *TaxEstimate {
	yearConstants, ok := TaxYears[income.Year]
	if !ok {
		panic(errors.Errorf("no tax constant info for year %d", income.Year))
	}
	statusConstants := yearConstants.ByStatus[income.Status]
	// TODO decision -- include deduction as first bracket?

	totalIncome := slice.Sum(slice.Map(func(i *IncomeSource) int { return i.Amount }, income.IncomeSources))
	capitalGains := slice.Sum(
		slice.Map(func(i *IncomeSource) int { return i.Amount },
			slice.Filter(func(i *IncomeSource) bool {
				switch i.IncomeType {
				case IncomeTypeW2:
					return false
				default:
					return true
				}
			}, income.IncomeSources)))
	wageIncome := totalIncome - capitalGains

	incomeAfterDeduction := totalIncome - income.Deduction

	var bracketTaxes []*TaxEstimateBracket
	for _, b := range statusConstants.OrdinaryIncomeBrackets.GetBrackets() {
		bracketTaxes = append(bracketTaxes, &TaxEstimateBracket{b.GetTax(incomeAfterDeduction), b})
	}

	// net investment tax -- 3.8% on investment income over $250K
	//   https://www.irs.gov/taxtopics/tc559
	niitInvestmentIncome := builtin.Max(0, incomeAfterDeduction-statusConstants.NetInvestmentTaxLimit)
	niitIncome := builtin.Min(capitalGains, niitInvestmentIncome)
	niitTax := int(yearConstants.NetInvestmentTaxRate * float32(niitIncome))

	//   note that NIIT is an alternative to 0.9%:
	//     NIIT applies to non-wage income only
	//     0.9% applies to wage income only
	// TODO is this calculated correctly?  using wageIncome
	additionalMedicareIncome := builtin.Max(0, wageIncome-statusConstants.MedicareAdditionalThreshold)
	additionalMedicareTax := int(yearConstants.MedicareAdditionalRate * float32(additionalMedicareIncome))

	// social security
	//   https://www.ssa.gov/oact/cola/cbb.html#:~:text=For%20earnings%20in%202024%2C%20this,for%20employees%20and%20employers%2C%20each.
	//   https://www.irs.gov/taxtopics/tc751

	// TODO LTCG

	return &TaxEstimate{
		Income:                   income,
		TaxYearConstants:         yearConstants,
		StatusConstants:          statusConstants,
		OrdinaryTaxes:            bracketTaxes,
		NiitIncome:               niitIncome,
		NiitTax:                  niitTax,
		AdditionalMedicareIncome: additionalMedicareIncome,
		AdditionalMedicareTax:    additionalMedicareTax,
		TotalIncome:              totalIncome,
		IncomeAfterDeduction:     incomeAfterDeduction,
	}
}

// TODO social security

// TODO medicare

// TODO investment income
