package taxes

import (
	"fmt"

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

func RunTaxes(incomes []*Income) {
	fmt.Println("ordinary income:")
	HackPrintOrdinaryIncomeBrackets()
	fmt.Printf("\nlong term capital gains:\n")
	HackPrintLTCGIncomeBrackets()

	var estimates []*TaxEstimate
	for _, inc := range incomes {
		estimate := EstimateTaxes(inc)
		// fmt.Printf("estimate: \n%s\n\n", json.MustMarshalToString(estimate))
		fmt.Println("estimate:")
		estimate.PrettyPrint()
		estimates = append(estimates, estimate)
	}

	PrettyPrintComparison(estimates)
}

func HackPrintOrdinaryIncomeBrackets() {
	statuses := []FilingStatus{FilingStatusSingle, FilingStatusMarriedJointly, FilingStatusMarriedSeparately, FilingStatusHeadOfHouseHold}
	rows := [][]string{}
	for _, status := range statuses {
		row := []string{}
		for _, b := range TaxYear2024.ByStatus[status].OrdinaryIncomeBrackets.GetBrackets() {
			end := b.End
			if end == nil {
				row = append(row, fmt.Sprintf("%7d -  <none>", b.Start))
			} else {
				row = append(row, fmt.Sprintf("%7d - %7d", b.Start, *end))
			}
		}
		rows = append(rows, row)
	}
	table := &Table{
		Headers: slice.Map(func(f FilingStatus) string { return f.ToString() }, statuses),
		Rows:    Transpose(rows),
	}
	fmt.Printf("%s\n", table.ToFormattedTable())
}

func HackPrintLTCGIncomeBrackets() {
	statuses := []FilingStatus{FilingStatusSingle, FilingStatusMarriedJointly, FilingStatusMarriedSeparately, FilingStatusHeadOfHouseHold}
	rows := [][]string{}
	for _, status := range statuses {
		row := []string{}
		for _, b := range TaxYear2024.ByStatus[status].LTCGIncomeBrackets.GetBrackets() {
			end := b.End
			if end == nil {
				row = append(row, fmt.Sprintf("%7d -  <none>", b.Start))
			} else {
				row = append(row, fmt.Sprintf("%7d - %7d", b.Start, *end))
			}
		}
		rows = append(rows, row)
	}
	table := &Table{
		Headers: slice.Map(func(f FilingStatus) string { return f.ToString() }, statuses),
		Rows:    Transpose(rows),
	}
	fmt.Printf("%s\n", table.ToFormattedTable())
}
