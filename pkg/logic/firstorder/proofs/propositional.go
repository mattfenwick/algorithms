package proofs

import . "github.com/mattfenwick/algorithms/pkg/logic/firstorder"

var propositionalProofSections = []*ProofsSection{
	NewProofsSection("theorems",
		RootProof("P v ~ P",
			ContraProof(
				Not(Or(P, Not(P))),
				ContraProof(
					P,
					IOr(P, Not(P), true),
					Reit(Not(Or(P, Not(P)))),
				),
				IOr(P, Not(P), false),
			),
		),
		RootProof("( P -> Q ) <-> ( ~ Q -> ~ P )",
			ArrowProof(
				Arrow(P, Q),
				ArrowProof(
					Not(Q),
					ContraProof(
						P,
						Reit(Arrow(P, Q)),
						EImply(P, Q),
						Reit(Not(Q)),
					),
				),
			),
			ArrowProof(
				Arrow(Not(Q), Not(P)),
				ArrowProof(
					P,
					ContraProof(
						Not(Q),
						Reit(Arrow(Not(Q), Not(P))),
						EImply(Not(Q), Not(P)),
						Reit(P),
					),
				),
			),
			IDArrow(Arrow(P, Q), Arrow(Not(Q), Not(P))),
		),
		RootProof("~ ( P -> Q ) <-> ( P ^ ~ Q )",
			ArrowProof(Not(Arrow(P, Q)),
				ContraProof(Not(P),
					ArrowProof(Not(Q),
						Reit(Not(P)),
					), // ~ Q -> ~ P
					ContrapositiveTheorem(Not(Q), Not(P)), // P -> Q
					Reit(Not(Arrow(P, Q))),
				), // P
				ContraProof(Q,
					Reit(P),
					IImply(P, Q),
					Reit(Not(Arrow(P, Q))),
				),
				IAnd(P, Not(Q)),
			),
			ArrowProof(And(P, Not(Q)),
				ContraProof(Arrow(P, Q),
					Reit(And(P, Not(Q))),
					EAnd(P, Not(Q), true),
					EAnd(P, Not(Q), false),
					EImply(P, Q),
				),
			),
			IDArrow(Not(Arrow(P, Q)), And(P, Not(Q))),
		),
		// RootProof("( P -> Q ) -> ( ~ P v Q )",
		// 	ArrowProof(Arrow(P, Q),
		// 		ArrowProof(P,
		// 			Reit( Arrow(P, Q)),
		// 			EImply(P, Q),
		// 			IOr(Not(P), Q, false),
		// 		),
		// 		ArrowProof(Not(P),
		// 			IOr(Not(P), Q, true),
		// 		),
		// 		// P v ~ P
		// 		EOr(P, Not(P), Or(Not(P), Q)),
		// 	),
		// ),
		// RootProof("( ~ P v Q ) -> ( P -> Q )"),
		RootProof("( P v Q ) <-> ( ~ P -> Q )",
			ArrowProof(Or(P, Q),
				ArrowProof(Q,
					ArrowProof(Not(P),
						Reit(Q),
					),
				), // Q -> ( ~ P -> Q )
				ArrowProof(P,
					ArrowProof(Not(Q),
						Reit(P),
					), // ~ Q -> P
					ContrapositiveTheorem(Not(Q), P), // ~ P -> Q
				), // P -> ( ~ P -> Q )
				EOr(P, Q, Arrow(Not(P), Q)),
			),
			ArrowProof(Arrow(Not(P), Q),
				ArrowProof(P,
					IOr(P, Q, true),
				),
				ArrowProof(Not(P),
					Reit(Arrow(Not(P), Q)),
					EImply(Not(P), Q),
					IOr(P, Q, false),
				),
				ExcludedMiddleTheorem(P),
				EOr(P, Not(P), Or(P, Q)),
			),
			IDArrow(Or(P, Q), Arrow(Not(P), Q)),
		),
		RootProof("~ ( P v Q ) <-> ( ~ P ^ ~ Q )",
			ArrowProof(Not(Or(P, Q)),
				ContraProof(P,
					IOr(P, Q, true),
					Reit(Not(Or(P, Q))),
				),
				ContraProof(Q,
					IOr(P, Q, false),
					Reit(Not(Or(P, Q))),
				),
				IAnd(Not(P), Not(Q)),
			),
			ArrowProof(And(Not(P), Not(Q)),
				ArrowProof(P,
					ContraProof(And(Not(P), Not(Q)),
						EAnd(Not(P), Not(Q), true),
						Reit(P),
					),
				),
				ArrowProof(Q,
					ContraProof(And(Not(P), Not(Q)),
						EAnd(Not(P), Not(Q), false),
						Reit(Q),
					),
				),
				ContraProof(Or(P, Q),
					Reit(Arrow(P, Not(And(Not(P), Not(Q))))),
					Reit(Arrow(Q, Not(And(Not(P), Not(Q))))),
					EOr(P, Q, Not(And(Not(P), Not(Q)))),
					Reit(And(Not(P), Not(Q))),
				),
			),
			IDArrow(Not(Or(P, Q)), And(Not(P), Not(Q))),
		),
		RootProof("~ ( P ^ Q ) <-> ( ~ P v ~ Q )",
			ArrowProof(Not(And(P, Q)),
				ContraProof(Not(Or(Not(P), Not(Q))),
					ContraProof(Not(P),
						IOr(Not(P), Not(Q), true),
						Reit(Not(Or(Not(P), Not(Q)))),
					),
					ContraProof(Not(Q),
						IOr(Not(P), Not(Q), false),
						Reit(Not(Or(Not(P), Not(Q)))),
					),
					IAnd(P, Q),
					Reit(Not(And(P, Q))),
				),
			),
			ArrowProof(Or(Not(P), Not(Q)),
				ArrowProof(Not(P),
					ContraProof(And(P, Q),
						EAnd(P, Q, true),
						Reit(Not(P)),
					),
				),
				ArrowProof(Not(Q),
					ContraProof(And(P, Q),
						EAnd(P, Q, false),
						Reit(Not(Q)),
					),
				),
				EOr(Not(P), Not(Q), Not(And(P, Q))),
			),
			IDArrow(Not(And(P, Q)), Or(Not(P), Not(Q))),
		),
		RootProof("( P <-> ~ Q ) <-> ~ ( P <-> Q )",
			ArrowProof(DArrow(P, Not(Q)),
				EDArrow(P, Not(Q), true),  // P -> ~ Q
				EDArrow(P, Not(Q), false), // ~ Q -> P
				ArrowProof(P,
					ContraProof(DArrow(P, Q),
						Reit(Arrow(P, Not(Q))), // P -> ~ Q
						EDArrow(P, Q, true),    // P -> Q
						Reit(P),                // P
						EImply(P, Q),           // Q
						EImply(P, Not(Q)),      // ~ Q
					), // ~ ( P <-> Q )
				), // P -> ~ ( P <-> Q )
				ArrowProof(Not(P),
					ContraProof(DArrow(P, Q),
						EDArrow(P, Q, false),        // Q -> P
						ContrapositiveTheorem(Q, P), // ~ P -> ~ Q
						Reit(Not(P)),                // ~ P
						EImply(Not(P), Not(Q)),      // ~ Q
						Reit(Arrow(Not(Q), P)),      // ~ Q -> P
						EImply(Not(Q), P),           // P
					), // ~ ( P <-> Q )
				), // ~ P -> ~ ( P <-> Q )
				ExcludedMiddleTheorem(P), // P v ~ P
				EOr(P, Not(P), Not(DArrow(P, Q))),
			),
			ArrowProof(Not(DArrow(P, Q)),
				ContraProof(Not(Arrow(P, Not(Q))),
					ArrowConjunctionTheorem(P, Not(Q), true),
					EAnd(P, Q, true),
					EAnd(P, Q, false),
					IImply(P, Q),
					IImply(Q, P),
					IDArrow(P, Q),
					Reit(Not(DArrow(P, Q))),
				),
				ContraProof(Not(Arrow(Not(Q), P)),
					ArrowConjunctionTheorem(Not(Q), P, true),
					EAnd(Not(Q), Not(P), true),
					EAnd(Not(Q), Not(P), false),
					IImply(Not(P), Not(Q)),
					ContrapositiveTheorem(Not(P), Not(Q)),
					IImply(Not(Q), Not(P)),
					ContrapositiveTheorem(Not(Q), Not(P)),
					IDArrow(P, Q),
					Reit(Not(DArrow(P, Q))),
				),
				IDArrow(P, Not(Q)),
			),
			IDArrow(DArrow(P, Not(Q)), Not(DArrow(P, Q))),
		),
	),
	NewProofsSection("reflexivity",
		RootProof(
			"P -> P",
			ArrowProof(P),
		),
		RootProof("( P ^ P ) <-> P",
			ArrowProof(And(P, P),
				EAnd(P, P, true),
			),
			ArrowProof(P,
				IAnd(P, P),
			),
			IDArrow(And(P, P), P),
		),
		RootProof("( P v P ) <-> P",
			ArrowProof(Or(P, P),
				ArrowProof(P),
				EOr(P, P, P),
			),
			ArrowProof(P,
				IOr(P, P, true),
			),
			IDArrow(Or(P, P), P),
		),
		RootProof("P <-> P",
			ArrowProof(P),
			IDArrow(P, P),
		),
	),
	NewProofsSection("basics",
		RootProof("~ ( P ^ ~ P )",
			ContraProof(
				And(P, Not(P)),
				EAnd(P, Not(P), true),
				EAnd(P, Not(P), false),
			)),
		RootProof("( P ^ Q ) -> ( P v Q )",
			ArrowProof(And(P, Q),
				EAnd(P, Q, true),
				IOr(P, Q, true),
			),
		),
		RootProof("( P -> ~ P ) -> ~ P",
			ArrowProof(Arrow(P, Not(P)),
				ContraProof(P,
					Reit(Arrow(P, Not(P))),
					EImply(P, Not(P)),
				),
			),
		),
		RootProof("( ( P -> Q ) -> P ) -> P",
			ArrowProof(Arrow(Arrow(P, Q), P),
				ArrowProof(Not(Arrow(P, Q)),
					ArrowConjunctionTheorem(P, Q, true), // P ^ ~ Q
					EAnd(P, Not(Q), true),               // P
				), // ( ~ ( P -> Q ) ) -> P
				ExcludedMiddleTheorem(Arrow(P, Q)),    // ( P -> Q ) v ~ ( P -> Q )
				EOr(Arrow(P, Q), Not(Arrow(P, Q)), P), // p
			),
		),
		RootProof("( P ^ ~ P ) -> Q",
			NonContradictionTheorem(P), // ~ ( P ^ ~ P )
			ArrowProof(Not(Q),
				Reit(Not(And(P, Not(P)))),
			), // ~ Q -> ~ ( P ^ ~ P )
			ContrapositiveTheorem(Not(Q), Not(And(P, Not(P)))), // ( P ^ ~ P ) -> Q
		),
	),
	NewProofsSection("arrows",
		RootProof("Q -> ( P -> Q )",
			ArrowProof(
				Q,
				ArrowProof(P,
					Reit(Q))),
		),
		RootProof("~ P -> ( P -> Q )",
			ArrowProof(Not(P),
				ArrowProof(Not(Q),
					Reit(Not(P)),
				),
				ContrapositiveTheorem(Not(Q), Not(P)),
			),
		),
		RootProof("( ( P ^ Q ) -> R ) <-> ( P -> ( Q -> R ) )",
			ArrowProof(Arrow(And(P, Q), R),
				ArrowProof(P,
					ArrowProof(Q,
						Reit(P),
						IAnd(P, Q),
						Reit(Arrow(And(P, Q), R)),
						EImply(And(P, Q), R),
					),
				),
			),
			ArrowProof(Arrow(P, Arrow(Q, R)),
				ArrowProof(And(P, Q),
					EAnd(P, Q, true),
					EAnd(P, Q, false),
					Reit(Arrow(P, Arrow(Q, R))),
					EImply(P, Arrow(Q, R)),
					EImply(Q, R),
				),
			),
			IDArrow(Arrow(And(P, Q), R), Arrow(P, Arrow(Q, R))),
		),
		RootProof("( ( P -> Q ) ^ ( P -> ~ Q ) ) -> ~ P",
			ArrowProof(And(Arrow(P, Q), Arrow(P, Not(Q))),
				ContraProof(P,
					Reit(And(Arrow(P, Q), Arrow(P, Not(Q)))),
					EAnd(Arrow(P, Q), Arrow(P, Not(Q)), true),
					EAnd(Arrow(P, Q), Arrow(P, Not(Q)), false),
					EImply(P, Q),
					EImply(P, Not(Q)),
				),
			),
		),
		RootProof("( ( P -> Q ) ^ ( ~ P -> Q ) ) -> Q",
			ArrowProof(And(Arrow(P, Q), Arrow(Not(P), Q)),
				EAnd(Arrow(P, Q), Arrow(Not(P), Q), true),
				EAnd(Arrow(P, Q), Arrow(Not(P), Q), false),
				ExcludedMiddleTheorem(P),
				EOr(P, Not(P), Q),
			),
		),
		RootProof("( P -> ( Q -> R ) ) -> ( Q -> ( P -> R ) )",
			ArrowProof(Arrow(P, Arrow(Q, R)),
				ArrowProof(Q,
					ArrowProof(P,
						Reit(Arrow(P, Arrow(Q, R))),
						EImply(P, Arrow(Q, R)),
						Reit(Q),
						EImply(Q, R),
					),
				),
			),
		),
	),
	NewProofsSection("transitivity",
		RootProof("( P -> Q ) -> ( ( Q -> R ) -> ( P -> R ) )",
			ArrowProof(Arrow(P, Q),
				ArrowProof(Arrow(Q, R),
					ArrowProof(P,
						Reit(Arrow(P, Q)),
						Reit(Arrow(Q, R)),
						EImply(P, Q),
						EImply(Q, R),
					),
				),
			),
		),
		RootProof("( P <-> Q ) -> ( ( Q <-> R ) -> ( P <-> R ) )",
			ArrowProof(DArrow(P, Q),
				ArrowProof(DArrow(Q, R),
					Reit(DArrow(P, Q)),
					EDArrow(P, Q, true),
					EDArrow(P, Q, false),
					EDArrow(Q, R, true),
					EDArrow(Q, R, false),
					ArrowProof(P,
						Reit(Arrow(P, Q)),
						Reit(Arrow(Q, R)),
						EImply(P, Q),
						EImply(Q, R),
					),
					ArrowProof(R,
						Reit(Arrow(R, Q)),
						Reit(Arrow(Q, P)),
						EImply(R, Q),
						EImply(Q, P),
					),
					IDArrow(P, R),
				),
			),
		),
	),
	NewProofsSection("commutativity",
		RootProof("( P ^ Q ) -> ( Q ^ P )",
			ArrowProof(
				And(P, Q),
				EAnd(P, Q, true),
				EAnd(P, Q, false),
				IAnd(Q, P),
			),
		),
		RootProof("( P v Q ) -> ( Q v P )",
			ArrowProof(
				Or(P, Q),
				ArrowProof(P,
					IOr(Q, P, false)),
				ArrowProof(Q,
					IOr(Q, P, true)),
				EOr(P, Q, Or(Q, P)),
			),
		),
		RootProof("( P <-> Q ) -> ( Q <-> P )",
			ArrowProof(DArrow(P, Q),
				EDArrow(P, Q, true),
				EDArrow(P, Q, false),
				IDArrow(Q, P),
			),
		),
	),
	NewProofsSection("associativity",
		RootProof("( ( P ^ Q ) ^ R ) <-> ( P ^ ( Q ^ R ) )",
			ArrowProof(And(And(P, Q), R),
				EAnd(And(P, Q), R, true),
				EAnd(And(P, Q), R, false),
				EAnd(P, Q, true),
				EAnd(P, Q, false),
				IAnd(Q, R),
				IAnd(P, And(Q, R)),
			),
			ArrowProof(And(P, And(Q, R)),
				EAnd(P, And(Q, R), true),
				EAnd(P, And(Q, R), false),
				EAnd(Q, R, true),
				EAnd(Q, R, false),
				IAnd(P, Q),
				IAnd(And(P, Q), R),
			),
			IDArrow(And(And(P, Q), R), And(P, And(Q, R))),
		),
		RootProof("( ( P v Q ) v R ) <-> ( P v ( Q v R ) )",
			ArrowProof(Or(Or(P, Q), R),
				ArrowProof(Or(P, Q),
					ArrowProof(P,
						IOr(P, Or(Q, R), true),
					),
					ArrowProof(Q,
						IOr(Q, R, true),
						IOr(P, Or(Q, R), false),
					),
					EOr(P, Q, Or(P, Or(Q, R))),
				),
				ArrowProof(R,
					IOr(Q, R, false),
					IOr(P, Or(Q, R), false),
				),
				EOr(Or(P, Q), R, Or(P, Or(Q, R))),
			),
			ArrowProof(Or(P, Or(Q, R)),
				ArrowProof(P,
					IOr(P, Q, true),
					IOr(Or(P, Q), R, true),
				),
				ArrowProof(Or(Q, R),
					ArrowProof(Q,
						IOr(P, Q, false),
						IOr(Or(P, Q), R, true),
					),
					ArrowProof(R,
						IOr(Or(P, Q), R, false),
					),
					EOr(Q, R, Or(Or(P, Q), R)),
				),
				EOr(P, Or(Q, R), Or(Or(P, Q), R)),
			),
			IDArrow(Or(Or(P, Q), R), Or(P, Or(Q, R))),
		),
		RootProof("( ( P <-> Q ) <-> R ) <-> ( P <-> ( Q <-> R ) )",
			ArrowProof(DArrow(DArrow(P, Q), R),
				EDArrow(DArrow(P, Q), R, true),  // ( P <-> Q ) -> R
				EDArrow(DArrow(P, Q), R, false), // R -> ( P <-> Q )
				ContraProof(Not(Arrow(P, DArrow(Q, R))),
					ArrowConjunctionTheorem(P, DArrow(Q, R), true), // P ^ ~ ( Q <-> R )
					EAnd(P, Not(DArrow(Q, R)), true),               // P
					EAnd(P, Not(DArrow(Q, R)), false),              // ~ ( Q <-> R )
					BiconditionalNegationTheorem(Q, R),             // Q <-> ~ R
					ContraProof(R,
						Reit(Arrow(R, DArrow(P, Q))), // R -> ( P <-> Q )
						EImply(R, DArrow(P, Q)),      // P <-> Q
						EDArrow(P, Q, true),          // P -> Q
						Reit(P),                      // P
						EImply(P, Q),                 // Q
						Reit(DArrow(Q, Not(R))),      // Q <-> ~ R
						EDArrow(Q, Not(R), true),     // Q -> ~ R
						EImply(Q, Not(R)),            // ~ R
					), // ~ R
					EDArrow(Q, Not(R), false),              // ~ R -> Q
					EImply(Not(R), Q),                      // Q
					Reit(Arrow(DArrow(P, Q), R)),           // ( P <-> Q ) -> R
					ContrapositiveTheorem(DArrow(P, Q), R), // ~ R -> ~ ( P <-> Q )
					EImply(Not(R), Not(DArrow(P, Q))),      // ~ ( P <-> Q )
					BiconditionalNegationTheorem(P, Q),     // P <-> ~ Q
					EDArrow(P, Not(Q), true),               // P -> ~ Q
					EImply(P, Not(Q)),                      // ~ Q
				), // ( P -> ( Q <-> R ) )
				ContraProof(Not(Arrow(DArrow(Q, R), P)),
					ArrowConjunctionTheorem(DArrow(Q, R), P, true), // ( Q <-> R ) ^ ~ P
					EAnd(DArrow(Q, R), Not(P), true),               // Q <-> R
					EAnd(DArrow(Q, R), Not(P), false),              // ~ P
					ContraProof(R,
						Reit(Arrow(R, DArrow(P, Q))), // R -> ( P <-> Q )
						Reit(DArrow(Q, R)),           // Q <-> R
						EDArrow(Q, R, false),         // R -> Q
						EImply(R, Q),                 // Q
						EImply(R, DArrow(P, Q)),      // P <-> Q
						EDArrow(P, Q, false),         // Q -> P
						EImply(Q, P),                 // P
						Reit(Not(P)),                 // ~ P
					), // ~ R
					EDArrow(Q, R, true),                    // Q -> R
					ContrapositiveTheorem(Q, R),            // ~ R -> ~ Q
					EImply(Not(R), Not(Q)),                 // ~ Q
					Reit(Arrow(DArrow(P, Q), R)),           // ( P <-> Q ) -> R
					ContrapositiveTheorem(DArrow(P, Q), R), // ~ R -> ~ ( P <-> Q)
					EImply(Not(R), Not(DArrow(P, Q))),      // ~ ( P <-> Q )
					BiconditionalNegationTheorem(P, Q),     // P <-> ~ Q
					EDArrow(P, Not(Q), false),              // ~ Q -> P
					EImply(Not(Q), P),                      // P
				), // ( ( Q <-> R ) -> P )
				IDArrow(P, DArrow(Q, R)), // P <-> ( Q <-> R )
			),
			ArrowProof(DArrow(P, DArrow(Q, R)),
				EDArrow(P, DArrow(Q, R), true),  // P -> ( Q <-> R )
				EDArrow(P, DArrow(Q, R), false), // ( Q <-> R ) -> P
				// prove: ( P <-> Q ) -> R
				ContraProof(Not(Arrow(DArrow(P, Q), R)),
					ArrowConjunctionTheorem(DArrow(P, Q), R, true), // ( P <-> Q ) ^ ~ R
					EAnd(DArrow(P, Q), Not(R), true),               // P <-> Q
					EAnd(DArrow(P, Q), Not(R), false),              // ~ R
					EDArrow(P, Q, true),                            // P -> Q
					ContraProof(P,
						Reit(Arrow(P, Q)),            // P -> Q
						EImply(P, Q),                 // Q
						Reit(Arrow(P, DArrow(Q, R))), // P -> ( Q <-> R )
						EImply(P, DArrow(Q, R)),      // Q <-> R
						EDArrow(Q, R, true),          // Q -> R
						EImply(Q, R),                 // R
						Reit(Not(R)),                 // ~ R
					), // ~ P
					Reit(Arrow(DArrow(Q, R), P)),           // ( Q <-> R ) -> P
					ContrapositiveTheorem(DArrow(Q, R), P), // ~ P -> ~ ( Q <-> R )
					EImply(Not(P), Not(DArrow(Q, R))),      // ~ ( Q <-> R )
					BiconditionalNegationTheorem(Q, R),     // Q <-> ~ R
					EDArrow(Q, Not(R), false),              // ~ R -> Q
					EImply(Not(R), Q),                      // Q
					EDArrow(P, Q, false),                   // Q -> P
					EImply(Q, P),                           // P
				), // ( ( P <-> Q ) -> R )
				ContraProof(Not(Arrow(R, DArrow(P, Q))),
					ArrowConjunctionTheorem(R, DArrow(P, Q), true), // R ^ ~ ( P <-> Q )
					EAnd(R, Not(DArrow(P, Q)), true),               // R
					EAnd(R, Not(DArrow(P, Q)), false),              // ~ ( P <-> Q )
					BiconditionalNegationTheorem(P, Q),             // P <-> ~ Q
					EDArrow(P, Not(Q), true),                       // P -> ~ Q
					ContraProof(P,
						Reit(Arrow(P, Not(Q))),       // P -> ~ Q
						EImply(P, Not(Q)),            // ~ Q
						Reit(Arrow(P, DArrow(Q, R))), // P -> ( Q <-> R )
						EImply(P, DArrow(Q, R)),      // Q <-> R
						EDArrow(Q, R, false),         // R -> Q
						Reit(R),                      // R
						EImply(R, Q),                 // Q
					), // ~ P
					EDArrow(P, Not(Q), false),              // ~ Q -> P
					ContrapositiveTheorem(Not(Q), P),       // ~ P -> Q
					EImply(Not(P), Q),                      // Q
					Reit(Arrow(DArrow(Q, R), P)),           // ( Q <-> R ) -> P
					ContrapositiveTheorem(DArrow(Q, R), P), // ~ P -> ~ ( Q <-> R )
					EImply(Not(P), Not(DArrow(Q, R))),      // ~ ( Q <-> R )
					BiconditionalNegationTheorem(Q, R),     // Q <-> ~ R
					EDArrow(Q, Not(R), true),               // Q -> ~ R
					EImply(Q, Not(R)),                      // ~ R
				), // ( R -> ( P <-> Q ) )
				IDArrow(DArrow(P, Q), R),
			),
			IDArrow(DArrow(DArrow(P, Q), R), DArrow(P, DArrow(Q, R))),
		),
	),
	NewProofsSection("distributivity",
		RootProof("( P -> ( Q -> R ) ) <-> ( ( P -> Q ) -> ( P -> R ) )",
			ArrowProof(Arrow(P, Arrow(Q, R)),
				ArrowProof(Arrow(P, Q),
					ArrowProof(P,
						Reit(Arrow(P, Arrow(Q, R))),
						EImply(P, Arrow(Q, R)),
						Reit(Arrow(P, Q)),
						EImply(P, Q),
						EImply(Q, R),
					),
				),
			),
			ArrowProof(Arrow(Arrow(P, Q), Arrow(P, R)),
				ArrowProof(P,
					ArrowProof(Q,
						Reit(Arrow(Arrow(P, Q), Arrow(P, R))),
						Reit(P),
						IImply(P, Q),
						EImply(Arrow(P, Q), Arrow(P, R)),
						EImply(P, R),
					),
				),
			),
			IDArrow(
				Arrow(P, Arrow(Q, R)),
				Arrow(Arrow(P, Q), Arrow(P, R)),
			),
		),
		RootProof("( P -> ( Q v R ) ) <-> ( ( P -> Q ) v ( P -> R ) )",
			ArrowProof(Arrow(P, Or(Q, R)),
				ArrowProof(P,
					Reit(Arrow(P, Or(Q, R))),
					EImply(P, Or(Q, R)),
					ArrowProof(Q,
						Reit(P),
						IImply(P, Q),
						IOr(Arrow(P, Q), Arrow(P, R), true),
					),
					ArrowProof(R,
						Reit(P),
						IImply(P, R),
						IOr(Arrow(P, Q), Arrow(P, R), false),
					),
					EOr(Q, R, Or(Arrow(P, Q), Arrow(P, R))),
				),
				ArrowProof(Not(P),
					ArrowProof(Not(Q),
						Reit(Not(P)),
					),
					ContrapositiveTheorem(Not(Q), Not(P)),
					IOr(Arrow(P, Q), Arrow(P, R), true),
				),
				ExcludedMiddleTheorem(P),
				EOr(P, Not(P), Or(Arrow(P, Q), Arrow(P, R))),
			),
			ArrowProof(Or(Arrow(P, Q), Arrow(P, R)),
				ArrowProof(P,
					ArrowProof(Arrow(P, Q),
						Reit(P),
						EImply(P, Q),
						IOr(Q, R, true),
					),
					ArrowProof(Arrow(P, R),
						Reit(P),
						EImply(P, R),
						IOr(Q, R, false),
					),
					Reit(Or(Arrow(P, Q), Arrow(P, R))),
					EOr(Arrow(P, Q), Arrow(P, R), Or(Q, R)),
				),
			),
			IDArrow(
				Arrow(P, Or(Q, R)),
				Or(Arrow(P, Q), Arrow(P, R)),
			),
		),
		RootProof("( P -> ( Q ^ R ) ) <-> ( ( P -> Q ) ^ ( P -> R ) )",
			ArrowProof(Arrow(P, And(Q, R)),
				ArrowProof(P,
					Reit(Arrow(P, And(Q, R))),
					EImply(P, And(Q, R)),
					EAnd(Q, R, true),
				),
				ArrowProof(P,
					Reit(Arrow(P, And(Q, R))),
					EImply(P, And(Q, R)),
					EAnd(Q, R, false),
				),
				IAnd(Arrow(P, Q), Arrow(P, R)),
			),
			ArrowProof(And(Arrow(P, Q), Arrow(P, R)),
				ArrowProof(P,
					Reit(And(Arrow(P, Q), Arrow(P, R))),
					EAnd(Arrow(P, Q), Arrow(P, R), true),
					EAnd(Arrow(P, Q), Arrow(P, R), false),
					EImply(P, Q),
					EImply(P, R),
					IAnd(Q, R),
				),
			),
			IDArrow(
				Arrow(P, And(Q, R)),
				And(Arrow(P, Q), Arrow(P, R)),
			),
		),
		RootProof("( P ^ ( Q v R ) ) <-> ( ( P ^ Q ) v ( P ^ R ) )",
			ArrowProof(And(P, Or(Q, R)),
				EAnd(P, Or(Q, R), true),
				EAnd(P, Or(Q, R), false),
				ArrowProof(Q,
					Reit(P),
					IAnd(P, Q),
					IOr(And(P, Q), And(P, R), true),
				),
				ArrowProof(R,
					Reit(P),
					IAnd(P, R),
					IOr(And(P, Q), And(P, R), false),
				),
				EOr(Q, R, Or(And(P, Q), And(P, R))),
			),
			ArrowProof(Or(And(P, Q), And(P, R)),
				ArrowProof(And(P, Q),
					EAnd(P, Q, true),
					EAnd(P, Q, false),
					IOr(Q, R, true),
					IAnd(P, Or(Q, R)),
				),
				ArrowProof(And(P, R),
					EAnd(P, R, true),
					EAnd(P, R, false),
					IOr(Q, R, false),
					IAnd(P, Or(Q, R)),
				),
				EOr(And(P, Q), And(P, R), And(P, Or(Q, R))),
			),
			IDArrow(
				And(P, Or(Q, R)),
				Or(And(P, Q), And(P, R)),
			),
		),
		RootProof("( P v ( Q ^ R ) ) <-> ( ( P v Q ) ^ ( P v R ) )",
			ArrowProof(Or(P, And(Q, R)),
				ArrowProof(P,
					IOr(P, Q, true),
					IOr(P, R, true),
					IAnd(Or(P, Q), Or(P, R)),
				),
				ArrowProof(And(Q, R),
					EAnd(Q, R, true),
					EAnd(Q, R, false),
					IOr(P, Q, false),
					IOr(P, R, false),
					IAnd(Or(P, Q), Or(P, R)),
				),
				EOr(P, And(Q, R), And(Or(P, Q), Or(P, R))),
			),
			ArrowProof(And(Or(P, Q), Or(P, R)),
				EAnd(Or(P, Q), Or(P, R), true),
				EAnd(Or(P, Q), Or(P, R), false),
				ArrowProof(P,
					IOr(P, And(Q, R), true),
				),
				ArrowProof(P,
					ArrowProof(Not(Q),
						Reit(P),
					),
					ContrapositiveTheorem(Not(Q), P),
				), // P -> ( ~ P -> Q )
				ArrowProof(P,
					ArrowProof(Not(R),
						Reit(P),
					),
					ContrapositiveTheorem(Not(R), P),
				), // P -> ( ~ P -> R )
				ArrowProof(Not(P),
					Reit(Or(P, Q)),
					Reit(Or(P, R)),
					ArrowProof(Q,
						Reit(Not(P)),
						IImply(Not(P), Q),
					), // Q -> ( ~ P -> Q )
					Reit(Arrow(P, Arrow(Not(P), Q))),
					EOr(P, Q, Arrow(Not(P), Q)),
					EImply(Not(P), Q),
					ArrowProof(R,
						Reit(Not(P)),
						IImply(Not(P), R),
					), // R -> ( ~ P -> R )
					Reit(Arrow(P, Arrow(Not(P), R))),
					EOr(P, R, Arrow(Not(P), R)),
					EImply(Not(P), R),
					IAnd(Q, R),
					IOr(P, And(Q, R), false),
				),
				ExcludedMiddleTheorem(P),
				EOr(P, Not(P), Or(P, And(Q, R))),
			),
			IDArrow(
				Or(P, And(Q, R)),
				And(Or(P, Q), Or(P, R)),
			),
		),
	),
	NewProofsSection("anti-distributivity",
		RootProof("( ( P v Q ) -> R ) <-> ( ( P -> R ) ^ ( Q -> R ) )",
			ArrowProof(Arrow(Or(P, Q), R),
				ArrowProof(P,
					IOr(P, Q, true),
					Reit(Arrow(Or(P, Q), R)),
					EImply(Or(P, Q), R),
				),
				ArrowProof(Q,
					IOr(P, Q, false),
					Reit(Arrow(Or(P, Q), R)),
					EImply(Or(P, Q), R),
				),
				IAnd(Arrow(P, R), Arrow(Q, R)),
			),
			ArrowProof(And(Arrow(P, R), Arrow(Q, R)),
				ArrowProof(Or(P, Q),
					Reit(And(Arrow(P, R), Arrow(Q, R))),
					EAnd(Arrow(P, R), Arrow(Q, R), true),
					EAnd(Arrow(P, R), Arrow(Q, R), false),
					EOr(P, Q, R),
				),
			),
			IDArrow(Arrow(Or(P, Q), R), And(Arrow(P, R), Arrow(Q, R))),
		),
		RootProof("( ( P ^ Q ) -> R ) <-> ( ( P -> R ) v ( Q -> R ) )",
			ArrowProof(Or(Arrow(P, R), Arrow(Q, R)),
				ArrowProof(Arrow(P, R),
					ArrowProof(And(P, Q),
						Reit(Arrow(P, R)),
						EAnd(P, Q, true),
						EImply(P, R),
					),
				),
				ArrowProof(Arrow(Q, R),
					ArrowProof(And(P, Q),
						Reit(Arrow(Q, R)),
						EAnd(P, Q, false),
						EImply(Q, R),
					),
				),
				EOr(Arrow(P, R), Arrow(Q, R), Arrow(And(P, Q), R)),
			),
			ArrowProof(Arrow(And(P, Q), R),
				ContraProof(Not(Or(Arrow(P, R), Arrow(Q, R))),
					DeMorgansOrToAndTheorem(
						Arrow(P, R),
						Arrow(Q, R), true), // ~ ( P -> R ) ^ ~ (Q -> R )
					EAnd(Not(Arrow(P, R)), Not(Arrow(Q, R)), true),  // ~ ( P -> R )
					EAnd(Not(Arrow(P, R)), Not(Arrow(Q, R)), false), // ~ ( Q -> R )
					ArrowConjunctionTheorem(P, R, true),             // P ^ ~ R
					EAnd(P, Not(R), true),                           // P
					EAnd(P, Not(R), false),                          // ~ R
					ArrowConjunctionTheorem(Q, R, true),             // Q ^ ~ R
					EAnd(Q, Not(R), true),                           // Q
					Reit(Arrow(And(P, Q), R)),                       // ( P ^ Q ) -> R
					IAnd(P, Q),                                      // P ^ Q
					EImply(And(P, Q), R),                            // R
				), // ( P -> R ) v ( Q -> R )
			),
			IDArrow(Arrow(And(P, Q), R), Or(Arrow(P, R), Arrow(Q, R))),
		),
	),
	NewProofsSection("double-distribution",
		// TODO https://en.wikipedia.org/wiki/Distributive_property#Truth_functional_connectives
		RootProof("( ( P v Q ) ^ ( R v S ) ) <-> ( ( ( P ^ R ) v ( P ^ S ) ) v ( ( Q ^ R ) v ( Q ^ S ) ) )",
			ArrowProof(And(Or(P, Q), Or(R, S)),
				EAnd(Or(P, Q), Or(R, S), true),
				EAnd(Or(P, Q), Or(R, S), false),
				ArrowProof(P,
					ArrowProof(R,
						Reit(P),
						IAnd(P, R),
						IOr(And(P, R), And(P, S), true),
						IOr(Or(And(P, R), And(P, S)), Or(And(Q, R), And(Q, S)), true),
					),
					ArrowProof(S,
						Reit(P),
						IAnd(P, S),
						IOr(And(P, R), And(P, S), false),
						IOr(Or(And(P, R), And(P, S)), Or(And(Q, R), And(Q, S)), true),
					),
					Reit(Or(R, S)),
					EOr(R, S, Or(Or(And(P, R), And(P, S)), Or(And(Q, R), And(Q, S)))),
				),
				ArrowProof(Q,
					ArrowProof(R,
						Reit(Q),
						IAnd(Q, R),
						IOr(And(Q, R), And(Q, S), true),
						IOr(Or(And(P, R), And(P, S)), Or(And(Q, R), And(Q, S)), false),
					),
					ArrowProof(S,
						Reit(Q),
						IAnd(Q, S),
						IOr(And(Q, R), And(Q, S), false),
						IOr(Or(And(P, R), And(P, S)), Or(And(Q, R), And(Q, S)), false),
					),
					Reit(Or(R, S)),
					EOr(R, S, Or(Or(And(P, R), And(P, S)), Or(And(Q, R), And(Q, S)))),
				),
				EOr(P, Q, Or(Or(And(P, R), And(P, S)), Or(And(Q, R), And(Q, S)))),
			),
			ArrowProof(Or(Or(And(P, R), And(P, S)), Or(And(Q, R), And(Q, S))),
				ArrowProof(Or(And(P, R), And(P, S)),
					ArrowProof(And(P, R),
						EAnd(P, R, true),
						EAnd(P, R, false),
						IOr(P, Q, true),
						IOr(R, S, true),
						IAnd(Or(P, Q), Or(R, S)),
					),
					ArrowProof(And(P, S),
						EAnd(P, S, true),
						EAnd(P, S, false),
						IOr(P, Q, true),
						IOr(R, S, false),
						IAnd(Or(P, Q), Or(R, S)),
					),
					EOr(And(P, R), And(P, S), And(Or(P, Q), Or(R, S))),
				),
				ArrowProof(Or(And(Q, R), And(Q, S)),
					ArrowProof(And(Q, R),
						EAnd(Q, R, true),
						EAnd(Q, R, false),
						IOr(P, Q, false),
						IOr(R, S, true),
						IAnd(Or(P, Q), Or(R, S)),
					),
					ArrowProof(And(Q, S),
						EAnd(Q, S, true),
						EAnd(Q, S, false),
						IOr(P, Q, false),
						IOr(R, S, false),
						IAnd(Or(P, Q), Or(R, S)),
					),
					EOr(And(Q, R), And(Q, S), And(Or(P, Q), Or(R, S))),
				),
				EOr(
					Or(And(P, R), And(P, S)),
					Or(And(Q, R), And(Q, S)),
					And(Or(P, Q), Or(R, S)),
				),
			),
			IDArrow(
				And(Or(P, Q), Or(R, S)),
				Or(Or(And(P, R), And(P, S)), Or(And(Q, R), And(Q, S))),
			),
		),
		RootProof("( ( P ^ Q ) v ( R ^ S ) ) <-> ( ( ( P v R ) ^ ( P v S ) ) ^ ( ( Q v R ) ^ ( Q v S ) ) )",
			ArrowProof(Or(And(P, Q), And(R, S)),
				ArrowProof(And(P, Q),
					EAnd(P, Q, true),
					EAnd(P, Q, false),
					IOr(P, R, true),
					IOr(P, S, true),
					IOr(Q, R, true),
					IOr(Q, S, true),
					IAnd(Or(P, R), Or(P, S)),
					IAnd(Or(Q, R), Or(Q, S)),
					IAnd(And(Or(P, R), Or(P, S)), And(Or(Q, R), Or(Q, S))),
				),
				ArrowProof(And(R, S),
					EAnd(R, S, true),
					EAnd(R, S, false),
					IOr(P, R, false),
					IOr(P, S, false),
					IOr(Q, R, false),
					IOr(Q, S, false),
					IAnd(Or(P, R), Or(P, S)),
					IAnd(Or(Q, R), Or(Q, S)),
					IAnd(And(Or(P, R), Or(P, S)), And(Or(Q, R), Or(Q, S))),
				),
				EOr(And(P, Q), And(R, S), And(And(Or(P, R), Or(P, S)), And(Or(Q, R), Or(Q, S)))),
			),
			ArrowProof(And(And(Or(P, R), Or(P, S)), And(Or(Q, R), Or(Q, S))),
				EAnd(And(Or(P, R), Or(P, S)), And(Or(Q, R), Or(Q, S)), true),
				EAnd(Or(P, R), Or(P, S), true),
				EAnd(Or(P, R), Or(P, S), false),
				EAnd(And(Or(P, R), Or(P, S)), And(Or(Q, R), Or(Q, S)), false),
				EAnd(Or(Q, R), Or(Q, S), true),
				EAnd(Or(Q, R), Or(Q, S), false),
				ArrowProof(And(P, Q),
					IOr(And(P, Q), And(R, S), true),
				), // ( P ^ Q ) -> ( ( P ^ Q ) v ( R ^ S ) )
				ArrowProof(Not(And(P, Q)), // ~ ( P ^ Q )
					DeMorgansAndToOrTheorem(P, Q, true), // ~ P v ~ Q
					ArrowProof(Not(P),
						Reit(Or(P, R)),
						DisjunctionArrowTheorem(P, R),
						Reit(Or(P, S)),
						DisjunctionArrowTheorem(P, S),
						EImply(Not(P), R),
						EImply(Not(P), S),
						IAnd(R, S),
						IOr(And(P, Q), And(R, S), false),
					),
					ArrowProof(Not(Q),
						Reit(Or(Q, R)),
						DisjunctionArrowTheorem(Q, R),
						Reit(Or(Q, S)),
						DisjunctionArrowTheorem(Q, S),
						EImply(Not(Q), R),
						EImply(Not(Q), S),
						IAnd(R, S),
						IOr(And(P, Q), And(R, S), false),
					),
					EOr(Not(P), Not(Q), Or(And(P, Q), And(R, S))),
				),
				ExcludedMiddleTheorem(And(P, Q)),
				EOr(And(P, Q), Not(And(P, Q)), Or(And(P, Q), And(R, S))),
			),
			IDArrow(
				Or(And(P, Q), And(R, S)),
				And(And(Or(P, R), Or(P, S)), And(Or(Q, R), Or(Q, S))),
			),
		),
	),
	NewProofsSection("disjunction",
		RootProof("( ( ( P -> R ) ^ ( Q -> S ) ) ^ ( P v Q ) ) -> ( R v S )",
			ArrowProof(And(And(Arrow(P, R), Arrow(Q, S)), Or(P, Q)),
				EAnd(And(Arrow(P, R), Arrow(Q, S)), Or(P, Q), true),
				EAnd(And(Arrow(P, R), Arrow(Q, S)), Or(P, Q), false),
				EAnd(Arrow(P, R), Arrow(Q, S), true),
				EAnd(Arrow(P, R), Arrow(Q, S), false),
				ArrowProof(P,
					Reit(Arrow(P, R)),
					EImply(P, R),
					IOr(R, S, true),
				),
				ArrowProof(Q,
					Reit(Arrow(Q, S)),
					EImply(Q, S),
					IOr(R, S, false),
				),
				EOr(P, Q, Or(R, S)),
			),
		),
		RootProof("( ( ( P -> R ) ^ ( Q -> S ) ) ^ ( ~ R v Q ) ) -> ( ~ P v S )",
			ArrowProof(And(And(Arrow(P, R), Arrow(Q, S)), Or(Not(R), Q)),
				EAnd(And(Arrow(P, R), Arrow(Q, S)), Or(Not(R), Q), true),
				EAnd(And(Arrow(P, R), Arrow(Q, S)), Or(Not(R), Q), false), // ~ R v Q
				EAnd(Arrow(P, R), Arrow(Q, S), true),                      // P -> R
				EAnd(Arrow(P, R), Arrow(Q, S), false),                     // Q -> S
				ArrowProof(Not(R),
					Reit(Arrow(P, R)),
					ContrapositiveTheorem(P, R), // ~ R -> ~ P
					EImply(Not(R), Not(P)),      // ~ P
					IOr(Not(P), S, true),
				), // ~ R -> ( ~ P v S )
				ArrowProof(Q,
					Reit(Arrow(Q, S)),
					EImply(Q, S),
					IOr(Not(P), S, false),
				), // Q -> ( ~ P v S )
				EOr(Not(R), Q, Or(Not(P), S)),
			),
		),
		RootProof("( ( ( P -> R ) ^ ( Q -> S ) ) ^ ( ~ R v ~ S ) ) -> ( ~ P v ~ Q )",
			ArrowProof(And(And(Arrow(P, R), Arrow(Q, S)), Or(Not(R), Not(S))),
				EAnd(And(Arrow(P, R), Arrow(Q, S)), Or(Not(R), Not(S)), true),
				EAnd(And(Arrow(P, R), Arrow(Q, S)), Or(Not(R), Not(S)), false), // ~ R v ~ S
				EAnd(Arrow(P, R), Arrow(Q, S), true),                           // P -> R
				EAnd(Arrow(P, R), Arrow(Q, S), false),                          // Q -> S
				ArrowProof(Not(R),
					Reit(Arrow(P, R)),
					ContrapositiveTheorem(P, R), // ~ R -> ~ P
					EImply(Not(R), Not(P)),      // ~ P
					IOr(Not(P), Not(Q), true),
				), // ~ R -> ( ~ P v ~ Q )
				ArrowProof(Not(S),
					Reit(Arrow(Q, S)),
					ContrapositiveTheorem(Q, S), // ~ S -> ~ Q
					EImply(Not(S), Not(Q)),      // ~ Q
					IOr(Not(P), Not(Q), false),
				), // ~ S -> ( ~ P v ~ Q )
				EOr(Not(R), Not(S), Or(Not(P), Not(Q))),
			),
		),
		RootProof("( ( ( P -> R ) v ( Q -> R ) ) ^ ( P ^ Q ) ) -> R",
			ArrowProof(And(Or(Arrow(P, R), Arrow(Q, R)), And(P, Q)),
				EAnd(Or(Arrow(P, R), Arrow(Q, R)), And(P, Q), true),  // ( P -> R ) v ( Q -> R)
				EAnd(Or(Arrow(P, R), Arrow(Q, R)), And(P, Q), false), // P ^ Q
				EAnd(P, Q, true),  // P
				EAnd(P, Q, false), // Q
				ArrowProof(Arrow(P, R),
					Reit(P),
					EImply(P, R),
				), // ( P -> R ) -> R
				ArrowProof(Arrow(Q, R),
					Reit(Q),
					EImply(Q, R),
				), // ( Q -> R ) -> R
				EOr(Arrow(P, R), Arrow(Q, R), R),
			),
		),
		// TODO ( ( ( P -> R ) v ( Q -> S ) ) ^ ( P ^ Q ) ) -> ( R v S )
		RootProof("( ( ( P -> R ) v ( Q -> R ) ) ^ ~ R ) -> ( ~ P v ~ Q )",
			ArrowProof(And(Or(Arrow(P, R), Arrow(Q, R)), Not(R)),
				EAnd(Or(Arrow(P, R), Arrow(Q, R)), Not(R), true),  // ( P -> R ) v ( Q -> R)
				EAnd(Or(Arrow(P, R), Arrow(Q, R)), Not(R), false), // ~ R
				ArrowProof(Arrow(P, R),
					Reit(Not(R)),
					ContrapositiveTheorem(P, R),
					EImply(Not(R), Not(P)),
					IOr(Not(P), Not(Q), true),
				), // ( P -> R ) -> ( ~ P v ~ Q )
				ArrowProof(Arrow(Q, R),
					Reit(Not(R)),
					ContrapositiveTheorem(Q, R),
					EImply(Not(R), Not(Q)),
					IOr(Not(P), Not(Q), false),
				), // ( Q -> R ) -> ( ~ P v ~ Q )
				EOr(Arrow(P, R), Arrow(Q, R), Or(Not(P), Not(Q))),
			),
		),
		// TODO ( ( ( P -> R ) v ( Q -> S ) ) ^ ( P ^ ~ S ) ) -> ( R v ~ Q )
		// TODO ( ( ( P -> R ) v ( Q -> S ) ) ^ ( ~ R ^ ~ S ) ) -> ( ~ P v ~ Q )
	),

	NewProofsSection("biconditional",
		RootProof("( ( P <-> Q ) ^ P ) -> Q",
			ArrowProof(And(DArrow(P, Q), P),
				EAnd(DArrow(P, Q), P, true),
				EAnd(DArrow(P, Q), P, false),
				EDArrow(P, Q, true),
				EImply(P, Q),
			),
		),
		RootProof("( ( P <-> Q ) ^ ~ P ) -> ~ Q",
			ArrowProof(And(DArrow(P, Q), Not(P)),
				EAnd(DArrow(P, Q), Not(P), true),
				EAnd(DArrow(P, Q), Not(P), false),
				EDArrow(P, Q, false), // Q -> P
				ContraProof(Q,
					Reit(Not(P)),
					Reit(Arrow(Q, P)),
					EImply(Q, P),
				),
			),
		),
		RootProof("( P <-> ~ Q ) -> ( ~ P <-> Q )",
			ArrowProof(DArrow(P, Not(Q)),
				EDArrow(P, Not(Q), true),         // P -> ~ Q
				ContrapositiveTheorem(P, Not(Q)), // Q -> ~ P
				EDArrow(P, Not(Q), false),        // ~ Q -> P
				ContrapositiveTheorem(Not(Q), P), // ~ P -> Q
				IDArrow(Not(P), Q),
			),
		),
		RootProof("( P <-> Q ) <-> ( ~ P <-> ~ Q )",
			ArrowProof(DArrow(P, Q),
				EDArrow(P, Q, true),
				ContrapositiveTheorem(P, Q),
				EDArrow(P, Q, false),
				ContrapositiveTheorem(Q, P),
				IDArrow(Not(P), Not(Q)),
			),
			ArrowProof(DArrow(Not(P), Not(Q)),
				EDArrow(Not(P), Not(Q), true),
				ContrapositiveTheorem(Not(P), Not(Q)),
				EDArrow(Not(P), Not(Q), false),
				ContrapositiveTheorem(Not(Q), Not(P)),
				IDArrow(P, Q),
			),
			IDArrow(DArrow(P, Q), DArrow(Not(P), Not(Q))),
		),
		RootProof("( P <-> Q ) v ~ ( P <-> Q )",
			ContraProof(Not(Or(DArrow(P, Q), Not(DArrow(P, Q)))),
				ContraProof(DArrow(P, Q),
					IOr(DArrow(P, Q), Not(DArrow(P, Q)), true),
					Reit(Not(Or(DArrow(P, Q), Not(DArrow(P, Q))))),
				),
				IOr(DArrow(P, Q), Not(DArrow(P, Q)), false),
			),
		),
		RootProof("( P <-> ~ Q ) <-> ( ( P v Q ) ^ ( ~ P v ~ Q ) )",
			ArrowProof(DArrow(P, Not(Q)),
				EDArrow(P, Not(Q), true),         // P -> ~ Q
				EDArrow(P, Not(Q), false),        // ~ Q -> P
				ContrapositiveTheorem(Not(Q), P), // ~ P -> Q
				ArrowProof(P,
					IOr(P, Q, true),                    // P v Q
					Reit(Arrow(P, Not(Q))),             // P -> ~ Q
					EImply(P, Not(Q)),                  // ~ Q
					IOr(Not(P), Not(Q), false),         // ~ P v ~ Q
					IAnd(Or(P, Q), Or(Not(P), Not(Q))), // ( P v Q ) ^ ( ~ P v ~ Q )
				), // P -> ( ( P v Q ) ^ ( ~ P v ~ Q ) )
				ArrowProof(Not(P),
					IOr(Not(P), Not(Q), true),          // ~ P v ~ Q
					Reit(Arrow(Not(P), Q)),             // ~ P -> Q
					EImply(Not(P), Q),                  // Q
					IOr(P, Q, false),                   // P v Q
					IAnd(Or(P, Q), Or(Not(P), Not(Q))), // ( P v Q ) ^ ( ~ P v ~ Q )
				), // ~ P -> ( ( P v Q ) ^ ( ~ P v ~ Q ) )
				ExcludedMiddleTheorem(P), // P v ~ P
				EOr(P, Not(P), And(Or(P, Q), Or(Not(P), Not(Q)))),
			),
			ArrowProof(And(Or(P, Q), Or(Not(P), Not(Q))),
				EAnd(Or(P, Q), Or(Not(P), Not(Q)), true),  // P v Q
				EAnd(Or(P, Q), Or(Not(P), Not(Q)), false), // ~ P v ~ Q
				ArrowProof(Not(P),
					ContraProof(And(P, Q),
						EAnd(P, Q, true), // P
						Reit(Not(P)),     // ~ P
					), // ~ ( P ^ Q )
				), // ~ P -> ~ ( P ^ Q )
				ArrowProof(Not(Q),
					ContraProof(And(P, Q),
						EAnd(P, Q, false), // Q
						Reit(Not(Q)),      // ~ Q
					), // ~ ( P ^ Q )
				), // ~ Q -> ~ ( P ^ Q )
				EOr(Not(P), Not(Q), Not(And(P, Q))), // ~ ( P ^ Q )
				ArrowProof(P,
					ContraProof(And(Not(P), Not(Q)),
						EAnd(Not(P), Not(Q), true), // ~ P
						Reit(P),                    // P
					), // ~ ( ~ P ^ ~ Q )
				), // P -> ~ ( ~ P ^ ~ Q )
				ArrowProof(Q,
					ContraProof(And(Not(P), Not(Q)),
						EAnd(Not(P), Not(Q), false), // ~ Q
						Reit(Q),                     // Q
					), // ~ ( ~ P ^ ~ Q )
				), // Q -> ~ ( ~ P ^ ~ Q )
				EOr(P, Q, Not(And(Not(P), Not(Q)))), // ~ ( ~ P ^ ~ Q )
				ArrowProof(P,
					ContraProof(Q,
						Reit(P),              // P
						IAnd(P, Q),           // P ^ Q
						Reit(Not(And(P, Q))), // ~ ( P ^ Q )
					), // ~ Q
				), // P -> ~ Q
				ArrowProof(Not(Q),
					ContraProof(Not(P),
						Reit(Not(Q)),                   // ~ Q
						IAnd(Not(P), Not(Q)),           // ~ P ^ ~ Q
						Reit(Not(And(Not(P), Not(Q)))), // ~ ( ~ P ^ ~ Q )
					), // P
				), // ~ Q -> P
				IDArrow(P, Not(Q)),
			),
			IDArrow(
				DArrow(P, Not(Q)),
				And(Or(P, Q), Or(Not(P), Not(Q))),
			),
		),
	// TODO no distributive
	),
	NewProofsSection("miscellaneous",
		// TODO may be wrong ?  this isn't the implication distributive formula according to wikipedia
		RootProof("( ( P -> Q ) -> R ) -> ( ( P -> Q ) -> ( P -> R ) )",
			ArrowProof(Arrow(Arrow(P, Q), R),
				ArrowProof(Arrow(P, Q),
					Reit(Arrow(Arrow(P, Q), R)),
					EImply(Arrow(P, Q), R),
					ArrowProof(P,
						Reit(R),
					),
				),
			),
		),
		// TODO may be wrong ?  this isn't the implication distributive formula according to wikipedia
		// RootProof("( ( P -> Q ) -> ( P -> R ) ) -> ( ( P -> Q ) -> R )",
		// 	ArrowProof(Arrow(Arrow(P, Q), Arrow(P, R)),
		// 		ArrowProof(Arrow(P, Q)),
		// 	),
		// ),
		RootProof("( P -> Q ) v ( Q -> R )",
			ContraProof(Not(Or(Arrow(P, Q), Arrow(Q, R))),
				DeMorgansOrToAndTheorem(Arrow(P, Q), Arrow(Q, R), true), // ~ ( P -> Q ) ^ ~ ( Q -> R )
				EAnd(Not(Arrow(P, Q)), Not(Arrow(Q, R)), true),          // ~ ( P -> Q )
				EAnd(Not(Arrow(P, Q)), Not(Arrow(Q, R)), false),         // ~ ( Q -> R )
				ArrowConjunctionTheorem(P, Q, true),                     // P ^ ~ Q
				EAnd(P, Not(Q), false),                                  // ~ Q
				ArrowConjunctionTheorem(Q, R, true),                     // Q ^ ~ R
				EAnd(Q, Not(R), true),                                   // Q
			), // ( P -> Q ) v ( Q -> R )
		),
		RootProof("( P -> Q ) v ( Q -> P )",
			ContraProof(Not(Or(Arrow(P, Q), Arrow(Q, P))),
				DeMorgansOrToAndTheorem(Arrow(P, Q), Arrow(Q, P), true), // ~ ( P -> Q ) ^ ~ ( Q -> P )
				EAnd(Not(Arrow(P, Q)), Not(Arrow(Q, P)), true),          // ~ ( P -> Q )
				EAnd(Not(Arrow(P, Q)), Not(Arrow(Q, P)), false),         // ~ ( Q -> P )
				ArrowConjunctionTheorem(P, Q, true),                     // P ^ ~ Q
				EAnd(P, Not(Q), false),                                  // ~ Q
				ArrowConjunctionTheorem(Q, P, true),                     // Q ^ ~ P
				EAnd(Q, Not(P), true),                                   // Q
			), // ( P -> Q ) v ( Q -> P )
		),
	),
}
