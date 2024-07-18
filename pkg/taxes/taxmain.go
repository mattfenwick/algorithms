package taxes

import (
	"fmt"
	"math"

	"github.com/mattfenwick/collections/pkg/slice"
)

func RunTaxes() {
	fmt.Println("ordinary income:")
	HackPrintOrdinaryIncomeBrackets()
	fmt.Printf("\nlong term capital gains:\n")
	HackPrintLTCGIncomeBrackets()
}

func HackPrintOrdinaryIncomeBrackets() {
	statuses := []FilingStatus{FilingStatusSingle, FilingStatusMarriedJointly, FilingStatusMarriedSeparately, FilingStatusHeadOfHouseHold}
	rows := [][]string{}
	for _, status := range statuses {
		start := 0
		row := []string{}
		for _, b := range TaxYear2024.ByStatus[status].OrdinaryIncomeBrackets {
			end := b.Max
			if end == math.MaxInt {
				row = append(row, fmt.Sprintf("%7d -  <none>", start))
			} else {
				row = append(row, fmt.Sprintf("%7d - %7d", start, end))
			}
			start = end + 1
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
		start := 0
		row := []string{}
		for _, b := range TaxYear2024.ByStatus[status].LTCGIncomeBrackets {
			end := b.Max
			if end == math.MaxInt {
				row = append(row, fmt.Sprintf("%7d -  <none>", start))
			} else {
				row = append(row, fmt.Sprintf("%7d - %7d", start, end))
			}
			start = end + 1
		}
		rows = append(rows, row)
	}
	table := &Table{
		Headers: slice.Map(func(f FilingStatus) string { return f.ToString() }, statuses),
		Rows:    Transpose(rows),
	}
	fmt.Printf("%s\n", table.ToFormattedTable())
}
