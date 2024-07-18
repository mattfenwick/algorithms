package main

import (
	"fmt"
	"os"

	"github.com/mattfenwick/algorithms/pkg"
	"github.com/mattfenwick/algorithms/pkg/taxes"
)

func main() {
	mode := "logs"
	fmt.Printf("args: %+v\n", os.Args)
	if len(os.Args) > 1 {
		mode = os.Args[1]
	}
	switch mode {
	case "logs":
		pkg.LogsMain()
	case "binary-search":
		pkg.BinarySearchMain()
	case "ugh":
		pkg.UghMain()
	case "kth-factor":
		pkg.KthFactorMain()
	case "count-complete-tree-nodes":
		pkg.CountNodes()
	case "count-negatives":
		pkg.CountNegativesMain()
	case "find-target-indices":
		pkg.FindTargetIndicesMain()
	case "k-weakest-rows":
		pkg.KWeakestRowsMain()
	case "maximum-count":
		pkg.MaximumCountMain()
	case "find-distance-value":
		pkg.FindDistanceValueMain()
	case "ship-within-days":
		pkg.ShipWithinDaysMain()
	case "max-consecutive-answers":
		pkg.MaxConsecutiveAnswersMain()
	case "search-suggestions":
		pkg.SuggestedProductsMain()
	case "taxes":
		taxes.RunTaxes()
	default:
		panic("missed case")
	}
}
