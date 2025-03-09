package pratt

import "fmt"

func Run(s string) {
	node := ParseString(s)
	fmt.Printf("%s\n", NodeString(node))
}
