package pratt

import "fmt"

func Run(s string) {
	node := Must(ParseString(s))
	fmt.Printf("%s\n", NodeString(node))
}
