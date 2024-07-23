package main

import (
	"fmt"
	"os"

	"github.com/mattfenwick/algorithms/pkg/algs"
)

func main() {
	mode := "logs"
	fmt.Printf("args: %+v\n", os.Args)
	if len(os.Args) > 1 {
		mode = os.Args[1]
	}
	switch mode {
	case "logs":
		algs.LogsMain()
	case "binary-search":
		algs.BinarySearchMain()
	case "ugh":
		algs.UghMain()
	case "kth-factor":
		algs.KthFactorMain()
	case "count-complete-tree-nodes":
		algs.CountNodes()
	case "count-negatives":
		algs.CountNegativesMain()
	case "find-target-indices":
		algs.FindTargetIndicesMain()
	case "k-weakest-rows":
		algs.KWeakestRowsMain()
	case "maximum-count":
		algs.MaximumCountMain()
	case "find-distance-value":
		algs.FindDistanceValueMain()
	case "ship-within-days":
		algs.ShipWithinDaysMain()
	case "max-consecutive-answers":
		algs.MaxConsecutiveAnswersMain()
	case "search-suggestions":
		algs.SuggestedProductsMain()
	default:
		panic("missed case")
	}
}
