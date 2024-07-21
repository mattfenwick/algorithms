package taxes

import (
	"fmt"

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
	Medicare             *MedicareTax
	SocialSecurity       []*SocialSecurityTax
	TotalIncome          int
	IncomeAfterDeduction int
}

func EstimateTaxes(income *Income) *TaxEstimate {
	yearConstants, ok := TaxYears[income.Year]
	if !ok {
		panic(errors.Errorf("no tax constant info for year %d", income.Year))
	}
	statusConstants := yearConstants.ByStatus[income.Status]

	totalIncome := slice.Sum(slice.Map(func(i *IncomeSource) int { return i.Amount }, income.IncomeSources))
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

	incomeAfterDeduction := totalIncome - income.Deduction

	var bracketTaxes []*TaxEstimateBracket
	for _, b := range statusConstants.OrdinaryIncomeBrackets.GetBrackets() {
		bracketTaxes = append(bracketTaxes, &TaxEstimateBracket{b.GetTax(incomeAfterDeduction), b})
	}

	// TODO LTCG

	return &TaxEstimate{
		Income:               income,
		TaxYearConstants:     yearConstants,
		StatusConstants:      statusConstants,
		OrdinaryTaxes:        bracketTaxes,
		Medicare:             EstimateMedicareTax(income),
		SocialSecurity:       EstimateSocialSecurityTax(income),
		TotalIncome:          totalIncome,
		IncomeAfterDeduction: incomeAfterDeduction,
	}
}

func (e *TaxEstimate) PrettyPrint() {
	// ordinary tax
	ordinaryTaxTable := NewTable([]string{"Range", "Rate", "Taxable amount", "Tax"})
	for _, t := range e.OrdinaryTaxes {
		end := "<none>"
		if t.Bracket.End != nil {
			end = fmt.Sprintf("%6d", *t.Bracket.End)
		}
		ordinaryTaxTable.AddRow([]string{
			fmt.Sprintf("%6d - %s", t.Bracket.Start, end),
			fmt.Sprintf("%d", t.Bracket.RawBracket.Rate),
			fmt.Sprintf("%d", t.BracketTax.TaxableAmount),
			fmt.Sprintf("%d", t.BracketTax.Tax),
		})
	}
	fmt.Printf("ordinary tax:\n%s\n\n", ordinaryTaxTable.ToFormattedTable())

	// medicare
	fmt.Println("medicare:")
	fmt.Printf(" - threshold: %d\n", -1) // TODO ???
	fmt.Printf(" - base income: %d\n", e.Medicare.BaseWageIncome)
	fmt.Printf(" - base tax: %d\n", e.Medicare.BaseWageTax)
	fmt.Printf(" - NIIT income: %d\n", e.Medicare.NiitIncome)
	fmt.Printf(" - NIIT tax: %d\n", e.Medicare.NiitTax)
	fmt.Printf(" - additional income: %d\n", e.Medicare.AdditionalWageIncome)
	fmt.Printf(" - additional tax: %d\n", e.Medicare.AdditionalWageTax)
	fmt.Printf("\n\n")

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

	// TODO LTCG

	// TODO income calcs

	fmt.Printf("\n\n")
}
