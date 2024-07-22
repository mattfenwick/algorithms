package taxes

import (
	"fmt"

	"github.com/mattfenwick/collections/pkg/builtin"
	"github.com/mattfenwick/collections/pkg/slice"
	"github.com/pkg/errors"
	"golang.org/x/exp/constraints"
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

	ordinaryIncome := income.WageIncome() + income.ShortTermCapitalGainIncome()
	totalIncome := slice.Sum(slice.Map(func(i *IncomeSource) int64 { return i.Amount }, income.IncomeSources))

	ordinaryIncomeAfterDeducion := ordinaryIncome - income.GetDeduction()
	var ordinaryTaxes []*TaxEstimateBracket
	for _, b := range statusConstants.OrdinaryIncomeBrackets.GetBrackets() {
		ordinaryTaxes = append(ordinaryTaxes, &TaxEstimateBracket{b.GetTax(ordinaryIncomeAfterDeducion), b})
	}

	// LTCG
	//   TODO question -- how does LTCG interact with deduction?
	totalIncomeAfterDeduction := totalIncome - income.GetDeduction()
	var ltcgTaxes []*TaxEstimateBracket
	for _, b := range statusConstants.LTCGIncomeBrackets.GetBrackets() {
		ltcgTaxes = append(ltcgTaxes,
			&TaxEstimateBracket{
				b.GetLongTermCapitalGainsTax(totalIncomeAfterDeduction, ordinaryIncomeAfterDeducion),
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
	e.Income.GetAdjustedGrossIncome()
	e.Income.GetTaxableIncome()
	e.Income.GetTotalIncome()
	inputTable := NewTable([]string{"Key", "Value"},
		[]string{"Status", e.Income.Status.ToString()},
		[]string{"Year", fmt.Sprintf("%d", e.Income.Year)},
		[]string{"Wages", fmt.Sprintf("%d", e.Income.WageIncome())},
		[]string{"Short term", fmt.Sprintf("%d", e.Income.ShortTermCapitalGainIncome())},
		[]string{"Long term", fmt.Sprintf("%d", e.Income.LongTermCapitalGainIncome())},
		[]string{"Total income", intToString(e.Income.GetTotalIncome())},
		[]string{"Adjustments", intToString(e.Income.Adjustments)},
		[]string{"AGI", intToString(e.Income.GetAdjustedGrossIncome())},
		[]string{"Deduction", fmt.Sprintf("%d", e.Income.GetDeduction())},
		[]string{"Taxable income", intToString(e.Income.GetTaxableIncome())},
		[]string{"Medicare base", intToString(medicareBase)},
		[]string{"Medicare additional", intToString(medicareAddnl)},
		[]string{"Medicare NIIT", intToString(medicareNiit)},
	)
	for _, s := range e.Income.IncomeSources {
		if s.IncomeType != IncomeTypeWage {
			continue
		}
		inputTable.AddRow([]string{fmt.Sprintf("Social security %s", s.Description), intToString(builtin.Min(e.TaxYearConstants.SocialSecurityLimit, s.Amount))})
	}
	fmt.Printf("income/input: \n%s\n\n", inputTable.ToFormattedTable())

	// taxes
	taxTable := NewTable([]string{"Description", "Rate (%)", "Taxable amount", "Tax"})

	// medicare tax
	taxTable.AddRow([]string{"Medicare base wage",
		fmt.Sprintf("%.2f", e.TaxYearConstants.MedicareBaseRate.ToDebugPercentage()),
		intToString(e.Medicare.BaseWageIncome),
		intToString(e.Medicare.BaseWageTax)})
	taxTable.AddRow([]string{"Medicare addnl wage",
		fmt.Sprintf("%.2f", e.TaxYearConstants.MedicareAdditionalRate.ToDebugPercentage()),
		intToString(e.Medicare.AdditionalWageIncome),
		intToString(e.Medicare.AdditionalWageTax)})
	taxTable.AddRow([]string{"Medicare NIIT",
		fmt.Sprintf("%.2f", e.TaxYearConstants.NetInvestmentTaxRate.ToDebugPercentage()),
		intToString(e.Medicare.NiitIncome),
		intToString(e.Medicare.NiitTax)})

	// social security tax
	ssTotal := int64(0)
	for _, t := range e.SocialSecurity {
		taxTable.AddRow([]string{
			fmt.Sprintf("Social security %s", t.Description),
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
			fmt.Sprintf("Ordinary: %6d - %s", t.Bracket.Start, end),
			fmt.Sprintf("%.0f", t.Bracket.RawBracket.Rate.ToDebugPercentage()),
			fmt.Sprintf("%d", t.BracketTax.TaxableAmount),
			fmt.Sprintf("%d", t.BracketTax.Tax),
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
			fmt.Sprintf("LTCG: %6d - %s", t.Bracket.Start, end),
			fmt.Sprintf("%.0f", t.Bracket.RawBracket.Rate.ToDebugPercentage()),
			fmt.Sprintf("%d", t.BracketTax.TaxableAmount),
			fmt.Sprintf("%d", t.BracketTax.Tax),
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

func intToString[A constraints.Integer](a A) string {
	return fmt.Sprintf("%d", a)
}
