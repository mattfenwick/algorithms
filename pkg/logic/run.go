package logic

import "fmt"

func Run() {
	tree := Not(
		And(
			Not(Var("P")),
			Not(Var("Q")),
		),
	)
	fmt.Println(PrettyPrint(tree))
	fmt.Println(Evaluate(tree, map[string]bool{"P": false, "Q": false}))
}
