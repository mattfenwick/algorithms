package logic

import (
	"fmt"

	"github.com/mattfenwick/collections/pkg/set"
)

func Run() {
	tree := Not(
		And(
			Not(Var("P")),
			Not(Var("Q"))))
	fmt.Println(tree.TermPrint(true))
	fmt.Println(Evaluate(tree, map[string]bool{"P": false, "Q": false}))

	parentEnv := &Environment{TrueTerms: set.FromSlice([]string{"Z"})}
	env := &Environment{
		Parent: parentEnv,
		TrueTerms: set.FromSlice([]string{
			"T",
			"T -> U",
			"A",
			"B",
			"Z -> X",
			"C -> X",
		}),
	}
	fmt.Println("before:")
	env.Print(0)

	var rules = []*Rule{
		ElimImplicationRule(Var("T"), Var("U")),
		IntroImplicationRule(Implication(Var("T"), Var("U")), Var("T")),
		IntroAndRule(Var("Z"), Var("U")),
		ElimAndRule(Var("A"), Var("B"), false),
		IntroOrRule(Var("Z"), Var("C"), true),
		ElimOrRule(Var("Z"), Var("C"), Var("X")),
	}
	for _, r := range rules {
		str := r.StandardForm().TermPrint(true)
		fmt.Printf("processing rule '%s'\n", str)
		if err := env.ApplyRule(r); err != nil {
			fmt.Printf("unable to apply rule '%s': %s\n", str, err)
		} else {
			fmt.Printf("successfully applied rule '%s'\n", str)
		}
		fmt.Println("after rule ", str)
		env.Print(0)

		if err := parentEnv.ApplyRule(r); err != nil {
			fmt.Printf("unable to apply rule '%s' in parent env: %s\n", str, err)
		} else {
			fmt.Printf("successfully applied rule '%s' in parent\n", str)
		}
		fmt.Printf("after rule %s in parent\n", str)
		env.Print(0)
		fmt.Println()
	}

	proof := NewProof(
		// NewSubProofImplication(Var("P"), Var("P")), // P -> P , umm actually not sure how to write this. let's circle back TODO
		// P -> P ^ P v ~P
		Must(NewSubProofImplication(Var("P"),
			IntroAndRule(Var("P"), Var("P")),
			IntroOrRule(And(Var("P"), Var("P")), Not(Var("P")), true),
		)),
	)
	env, err := CheckProof(proof)
	fmt.Printf("result from proof: %s\n", err)
	env.Print(0)
}

func Must[A any](a A, err error) A {
	if err != nil {
		panic(err)
	}
	return a
}
