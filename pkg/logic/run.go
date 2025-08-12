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
	fmt.Println(PrettyPrint(tree))
	fmt.Println(Evaluate(tree, map[string]bool{"P": false, "Q": false}))

	parentEnv := &Environment{TrueTerms: set.FromSlice([]string{"Z"})}
	env := &Environment{
		Parent:    parentEnv,
		TrueTerms: set.FromSlice([]string{"T", "T -> U"}),
	}
	fmt.Println("before:")
	env.Print(0)

	var rules = []Rule{
		&ElimImplicationRule{If: Var("T"), Then: Var("U")},
		&IntroImplicationRule{If: Implication(Var("T"), Var("U")), Then: Var("T")},
		&ReiterateRule{Term: Var("Z")},
	}
	for _, r := range rules {
		str := PrettyPrint(StandardForm(r))
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
