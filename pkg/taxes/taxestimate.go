package taxes

import (
	"fmt"

	"github.com/mattfenwick/collections/pkg/slice"
	"github.com/pkg/errors"
	"golang.org/x/exp/constraints"
)

type TaxEstimateBracket struct {
	BracketTax *BracketTax
	Bracket    *Bracket
}

type TaxEstimate struct {
	Income           *Income
	TaxYearConstants *TaxYearConstants
	StatusConstants  *TaxStatusConstants
	OrdinaryTaxes    []*TaxEstimateBracket
	LTCGTaxes        []*TaxEstimateBracket
	Medicare         *MedicareTax
	SocialSecurity   []*SocialSecurityTax
	TotalIncome      int64
	// IncomeAfterDeduction int
}

func EstimateTaxes(income *Income) *TaxEstimate {
	yearConstants, ok := TaxYears[income.Year]
	if !ok {
		panic(errors.Errorf("no tax constant info for year %d", income.Year))
	}
	statusConstants := yearConstants.ByStatus[income.Status]

	ordinaryIncome := income.WageIncome() + income.ShortTermCapitalGainIncome()
	totalIncome := slice.Sum(slice.Map(func(i *IncomeSource) int64 { return i.Amount }, income.IncomeSources))
	// capitalGains := slice.Sum(
	// 	slice.Map(func(i *IncomeSource) int { return i.Amount },
	// 		slice.Filter(func(i *IncomeSource) bool {
	// 			switch i.IncomeType {
	// 			case IncomeTypeWage:
	// 				return false
	// 			default:
	// 				return true
	// 			}
	// 		}, income.IncomeSources)))

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
	inputTable := NewTable([]string{"Key", "Value"},
		[]string{"Status", e.Income.Status.ToString()},
		[]string{"Year", fmt.Sprintf("%d", e.Income.Year)},
		[]string{"Deduction", fmt.Sprintf("%d", e.Income.GetDeduction())},
		[]string{"Wages", fmt.Sprintf("%d", e.Income.WageIncome())},
		[]string{"Short term", fmt.Sprintf("%d", e.Income.ShortTermCapitalGainIncome())},
		[]string{"Long term", fmt.Sprintf("%d", e.Income.LongTermCapitalGainIncome())},
	)
	fmt.Printf("income/input: \n%s\n\n", inputTable.ToFormattedTable())

	// ordinary tax
	ordinaryTaxTable := NewTable([]string{"Range", "Rate", "Taxable amount", "Tax"})
	for _, t := range e.OrdinaryTaxes {
		end := "<none>"
		if t.Bracket.End != nil {
			end = fmt.Sprintf("%6d", *t.Bracket.End)
		}
		ordinaryTaxTable.AddRow([]string{
			fmt.Sprintf("%6d - %s", t.Bracket.Start, end),
			fmt.Sprintf("%.0f%%", t.Bracket.RawBracket.Rate.ToDebugPercentage()),
			fmt.Sprintf("%d", t.BracketTax.TaxableAmount),
			fmt.Sprintf("%d", t.BracketTax.Tax),
		})
	}
	fmt.Printf("ordinary tax:\n%s\n\n", ordinaryTaxTable.ToFormattedTable())

	// medicare
	medicareTable := NewTable([]string{"", "Taxable amount", "Rate", "Tax"},
		[]string{"Base wage",
			intToString(e.Medicare.BaseWageIncome),
			fmt.Sprintf("%.2f", e.TaxYearConstants.MedicareBaseRate.ToDebugPercentage()),
			intToString(e.Medicare.BaseWageTax)},
		[]string{"Addnl wage",
			intToString(e.Medicare.AdditionalWageIncome),
			fmt.Sprintf("%.2f + %.2f",
				e.TaxYearConstants.MedicareBaseRate.ToDebugPercentage(),
				e.TaxYearConstants.MedicareAdditionalRate.ToDebugPercentage()),
			intToString(e.Medicare.AdditionalWageTax)},
		[]string{"NIIT",
			intToString(e.Medicare.NiitIncome),
			fmt.Sprintf("%.2f", e.TaxYearConstants.NetInvestmentTaxRate.ToDebugPercentage()),
			intToString(e.Medicare.NiitTax)},
	)
	// fmt.Printf(" - threshold: %d\n", -1) // TODO ???
	fmt.Printf("medicare: \n%s\n\n", medicareTable.ToFormattedTable())

	// social security
	socialSecurityTable := NewTable([]string{"Job", "Income", "Taxable amount", "Tax"})
	for _, t := range e.SocialSecurity {
		socialSecurityTable.AddRow([]string{
			t.Description,
			fmt.Sprintf("%d", t.WageIncome),
			fmt.Sprintf("%d", t.TaxableAmount),
			fmt.Sprintf("%d", t.Tax),
		})
	}
	fmt.Printf("social security tax:\n%s\n", socialSecurityTable.ToFormattedTable())

	// LTCG
	ltcgTaxTable := NewTable([]string{"Range", "Rate", "Taxable amount", "Tax"})
	for _, t := range e.LTCGTaxes {
		end := "<none>"
		if t.Bracket.End != nil {
			end = fmt.Sprintf("%6d", *t.Bracket.End)
		}
		ltcgTaxTable.AddRow([]string{
			fmt.Sprintf("%6d - %s", t.Bracket.Start, end),
			fmt.Sprintf("%.0f%%", t.Bracket.RawBracket.Rate.ToDebugPercentage()),
			fmt.Sprintf("%d", t.BracketTax.TaxableAmount),
			fmt.Sprintf("%d", t.BracketTax.Tax),
		})
	}
	fmt.Printf("ltcg tax:\n%s\n\n", ltcgTaxTable.ToFormattedTable())

	// totals
	totalsTable := NewTable([]string{"Description", "Tax"})
	for _, t := range e.OrdinaryTaxes {
		totalsTable.AddRow([]string{
			fmt.Sprintf("Wage+STCG: %.0f%%", t.Bracket.RawBracket.Rate.ToDebugPercentage()),
			intToString(t.BracketTax.Tax),
		})
	}
	totalsTable.AddRow([]string{"Wage+STCG total",
		intToString(slice.Sum(slice.Map(func(t *TaxEstimateBracket) int64 { return t.BracketTax.Tax }, e.OrdinaryTaxes)))})
	for _, t := range e.LTCGTaxes {
		totalsTable.AddRow([]string{
			fmt.Sprintf("LTCG: %.0f%%", t.Bracket.RawBracket.Rate.ToDebugPercentage()),
			intToString(t.BracketTax.Tax),
		})
	}
	totalsTable.AddRow([]string{"LTCG total",
		intToString(slice.Sum(slice.Map(func(t *TaxEstimateBracket) int64 { return t.BracketTax.Tax }, e.LTCGTaxes)))})
	totalsTable.AddRow([]string{"Medicare base wage", intToString(e.Medicare.BaseWageTax)})
	totalsTable.AddRow([]string{"Medicare addnl wage", intToString(e.Medicare.AdditionalWageTax)})
	totalsTable.AddRow([]string{"Medicare NIIT", intToString(e.Medicare.NiitTax)})
	totalsTable.AddRow([]string{"Medicare total",
		intToString(e.Medicare.BaseWageTax + e.Medicare.AdditionalWageTax + e.Medicare.NiitTax)})
	for _, s := range e.SocialSecurity {
		totalsTable.AddRow([]string{"Social security " + s.Description, intToString(s.Tax)})
	}
	totalsTable.AddRow([]string{"Social security total",
		intToString(slice.Sum(slice.Map(func(s *SocialSecurityTax) int64 { return s.Tax }, e.SocialSecurity)))})

	fmt.Printf("totals:\n%s\n\n", totalsTable.ToFormattedTable())

	fmt.Printf("\n\n")
}

func intToString[A constraints.Integer](a A) string {
	return fmt.Sprintf("%d", a)
}
