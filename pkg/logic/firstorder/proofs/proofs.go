package proofs

import . "github.com/mattfenwick/algorithms/pkg/logic/firstorder"

var (
	P = Prop("P")
	Q = Prop("Q")
	R = Prop("R")
	S = Prop("S")
)

type ProofsSection struct {
	Name   string
	Proofs []*Proof
}

func NewProofsSection(name string, proofs ...*Proof) *ProofsSection {
	return &ProofsSection{Name: name, Proofs: proofs}
}

var proofs = []*ProofsSection{
	NewProofsSection("basics",
		NewRootProof("P v ~ P",
			NewProofContradiction(Not(Or(P, Not(P))),
				NewProofContradiction(P,
					IOr(P, Not(P), true),
					&Reiterate{Term: Not(Or(P, Not(P)))},
				),
				IOr(P, Not(P), false),
			),
		),
		NewRootProof("∀x.( Q(x) ) -> Q(a)",
			NewProofImplication(Forall("x", Prop("Q", "x")),
				EForall(Prop("Q", "x"), "x", "a"),
			),
		),
		NewRootProof("Q(a) -> ∃x.( Q(x) )",
			NewProofImplication(Prop("Q", "a"),
				IExistential(Prop("Q", "a"), "a", "x"),
			),
		),
		NewRootProof("∃x.( Q(x) ^ ( Q(x) -> R ) ) -> R",
			NewProofImplication(Existential("x", And(Prop("Q", "x"), Implication(Prop("Q", "x"), R))),
				NewProofExistentialElim("a",
					Existential("x", And(Prop("Q", "x"), Implication(Prop("Q", "x"), R))), // Q(a) ^ Q(a) -> R
					EAnd(Prop("Q", "a"), Implication(Prop("Q", "a"), R), true),            // Q(a)
					EAnd(Prop("Q", "a"), Implication(Prop("Q", "a"), R), false),           // Q(a) -> R
					EImply(Prop("Q", "a"), R),                                             // R
				),
			),
		),
		NewRootProof("( ∀y.( Q(y) ) ^ ∃x.( Q(x) -> R ) ) -> R",
			NewProofImplication(And(Forall("y", Prop("Q", "y")), Existential("x", Implication(Prop("Q", "x"), R))),
				EAnd(Forall("y", Prop("Q", "y")), Existential("x", Implication(Prop("Q", "x"), R)), true),  // ∀y.( Q(y) )
				EAnd(Forall("y", Prop("Q", "y")), Existential("x", Implication(Prop("Q", "x"), R)), false), // ∃x.( Q(x) -> R )
				NewProofExistentialElim("a",
					Existential("x", Implication(Prop("Q", "x"), R)), // Q(a) -> R
					&Reiterate{Term: Forall("y", Prop("Q", "y"))},    // ∀y.( Q(y) )
					EForall(Prop("Q", "y"), "y", "a"),                // Q(a)
					EImply(Prop("Q", "a"), R),                        // R
				),
			),
		),
		// example of using ∃ hypothesis in conclusion
		// NewRootProof("∃x.( Q(x) ) -> Q(a)",
		// 	NewProofImplication(Existential("x", Prop("Q", "x")),
		// 		NewProofExistentialElim("a",
		// 			Existential("x", Prop("Q", "x")), // Q(a)
		// 		),
		// 	),
		// ),
		// example of using ∃ which isn't available
		// NewRootProof("R",
		// 	NewProofExistentialElim("a",
		// 		Existential("x", And(R, Prop("Q", "x"))), // R ^ Q(a)
		// 		EAnd(R, Prop("Q", "a"), true),            // R
		// 	),
		// ),
	),
}
