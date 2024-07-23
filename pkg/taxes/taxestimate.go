package taxes

import (
	"fmt"
	"strings"

	"github.com/mattfenwick/collections/pkg/builtin"
	"github.com/mattfenwick/collections/pkg/slice"
	"github.com/pkg/errors"
)

type TaxEstimateBracket struct {
	BracketTax *BracketTax
	Bracket    *Bracket
}

type TaxEstimate struct {
	Income               *Income
	TaxYearConstants     *TaxYearConstants
	StatusConstants      *TaxStatusConstants
	OrdinaryTaxes        []*TaxEstimateBracket
	LTCGTaxes            []*TaxEstimateBracket
	Medicare             *MedicareTax
	SocialSecurity       []*SocialSecurityTax
	TotalIncome          int64
	MedicareIncome       int64
	SocialSecurityIncome int64
	AdjustedGrossIncome  int64
}

func EstimateTaxes(income *Income) *TaxEstimate {
	yearConstants, ok := TaxYears[income.Year]
	if !ok {
		panic(errors.Errorf("no tax constant info for year %d", income.Year))
	}
	statusConstants := yearConstants.ByStatus[income.Status]

	ordinaryIncome := income.GetTaxableIncome() - income.LongTermCapitalGainIncome()
	totalIncome := slice.Sum(slice.Map(func(i *IncomeSource) int64 { return i.Amount }, income.IncomeSources))

	var ordinaryTaxes []*TaxEstimateBracket
	for _, b := range statusConstants.OrdinaryIncomeBrackets.GetBrackets() {
		ordinaryTaxes = append(ordinaryTaxes, &TaxEstimateBracket{b.GetTax(ordinaryIncome), b})
	}

	// LTCG
	//   TODO question -- how does LTCG interact with deduction?
	var ltcgTaxes []*TaxEstimateBracket
	for _, b := range statusConstants.LTCGIncomeBrackets.GetBrackets() {
		ltcgTaxes = append(ltcgTaxes,
			&TaxEstimateBracket{
				b.GetLongTermCapitalGainsTax(income.GetTaxableIncome(), ordinaryIncome),
				b})
	}

	return &TaxEstimate{
		Income:           income,
		TaxYearConstants: yearConstants,
		StatusConstants:  statusConstants,
		OrdinaryTaxes:    ordinaryTaxes,
		LTCGTaxes:        ltcgTaxes,
		Medicare:         EstimateMedicareTax(income),
		SocialSecurity:   EstimateSocialSecurityTax(income),
		TotalIncome:      totalIncome,
		// IncomeAfterDeduction: incomeAfterDeduction,
	}
}

func (e *TaxEstimate) PrettyPrint() {
	// income/input table
	medicareBase, medicareAddnl, medicareNiit := e.Income.MedicareIncome()
	var socialSecurity, wages, nonTaxableWages, nonTaxablePayroll []int64
	for _, s := range e.Income.IncomeSources {
		if s.IncomeType != IncomeTypeWage {
			continue
		}
		socialSecurity = append(socialSecurity, builtin.Min(e.TaxYearConstants.SocialSecurityLimit, s.Amount))
		wages = append(wages, s.Amount)
		nonTaxableWages = append(nonTaxableWages, s.NonTaxableWages)
		nonTaxablePayroll = append(nonTaxablePayroll, s.NonTaxablePayroll)
	}
	inputTable := NewTable([]string{"Key", "Value"},
		[]string{"Status", e.Income.Status.ToString()},
		[]string{"Year", intToString(e.Income.Year)},
		[]string{"Wages", strings.Join(slice.Map(intToString, wages), "\n")},
		[]string{"Nontaxable wages", strings.Join(slice.Map(intToString, nonTaxableWages), "\n")},
		[]string{"Short term", intToString(e.Income.ShortTermCapitalGainIncome())},
		[]string{"Long term", intToString(e.Income.LongTermCapitalGainIncome())},
		[]string{"Gross", intToString(e.Income.GetGrossIncome())},
		[]string{"Gross less nontaxable", intToString(e.Income.GetGrossIncomeLessNonTaxable())},
		[]string{"Adjustments", intToString(e.Income.Adjustments)},
		[]string{"AGI", intToString(e.Income.GetAdjustedGrossIncome())},
		[]string{"Deduction", intToString(e.Income.GetDeduction())},
		[]string{"Taxable income", intToString(e.Income.GetTaxableIncome())},
		[]string{"Nontaxable payroll", strings.Join(slice.Map(intToString, nonTaxablePayroll), "\n")},
		[]string{"Medicare base", intToString(medicareBase)},
		[]string{"Medicare additional", intToString(medicareAddnl)},
		[]string{"Medicare NIIT", intToString(medicareNiit)},
		[]string{"Social security", strings.Join(slice.Map(intToString, socialSecurity), "\n")},
	)
	fmt.Printf("income/input: \n%s\n\n", inputTable.ToFormattedTable())

	// taxes
	taxTable := NewTable([]string{"Tax", "Description", "Rate (%)", "Taxable amount", "Tax"})

	// medicare tax
	taxTable.AddRow([]string{"Medicare", "Base wage",
		fmt.Sprintf("%.2f", e.TaxYearConstants.MedicareBaseRate.ToDebugPercentage()),
		intToString(e.Medicare.BaseWageIncome),
		intToString(e.Medicare.BaseWageTax)})
	taxTable.AddRow([]string{"Medicare", "Addnl wage",
		fmt.Sprintf("%.2f", e.TaxYearConstants.MedicareAdditionalRate.ToDebugPercentage()),
		intToString(e.Medicare.AdditionalWageIncome),
		intToString(e.Medicare.AdditionalWageTax)})
	taxTable.AddRow([]string{"Medicare", "NIIT",
		fmt.Sprintf("%.2f", e.TaxYearConstants.NetInvestmentTaxRate.ToDebugPercentage()),
		intToString(e.Medicare.NiitIncome),
		intToString(e.Medicare.NiitTax)})

	// social security tax
	ssTotal := int64(0)
	for _, t := range e.SocialSecurity {
		taxTable.AddRow([]string{
			"Social security", t.Description,
			fmt.Sprintf("%.1f", e.TaxYearConstants.SocialSecurityRate.ToDebugPercentage()),
			intToString(t.TaxableAmount),
			intToString(t.Tax),
		})
		ssTotal += t.Tax
	}

	// ordinary tax
	ordinaryTotal := int64(0)
	ordinaryMarginalRate := Rate_0Percent
	for _, t := range e.OrdinaryTaxes {
		end := "<none>"
		if t.Bracket.End != nil {
			end = fmt.Sprintf("%6d", *t.Bracket.End)
		}
		taxTable.AddRow([]string{
			"Ordinary",
			fmt.Sprintf("%6d - %s", t.Bracket.Start, end),
			fmt.Sprintf("%.0f", t.Bracket.RawBracket.Rate.ToDebugPercentage()),
			intToString(t.BracketTax.TaxableAmount),
			intToString(t.BracketTax.Tax),
		})
		ordinaryTotal += t.BracketTax.Tax
		if t.BracketTax.TaxableAmount > 0 {
			ordinaryMarginalRate = t.Bracket.RawBracket.Rate
		}
	}

	// LTCG tax
	ltcgTotal := int64(0)
	ltcgMarginalRate := Rate_0Percent
	for _, t := range e.LTCGTaxes {
		end := "<none>"
		if t.Bracket.End != nil {
			end = fmt.Sprintf("%6d", *t.Bracket.End)
		}
		taxTable.AddRow([]string{
			"LTCG",
			fmt.Sprintf("%6d - %s", t.Bracket.Start, end),
			fmt.Sprintf("%.0f", t.Bracket.RawBracket.Rate.ToDebugPercentage()),
			intToString(t.BracketTax.TaxableAmount),
			intToString(t.BracketTax.Tax),
		})
		ltcgTotal += t.BracketTax.Tax
		if t.BracketTax.TaxableAmount > 0 {
			ltcgMarginalRate = t.Bracket.RawBracket.Rate
		}
	}

	medicareTotal := e.Medicare.BaseWageTax + e.Medicare.AdditionalWageTax + e.Medicare.NiitTax
	grandTotal := medicareTotal + ssTotal + ordinaryTotal + ltcgTotal
	taxTotalTable := NewTable([]string{"Description", "Tax", "Marginal rate", "Effective rate"})
	taxTotalTable.AddRow([]string{"Medicare",
		intToString(medicareTotal),
		fmt.Sprintf("%.2f", e.Medicare.MarginalRate.ToDebugPercentage()),
		fmt.Sprintf("%.2f", 100*float32(medicareTotal)/float32(e.TotalIncome))})
	taxTotalTable.AddRow([]string{"Social security",
		intToString(ssTotal),
		"?",
		fmt.Sprintf("%.2f", 100*float32(ssTotal)/float32(e.TotalIncome))})
	taxTotalTable.AddRow([]string{"Ordinary",
		intToString(ordinaryTotal),
		fmt.Sprintf("%.2f", ordinaryMarginalRate.ToDebugPercentage()),
		fmt.Sprintf("%.2f", 100*float32(ordinaryTotal)/float32(e.TotalIncome))})
	taxTotalTable.AddRow([]string{"LTCG",
		intToString(ltcgTotal),
		fmt.Sprintf("%.2f", ltcgMarginalRate.ToDebugPercentage()),
		fmt.Sprintf("%.2f", 100*float32(ltcgTotal)/float32(e.TotalIncome))})
	taxTotalTable.AddRow([]string{"Grand total",
		intToString(grandTotal),
		"?",
		fmt.Sprintf("%.2f", 100*float32(grandTotal)/float32(e.TotalIncome))})

	fmt.Printf("tax breakdown:\n%s\n\n", taxTable.ToFormattedTable())
	fmt.Printf("totals:\n%s\n\n", taxTotalTable.ToFormattedTable())

	fmt.Printf("\n\n")
}
