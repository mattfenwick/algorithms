package pratt

import "fmt"

func Run(s string) {
	node := Must(TestOperators.ParseString(s))
	fmt.Printf("%s\n\n", NodeString(node))
	fmt.Printf("%s\n\n", Parens(node))
}
