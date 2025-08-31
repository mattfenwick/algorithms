package proofs

import . "github.com/mattfenwick/algorithms/pkg/logic/firstorder"

var (
	P = Pred("P")
	Q = Pred("Q")
	R = Pred("R")
	S = Pred("S")

	T = Pred("T") // True
	F = Pred("F") // False

	Qa = Pred("Q", &TermVar{Name: "a", IsBound: false})
	Qb = Pred("Q", &TermVar{Name: "b", IsBound: false})
	Qx = Pred("Q", &TermVar{Name: "x", IsBound: true})
	Qy = Pred("Q", &TermVar{Name: "y", IsBound: true})
	Qz = Pred("Q", &TermVar{Name: "z", IsBound: true})

	Pa = Pred("P", &TermVar{Name: "a", IsBound: false})
	Pb = Pred("P", &TermVar{Name: "b", IsBound: false})
	Px = Pred("P", &TermVar{Name: "x", IsBound: true})
	Py = Pred("P", &TermVar{Name: "y", IsBound: true})
	Pz = Pred("P", &TermVar{Name: "z", IsBound: true})
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
		RootProof("P v ~ P",
			ContraProof(Not(Or(P, Not(P))),
				ContraProof(P,
					IOr(P, Not(P), true),
					&Reiterate{Formula: Not(Or(P, Not(P)))},
				),
				IOr(P, Not(P), false),
			),
		),

		// the following two are not allowed -- a is free and not in scope
		// NewRootProof("∀x.( Q(x) ) -> Q(a)",
		// 	NewProofImplication(Forall("x", Prop("Q", "x")),
		// 		EForall(Prop("Q", "x"), "x", "a"),
		// 	),
		// ),
		// RootProof("Q(a) -> ∃x.( Q(x) )",
		// 	ArrowProof(Qa,
		// 		IExist(Qa, "a", "x"),
		// 	),
		// ),

		RootProof("( ∃x.( T ) ^ ( P -> ∃x.( Q(x) ) ) ) <-> ∃x.( P -> Q(x) )",
			ArrowProof(And(Exist("x", T), Arrow(P, Exist("x", Qx))),
				EAnd(Exist("x", T), Arrow(P, Exist("x", Qx)), true),  // ∃x.( T )
				EAnd(Exist("x", T), Arrow(P, Exist("x", Qx)), false), // P -> ∃x.( Q(x) )
				ExistElimProof("a", Exist("x", T),
					ArrowProof(P,
						&Reiterate{Formula: Arrow(P, Exist("x", Qx))}, // P -> ∃x.( Q(x) )
						EImply(P, Exist("x", Qx)),                     // ∃x.( Q(x)
						ExistElimProof("b", Exist("x", Qx),
							&Reiterate{Formula: P},         // P
							IImply(P, Qb),                  // P -> Q(b)
							IExist(Arrow(P, Qb), "b", "x"), // ∃x.( P -> Q(x) )
						), // ∃x.( P -> Q(x) )
					), // P -> ∃x.( P -> Q(x) )
					ArrowProof(Not(P),
						ArrowProof(Not(Qa),
							&Reiterate{Formula: Not(P)},
						), // ~ Q(a) -> ~ P
						ContrapositiveTheorem(Not(Qa), Not(P)), // P -> Q(a)
						IExist(Arrow(P, Qa), "a", "x"),         // ∃x.( P -> Q(x) )
					), // ~ P -> ∃x.( P -> Q(x) )
					ExcludedMiddleTheorem(P), // P v ~ P
					EOr(P, Not(P), Exist("x", Arrow(P, Qx))),
				),
			),
			ArrowProof(Exist("x", Arrow(P, Qx)),
				ArrowProof(P,
					&Reiterate{Formula: Exist("x", Arrow(P, Qx))},
					ExistElimProof("a", Exist("x", Arrow(P, Qx)),
						&Reiterate{Formula: P}, // P
						EImply(P, Qa),          // Q(a)
						IExist(Qa, "a", "x"),   // ∃x.( Q(x) )
					), // ∃x.( Q(x) )
				), // P -> ∃x.( Q(x) )
				ExistElimProof("a", Exist("x", Arrow(P, Qx)),
					&Reiterate{Formula: T}, // T
					IExist(T, "a", "x"),    // ∃x.( T )
				), // ∃x.( T )
				IAnd(Exist("x", T), Arrow(P, Exist("x", Qx))),
			),
			IDArrow(
				And(Exist("x", T), Arrow(P, Exist("x", Qx))),
				Exist("x", Arrow(P, Qx)),
			),
		),
		RootProof("∃x.( Q(x) ^ ( Q(x) -> R ) ) -> R",
			ArrowProof(Exist("x", And(Qx, Arrow(Qx, R))),
				ExistElimProof("a",
					Exist("x", And(Qx, Arrow(Qx, R))), // Q(a) ^ Q(a) -> R
					EAnd(Qa, Arrow(Qa, R), true),      // Q(a)
					EAnd(Qa, Arrow(Qa, R), false),     // Q(a) -> R
					EImply(Qa, R),                     // R
				),
			),
		),
		RootProof("( ∀y.( Q(y) ) ^ ∃x.( Q(x) -> R ) ) -> R",
			ArrowProof(And(Forall("y", Qy), Exist("x", Arrow(Qx, R))),
				EAnd(Forall("y", Qy), Exist("x", Arrow(Qx, R)), true),  // ∀y.( Q(y) )
				EAnd(Forall("y", Qy), Exist("x", Arrow(Qx, R)), false), // ∃x.( Q(x) -> R )
				ExistElimProof("a",
					Exist("x", Arrow(Qx, R)),             // Q(a) -> R
					&Reiterate{Formula: Forall("y", Qy)}, // ∀y.( Q(y) )
					EForall(Qy, "y", "a"),                // Q(a)
					EImply(Qa, R),                        // R
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
		RootProof("∀x.( P(x) ^ Q(x) ) -> ( ∀y.( P(y) ) ^ ∀z.( Q(z) ) )",
			ArrowProof(Forall("x", And(Px, Qx)),
				ForallIntroProof("y", "a",
					&Reiterate{Formula: Forall("x", And(Px, Qx))}, // ∀x.( P(x) ^ Q(x) )
					// TODO are we actually checking if 'a' is "in scope" before doing this?
					EForall(And(Px, Qx), "x", "a"), // P(a) ^ Q(a)
					EAnd(Pa, Qa, true),             // P(a)
				), // ∀y.( P(y)
				ForallIntroProof("z", "a",
					&Reiterate{Formula: Forall("x", And(Px, Qx))}, // ∀x.( P(x) ^ Q(x) )
					// TODO are we actually checking if 'a' is "in scope" before doing this?
					EForall(And(Px, Qx), "x", "a"), // P(a) ^ Q(a)
					EAnd(Pa, Qa, false),            // Q(a)
				), // ∀z.( Q(z)
				IAnd(Forall("y", Py), Forall("z", Qz)),
			),
		),

		// RootProof("∃x.( R ^ Q(x) ) -> R",
		// 	ArrowProof(Exist("x", And(R, Qx)), // ∃x.( R ^ Q(x) )
		// 		ExistElimProof("a",
		// 			Exist("x", And(R, Qx)), // R ^ Q(a)
		// 			EAnd(R, Qa, true),      // R
		// 		),
		// 	),
		// ),

		// RootProof("( ∀x.( Q(x) ) ^ ∃x.( T ) ) -> Q(a)",
		// 	ArrowProof(And(Forall("x", Qx), Exist("x", T)),
		// 		ExistElimProof("a",
		// 			Exist("x", T),
		// 			EForall(Qx, "x", "a"),
		// 		),
		// 	),
		// ),

		RootProof("∀x.( ~ Q(x) ) <-> ~ ∃x.( Q(x) )",
			ArrowProof(Forall("x", Not(Qx)),
				ExistContraProof("a", Exist("x", Qx),
					&Reiterate{Formula: Forall("x", Not(Qx))}, // ∀x.( ~ Q(x) )
					EForall(Not(Qx), "x", "a"),                // ~ Q(a)
				),
			),
			ArrowProof(Not(Exist("x", Qx)),
				ForallIntroProof("x", "a",
					ContraProof(Qa,
						IExist(Qa, "a", "x"),                     // ∃x.( Q(x) )
						&Reiterate{Formula: Not(Exist("x", Qx))}, // ~ ∃x.( Q(x) )
					), // ~ Q(a)
				), // ∀x.( ~ Q(x) )
			),
			IDArrow(Forall("x", Not(Qx)), Not(Exist("x", Qx))),
		),
		RootProof("~ ∀x.( Q(x) ) <-> ∃x.( ~ Q(x) )",
			ArrowProof(Not(Forall("x", Qx)), // ~ ∀x.( Q(x) )
				ContraProof(Not(Exist("x", Not(Qx))), // ~ ∃x.( ~ Q(x) )
					ForallIntroProof("x", "a",
						ContraProof(Not(Qa), // ~ Q(a)
							IExist(Not(Qa), "a", "x"),                     // ∃x.( ~ Q(x) )
							&Reiterate{Formula: Not(Exist("x", Not(Qx)))}, // ~ ∃x.( ~ Q(x) )
						), // Q(a)
					), // ∀x.( Q(x) )
					&Reiterate{Formula: Not(Forall("x", Qx))}, // ~ ∀x.( Q(x) )
				), // ∃x.( ~ Q(x) )
			),
			ArrowProof(Exist("x", Not(Qx)), // ∃x.( ~ Q(x) )
				ExistElimProof("a", Exist("x", Not(Qx)), // ~ Q(a)
					ContraProof(Forall("x", Qx), // ∀x.( Q(x) )
						EForall(Qx, "x", "a"),        // Q(a)
						&Reiterate{Formula: Not(Qa)}, // ~ Q(a)
					), // ~ ∀x.( Q(x) )
				), // ~ // ∀x.( Q(x) )
			),
			IDArrow(Not(Forall("x", Qx)), Exist("x", Not(Qx))),
		),
	),
}
