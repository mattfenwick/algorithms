package logic

import (
	"fmt"
)

func Run() {
	tree := Not(
		And(
			Not(Var("P")),
			Not(Var("Q"))))
	fmt.Println(tree.TermPrint(true))
	fmt.Println(Evaluate(tree, map[string]bool{"P": false, "Q": false}))

	// parentEnv := &Scope{TrueTerms: set.FromSlice([]string{"Z"})}
	// env := &Scope{
	// 	Parent: parentEnv,
	// 	TrueTerms: set.FromSlice([]string{
	// 		"T",
	// 		"T -> U",
	// 		"A",
	// 		"B",
	// 		"Z -> X",
	// 		"C -> X",
	// 	}),
	// }
	// fmt.Println("before:")
	// env.Print(0)

	// var rules = []*Rule{
	// 	EImply(Var("T"), Var("U")),
	// 	IImply(Implication(Var("T"), Var("U")), Var("T")),
	// 	IAnd(Var("Z"), Var("U")),
	// 	EAnd(Var("A"), Var("B"), false),
	// 	IOr(Var("Z"), Var("C"), true),
	// 	EOr(Var("Z"), Var("C"), Var("X")),
	// }
	// for _, r := range rules {
	// 	str := r.StandardForm().TermPrint(true)
	// 	fmt.Printf("processing rule '%s'\n", str)
	// 	if err := env.ApplyRule(r); err != nil {
	// 		fmt.Printf("unable to apply rule '%s': %s\n", str, err)
	// 	} else {
	// 		fmt.Printf("successfully applied rule '%s'\n", str)
	// 	}
	// 	fmt.Println("after rule ", str)
	// 	env.Print(0)

	// 	if err := parentEnv.ApplyRule(r); err != nil {
	// 		fmt.Printf("unable to apply rule '%s' in parent env: %s\n", str, err)
	// 	} else {
	// 		fmt.Printf("successfully applied rule '%s' in parent\n", str)
	// 	}
	// 	fmt.Printf("after rule %s in parent\n", str)
	// 	env.Print(0)
	// 	fmt.Println()
	// }

	proof := NewRootProof(
		"P -> P",
		NewProofImplication(Var("P"), &Repeat{Var("P")}),
		// P -> P ^ P v ~P
		NewProofImplication(Var("P"),
			IAnd(Var("P"), Var("P")),
			IOr(And(Var("P"), Var("P")), Not(Var("P")), true),
		),
	)
	env, err := CheckProof(proof)
	fmt.Printf("result from proof '%s': %s\n", proof.Name, err)
	env.Print(0)

	for _, eg := range examples {
		env, err := CheckProof(eg)
		fmt.Printf("\n\nresult from proof '%s': %s\n", eg.Name, err)
		env.Print(0)
	}
}

func Must[A any](a A, err error) A {
	if err != nil {
		panic(err)
	}
	return a
}
