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

	var rules = []Rule{
		&ElimImplicationRule{If: Var("T"), Then: Var("U")},
		&IntroImplicationRule{If: Implication(Var("T"), Var("U")), Then: Var("T")},
		&ReiterateRule{Term: Var("Z")},
		&IntroAndRule{Left: Var("Z"), Right: Var("U")},
		&ElimAndRule{Left: Var("A"), Right: Var("B"), IsLeft: false},
		&IntroOrRule{Left: Var("Z"), Right: Var("C"), IsLeft: true},
		&ElimOrRule{If1: Var("Z"), If2: Var("C"), Then: Var("X")},
	}
	for _, r := range rules {
		str := StandardForm(r).TermPrint(true)
		fmt.Printf("processing rule '%s'\n", str)
		if err := env.Apply(r); err != nil {
			fmt.Printf("unable to apply rule '%s': %s\n", str, err)
		} else {
			fmt.Printf("successfully applied rule '%s'\n", str)
		}
		fmt.Println("after rule ", str)
		env.Print(0)

		if err := parentEnv.Apply(r); err != nil {
			fmt.Printf("unable to apply rule '%s' in parent env: %s\n", str, err)
		} else {
			fmt.Printf("successfully applied rule '%s' in parent\n", str)
		}
		fmt.Printf("after rule %s in parent\n", str)
		env.Print(0)
		fmt.Println()
	}
}
