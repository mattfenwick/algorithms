package taxes

import (
	"fmt"

	"github.com/mattfenwick/collections/pkg/json"
	"github.com/mattfenwick/collections/pkg/slice"
)

func RunTaxes() {
	fmt.Println("ordinary income:")
	HackPrintOrdinaryIncomeBrackets()
	fmt.Printf("\nlong term capital gains:\n")
	HackPrintLTCGIncomeBrackets()

	estimate := EstimateTaxes(&Income{
		Year:   2024,
		Status: FilingStatusSingle,
		IncomeSources: []*IncomeSource{
			{Description: "job1", IncomeType: IncomeTypeW2, Amount: 50000},
		},
		Deduction: TaxYear2024.ByStatus[FilingStatusSingle].StandardDeduction,
	})
	fmt.Printf("estimate: \n%s\n\n", json.MustMarshalToString(estimate))
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
