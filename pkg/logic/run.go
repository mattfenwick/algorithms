package logic

import (
	"fmt"

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
	parentEnv := &Environment{TrueTerms: set.FromSlice([]string{"Z"})}
	env := &Environment{
		Parent:    parentEnv,
		TrueTerms: set.FromSlice([]string{"T", "T -> U"}),
	}
	fmt.Println("before:")
	env.Print(0)
	if err := env.Apply(elimImp); err != nil {
		panic(err)
	}
	fmt.Println("after:")
	env.Print(0)

	intImp := &IntroImplicationRule{If: Implication(Var("T"), Var("U")), Then: Var("T")}
	if err := env.Apply(intImp); err != nil {
		panic(err)
	}
	fmt.Println("after intro:")
	env.Print(0)

	reiterate := &ReiterateRule{Term: Var("Z")}
	if err := env.Apply(reiterate); err != nil {
		panic(err)
	}
	fmt.Println("after reiterate:")
	env.Print(0)

	if err := parentEnv.Apply(reiterate); err != nil {
		fmt.Println(err)
	}
	fmt.Println("after reiterate on parent:")
	env.Print(0)
}
