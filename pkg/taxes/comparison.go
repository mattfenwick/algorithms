package taxes

import (
	"fmt"
	"strings"

	"github.com/mattfenwick/collections/pkg/builtin"
	"github.com/mattfenwick/collections/pkg/slice"
)

func PrettyPrintComparison(estimates []*TaxEstimate) {
	// income/input table
	inputHeader := []string{""}
	inputColumns := [][]string{
		{"Status", "Year", "Wages", "Short term", "Long term", "Gross", "Gross less nontaxable", "Medicare base", "Medicare additional", "Medicare NIIT", "Social security", "Adjustments", "AGI", "Deduction", "Taxable income"},
	}

	for _, e := range estimates {
		medicareBase, medicareAddnl, medicareNiit := e.Income.MedicareIncome()
		var socialSecurity []int64
		for _, s := range e.Income.IncomeSources {
			if s.IncomeType != IncomeTypeWage {
				continue
			}
			socialSecurity = append(socialSecurity, builtin.Min(e.TaxYearConstants.SocialSecurityLimit, s.Amount))
		}
		socialSecurityString := strings.Join(slice.Map(intToString, socialSecurity), "\n")
		wageString := strings.Join(
			slice.Map(
				func(i *IncomeSource) string { return intToString(i.Amount) },
				slice.Filter(func(i *IncomeSource) bool { return i.IncomeType == IncomeTypeWage }, e.Income.IncomeSources)),
			"\n")
		column := []string{
			e.Income.Status.ToString(),
			intToString(e.Income.Year),
			wageString,
			intToString(e.Income.ShortTermCapitalGainIncome()),
			intToString(e.Income.LongTermCapitalGainIncome()),
			intToString(e.Income.GetGrossIncome()),
			intToString(e.Income.GetGrossIncomeLessNonTaxable()),
			intToString(medicareBase),
			intToString(medicareAddnl),
			intToString(medicareNiit),
			socialSecurityString,
			intToString(e.Income.Adjustments),
			intToString(e.Income.GetAdjustedGrossIncome()),
			intToString(e.Income.GetDeduction()),
			intToString(e.Income.GetTaxableIncome()),
		}
		inputColumns = append(inputColumns, column)
		inputHeader = append(inputHeader, e.Income.Description)
	}
	inputTable := NewTable(inputHeader, Transpose(inputColumns)...)
	fmt.Printf("income/input: \n%s\n\n", inputTable.ToFormattedTable())

	// // taxes
	// taxHeader := []string{""}
	// taxColumns := []string{
	// 	"Medicare base wage",
	// 	"Medicare addnl wage",
	// 	"Medicare NIIT",
	// 	"Social security",
	// 	// TODO what if brackets change year to year?
	// }
	// taxTable := NewTable([]string{"Description", "Tax"})

	// // medicare tax
	// taxTable.AddRow([]string{"Medicare base wage",
	// 	fmt.Sprintf("%.2f", e.TaxYearConstants.MedicareBaseRate.ToDebugPercentage()),
	// 	intToString(e.Medicare.BaseWageIncome),
	// 	intToString(e.Medicare.BaseWageTax)})
	// taxTable.AddRow([]string{"Medicare addnl wage",
	// 	fmt.Sprintf("%.2f", e.TaxYearConstants.MedicareAdditionalRate.ToDebugPercentage()),
	// 	intToString(e.Medicare.AdditionalWageIncome),
	// 	intToString(e.Medicare.AdditionalWageTax)})
	// taxTable.AddRow([]string{"Medicare NIIT",
	// 	fmt.Sprintf("%.2f", e.TaxYearConstants.NetInvestmentTaxRate.ToDebugPercentage()),
	// 	intToString(e.Medicare.NiitIncome),
	// 	intToString(e.Medicare.NiitTax)})

	// // social security tax
	// ssTotal := int64(0)
	// for _, t := range e.SocialSecurity {
	// 	taxTable.AddRow([]string{
	// 		fmt.Sprintf("Social security %s", t.Description),
	// 		fmt.Sprintf("%.1f", e.TaxYearConstants.SocialSecurityRate.ToDebugPercentage()),
	// 		intToString(t.TaxableAmount),
	// 		intToString(t.Tax),
	// 	})
	// 	ssTotal += t.Tax
	// }

	// // ordinary tax
	// ordinaryTotal := int64(0)
	// ordinaryMarginalRate := Rate_0Percent
	// for _, t := range e.OrdinaryTaxes {
	// 	end := "<none>"
	// 	if t.Bracket.End != nil {
	// 		end = fmt.Sprintf("%6d", *t.Bracket.End)
	// 	}
	// 	taxTable.AddRow([]string{
	// 		fmt.Sprintf("Ordinary: %6d - %s", t.Bracket.Start, end),
	// 		fmt.Sprintf("%.0f", t.Bracket.RawBracket.Rate.ToDebugPercentage()),
	// 		intToString(t.BracketTax.TaxableAmount),
	// 		intToString(t.BracketTax.Tax),
	// 	})
	// 	ordinaryTotal += t.BracketTax.Tax
	// 	if t.BracketTax.TaxableAmount > 0 {
	// 		ordinaryMarginalRate = t.Bracket.RawBracket.Rate
	// 	}
	// }

	// // LTCG tax
	// ltcgTotal := int64(0)
	// ltcgMarginalRate := Rate_0Percent
	// for _, t := range e.LTCGTaxes {
	// 	end := "<none>"
	// 	if t.Bracket.End != nil {
	// 		end = fmt.Sprintf("%6d", *t.Bracket.End)
	// 	}
	// 	taxTable.AddRow([]string{
	// 		fmt.Sprintf("LTCG: %6d - %s", t.Bracket.Start, end),
	// 		fmt.Sprintf("%.0f", t.Bracket.RawBracket.Rate.ToDebugPercentage()),
	// 		intToString(t.BracketTax.TaxableAmount),
	// 		intToString(t.BracketTax.Tax),
	// 	})
	// 	ltcgTotal += t.BracketTax.Tax
	// 	if t.BracketTax.TaxableAmount > 0 {
	// 		ltcgMarginalRate = t.Bracket.RawBracket.Rate
	// 	}
	// }

	totalHeader := []string{""}
	totalColumns := [][]string{
		{"Medicare", "Social security", "Ordinary", "LTCG", "Total tax", "Effective rate"},
	}
	for _, e := range estimates {
		ssTotal := slice.Sum(slice.Map(func(s *SocialSecurityTax) int64 { return s.Tax }, e.SocialSecurity))
		ordinaryTotal := slice.Sum(slice.Map(func(t *TaxEstimateBracket) int64 { return t.BracketTax.Tax }, e.OrdinaryTaxes))
		ltcgTotal := slice.Sum(slice.Map(func(t *TaxEstimateBracket) int64 { return t.BracketTax.Tax }, e.LTCGTaxes))
		medicareTotal := e.Medicare.BaseWageTax + e.Medicare.AdditionalWageTax + e.Medicare.NiitTax
		grandTotal := medicareTotal + ssTotal + ordinaryTotal + ltcgTotal

		totalColumns = append(totalColumns, []string{
			intToString(medicareTotal),
			intToString(ssTotal),
			intToString(ordinaryTotal),
			intToString(ltcgTotal),
			intToString(grandTotal),
			fmt.Sprintf("%.2f", 100*float32(grandTotal)/float32(e.Income.GetGrossIncome())),
		})
		totalHeader = append(totalHeader, e.Income.Description)
	}
	totalTable := NewTable(totalHeader, Transpose(totalColumns)...)
	fmt.Printf("totals:\n%s\n\n", totalTable.ToFormattedTable())

	fmt.Printf("\n\n")
}
