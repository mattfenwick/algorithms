package taxes

import (
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
	Description string
	IncomeType  IncomeType
	Amount      int64
}

type Income struct {
	Year          int
	Status        FilingStatus
	IncomeSources []*IncomeSource
	Deduction     *int64
}

func (i *Income) GetDeduction() int64 {
	standardDeduction := TaxYears[i.Year].ByStatus[i.Status].StandardDeduction
	if i.Deduction != nil {
		deduction := *i.Deduction
		if deduction < standardDeduction {
			logrus.Warningf("nonsenical deduction -- claimed deduction is less than standard deduction (%d < %d)", deduction, standardDeduction)
		}
		return deduction
	}
	return standardDeduction
}

func (i *Income) WageIncome() int64 {
	out := int64(0)
	for _, s := range i.IncomeSources {
		if s.IncomeType == IncomeTypeWage {
			out += s.Amount
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
