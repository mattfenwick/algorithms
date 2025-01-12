package taxes

import (
	"fmt"

	"github.com/mattfenwick/algorithms/pkg/utils"
	"github.com/mattfenwick/collections/pkg/slice"
)

var (
	defaultIncomes = []*Income{
		{
			Year:   2024,
			Status: FilingStatusSingle,
			IncomeSources: []*IncomeSource{
				{Description: "job1", IncomeType: IncomeTypeWage, Amount: 25_000},
				{Description: "inv1", IncomeType: IncomeTypeLongTerm, Amount: 250_000},
			},
			Deduction: nil,
		},
		{
			Year:   2024,
			Status: FilingStatusSingle,
			IncomeSources: []*IncomeSource{
				{Description: "job1", IncomeType: IncomeTypeWage, Amount: 250_000},
				{Description: "inv1", IncomeType: IncomeTypeLongTerm, Amount: 25_000},
			},
			Deduction: nil,
		},
		{
			Year:   2024,
			Status: FilingStatusSingle,
			IncomeSources: []*IncomeSource{
				{Description: "job1", IncomeType: IncomeTypeWage, Amount: 30_000},
			},
			Deduction: nil,
		},
		{
			Year:   2024,
			Status: FilingStatusSingle,
			IncomeSources: []*IncomeSource{
				{Description: "inv1", IncomeType: IncomeTypeLongTerm, Amount: 30_000},
			},
			Deduction: nil,
		},
		{
			Year:   2024,
			Status: FilingStatusSingle,
			IncomeSources: []*IncomeSource{
				{Description: "job1", IncomeType: IncomeTypeWage, Amount: 25_000},
				{Description: "inv1", IncomeType: IncomeTypeLongTerm, Amount: 25_000},
			},
			Deduction: nil,
		},
		{
			Year:   2024,
			Status: FilingStatusSingle,
			IncomeSources: []*IncomeSource{
				{Description: "job1", IncomeType: IncomeTypeWage, Amount: 5_000},
				{Description: "inv1", IncomeType: IncomeTypeLongTerm, Amount: 25_000},
			},
			Deduction: nil,
		},
		{
			Year:   2024,
			Status: FilingStatusSingle,
			IncomeSources: []*IncomeSource{
				{Description: "job1", IncomeType: IncomeTypeWage, Amount: 65_000},
				{Description: "inv1", IncomeType: IncomeTypeLongTerm, Amount: 35_000},
			},
			Deduction: nil,
		},
	}
)

func RunBrackets(taxYear []int) {
	for _, year := range taxYear {
		taxYear, ok := TaxYears[year]
		if !ok {
			fmt.Printf("tax year constants not found for <%d>\n", year)
			continue
		}
		fmt.Printf("tax year %d ordinary income:\n", year)
		PrintOrdinaryIncomeBrackets(taxYear)
		fmt.Printf("\nlong term %d capital gains:\n", year)
		PrintLTCGIncomeBrackets(taxYear)
		fmt.Printf("\n\n")

		// // TODO med
		// taxYear.MedicareAdditionalRate
		// taxYear.MedicareBaseRate
		// taxYear.NetInvestmentTaxRate
		// // TODO social
		// taxYear.SocialSecurityLimit
		// taxYear.SocialSecurityRate
		// // TODO deduction
		// taxYear.ByStatus[FilingStatusMarriedJointly].StandardDeduction
	}
}

func RunTaxes(incomes []*Income) {
	var estimates []*TaxEstimate
	for _, inc := range incomes {
		estimate := EstimateTaxes(inc)
		// fmt.Printf("estimate: \n%s\n\n", json.MustMarshalToString(estimate))
		fmt.Printf("estimate for %s:\n", inc.Description)
		estimate.PrettyPrint()
		estimates = append(estimates, estimate)
	}

	fmt.Printf("comparison:\n")
	PrettyPrintComparison(estimates)
}

func PrintOrdinaryIncomeBrackets(taxYear *TaxYearConstants) {
	statuses := []FilingStatus{FilingStatusSingle, FilingStatusMarriedJointly, FilingStatusMarriedSeparately, FilingStatusHeadOfHouseHold}
	rows := [][]string{}
	var rateRow []string
	for _, b := range taxYear.ByStatus[FilingStatusMarriedJointly].OrdinaryIncomeBrackets.GetBrackets() {
		rateRow = append(rateRow, fmt.Sprintf("%.0f", b.RawBracket.Rate.ToDebugPercentage()))
	}
	rows = append(rows, rateRow)
	for _, status := range statuses {
		row := []string{}
		for _, b := range taxYear.ByStatus[status].OrdinaryIncomeBrackets.GetBrackets() {
			end := b.End
			if end == nil {
				row = append(row, fmt.Sprintf("%7d -  <none>", b.Start))
			} else {
				row = append(row, fmt.Sprintf("%7d - %7d", b.Start, *end))
			}
		}
		rows = append(rows, row)
	}
	table := &utils.Table{
		Headers: slice.Cons("Rate %", slice.Map(func(f FilingStatus) string { return f.ToString() }, statuses)),
		Rows:    Transpose(rows),
	}
	fmt.Printf("%s\n", table.ToFormattedTable())
}

func PrintLTCGIncomeBrackets(taxYear *TaxYearConstants) {
	statuses := []FilingStatus{FilingStatusSingle, FilingStatusMarriedJointly, FilingStatusMarriedSeparately, FilingStatusHeadOfHouseHold}
	rows := [][]string{}
	var rateRow []string
	for _, b := range taxYear.ByStatus[FilingStatusMarriedJointly].LTCGIncomeBrackets.GetBrackets() {
		rateRow = append(rateRow, fmt.Sprintf("%.0f", b.RawBracket.Rate.ToDebugPercentage()))
	}
	rows = append(rows, rateRow)
	for _, status := range statuses {
		row := []string{}
		for _, b := range taxYear.ByStatus[status].LTCGIncomeBrackets.GetBrackets() {
			end := b.End
			if end == nil {
				row = append(row, fmt.Sprintf("%7d -  <none>", b.Start))
			} else {
				row = append(row, fmt.Sprintf("%7d - %7d", b.Start, *end))
			}
		}
		rows = append(rows, row)
	}
	table := &utils.Table{
		Headers: slice.Cons("Rate %", slice.Map(func(f FilingStatus) string { return f.ToString() }, statuses)),
		Rows:    Transpose(rows),
	}
	fmt.Printf("%s\n", table.ToFormattedTable())
}
