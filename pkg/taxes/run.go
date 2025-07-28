package taxes

import (
	"fmt"

	"github.com/mattfenwick/algorithms/pkg/utils"
	"github.com/mattfenwick/collections/pkg/slice"
)

var (
	defaultIncomes = []*Income{
		{
			Year:   2025,
			Status: FilingStatusSingle,
			IncomeSources: []*IncomeSource{
				{Description: "job1", IncomeType: IncomeTypeWage, Amount: 25_000},
				{Description: "inv1", IncomeType: IncomeTypeLongTerm, Amount: 250_000},
			},
			Deduction: nil,
		},
		{
			Year:   2025,
			Status: FilingStatusSingle,
			IncomeSources: []*IncomeSource{
				{Description: "job1", IncomeType: IncomeTypeWage, Amount: 250_000},
				{Description: "inv1", IncomeType: IncomeTypeLongTerm, Amount: 25_000},
			},
			Deduction: nil,
		},
		{
			Year:   2025,
			Status: FilingStatusSingle,
			IncomeSources: []*IncomeSource{
				{Description: "job1", IncomeType: IncomeTypeWage, Amount: 30_000},
			},
			Deduction: nil,
		},
		{
			Year:   2025,
			Status: FilingStatusSingle,
			IncomeSources: []*IncomeSource{
				{Description: "inv1", IncomeType: IncomeTypeLongTerm, Amount: 30_000},
			},
			Deduction: nil,
		},
		{
			Year:   2025,
			Status: FilingStatusSingle,
			IncomeSources: []*IncomeSource{
				{Description: "job1", IncomeType: IncomeTypeWage, Amount: 25_000},
				{Description: "inv1", IncomeType: IncomeTypeLongTerm, Amount: 25_000},
			},
			Deduction: nil,
		},
		{
			Year:   2025,
			Status: FilingStatusSingle,
			IncomeSources: []*IncomeSource{
				{Description: "job1", IncomeType: IncomeTypeWage, Amount: 5_000},
				{Description: "inv1", IncomeType: IncomeTypeLongTerm, Amount: 25_000},
			},
			Deduction: nil,
		},
		{
			Year:   2025,
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
		fmt.Printf("\n%d medicare and standard deduction:\n", year)
		PrintMedicareThresholdAndStandardDeductions(taxYear)
		fmt.Printf("\n%d medicare and social security rates and limits:\n", year)
		PrintMedicareAndSocialSecurityInfo(taxYear)
		fmt.Printf("\n\n")
	}
}

func RunTaxes(incomes []*Income, showComparison bool) {
	var estimates []*TaxEstimate
	for _, inc := range incomes {
		estimate := EstimateTaxes(inc)
		// fmt.Printf("estimate: \n%s\n\n", json.MustMarshalToString(estimate))
		fmt.Printf("estimate for %s:\n", inc.Description)
		estimate.PrettyPrint()
		estimates = append(estimates, estimate)
	}

	if showComparison && len(estimates) > 1 {
		fmt.Printf("comparison:\n")
		PrettyPrintComparison(estimates)
	}
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

func PrintMedicareThresholdAndStandardDeductions(taxYear *TaxYearConstants) {
	statuses := []FilingStatus{FilingStatusSingle, FilingStatusMarriedJointly, FilingStatusMarriedSeparately, FilingStatusHeadOfHouseHold}
	medicareRow := []string{"Medicare additional threshold"}
	standardDeductionRow := []string{"Standard deduction"}
	for _, s := range statuses {
		medicareRow = append(medicareRow, fmt.Sprintf("%d", taxYear.ByStatus[s].MedicareAdditionalThreshold))
		standardDeductionRow = append(standardDeductionRow, fmt.Sprintf("%d", taxYear.ByStatus[s].StandardDeduction))
	}
	table := &utils.Table{
		Headers: slice.Cons("", slice.Map(func(f FilingStatus) string { return f.ToString() }, statuses)),
		Rows:    [][]string{medicareRow, standardDeductionRow},
	}
	fmt.Printf("%s\n", table.ToFormattedTable())
}

func PrintMedicareAndSocialSecurityInfo(taxYear *TaxYearConstants) {
	table := &utils.Table{
		Headers: []string{"", ""},
		Rows: [][]string{
			{"Medicare base rate", fmt.Sprintf("%.2f %%", taxYear.MedicareBaseRate.ToDebugPercentage())},
			{"Medicare additional rate", fmt.Sprintf("%.2f %%", taxYear.MedicareAdditionalRate.ToDebugPercentage())},
			{"Medicare NIIT rate", fmt.Sprintf("%.1f %%", taxYear.NetInvestmentTaxRate.ToDebugPercentage())},
			{"Social security rate", fmt.Sprintf("%.1f %%", taxYear.SocialSecurityRate.ToDebugPercentage())},
			{"Social security limit", intToString(taxYear.SocialSecurityLimit)},
		},
	}
	fmt.Printf("%s\n", table.ToFormattedTable())
}
