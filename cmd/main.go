package main

import (
	"fmt"
	"os"

	"github.com/mattfenwick/algorithms/pkg"
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
	default: panic("missed case")
	}
}
