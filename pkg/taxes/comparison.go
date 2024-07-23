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
		{
			"Status", "Year", "Wages", "Nontaxable wages",
			"Short term", "Long term", "Gross", "Gross less nontaxable",
			"Adjustments", "AGI", "Deduction", "Taxable income",
			"Nontaxable payroll", "Medicare base", "Medicare additional", "Medicare NIIT", "Social security",
		},
	}

	for _, e := range estimates {
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
		socialSecurityString := strings.Join(slice.Map(intToString, socialSecurity), "\n")
		column := []string{
			e.Income.Status.ToString(),
			intToString(e.Income.Year),
			strings.Join(slice.Map(intToString, wages), "\n"),
			strings.Join(slice.Map(intToString, nonTaxableWages), "\n"),
			intToString(e.Income.ShortTermCapitalGainIncome()),
			intToString(e.Income.LongTermCapitalGainIncome()),
			intToString(e.Income.GetGrossIncome()),
			intToString(e.Income.GetGrossIncomeLessNonTaxable()),
			intToString(e.Income.Adjustments),
			intToString(e.Income.GetAdjustedGrossIncome()),
			intToString(e.Income.GetDeduction()),
			intToString(e.Income.GetTaxableIncome()),
			strings.Join(slice.Map(intToString, nonTaxablePayroll), "\n"),
			intToString(medicareBase),
			intToString(medicareAddnl),
			intToString(medicareNiit),
			socialSecurityString,
		}
		inputColumns = append(inputColumns, column)
		inputHeader = append(inputHeader, e.Income.Description)
	}
	inputTable := NewTable(inputHeader, Transpose(inputColumns)...)
	fmt.Printf("income/input: \n%s\n\n", inputTable.ToFormattedTable())

	breakdownHeader := []string{"", ""}
	breakdownColumns := [][]string{
		{
			"Medicare", "Medicare", "Medicare",
			"Social security",
			"Ordinary", "Ordinary", "Ordinary", "Ordinary", "Ordinary", "Ordinary", "Ordinary",
			"LTCG", "LTCG", "LTCG",
		},
		{
			"Base wage", "Addnl wage", "NIIT",
			"",
			"10%", "12%", "22%", "24%", "32%", "35%", "37%",
			"0%", "15%", "20%",
		},
	}
	for _, e := range estimates {
		// medicare tax
		row := []string{
			fmt.Sprintf("%d - %d", e.Medicare.BaseWageIncome, e.Medicare.BaseWageTax),
			fmt.Sprintf("%d - %d", e.Medicare.AdditionalWageIncome, e.Medicare.AdditionalWageTax),
			fmt.Sprintf("%d - %d", e.Medicare.NiitIncome, e.Medicare.NiitTax),
			strings.Join(slice.Map(func(s *SocialSecurityTax) string {
				return fmt.Sprintf("%d - %d", s.TaxableAmount, s.Tax)
			}, e.SocialSecurity), "\n"),
		}

		// ordinary tax
		for _, t := range e.OrdinaryTaxes {
			row = append(row, fmt.Sprintf("%d - %d", t.BracketTax.TaxableAmount, t.BracketTax.Tax))
		}

		// LTCG tax
		for _, t := range e.LTCGTaxes {
			row = append(row, fmt.Sprintf("%d - %d", t.BracketTax.TaxableAmount, t.BracketTax.Tax))
		}
		breakdownColumns = append(breakdownColumns, row)
		breakdownHeader = append(breakdownHeader, e.Income.Description)
	}
	breakdownTable := NewTable(breakdownHeader, Transpose(breakdownColumns)...)
	fmt.Printf("bracket breakdown: \n%s\n\n", breakdownTable.ToFormattedTable())

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
