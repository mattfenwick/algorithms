package logic

import (
	"fmt"
	"strings"

	"github.com/mattfenwick/collections/pkg/set"
)

func Run() {
	tree := Not(
		And(
			Not(Var("P")),
			Not(Var("Q")),
		),
	)
	fmt.Println(PrettyPrint(tree))
	fmt.Println(Evaluate(tree, map[string]bool{"P": false, "Q": false}))

	elimImp := &ElimImplicationRule{
		If:   Var("T"),
		Then: Var("U"),
	}
	fmt.Printf("%s\n", PrettyPrint(StandardForm(elimImp)))
	env := &Environment{
		TrueNodes: set.FromSlice([]string{"T", "T -> U"}),
	}
	fmt.Printf("before: %+v\n", strings.Join(env.TrueNodes.ToSlice(), ","))
	if err := env.Apply(elimImp); err != nil {
		panic(err)
	}
	fmt.Printf("after: %+v\n", strings.Join(env.TrueNodes.ToSlice(), ","))

	intImp := &IntroImplicationRule{If: Implication(Var("T"), Var("U")), Then: Var("T")}
	if err := env.Apply(intImp); err != nil {
		panic(err)
	}
	fmt.Printf("after intro ->: %+v\n", strings.Join(env.TrueNodes.ToSlice(), ","))
}
