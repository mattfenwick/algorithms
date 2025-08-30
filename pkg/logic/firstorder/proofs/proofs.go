package proofs

import . "github.com/mattfenwick/algorithms/pkg/logic/firstorder"

var (
	P = Prop("P")
	Q = Prop("Q")
	R = Prop("R")
	S = Prop("S")

	T = Prop("T") // True
	F = Prop("F") // False

	Qa = Prop("Q", "a")
	Qb = Prop("Q", "b")
	Qx = Prop("Q", "x")
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
					&Reiterate{Formula: Not(Or(P, Not(P)))},
				),
				IOr(P, Not(P), false),
			),
		),
		// NewRootProof("∀x.( Q(x) ) -> Q(a)",
		// 	NewProofImplication(Forall("x", Prop("Q", "x")),
		// 		EForall(Prop("Q", "x"), "x", "a"),
		// 	),
		// ),
		NewRootProof("Q(a) -> ∃x.( Q(x) )",
			NewProofImplication(Prop("Q", "a"),
				IExistential(Prop("Q", "a"), "a", "x"),
			),
		),
		NewRootProof("( ∃x.( T ) ^ ( P -> ∃x.( Q(x) ) ) ) <-> ∃x.( P -> Q(x) )",
			NewProofImplication(And(Existential("x", T), Implication(P, Existential("x", Qx))),
				EAnd(Existential("x", T), Implication(P, Existential("x", Qx)), true),  // ∃x.( T )
				EAnd(Existential("x", T), Implication(P, Existential("x", Qx)), false), // P -> ∃x.( Q(x) )
				NewProofExistentialElim("a", Existential("x", T),
					NewProofImplication(P,
						&Reiterate{Formula: Implication(P, Existential("x", Qx))}, // P -> ∃x.( Q(x) )
						EImply(P, Existential("x", Qx)),                           // ∃x.( Q(x)
						NewProofExistentialElim("b", Existential("x", Qx),
							&Reiterate{Formula: P},                     // P
							IImply(P, Qb),                              // P -> Q(b)
							IExistential(Implication(P, Qb), "b", "x"), // ∃x.( P -> Q(x) )
						), // ∃x.( P -> Q(x) )
					), // P -> ∃x.( P -> Q(x) )
					NewProofImplication(Not(P),
						NewProofImplication(Not(Qa),
							&Reiterate{Formula: Not(P)},
						), // ~ Q(a) -> ~ P
						ContrapositiveTheorem(Not(Qa), Not(P)),     // P -> Q(a)
						IExistential(Implication(P, Qa), "a", "x"), // ∃x.( P -> Q(x) )
					), // ~ P -> ∃x.( P -> Q(x) )
					ExcludedMiddleTheorem(P), // P v ~ P
					EOr(P, Not(P), Existential("x", Implication(P, Qx))),
				),
			),
			NewProofImplication(Existential("x", Implication(P, Qx)),
				NewProofImplication(P,
					&Reiterate{Formula: Existential("x", Implication(P, Qx))},
					NewProofExistentialElim("a", Existential("x", Implication(P, Qx)),
						&Reiterate{Formula: P},     // P
						EImply(P, Qa),              // Q(a)
						IExistential(Qa, "a", "x"), // ∃x.( Q(x) )
					), // ∃x.( Q(x) )
				), // P -> ∃x.( Q(x) )
				NewProofExistentialElim("a", Existential("x", Implication(P, Qx)),
					&Reiterate{Formula: T},    // T
					IExistential(T, "a", "x"), // ∃x.( T )
				), // ∃x.( T )
				IAnd(Existential("x", T), Implication(P, Existential("x", Qx))),
			),
			IBiconditional(
				And(Existential("x", T), Implication(P, Existential("x", Qx))),
				Existential("x", Implication(P, Qx)),
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
					&Reiterate{Formula: Forall("y", Prop("Q", "y"))}, // ∀y.( Q(y) )
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
					&Reiterate{Formula: Forall("x", And(Prop("P", "x"), Prop("Q", "x")))}, // ∀x.( P(x) ^ Q(x) )
					// TODO are we actually checking if 'a' is "in scope" before doing this?
					EForall(And(Prop("P", "x"), Prop("Q", "x")), "x", "a"), // P(a) ^ Q(a)
					EAnd(Prop("P", "a"), Prop("Q", "a"), true),             // P(a)
				), // ∀y.( P(y)
				NewProofForallIntro("z", "a",
					&Reiterate{Formula: Forall("x", And(Prop("P", "x"), Prop("Q", "x")))}, // ∀x.( P(x) ^ Q(x) )
					// TODO are we actually checking if 'a' is "in scope" before doing this?
					EForall(And(Prop("P", "x"), Prop("Q", "x")), "x", "a"), // P(a) ^ Q(a)
					EAnd(Prop("P", "a"), Prop("Q", "a"), false),            // Q(a)
				), // ∀z.( Q(z)
				IAnd(Forall("y", Prop("P", "y")), Forall("z", Prop("Q", "z"))),
			),
		),
	),
}
