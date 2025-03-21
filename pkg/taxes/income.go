package taxes

import (
	"github.com/mattfenwick/collections/pkg/builtin"
	"github.com/mattfenwick/collections/pkg/slice"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
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
	IncomeTypeWage      IncomeType = "IncomeTypeWage"
	IncomeTypeShortTerm IncomeType = "IncomeTypeShortTerm"
	IncomeTypeLongTerm  IncomeType = "IncomeTypeLongTerm"
)

type IncomeSource struct {
	Description       string
	IncomeType        IncomeType
	Amount            int64
	NonTaxableWages   int64
	NonTaxablePayroll int64
}

type Income struct {
	Description   string
	Year          int
	Status        FilingStatus
	IncomeSources []*IncomeSource
	Adjustments   int64 // see https://www.irs.gov/pub/irs-pdf/f1040s1.pdf => 401K, HSA, etc.
	Deduction     *int64
}

func (i *Income) GetTaxConstants() (*TaxYearConstants, *TaxStatusConstants) {
	yearConstants := TaxYears[i.Year]
	statusConstants := yearConstants.ByStatus[i.Status]
	return yearConstants, statusConstants
}

func (i *Income) MedicareIncome() (int64, int64, int64) {
	wageIncome := int64(0)
	for _, s := range i.IncomeSources {
		if s.IncomeType != IncomeTypeWage {
			continue
		}
		wageIncome += builtin.Max(0, s.Amount-s.NonTaxablePayroll)
	}
	capitalGains := i.InvestmentIncome()

	_, statusConstants := i.GetTaxConstants()
	baseWageIncome := builtin.Min(statusConstants.MedicareAdditionalThreshold, wageIncome)

	// NIIT income -- investments go "on top"
	// so start by figuring out how much income is over the threshold
	niitIncome := builtin.Max(0, wageIncome+capitalGains-statusConstants.MedicareAdditionalThreshold)
	// then figure out how much over-the-threshold is investment income
	niitIncome = builtin.Min(niitIncome, capitalGains)

	additionalWageIncome := builtin.Max(0, wageIncome-statusConstants.MedicareAdditionalThreshold)
	return baseWageIncome, additionalWageIncome, niitIncome
}

func (i *Income) SocialSecurityIncome() []int64 {
	var out []int64
	for _, s := range i.IncomeSources {
		if s.IncomeType == IncomeTypeWage {
			out = append(out, builtin.Min(s.Amount, TaxYears[i.Year].SocialSecurityLimit))
		}
	}
	return out
}

func (i *Income) GetGrossIncome() int64 {
	return slice.Sum(slice.Map(func(is *IncomeSource) int64 {
		return is.Amount
	}, i.IncomeSources))
}

func (i *Income) GetGrossIncomeLessNonTaxable() int64 {
	return slice.Sum(slice.Map(func(is *IncomeSource) int64 {
		// TODO is this adjustment right?
		return builtin.Max(0, is.Amount-is.NonTaxableWages)
	}, i.IncomeSources))
}

func (i *Income) GetAdjustedGrossIncome() int64 {
	return i.GetGrossIncomeLessNonTaxable() - i.Adjustments
}

func (i *Income) GetTaxableIncome() int64 {
	return builtin.Max(0, i.GetAdjustedGrossIncome()-i.GetDeduction())
}

func (i *Income) GetDeduction() int64 {
	standardDeduction := TaxYears[i.Year].ByStatus[i.Status].StandardDeduction
	if i.Deduction != nil {
		deduction := *i.Deduction
		if deduction < standardDeduction {
			logrus.Warningf("nonsensical deduction -- claimed deduction is less than standard deduction (%d < %d)", deduction, standardDeduction)
		}
		return deduction
	}
	return standardDeduction
}

func (i *Income) WageIncome() int64 {
	out := int64(0)
	for _, s := range i.IncomeSources {
		if s.IncomeType == IncomeTypeWage {
			out += builtin.Max(0, s.Amount-s.NonTaxableWages)
		}
	}
	return out
}

func (i *Income) ShortTermCapitalGainIncome() int64 {
	out := int64(0)
	for _, s := range i.IncomeSources {
		if s.IncomeType == IncomeTypeShortTerm {
			out += s.Amount
		}
	}
	return out
}

func (i *Income) LongTermCapitalGainIncome() int64 {
	out := int64(0)
	for _, s := range i.IncomeSources {
		if s.IncomeType == IncomeTypeLongTerm {
			out += s.Amount
		}
	}
	return out
}

func (i *Income) InvestmentIncome() int64 {
	out := int64(0)
	for _, s := range i.IncomeSources {
		if s.IncomeType == IncomeTypeShortTerm || s.IncomeType == IncomeTypeLongTerm {
			out += s.Amount
		}
	}
	return out
}
