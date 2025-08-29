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
		// TODO might not be possible to prove?
		// NewRootProof("( P -> ∃x.( Q(x) ) ) -> ∃x.( P -> Q(x) )",
		// 	NewProofImplication(Implication(P, Existential("x", Prop("Q", "x")),
		// 		IExistential(Prop("Q", "a"), "a", "x"),
		// 	),
		// ),
		NewRootProof("∃x.( P -> Q(x) ) -> ( P -> ∃y.( Q(y) ) )",
			NewProofImplication(Existential("x", Implication(P, Prop("Q", "x"))),
				NewProofImplication(P,
					&Reiterate{Term: Existential("x", Implication(P, Prop("Q", "x")))},
					NewProofExistentialElim("a", Existential("x", Implication(P, Prop("Q", "x"))),
						&Reiterate{Term: P},                    // P
						EImply(P, Prop("Q", "a")),              // Q(a)
						IExistential(Prop("Q", "a"), "a", "y"), // ∃y.( Q(y) )
					), // ∃y.( Q(y) )
				), // P -> ∃y.( Q(y) )
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
		NewRootProof("∀x.( P(x) ^ Q(x) ) -> ( ∀y.( P(y) ) ^ ∀z.( Q(z) ) )",
			NewProofImplication(Forall("x", And(Prop("P", "x"), Prop("Q", "x"))),
				NewProofForallIntro("y", "a",
					&Reiterate{Term: Forall("x", And(Prop("P", "x"), Prop("Q", "x")))}, // ∀x.( P(x) ^ Q(x) )
					// TODO are we actually checking if 'a' is "in scope" before doing this?
					EForall(And(Prop("P", "x"), Prop("Q", "x")), "x", "a"), // P(a) ^ Q(a)
					EAnd(Prop("P", "a"), Prop("Q", "a"), true),             // P(a)
				), // ∀y.( P(y)
				NewProofForallIntro("z", "a",
					&Reiterate{Term: Forall("x", And(Prop("P", "x"), Prop("Q", "x")))}, // ∀x.( P(x) ^ Q(x) )
					// TODO are we actually checking if 'a' is "in scope" before doing this?
					EForall(And(Prop("P", "x"), Prop("Q", "x")), "x", "a"), // P(a) ^ Q(a)
					EAnd(Prop("P", "a"), Prop("Q", "a"), false),            // Q(a)
				), // ∀z.( Q(z)
				IAnd(Forall("y", Prop("P", "y")), Forall("z", Prop("Q", "z"))),
			),
		),
	),
}
