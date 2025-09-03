package proofs

import . "github.com/mattfenwick/algorithms/pkg/logic/firstorder"

var (
	P = Pred("P")
	Q = Pred("Q")
	R = Pred("R")
	S = Pred("S")

	T = Pred("T") // True
	F = Pred("F") // False

	a = FreeTermVar("a")
	b = FreeTermVar("b")
	c = FreeTermVar("c")

	x = BoundTermVar("x")
	y = BoundTermVar("y")
	z = BoundTermVar("z")

	Qa = Pred("Q", a)
	Qb = Pred("Q", b)
	Qx = Pred("Q", x)
	Qy = Pred("Q", y)
	Qz = Pred("Q", z)

	Pa = Pred("P", a)
	Pb = Pred("P", b)
	Px = Pred("P", x)
	Py = Pred("P", y)
	Pz = Pred("P", z)

	Pab = Pred("P", a, b)
	Pax = Pred("P", a, x)
	Pay = Pred("P", a, y)
	Pba = Pred("P", b, a)
	Pby = Pred("P", b, y)
	Pxy = Pred("P", x, y)
	Pxb = Pred("P", x, b)
)

type ProofsSection struct {
	Name   string
	Proofs []*Proof
}

func NewProofsSection(name string, proofs ...*Proof) *ProofsSection {
	return &ProofsSection{Name: name, Proofs: proofs}
}

var quantifierProofs = []*ProofsSection{
	NewProofsSection("basics",
		RootProof("∀x.( P(x) v ~ P(x) )",
			ForallIntroProof("x", "a",
				ContraProof(Not(Or(Pa, Not(Pa))), // ~ ( P(a) v ~ P(a) )
					ContraProof(Pa, // P(a)
						IOr(Pa, Not(Pa), true),                    // P(a) v ~ P(a)
						&Reiterate{Formula: Not(Or(Pa, Not(Pa)))}, // ~ ( P(a) v ~ P(a) )
					), // ~ P(a)
					IOr(Pa, Not(Pa), false), // P(a) v ~ P(a)
				), // P(a) v ~ P(a)
			),
		),
		RootProof("~ ∃x.( ~ ( P(x) v ~ P(x) ) )",
			ExistContraProof("a",
				Exist("x", Not(Or(Px, Not(Px)))), // ~ ( P(a) v ~ P(a) )
				ContraProof(Pa, // P(a)
					IOr(Pa, Not(Pa), true),                    // P(a) v ~ P(a)
					&Reiterate{Formula: Not(Or(Pa, Not(Pa)))}, // ~ ( P(a) v ~ P(a) )
				), // ~ P(a)
				IOr(Pa, Not(Pa), false), // P(a) v ~ P(a)
			),
		),
		RootProof("∀x.( ~ ( P(x) ^ ~ P(x) ) )",
			ForallIntroProof("x", "a",
				ContraProof(
					And(Pa, Not(Pa)),
					EAnd(Pa, Not(Pa), true),
					EAnd(Pa, Not(Pa), false),
				),
			),
		),
		RootProof("~ ∃x.( P(x) ^ ~ P(x) )",
			ExistContraProof("a",
				Exist("x", And(Px, Not(Px))), // P(a) ^ ~ P(a)
				EAnd(Pa, Not(Pa), true),      // P(a)
				EAnd(Pa, Not(Pa), false),     // ~ P(a)
			),
		),
	),
	NewProofsSection("DeMorgan's",
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
	NewProofsSection("exist-or-empty",
		RootProof("( ∀x.( Q(x) ) ^ ∀x.( ~ Q(x) ) ) <-> ( ~ ∃x.( Q(x) ) ^ ~ ∃x.( ~ Q(x) ) )",
			ArrowProof(And(Forall("x", Qx), Forall("x", Not(Qx))),
				EAnd(Forall("x", Qx), Forall("x", Not(Qx)), true),
				EAnd(Forall("x", Qx), Forall("x", Not(Qx)), false),
				DeMorgansForallToExistTheorem(Forall("x", Qx), false),
				DeMorgansForallToExistTheorem(Forall("x", Not(Qx)), false),
				IAnd(Not(Exist("x", Qx)), Not(Exist("x", Not(Qx)))),
			),
			ArrowProof(And(Not(Exist("x", Qx)), Not(Exist("x", Not(Qx)))),
				EAnd(Not(Exist("x", Qx)), Not(Exist("x", Not(Qx))), true),
				EAnd(Not(Exist("x", Qx)), Not(Exist("x", Not(Qx))), false),
				DeMorgansExistToForallTheorem(Exist("x", Qx), true),
				DeMorgansExistToForallTheorem(Exist("x", Not(Qx)), true),
				IAnd(Forall("x", Qx), Forall("x", Not(Qx))),
			),
			IDArrow(
				And(Forall("x", Qx), Forall("x", Not(Qx))),
				And(Not(Exist("x", Qx)), Not(Exist("x", Not(Qx)))),
			),
		),
		RootProof("( ∃x.( Q(x) ) ^ ∃x.( ~ Q(x) ) ) <-> ( ~ ∀x.( Q(x) ) ^ ~ ∀x.( ~ Q(x) ) )",
			ArrowProof(And(Not(Forall("x", Qx)), Not(Forall("x", Not(Qx)))),
				EAnd(Not(Forall("x", Qx)), Not(Forall("x", Not(Qx))), true),
				EAnd(Not(Forall("x", Qx)), Not(Forall("x", Not(Qx))), false),
				DeMorgansForallToExistTheorem(Forall("x", Qx), true),
				DeMorgansForallToExistTheorem(Forall("x", Not(Qx)), true),
				IAnd(Exist("x", Qx), Exist("x", Not(Qx))),
			),
			ArrowProof(And(Exist("x", Qx), Exist("x", Not(Qx))),
				EAnd(Exist("x", Qx), Exist("x", Not(Qx)), true),
				EAnd(Exist("x", Qx), Exist("x", Not(Qx)), false),
				DeMorgansExistToForallTheorem(Exist("x", Qx), false),
				DeMorgansExistToForallTheorem(Exist("x", Not(Qx)), false),
				IAnd(Not(Forall("x", Qx)), Not(Forall("x", Not(Qx)))),
			),
			IDArrow(
				And(Exist("x", Qx), Exist("x", Not(Qx))),
				And(Not(Forall("x", Qx)), Not(Forall("x", Not(Qx)))),
			),
		),
		RootProof("( ∃x.( Q(x) ) v ∃x.( ~ Q(x) ) ) <-> ∃x.( T )",
			ArrowProof(Or(Exist("x", Qx), Exist("x", Not(Qx))),
				ArrowProof(Exist("x", Qx),
					ExistElimProof("a",
						Exist("x", Qx),
						&Reiterate{Formula: T},
						IExist(T, "a", "x"),
					), // ∃x.( T )
				), // ∃x.( Q(x) ) -> ∃x.( T )
				ArrowProof(Exist("x", Not(Qx)),
					ExistElimProof("a",
						Exist("x", Not(Qx)),
						&Reiterate{Formula: T},
						IExist(T, "a", "x"),
					), // ∃x.( T )
				), // ∃x.( ~ Q(x) ) -> ∃x.( T )
				EOr(Exist("x", Qx), Exist("x", Not(Qx)), Exist("x", T)),
			),
			ArrowProof(Exist("x", T),
				ExistElimProof("a",
					Exist("x", T),             // ∃x.( T )
					ExcludedMiddleTheorem(Qa), // Q(a) v ~ Q(a)
					ArrowProof(Qa,
						IExist(Qa, "a", "x"),                           // ∃x.( Q(x) )
						IOr(Exist("x", Qx), Exist("x", Not(Qx)), true), // ∃x.( Q(x) ) v ∃x.( ~ Q(x) )
					), // Q(a) -> ∃x.( Q(x) ) v ∃x.( ~ Q(x) )
					ArrowProof(Not(Qa),
						IExist(Not(Qa), "a", "x"),                       // ∃x.( ~ Q(x) )
						IOr(Exist("x", Qx), Exist("x", Not(Qx)), false), // ∃x.( Q(x) ) v ∃x.( ~ Q(x) )
					), // ~ Q(a) -> ∃x.( Q(x) ) v ∃x.( ~ Q(x) )
					EOr(Qa, Not(Qa), Or(Exist("x", Qx), Exist("x", Not(Qx)))), // ∃x.( Q(x) ) v ∃x.( ~ Q(x) )
				), // ∃x.( Q(x) ) v ∃x.( ~ Q(x) )
			),
			IDArrow(
				Or(Exist("x", Qx), Exist("x", Not(Qx))),
				Exist("x", T),
			),
		),
		RootProof("( ~ ∃x.( Q(x) ) ^ ~ ∃x.( ~ Q(x) ) ) <-> ~ ∃x.( T )",
			ArrowProof(And(Not(Exist("x", Qx)), Not(Exist("x", Not(Qx)))),
				EAnd(Not(Exist("x", Qx)), Not(Exist("x", Not(Qx))), true),  // ~ ∃x.( Q(x) )
				EAnd(Not(Exist("x", Qx)), Not(Exist("x", Not(Qx))), false), // ~ ∃x.( ~ Q(x) )
				ExistContraProof("a",
					Exist("x", T), // T
					ContraProof(Qa,
						&Reiterate{Formula: Not(Exist("x", Qx))}, // ~ ∃x.( Q(x) )
						IExist(Qa, "a", "x"),                     // ∃x.( Q(x) )
					), // ~ Q(a)
					&Reiterate{Formula: Not(Exist("x", Not(Qx)))}, // ~ ∃x.( ~ Q(x) )
					IExist(Not(Qa), "a", "x"),                     // ∃x.( ~ Q(x) )
				), // ~ ∃x.( T )
			),
			ArrowProof(Not(Exist("x", T)),
				ExistContraProof("a",
					Exist("x", Qx),                          // Q(a)
					&Reiterate{Formula: T},                  // T
					IExist(T, "a", "x"),                     // ∃x.( T )
					&Reiterate{Formula: Not(Exist("x", T))}, // ~ ∃x.( T )
				), // ~ ∃x.( Q(x) )
				ExistContraProof("a",
					Exist("x", Not(Qx)),                     // ~ Q(a)
					&Reiterate{Formula: T},                  // T
					IExist(T, "a", "x"),                     // ∃x.( T )
					&Reiterate{Formula: Not(Exist("x", T))}, // ~ ∃x.( T )
				), // ~ ∃x.( ~ Q(x) )
				IAnd(Not(Exist("x", Qx)), Not(Exist("x", Not(Qx)))), // ~ ∃x.( Q(x) ) ^ ~ ∃x.( ~ Q(x) )
			),
			IDArrow(
				And(Not(Exist("x", Qx)), Not(Exist("x", Not(Qx)))),
				Not(Exist("x", T)),
			),
		),
	),
	NewProofsSection("distributive",
		RootProof("∀x.( P(x) ^ Q(x) ) <-> ( ∀x.( P(x) ) ^ ∀x.( Q(x) ) )",
			ArrowProof(Forall("x", And(Px, Qx)),
				ForallIntroProof("x", "a",
					&Reiterate{Formula: Forall("x", And(Px, Qx))},
					EForall(And(Px, Qx), "x", "a"),
					EAnd(Pa, Qa, true), // P(a)
				), // ∀x.( P(x) )
				ForallIntroProof("x", "a",
					&Reiterate{Formula: Forall("x", And(Px, Qx))},
					EForall(And(Px, Qx), "x", "a"),
					EAnd(Pa, Qa, false), // Q(a)
				), // ∀x.( Q(x) )
				IAnd(Forall("x", Px), Forall("x", Qx)),
			),
			ArrowProof(And(Forall("x", Px), Forall("x", Qx)),
				ForallIntroProof("x", "a",
					&Reiterate{Formula: And(Forall("x", Px), Forall("x", Qx))},
					EAnd(Forall("x", Px), Forall("x", Qx), true),  // ∀x.( P(x) )
					EAnd(Forall("x", Px), Forall("x", Qx), false), // ∀x.( Q(x) )
					EForall(Px, "x", "a"),                         // P(a)
					EForall(Qx, "x", "a"),                         // Q(a)
					IAnd(Pa, Qa),
				),
			),
			IDArrow(
				Forall("x", And(Px, Qx)),
				And(Forall("x", Px), Forall("x", Qx)),
			),
		),
		RootProof("∃x.( P(x) v Q(x) ) <-> ( ∃x.( P(x) ) v ∃x.( Q(x) ) )",
			ArrowProof(Exist("x", Or(Px, Qx)),
				ExistElimProof("a",
					Exist("x", Or(Px, Qx)), // P(a) v Q(a)
					ArrowProof(Pa, // P(a)
						IExist(Pa, "a", "x"),                      // ∃x.( P(x)
						IOr(Exist("x", Px), Exist("x", Qx), true), // ∃x.( P(x) ) v ∃x.( Q(x) )
					), // P(a) -> ( ∃x.( P(x) ) v ∃x.( Q(x) ) )
					ArrowProof(Qa, // Q(a)
						IExist(Qa, "a", "x"),                       // ∃x.( Q(x)
						IOr(Exist("x", Px), Exist("x", Qx), false), // ∃x.( P(x) ) v ∃x.( Q(x) )
					), // Q(a) -> ( ∃x.( P(x) ) v ∃x.( Q(x) ) )
					EOr(Pa, Qa, Or(Exist("x", Px), Exist("x", Qx))), // ∃x.( P(x) ) v ∃x.( Q(x) )
				), // ∃x.( P(x) ) v ∃x.( Q(x) )
			),
			ArrowProof(Or(Exist("x", Px), Exist("x", Qx)),
				ArrowProof(Exist("x", Px),
					ExistElimProof("a",
						Exist("x", Px),               // P(a)
						IOr(Pa, Qa, true),            // P(a) v Q(a)
						IExist(Or(Pa, Qa), "a", "x"), // ∃x.( P(x) v Q(x) )
					), // ∃x.( P(x) v Q(x) )
				), // ∃x.( P(x) ) -> ∃x.( P(x) v Q(x) )
				ArrowProof(Exist("x", Qx),
					ExistElimProof("a",
						Exist("x", Qx),               // Q(a)
						IOr(Pa, Qa, false),           // P(a) v Q(a)
						IExist(Or(Pa, Qa), "a", "x"), // ∃x.( P(x) v Q(x) )
					), // ∃x.( P(x) v Q(x) )
				), // ∃x.( Q(x) ) -> ∃x.( P(x) v Q(x) )
				EOr(Exist("x", Px), Exist("x", Qx), Exist("x", Or(Px, Qx))), // ∃x.( P(x) v Q(x) )
			),
			IDArrow(
				Exist("x", Or(Px, Qx)),
				Or(Exist("x", Px), Exist("x", Qx)),
			),
		),
		RootProof("∀x.( P(x) -> Q(x) ) -> ( ∀x.( P(x) ) -> ∀x.( Q(x) ) )",
			ArrowProof(Forall("x", Arrow(Px, Qx)),
				ArrowProof(Forall("x", Px),
					ForallIntroProof("x", "a",
						&Reiterate{Formula: Forall("x", Arrow(Px, Qx))}, // ∀x.( Q(x) -> P(x) )
						&Reiterate{Formula: Forall("x", Px)},            // ∀x.( P(x) )
						EForall(Arrow(Px, Qx), "x", "a"),                // P(a) -> Q(a)
						EForall(Px, "x", "a"),                           // P(a)
						EImply(Pa, Qa),                                  // Q(a)
					),
				),
			),
		),
		RootProof("( ∃x.( P(x) ) -> ∃x.( Q(x) ) ) -> ( ∃x.( T ) -> ∃x.( P(x) -> Q(x) ) )",
			ArrowProof(Arrow(Exist("x", Px), Exist("x", Qx)),
				ArrowProof(Exist("x", T),
					ArrowProof(Exist("x", Px),
						&Reiterate{Formula: Arrow(Exist("x", Px), Exist("x", Qx))},
						EImply(Exist("x", Px), Exist("x", Qx)), // ∃x.( Q(x) )
						ExistElimProof("a",
							Exist("x", Qx), // Q(a)
							ArrowProof(Pa, // P(a)
								&Reiterate{Formula: Qa}, // Q(a)
							), // P(a) -> Q(a)
							IExist(Arrow(Pa, Qa), "a", "x"), // ∃x.( P(x) -> Q(x) )
						), // ∃x.( P(x) -> Q(x) )
					), // ∃x.( P(x) ) -> ∃x.( P(x) -> Q(x) )
					ArrowProof(Not(Exist("x", Px)),
						ContraProof(Not(Exist("x", Not(Px))),
							ExistContraProof("a",
								Exist("x", T), // T
								ContraProof(Pa,
									IExist(Pa, "a", "x"),                     // ∃x.( P(x) )
									&Reiterate{Formula: Not(Exist("x", Px))}, // ~ ∃x.( P(x) )
								), // ~ P(a)
								&Reiterate{Formula: Not(Exist("x", Not(Px)))}, // ~ ∃x.( ~ P(x) )
								IExist(Not(Pa), "a", "x"),                     // ∃x.( ~ P(x) )
							), // ~ ∃x.( T )
							&Reiterate{Formula: Exist("x", T)}, // ∃x.( T )
						), // ∃x.( ~ P(x) )
						ExistElimProof("a",
							Exist("x", Not(Px)),                  // ~ P(a)
							IOr(Not(Pa), Qa, true),               // ~ P(a) v Q(a)
							DisjunctionArrowTheorem(Not(Pa), Qa), // P(a) -> Q(a)
							IExist(Arrow(Pa, Qa), "a", "x"),      // ∃x.( P(x) -> Q(x) )
						), // ∃x.( P(x) -> Q(x) )
					), // ~ ∃x.( P(x) ) -> ∃x.( P(x) -> Q(x) )
					ExcludedMiddleTheorem(Exist("x", Px)), // ∃x.( P(x) ) v ~ ∃x.( P(x) )
					EOr(Exist("x", Px), Not(Exist("x", Px)), Exist("x", Arrow(Px, Qx))),
				),
			),
		),
	),
	NewProofsSection("commutativity",
		RootProof("∃x.( ∃y.( P(x,y) ) ) -> ∃y.( ∃x.( P(x,y) ) )",
			ArrowProof(Exist("x", Exist("y", Pxy)),
				ExistElimProof("a",
					Exist("x", Exist("y", Pxy)), // ∃y.( P(a,y) )
					ExistElimProof("b",
						Exist("y", Pay),                   // P(a,b)
						IExist(Pab, "a", "x"),             // ∃x.( P(x,b) )
						IExist(Exist("x", Pxb), "b", "y"), // ∃y.( ∃x.( P(x,y) ) )
					), // ∃y.( ∃x.( P(x,y) ) )
				), // ∃y.( ∃x.( P(x,y) ) )
			),
		),
		RootProof("∀x.( ∀y.( P(x,y) ) ) -> ∀y.( ∀x.( P(x,y) ) )",
			ArrowProof(Forall("x", Forall("y", Pxy)),
				ForallIntroProof("y", "b",
					ForallIntroProof("x", "a",
						&Reiterate{Formula: Forall("x", Forall("y", Pxy))}, // ∀x.( ∀y.( P(x,y) ) )
						EForall(Forall("y", Pxy), "x", "a"),                // ∀y.( P(a,y) )
						EForall(Pay, "y", "b"),                             // P(a,b)
					), // ∀x.( P(x,b) ) )
				), // ∀y .( ∀x.( P(x,y) ) )
			),
		),
		RootProof("∃x.( ∀y.( P(x,y) ) ) -> ∀y.( ∃x.( P(x,y) ) )",
			ArrowProof(Exist("x", Forall("y", Pxy)),
				ForallIntroProof("y", "a",
					&Reiterate{Formula: Exist("x", Forall("y", Pxy))}, // ∃x.( ∀y.( P(x,y) ) )
					ExistElimProof("b",
						Exist("x", Forall("y", Pxy)), // ∀y.( P(b,y) )
						EForall(Pby, "y", "a"),       // P(b,a)
						IExist(Pba, "b", "x"),        // ∃x.( P(x,a) )
					), // ∃x.( P(x,a) )
				), // ∀y.( ∃x.( P(x,y) ) )
			),
		),
		RootProof("∃x.( ∀y.( P(x,y) ) ) -> ∀y.( ∃x.( P(x,y) ) )",
			ArrowProof(Exist("x", Forall("y", Pxy)),
				ExistElimProof("b",
					Exist("x", Forall("y", Pxy)), // ∀y.( P(b,y) )
					ForallIntroProof("y", "a",
						&Reiterate{Formula: Forall("y", Pby)}, // ∀y.( P(b,y) )
						EForall(Pby, "y", "a"),                // P(b,a)
						IExist(Pba, "b", "x"),                 // ∃x.( P(x,a) )
					), // ∀y.( ∃x.( P(x,a) ) )
				), // ∀y.( ∃x.( P(x,y) ) )
			),
		),
	),
	NewProofsSection("random",
		RootProof("( P ^ ∃x.( Q(x) ) ) <-> ∃x.( P ^ Q(x) )",
			ArrowProof(And(P, Exist("x", Qx)),
				EAnd(P, Exist("x", Qx), true),
				EAnd(P, Exist("x", Qx), false),
				ExistElimProof("a",
					Exist("x", Qx),               // Q(a)
					&Reiterate{Formula: P},       // P
					IAnd(P, Qa),                  // P ^ Q(a)
					IExist(And(P, Qa), "a", "x"), // ∃x.( P ^ Q(x) )
				), // ∃x.( P ^ Q(x) )
			),
			ArrowProof(Exist("x", And(P, Qx)),
				ExistElimProof("a",
					Exist("x", And(P, Qx)),  // P ^ Q(a)
					EAnd(P, Qa, true),       // P
					EAnd(P, Qa, false),      // Q(a)
					IExist(Qa, "a", "x"),    // ∃x.( Q(x) )
					IAnd(P, Exist("x", Qx)), // P ^ ∃x.( Q(x) )
				), // P ^ ∃x.( Q(x) )
			),
			IDArrow(
				And(P, Exist("x", Qx)),
				Exist("x", And(P, Qx)),
			),
		),
		RootProof("( ( P v ∃x.( Q(x) ) ) ^ ∃x.( T ) ) <-> ∃x.( P v Q(x) )",
			ArrowProof(And(Or(P, Exist("x", Qx)), Exist("x", T)),
				EAnd(Or(P, Exist("x", Qx)), Exist("x", T), true),  // ∃x.T
				EAnd(Or(P, Exist("x", Qx)), Exist("x", T), false), // P v ∃x.( Q(x) )
				ArrowProof(P,
					&Reiterate{Formula: Exist("x", T)}, // ∃x.T
					ExistElimProof("a",
						Exist("x", T),               // T
						&Reiterate{Formula: P},      // P
						IOr(P, Qa, true),            // P v Q(a)
						IExist(Or(P, Qa), "a", "x"), // ∃x.( P v Q(x) )
					), // ∃x.( P v Q(x) )
				), // P -> // ∃x.( P v Q(x) )
				ArrowProof(Exist("x", Qx),
					ExistElimProof("a",
						Exist("x", Qx),              // Q(a)
						IOr(P, Qa, false),           // P v Q(a)
						IExist(Or(P, Qa), "a", "x"), // ∃x.( P v Q(x) )
					), // ∃x.( P v Q(x) )
				), // ∃x.( Q(x) ) -> ∃x.( P v Q(x) )
				EOr(P, Exist("x", Qx), Exist("x", Or(P, Qx))),
			),
			ArrowProof(Exist("x", Or(P, Qx)),
				ExistElimProof("a",
					Exist("x", Or(P, Qx)), // P v Q(a)
					ArrowProof(P,
						IOr(P, Exist("x", Qx), true), // P v ∃x.( Q(x) )
					), // P -> P v ∃x.( Q(x) )
					ArrowProof(Qa,
						IExist(Qa, "a", "x"),          // ∃x.( Q(x) )
						IOr(P, Exist("x", Qx), false), // P v ∃x.( Q(x) )
					), // Q(a) -> P v ∃x.( Q(x) )
					EOr(P, Qa, Or(P, Exist("x", Qx))),          // P v ∃x.( Q(x) )
					&Reiterate{Formula: T},                     // T
					IExist(T, "a", "x"),                        // ∃x.( T )
					IAnd(Or(P, Exist("x", Qx)), Exist("x", T)), // ( P v ∃x.( Q(x) ) ) ^ ∃x.( T )
				), // ( P v ∃x.( Q(x) ) ) ^ ∃x.( T )
			),
			IDArrow(
				And(Or(P, Exist("x", Qx)), Exist("x", T)),
				Exist("x", Or(P, Qx)),
			),
		),

		RootProof("( P v ∀x.( Q(x) ) ) <-> ∀x.( P v Q(x) )",
			ArrowProof(Or(P, Forall("x", Qx)),
				ArrowProof(P,
					ForallIntroProof("x", "a",
						&Reiterate{Formula: P}, // P
						IOr(P, Qa, true),       // P v Q(a)
					), // ∀x.( P v Q(x) )
				), // P -> ∀x.( P v Q(x) )
				ArrowProof(Forall("x", Qx),
					ForallIntroProof("x", "a",
						&Reiterate{Formula: Forall("x", Qx)}, // ∀x.( Q(x) )
						EForall(Qx, "x", "a"),                // Q(a)
						IOr(P, Qa, false),                    // P v Q(a)
					), // ∀x.( P v Q(x) )
				), // ∀x.( Q(x) ) -> ∀x.( P v Q(x) )
				EOr(P, Forall("x", Qx), Forall("x", Or(P, Qx))), // ∀x.( P v Q(x) )
			),
			ArrowProof(Forall("x", Or(P, Qx)),
				ArrowProof(P,
					IOr(P, Forall("x", Qx), true), // P v ∀x.( Q(x) )
				), // P -> ( P v ∀x.( Q(x) ) )
				ArrowProof(Not(P),
					ForallIntroProof("x", "a",
						&Reiterate{Formula: Forall("x", Or(P, Qx))}, // ∀x.( P v Q(x) )
						EForall(Or(P, Qx), "x", "a"),                // P v Q(a)
						DisjunctionArrowTheorem(P, Qa),              // ~ P -> Q(a)
						&Reiterate{Formula: Not(P)},                 // ~ P
						EImply(Not(P), Qa),                          // Q(a)
					), // ∀x.( Q(x) )
					IOr(P, Forall("x", Qx), false), // P v ∀x.( Q(x) )
				), // ~ P -> ( P v ∀x.( Q(x) ) )
				ExcludedMiddleTheorem(P),               // P v ~ P
				EOr(P, Not(P), Or(P, Forall("x", Qx))), // P v ∀x.( Q(x) )
			),
			IDArrow(
				Or(P, Forall("x", Qx)),
				Forall("x", Or(P, Qx)),
			),
		),
		RootProof("( ( P ^ ∀x.( Q(x) ) ) ^ ∃x.( T ) ) <-> ( ∀x.( P ^ Q(x) ) ^ ∃x.( T ) )",
			ArrowProof(
				And(And(P, Forall("x", Qx)), Exist("x", T)),
				EAnd(And(P, Forall("x", Qx)), Exist("x", T), true),  // P ^ ∀x.( Q(x) )
				EAnd(P, Forall("x", Qx), true),                      // P
				EAnd(P, Forall("x", Qx), false),                     // ∀x.( Q(x) )
				EAnd(And(P, Forall("x", Qx)), Exist("x", T), false), // ∃x.( T )
				ForallIntroProof("x", "a",
					&Reiterate{Formula: Forall("x", Qx)}, // ∀x.( Q(x) )
					EForall(Qx, "x", "a"),                // Q(a)
					&Reiterate{Formula: P},               // P
					IAnd(P, Qa),                          // P ^ Q(a)
				), // ∀x.( P ^ Q(x) )
				IAnd(Forall("x", And(P, Qx)), Exist("x", T)), // ∀x.( P ^ Q(x) ) ^ ∃x.( T )
			),
			ArrowProof(And(Forall("x", And(P, Qx)), Exist("x", T)),
				EAnd(Forall("x", And(P, Qx)), Exist("x", T), true),  // ∀x.( P ^ Q(x) )
				EAnd(Forall("x", And(P, Qx)), Exist("x", T), false), // ∃x.( T )
				ExistElimProof("a",
					Exist("x", T), // T
					&Reiterate{Formula: Forall("x", And(P, Qx))}, // ∀x.( P ^ Q(x) )
					EForall(And(P, Qx), "x", "a"),                // P ^ Q(a)
					EAnd(P, Qa, true),                            // P
				), // P
				ForallIntroProof("x", "a",
					&Reiterate{Formula: Forall("x", And(P, Qx))}, // ∀x.( P ^ Q(x) )
					EForall(And(P, Qx), "x", "a"),                // P ^ Q(a)
					EAnd(P, Qa, false),                           // Q(a)
				), // ∀x.( Q(x) )
				IAnd(P, Forall("x", Qx)),                     // P ^ ∀x.( Q(x) )
				IAnd(And(P, Forall("x", Qx)), Exist("x", T)), // ( P ^ ∀x.( Q(x) ) ) ^ ∃x.( T )
			),
			IDArrow(
				And(And(P, Forall("x", Qx)), Exist("x", T)),
				And(Forall("x", And(P, Qx)), Exist("x", T)),
			),
		),

		RootProof("P -> ( ∀x.( ~ P ) -> ~ ∃x.( T ) )",
			ArrowProof(P,
				ArrowProof(Forall("x", Not(P)),
					ExistContraProof("a",
						Exist("x", T),                            // T
						&Reiterate{Formula: Forall("x", Not(P))}, // ∀x.( ~ P )
						EForall(Not(P), "x", "a"),                // ~ P
						&Reiterate{Formula: P},                   // P
					), // ~ ∃x.( T )
				),
			),
		),
		RootProof("P -> ( ∃x.( T ) -> ~ ∀x.( ~ P ) )",
			ArrowProof(P,
				ArrowProof(Exist("x", T),
					ExistElimProof("a",
						Exist("x", T), // T
						ContraProof(Forall("x", Not(P)),
							EForall(Not(P), "x", "a"), // ~ P
							&Reiterate{Formula: P},    // P
						), // ~ ∀x.( ~ P )
					), // ~ ∀x.( ~ P )
				),
			),
		),
		RootProof("∃x.( T ) -> ( ∀x.( ~ P ) -> ~ P )",
			ArrowProof(Exist("x", T),
				ArrowProof(Forall("x", Not(P)),
					&Reiterate{Formula: Exist("x", T)}, // ∃x.( T )
					ExistElimProof("a",
						Exist("x", T),                            // T
						&Reiterate{Formula: Forall("x", Not(P))}, // ∀x.( ~ P )
						EForall(Not(P), "x", "a"),                // ~ P
					), // ~ P
				),
			),
		),

		RootProof("∃x.( R ) -> R",
			ArrowProof(Exist("x", R), // ∃x.( R )
				ExistElimProof("a",
					Exist("x", R), // R
				),
			),
		),
		RootProof("R -> ∀x.( R )",
			ArrowProof(R,
				ForallIntroProof("x", "a",
					&Reiterate{Formula: R},
				),
			),
		),
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
		RootProof("( ∀x.( Q(x) ) ^ ∃x.( Q(x) -> R ) ) -> R",
			ArrowProof(And(Forall("x", Qx), Exist("x", Arrow(Qx, R))),
				EAnd(Forall("x", Qx), Exist("x", Arrow(Qx, R)), true),  // ∀x.( Q(x) )
				EAnd(Forall("x", Qx), Exist("x", Arrow(Qx, R)), false), // ∃x.( Q(x) -> R )
				ExistElimProof("a",
					Exist("x", Arrow(Qx, R)),             // Q(a) -> R
					&Reiterate{Formula: Forall("x", Qx)}, // ∀x.( Q(x) )
					EForall(Qx, "x", "a"),                // Q(a)
					EImply(Qa, R),                        // R
				),
			),
		),
		RootProof("( ∀x.( Q(x) ) ^ ∃x.( T ) ) -> ∃x.( Q(x) )",
			ArrowProof(And(Forall("x", Qx), Exist("x", T)),
				EAnd(Forall("x", Qx), Exist("x", T), true),  // ∀x.( Q(x) )
				EAnd(Forall("x", Qx), Exist("x", T), false), // ∃x.( T )
				ExistElimProof("a",
					Exist("x", T),                        // T
					&Reiterate{Formula: Forall("x", Qx)}, // ∀x.( Q(x) )
					EForall(Qx, "x", "a"),                // Q(a)
					IExist(Qa, "a", "x"),                 // ∃x.( Q(x) )
				), // ∃x.( Q(x) )
			), // ( ∀x.( Q(x) ) ^ ∃x.( T ) ) -> ∃x.( Q(x) )
		),
	),

	// https://en.wikipedia.org/wiki/First-order_logic#Provable_identities
}
