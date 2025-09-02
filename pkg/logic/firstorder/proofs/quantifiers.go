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
	// TODO ( ∃x.( Q(x) ) v ∃x.( ~ Q(x) ) ) <-> ∃x.( T )
	// TODO ( ∀x.( Q(x) ) ^ ( ∀x.( ~ Q(x) ) <-> ~ ∃x.( Q(x) ) ^ ~ ∃x.( ~ Q(x) ) <-> ~ ∃x.( T )
	// TODO ( ∃x.( Q(x) ) ^ ( ∃x.( ~ Q(x) ) <-> ~ ∀x.( Q(x) ) ^ ~ ∀x.( ~ Q(x) )
	// TODO ~ ∃x.( T ) -> ~ ∃x.( P(x) )
	// TODO P ^ ∃x.( Q(x) ) <-> ∃x.( P ^ Q(x) )
	// TODO P v ∀x.( Q(x) ) <-> ∀x.( P v Q(x) )
}
