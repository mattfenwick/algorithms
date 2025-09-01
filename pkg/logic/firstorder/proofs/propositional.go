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
					&Reiterate{Formula: Not(Or(P, Not(P)))},
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
						&Reiterate{Formula: Arrow(P, Q)},
						EImply(P, Q),
						&Reiterate{Formula: Not(Q)},
					),
				),
			),
			ArrowProof(
				Arrow(Not(Q), Not(P)),
				ArrowProof(
					P,
					ContraProof(
						Not(Q),
						&Reiterate{Formula: Arrow(Not(Q), Not(P))},
						EImply(Not(Q), Not(P)),
						&Reiterate{Formula: P},
					),
				),
			),
			IDArrow(Arrow(P, Q), Arrow(Not(Q), Not(P))),
		),
		RootProof("~ ( P -> Q ) <-> ( P ^ ~ Q )",
			ArrowProof(Not(Arrow(P, Q)),
				ContraProof(Not(P),
					ArrowProof(Not(Q),
						&Reiterate{Formula: Not(P)},
					), // ~ Q -> ~ P
					ContrapositiveTheorem(Not(Q), Not(P)), // P -> Q
					&Reiterate{Formula: Not(Arrow(P, Q))},
				), // P
				ContraProof(Q,
					&Reiterate{Formula: P},
					IImply(P, Q),
					&Reiterate{Formula: Not(Arrow(P, Q))},
				),
				IAnd(P, Not(Q)),
			),
			ArrowProof(And(P, Not(Q)),
				ContraProof(Arrow(P, Q),
					&Reiterate{Formula: And(P, Not(Q))},
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
		// 			&Reiterate{Formula: Arrow(P, Q)},
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
						&Reiterate{Formula: Q},
					),
				), // Q -> ( ~ P -> Q )
				ArrowProof(P,
					ArrowProof(Not(Q),
						&Reiterate{Formula: P},
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
					&Reiterate{Formula: Arrow(Not(P), Q)},
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
					&Reiterate{Formula: Not(Or(P, Q))},
				),
				ContraProof(Q,
					IOr(P, Q, false),
					&Reiterate{Formula: Not(Or(P, Q))},
				),
				IAnd(Not(P), Not(Q)),
			),
			ArrowProof(And(Not(P), Not(Q)),
				ArrowProof(P,
					ContraProof(And(Not(P), Not(Q)),
						EAnd(Not(P), Not(Q), true),
						&Reiterate{Formula: P},
					),
				),
				ArrowProof(Q,
					ContraProof(And(Not(P), Not(Q)),
						EAnd(Not(P), Not(Q), false),
						&Reiterate{Formula: Q},
					),
				),
				ContraProof(Or(P, Q),
					&Reiterate{Formula: Arrow(P, Not(And(Not(P), Not(Q))))},
					&Reiterate{Formula: Arrow(Q, Not(And(Not(P), Not(Q))))},
					EOr(P, Q, Not(And(Not(P), Not(Q)))),
					&Reiterate{Formula: And(Not(P), Not(Q))},
				),
			),
			IDArrow(Not(Or(P, Q)), And(Not(P), Not(Q))),
		),
		RootProof("~ ( P ^ Q ) <-> ( ~ P v ~ Q )",
			ArrowProof(Not(And(P, Q)),
				ContraProof(Not(Or(Not(P), Not(Q))),
					ContraProof(Not(P),
						IOr(Not(P), Not(Q), true),
						&Reiterate{Formula: Not(Or(Not(P), Not(Q)))},
					),
					ContraProof(Not(Q),
						IOr(Not(P), Not(Q), false),
						&Reiterate{Formula: Not(Or(Not(P), Not(Q)))},
					),
					IAnd(P, Q),
					&Reiterate{Formula: Not(And(P, Q))},
				),
			),
			ArrowProof(Or(Not(P), Not(Q)),
				ArrowProof(Not(P),
					ContraProof(And(P, Q),
						EAnd(P, Q, true),
						&Reiterate{Formula: Not(P)},
					),
				),
				ArrowProof(Not(Q),
					ContraProof(And(P, Q),
						EAnd(P, Q, false),
						&Reiterate{Formula: Not(Q)},
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
						&Reiterate{Formula: Arrow(P, Not(Q))}, // P -> ~ Q
						EDArrow(P, Q, true),               // P -> Q
						&Reiterate{Formula: P},                      // P
						EImply(P, Q),                             // Q
						EImply(P, Not(Q)),                        // ~ Q
					), // ~ ( P <-> Q )
				), // P -> ~ ( P <-> Q )
				ArrowProof(Not(P),
					ContraProof(DArrow(P, Q),
						EDArrow(P, Q, false),              // Q -> P
						ContrapositiveTheorem(Q, P),              // ~ P -> ~ Q
						&Reiterate{Formula: Not(P)},                 // ~ P
						EImply(Not(P), Not(Q)),                   // ~ Q
						&Reiterate{Formula: Arrow(Not(Q), P)}, // ~ Q -> P
						EImply(Not(Q), P),                        // P
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
					&Reiterate{Formula: Not(DArrow(P, Q))},
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
					&Reiterate{Formula: Not(DArrow(P, Q))},
				),
				IDArrow(P, Not(Q)),
			),
			IDArrow(DArrow(P, Not(Q)), Not(DArrow(P, Q))),
		),
	),
	NewProofsSection("basics",
		RootProof(
			"P -> P",
			ArrowProof(P)),
		RootProof("~ ( P ^ ~ P )",
			ContraProof(
				And(P, Not(P)),
				EAnd(P, Not(P), true),
				EAnd(P, Not(P), false),
			)),
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
		RootProof("( P ^ Q ) -> ( P v Q )",
			ArrowProof(And(P, Q),
				EAnd(P, Q, true),
				IOr(P, Q, true),
			),
		),
		RootProof("( P -> ~ P ) -> ~ P",
			ArrowProof(Arrow(P, Not(P)),
				ContraProof(P,
					&Reiterate{Formula: Arrow(P, Not(P))},
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
				ExcludedMiddleTheorem(Arrow(P, Q)),          // ( P -> Q ) v ~ ( P -> Q )
				EOr(Arrow(P, Q), Not(Arrow(P, Q)), P), // p
			),
		),
		RootProof("( P ^ ~ P ) -> Q",
			NonContradictionTheorem(P), // ~ ( P ^ ~ P )
			ArrowProof(Not(Q),
				&Reiterate{Formula: Not(And(P, Not(P)))},
			), // ~ Q -> ~ ( P ^ ~ P )
			ContrapositiveTheorem(Not(Q), Not(And(P, Not(P)))), // ( P ^ ~ P ) -> Q
		),
	),
	NewProofsSection("arrows",
		RootProof("Q -> ( P -> Q )",
			ArrowProof(
				Q,
				ArrowProof(P,
					&Reiterate{Formula: Q})),
		),
		RootProof("~ P -> ( P -> Q )",
			ArrowProof(Not(P),
				ArrowProof(Not(Q),
					&Reiterate{Formula: Not(P)},
				),
				ContrapositiveTheorem(Not(Q), Not(P)),
			),
		),
		RootProof("( ( P -> Q ) ^ ( Q -> R ) ) -> ( P -> R )",
			ArrowProof(And(Arrow(P, Q), Arrow(Q, R)),
				ArrowProof(P,
					&Reiterate{Formula: And(Arrow(P, Q), Arrow(Q, R))},
					EAnd(Arrow(P, Q), Arrow(Q, R), true),
					EAnd(Arrow(P, Q), Arrow(Q, R), false),
					EImply(P, Q),
					EImply(Q, R),
				)),
		),
		RootProof("( ( P ^ Q ) -> R ) <-> ( P -> ( Q -> R ) )",
			ArrowProof(Arrow(And(P, Q), R),
				ArrowProof(P,
					ArrowProof(Q,
						&Reiterate{Formula: P},
						IAnd(P, Q),
						&Reiterate{Formula: Arrow(And(P, Q), R)},
						EImply(And(P, Q), R),
					),
				),
			),
			ArrowProof(Arrow(P, Arrow(Q, R)),
				ArrowProof(And(P, Q),
					EAnd(P, Q, true),
					EAnd(P, Q, false),
					&Reiterate{Formula: Arrow(P, Arrow(Q, R))},
					EImply(P, Arrow(Q, R)),
					EImply(Q, R),
				),
			),
			IDArrow(Arrow(And(P, Q), R), Arrow(P, Arrow(Q, R))),
		),
		RootProof("( ( P -> Q ) ^ ( P -> ~ Q ) ) -> ~ P",
			ArrowProof(And(Arrow(P, Q), Arrow(P, Not(Q))),
				ContraProof(P,
					&Reiterate{Formula: And(Arrow(P, Q), Arrow(P, Not(Q)))},
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
						&Reiterate{Formula: Arrow(P, Arrow(Q, R))},
						EImply(P, Arrow(Q, R)),
						&Reiterate{Formula: Q},
						EImply(Q, R),
					),
				),
			),
		),
		RootProof("( ( P v Q ) -> R ) <-> ( ( P -> R ) ^ ( Q -> R ) )",
			ArrowProof(Arrow(Or(P, Q), R),
				ArrowProof(P,
					IOr(P, Q, true),
					&Reiterate{Formula: Arrow(Or(P, Q), R)},
					EImply(Or(P, Q), R),
				),
				ArrowProof(Q,
					IOr(P, Q, false),
					&Reiterate{Formula: Arrow(Or(P, Q), R)},
					EImply(Or(P, Q), R),
				),
				IAnd(Arrow(P, R), Arrow(Q, R)),
			),
			ArrowProof(And(Arrow(P, R), Arrow(Q, R)),
				ArrowProof(Or(P, Q),
					&Reiterate{Formula: And(Arrow(P, R), Arrow(Q, R))},
					EAnd(Arrow(P, R), Arrow(Q, R), true),
					EAnd(Arrow(P, R), Arrow(Q, R), false),
					EOr(P, Q, R),
				),
			),
			IDArrow(Arrow(Or(P, Q), R), And(Arrow(P, R), Arrow(Q, R))),
		),
		RootProof("( ( P -> R ) v ( Q -> R ) ) <-> ( ( P ^ Q ) -> R )",
			ArrowProof(Or(Arrow(P, R), Arrow(Q, R)),
				ArrowProof(Arrow(P, R),
					ArrowProof(And(P, Q),
						&Reiterate{Formula: Arrow(P, R)},
						EAnd(P, Q, true),
						EImply(P, R),
					),
				),
				ArrowProof(Arrow(Q, R),
					ArrowProof(And(P, Q),
						&Reiterate{Formula: Arrow(Q, R)},
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
					ArrowConjunctionTheorem(P, R, true),                         // P ^ ~ R
					EAnd(P, Not(R), true),                                       // P
					EAnd(P, Not(R), false),                                      // ~ R
					ArrowConjunctionTheorem(Q, R, true),                         // Q ^ ~ R
					EAnd(Q, Not(R), true),                                       // Q
					&Reiterate{Formula: Arrow(And(P, Q), R)},                 // ( P ^ Q ) -> R
					IAnd(P, Q),           // P ^ Q
					EImply(And(P, Q), R), // R
				), // ( P -> R ) v ( Q -> R )
			),
			IDArrow(Or(Arrow(P, R), Arrow(Q, R)), Arrow(And(P, Q), R)),
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
					BiconditionalNegationTheorem(Q, R),                    // Q <-> ~ R
					ContraProof(R,
						&Reiterate{Formula: Arrow(R, DArrow(P, Q))}, // R -> ( P <-> Q )
						EImply(R, DArrow(P, Q)),                        // P <-> Q
						EDArrow(P, Q, true),                            // P -> Q
						&Reiterate{Formula: P},                                   // P
						EImply(P, Q),                                          // Q
						&Reiterate{Formula: DArrow(Q, Not(R))},            // Q <-> ~ R
						EDArrow(Q, Not(R), true),                       // Q -> ~ R
						EImply(Q, Not(R)),                                     // ~ R
					), // ~ R
					EDArrow(Q, Not(R), false),                      // ~ R -> Q
					EImply(Not(R), Q),                                     // Q
					&Reiterate{Formula: Arrow(DArrow(P, Q), R)}, // ( P <-> Q ) -> R
					ContrapositiveTheorem(DArrow(P, Q), R),         // ~ R -> ~ ( P <-> Q )
					EImply(Not(R), Not(DArrow(P, Q))),              // ~ ( P <-> Q )
					BiconditionalNegationTheorem(P, Q),                    // P <-> ~ Q
					EDArrow(P, Not(Q), true),                       // P -> ~ Q
					EImply(P, Not(Q)),                                     // ~ Q
				), // ( P -> ( Q <-> R ) )
				ContraProof(Not(Arrow(DArrow(Q, R), P)),
					ArrowConjunctionTheorem(DArrow(Q, R), P, true), // ( Q <-> R ) ^ ~ P
					EAnd(DArrow(Q, R), Not(P), true),               // Q <-> R
					EAnd(DArrow(Q, R), Not(P), false),              // ~ P
					ContraProof(R,
						&Reiterate{Formula: Arrow(R, DArrow(P, Q))}, // R -> ( P <-> Q )
						&Reiterate{Formula: DArrow(Q, R)},                 // Q <-> R
						EDArrow(Q, R, false),                           // R -> Q
						EImply(R, Q),                                          // Q
						EImply(R, DArrow(P, Q)),                        // P <-> Q
						EDArrow(P, Q, false),                           // Q -> P
						EImply(Q, P),                                          // P
						&Reiterate{Formula: Not(P)},                              // ~ P
					), // ~ R
					EDArrow(Q, R, true),                            // Q -> R
					ContrapositiveTheorem(Q, R),                           // ~ R -> ~ Q
					EImply(Not(R), Not(Q)),                                // ~ Q
					&Reiterate{Formula: Arrow(DArrow(P, Q), R)}, // ( P <-> Q ) -> R
					ContrapositiveTheorem(DArrow(P, Q), R),         // ~ R -> ~ ( P <-> Q)
					EImply(Not(R), Not(DArrow(P, Q))),              // ~ ( P <-> Q )
					BiconditionalNegationTheorem(P, Q),                    // P <-> ~ Q
					EDArrow(P, Not(Q), false),                      // ~ Q -> P
					EImply(Not(Q), P),                                     // P
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
						&Reiterate{Formula: Arrow(P, Q)}, // P -> Q
						EImply(P, Q),                        // Q
						&Reiterate{Formula: Arrow(P, DArrow(Q, R))}, // P -> ( Q <-> R )
						EImply(P, DArrow(Q, R)),                        // Q <-> R
						EDArrow(Q, R, true),                            // Q -> R
						EImply(Q, R),                                          // R
						&Reiterate{Formula: Not(R)},                              // ~ R
					), // ~ P
					&Reiterate{Formula: Arrow(DArrow(Q, R), P)}, // ( Q <-> R ) -> P
					ContrapositiveTheorem(DArrow(Q, R), P),         // ~ P -> ~ ( Q <-> R )
					EImply(Not(P), Not(DArrow(Q, R))),              // ~ ( Q <-> R )
					BiconditionalNegationTheorem(Q, R),                    // Q <-> ~ R
					EDArrow(Q, Not(R), false),                      // ~ R -> Q
					EImply(Not(R), Q),                                     // Q
					EDArrow(P, Q, false),                           // Q -> P
					EImply(Q, P),                                          // P
				), // ( ( P <-> Q ) -> R )
				ContraProof(Not(Arrow(R, DArrow(P, Q))),
					ArrowConjunctionTheorem(R, DArrow(P, Q), true), // R ^ ~ ( P <-> Q )
					EAnd(R, Not(DArrow(P, Q)), true),               // R
					EAnd(R, Not(DArrow(P, Q)), false),              // ~ ( P <-> Q )
					BiconditionalNegationTheorem(P, Q),                    // P <-> ~ Q
					EDArrow(P, Not(Q), true),                       // P -> ~ Q
					ContraProof(P,
						&Reiterate{Formula: Arrow(P, Not(Q))},              // P -> ~ Q
						EImply(P, Not(Q)),                                     // ~ Q
						&Reiterate{Formula: Arrow(P, DArrow(Q, R))}, // P -> ( Q <-> R )
						EImply(P, DArrow(Q, R)),                        // Q <-> R
						EDArrow(Q, R, false),                           // R -> Q
						&Reiterate{Formula: R},                                   // R
						EImply(R, Q),                                          // Q
					), // ~ P
					EDArrow(P, Not(Q), false),                      // ~ Q -> P
					ContrapositiveTheorem(Not(Q), P),                      // ~ P -> Q
					EImply(Not(P), Q),                                     // Q
					&Reiterate{Formula: Arrow(DArrow(Q, R), P)}, // ( Q <-> R ) -> P
					ContrapositiveTheorem(DArrow(Q, R), P),         // ~ P -> ~ ( Q <-> R )
					EImply(Not(P), Not(DArrow(Q, R))),              // ~ ( Q <-> R )
					BiconditionalNegationTheorem(Q, R),                    // Q <-> ~ R
					EDArrow(Q, Not(R), true),                       // Q -> ~ R
					EImply(Q, Not(R)),                                     // ~ R
				), // ( R -> ( P <-> Q ) )
				IDArrow(DArrow(P, Q), R),
			),
			IDArrow(DArrow(DArrow(P, Q), R), DArrow(P, DArrow(Q, R))),
		),
	),
	NewProofsSection("distributivity",
		RootProof("( P -> ( Q -> R ) ) -> ( ( P -> Q ) -> ( P -> R ) )",
			ArrowProof(Arrow(P, Arrow(Q, R)),
				ArrowProof(Arrow(P, Q),
					ArrowProof(P,
						&Reiterate{Formula: Arrow(P, Arrow(Q, R))},
						EImply(P, Arrow(Q, R)),
						&Reiterate{Formula: Arrow(P, Q)},
						EImply(P, Q),
						EImply(Q, R),
					),
				),
			),
		),
		RootProof("( ( P -> Q ) -> ( P -> R ) ) -> ( P -> ( Q -> R ) )",
			ArrowProof(Arrow(Arrow(P, Q), Arrow(P, R)),
				ArrowProof(P,
					ArrowProof(Q,
						&Reiterate{Formula: Arrow(Arrow(P, Q), Arrow(P, R))},
						&Reiterate{Formula: P},
						IImply(P, Q),
						EImply(Arrow(P, Q), Arrow(P, R)),
						EImply(P, R),
					),
				),
			),
		),
		RootProof("( P -> ( Q v R ) ) -> ( ( P -> Q ) v ( P -> R ) )",
			ArrowProof(Arrow(P, Or(Q, R)),
				ArrowProof(P,
					&Reiterate{Formula: Arrow(P, Or(Q, R))},
					EImply(P, Or(Q, R)),
					ArrowProof(Q,
						&Reiterate{Formula: P},
						IImply(P, Q),
						IOr(Arrow(P, Q), Arrow(P, R), true),
					),
					ArrowProof(R,
						&Reiterate{Formula: P},
						IImply(P, R),
						IOr(Arrow(P, Q), Arrow(P, R), false),
					),
					EOr(Q, R, Or(Arrow(P, Q), Arrow(P, R))),
				),
				ArrowProof(Not(P),
					ArrowProof(Not(Q),
						&Reiterate{Formula: Not(P)},
					),
					ContrapositiveTheorem(Not(Q), Not(P)),
					IOr(Arrow(P, Q), Arrow(P, R), true),
				),
				ExcludedMiddleTheorem(P),
				EOr(P, Not(P), Or(Arrow(P, Q), Arrow(P, R))),
			),
		),
		RootProof("( ( P -> Q ) v ( P -> R ) ) -> ( P -> ( Q v R ) )",
			ArrowProof(Or(Arrow(P, Q), Arrow(P, R)),
				ArrowProof(P,
					ArrowProof(Arrow(P, Q),
						&Reiterate{Formula: P},
						EImply(P, Q),
						IOr(Q, R, true),
					),
					ArrowProof(Arrow(P, R),
						&Reiterate{Formula: P},
						EImply(P, R),
						IOr(Q, R, false),
					),
					&Reiterate{Formula: Or(Arrow(P, Q), Arrow(P, R))},
					EOr(Arrow(P, Q), Arrow(P, R), Or(Q, R)),
				),
			),
		),
		RootProof("( P -> ( Q ^ R ) ) -> ( ( P -> Q ) ^ ( P -> R ) )",
			ArrowProof(Arrow(P, And(Q, R)),
				ArrowProof(P,
					&Reiterate{Formula: Arrow(P, And(Q, R))},
					EImply(P, And(Q, R)),
					EAnd(Q, R, true),
				),
				ArrowProof(P,
					&Reiterate{Formula: Arrow(P, And(Q, R))},
					EImply(P, And(Q, R)),
					EAnd(Q, R, false),
				),
				IAnd(Arrow(P, Q), Arrow(P, R)),
			),
		),
		RootProof("( ( P -> Q ) ^ ( P -> R ) ) -> ( P -> ( Q ^ R ) )",
			ArrowProof(And(Arrow(P, Q), Arrow(P, R)),
				ArrowProof(P,
					&Reiterate{Formula: And(Arrow(P, Q), Arrow(P, R))},
					EAnd(Arrow(P, Q), Arrow(P, R), true),
					EAnd(Arrow(P, Q), Arrow(P, R), false),
					EImply(P, Q),
					EImply(P, R),
					IAnd(Q, R),
				),
			),
		),
		RootProof("( P ^ ( Q v R ) ) -> ( ( P ^ Q ) v ( P ^ R ) )",
			ArrowProof(And(P, Or(Q, R)),
				EAnd(P, Or(Q, R), true),
				EAnd(P, Or(Q, R), false),
				ArrowProof(Q,
					&Reiterate{Formula: P},
					IAnd(P, Q),
					IOr(And(P, Q), And(P, R), true),
				),
				ArrowProof(R,
					&Reiterate{Formula: P},
					IAnd(P, R),
					IOr(And(P, Q), And(P, R), false),
				),
				EOr(Q, R, Or(And(P, Q), And(P, R))),
			),
		),
		RootProof("( ( P ^ Q ) v ( P ^ R ) ) -> ( P ^ ( Q v R ) )",
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
		),
		RootProof("( P v ( Q ^ R ) ) -> ( ( P v Q ) ^ ( P v R ) )",
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
		),
		RootProof("( ( P v Q ) ^ ( P v R ) ) -> ( P v ( Q ^ R ) )",
			ArrowProof(And(Or(P, Q), Or(P, R)),
				EAnd(Or(P, Q), Or(P, R), true),
				EAnd(Or(P, Q), Or(P, R), false),
				ArrowProof(P,
					IOr(P, And(Q, R), true),
				),
				ArrowProof(P,
					ArrowProof(Not(Q),
						&Reiterate{Formula: P},
					),
					ContrapositiveTheorem(Not(Q), P),
				), // P -> ( ~ P -> Q )
				ArrowProof(P,
					ArrowProof(Not(R),
						&Reiterate{Formula: P},
					),
					ContrapositiveTheorem(Not(R), P),
				), // P -> ( ~ P -> R )
				ArrowProof(Not(P),
					&Reiterate{Formula: Or(P, Q)},
					&Reiterate{Formula: Or(P, R)},
					ArrowProof(Q,
						&Reiterate{Formula: Not(P)},
						IImply(Not(P), Q),
					), // Q -> ( ~ P -> Q )
					&Reiterate{Formula: Arrow(P, Arrow(Not(P), Q))},
					EOr(P, Q, Arrow(Not(P), Q)),
					EImply(Not(P), Q),
					ArrowProof(R,
						&Reiterate{Formula: Not(P)},
						IImply(Not(P), R),
					), // R -> ( ~ P -> R )
					&Reiterate{Formula: Arrow(P, Arrow(Not(P), R))},
					EOr(P, R, Arrow(Not(P), R)),
					EImply(Not(P), R),
					IAnd(Q, R),
					IOr(P, And(Q, R), false),
				),
				ExcludedMiddleTheorem(P),
				EOr(P, Not(P), Or(P, And(Q, R))),
			),
		),
		// TODO https://en.wikipedia.org/wiki/Distributive_property#Truth_functional_connectives
		// TODO double distribution
		// TODO ( ( P ^ Q ) v ( R ^ S ) ) <-> ( ( ( P v R ) ^ ( P v S ) ) ^ ( ( Q v R ) ^ ( Q v S ) ) )
		// TODO ( ( P v Q ) ^ ( R v S ) ) <-> ( ( ( P ^ R ) v ( P ^ S ) ) v ( ( Q ^ R ) v ( Q ^ S ) ) )
	),
	NewProofsSection("disjunction",
		RootProof("( ( ( P -> R ) ^ ( Q -> S ) ) ^ ( P v Q ) ) -> ( R v S )",
			ArrowProof(And(And(Arrow(P, R), Arrow(Q, S)), Or(P, Q)),
				EAnd(And(Arrow(P, R), Arrow(Q, S)), Or(P, Q), true),
				EAnd(And(Arrow(P, R), Arrow(Q, S)), Or(P, Q), false),
				EAnd(Arrow(P, R), Arrow(Q, S), true),
				EAnd(Arrow(P, R), Arrow(Q, S), false),
				ArrowProof(P,
					&Reiterate{Formula: Arrow(P, R)},
					EImply(P, R),
					IOr(R, S, true),
				),
				ArrowProof(Q,
					&Reiterate{Formula: Arrow(Q, S)},
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
					&Reiterate{Formula: Arrow(P, R)},
					ContrapositiveTheorem(P, R), // ~ R -> ~ P
					EImply(Not(R), Not(P)),      // ~ P
					IOr(Not(P), S, true),
				), // ~ R -> ( ~ P v S )
				ArrowProof(Q,
					&Reiterate{Formula: Arrow(Q, S)},
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
					&Reiterate{Formula: Arrow(P, R)},
					ContrapositiveTheorem(P, R), // ~ R -> ~ P
					EImply(Not(R), Not(P)),      // ~ P
					IOr(Not(P), Not(Q), true),
				), // ~ R -> ( ~ P v ~ Q )
				ArrowProof(Not(S),
					&Reiterate{Formula: Arrow(Q, S)},
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
					&Reiterate{Formula: P},
					EImply(P, R),
				), // ( P -> R ) -> R
				ArrowProof(Arrow(Q, R),
					&Reiterate{Formula: Q},
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
					&Reiterate{Formula: Not(R)},
					ContrapositiveTheorem(P, R),
					EImply(Not(R), Not(P)),
					IOr(Not(P), Not(Q), true),
				), // ( P -> R ) -> ( ~ P v ~ Q )
				ArrowProof(Arrow(Q, R),
					&Reiterate{Formula: Not(R)},
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
					&Reiterate{Formula: Not(P)},
					&Reiterate{Formula: Arrow(Q, P)},
					EImply(Q, P),
				),
			),
		),
		RootProof("( P <-> ~ Q ) -> ( ~ P <-> Q )",
			ArrowProof(DArrow(P, Not(Q)),
				EDArrow(P, Not(Q), true),  // P -> ~ Q
				ContrapositiveTheorem(P, Not(Q)), // Q -> ~ P
				EDArrow(P, Not(Q), false), // ~ Q -> P
				ContrapositiveTheorem(Not(Q), P), // ~ P -> Q
				IDArrow(Not(P), Q),
			),
		),
		RootProof("( P <-> Q ) -> ( ~ P <-> ~ Q )",
			ArrowProof(DArrow(P, Q),
				EDArrow(P, Q, true),
				ContrapositiveTheorem(P, Q),
				EDArrow(P, Q, false),
				ContrapositiveTheorem(Q, P),
				IDArrow(Not(P), Not(Q)),
			),
		),
		RootProof("( ~ P <-> ~ Q ) -> ( P <-> Q )",
			ArrowProof(DArrow(Not(P), Not(Q)),
				EDArrow(Not(P), Not(Q), true),
				ContrapositiveTheorem(Not(P), Not(Q)),
				EDArrow(Not(P), Not(Q), false),
				ContrapositiveTheorem(Not(Q), Not(P)),
				IDArrow(P, Q),
			),
		),
		RootProof("( P <-> Q ) v ~ ( P <-> Q )",
			ContraProof(Not(Or(DArrow(P, Q), Not(DArrow(P, Q)))),
				ContraProof(DArrow(P, Q),
					IOr(DArrow(P, Q), Not(DArrow(P, Q)), true),
					&Reiterate{Formula: Not(Or(DArrow(P, Q), Not(DArrow(P, Q))))},
				),
				IOr(DArrow(P, Q), Not(DArrow(P, Q)), false),
			),
		),
		RootProof("( P <-> ~ Q ) -> ( ( P v Q ) ^ ( ~ P v ~ Q ) )",
			ArrowProof(DArrow(P, Not(Q)),
				EDArrow(P, Not(Q), true),  // P -> ~ Q
				EDArrow(P, Not(Q), false), // ~ Q -> P
				ContrapositiveTheorem(Not(Q), P), // ~ P -> Q
				ArrowProof(P,
					IOr(P, Q, true),                          // P v Q
					&Reiterate{Formula: Arrow(P, Not(Q))}, // P -> ~ Q
					EImply(P, Not(Q)),                        // ~ Q
					IOr(Not(P), Not(Q), false),               // ~ P v ~ Q
					IAnd(Or(P, Q), Or(Not(P), Not(Q))),       // ( P v Q ) ^ ( ~ P v ~ Q )
				), // P -> ( ( P v Q ) ^ ( ~ P v ~ Q ) )
				ArrowProof(Not(P),
					IOr(Not(P), Not(Q), true),                // ~ P v ~ Q
					&Reiterate{Formula: Arrow(Not(P), Q)}, // ~ P -> Q
					EImply(Not(P), Q),                        // Q
					IOr(P, Q, false),                         // P v Q
					IAnd(Or(P, Q), Or(Not(P), Not(Q))),       // ( P v Q ) ^ ( ~ P v ~ Q )
				), // ~ P -> ( ( P v Q ) ^ ( ~ P v ~ Q ) )
				ExcludedMiddleTheorem(P), // P v ~ P
				EOr(P, Not(P), And(Or(P, Q), Or(Not(P), Not(Q)))),
			),
		),
		RootProof("( ( P v Q ) ^ ( ~ P v ~ Q ) ) -> ( P <-> ~ Q )",
			ArrowProof(And(Or(P, Q), Or(Not(P), Not(Q))),
				EAnd(Or(P, Q), Or(Not(P), Not(Q)), true),  // P v Q
				EAnd(Or(P, Q), Or(Not(P), Not(Q)), false), // ~ P v ~ Q
				ArrowProof(Not(P),
					ContraProof(And(P, Q),
						EAnd(P, Q, true),         // P
						&Reiterate{Formula: Not(P)}, // ~ P
					), // ~ ( P ^ Q )
				), // ~ P -> ~ ( P ^ Q )
				ArrowProof(Not(Q),
					ContraProof(And(P, Q),
						EAnd(P, Q, false),        // Q
						&Reiterate{Formula: Not(Q)}, // ~ Q
					), // ~ ( P ^ Q )
				), // ~ Q -> ~ ( P ^ Q )
				EOr(Not(P), Not(Q), Not(And(P, Q))), // ~ ( P ^ Q )
				ArrowProof(P,
					ContraProof(And(Not(P), Not(Q)),
						EAnd(Not(P), Not(Q), true), // ~ P
						&Reiterate{Formula: P},        // P
					), // ~ ( ~ P ^ ~ Q )
				), // P -> ~ ( ~ P ^ ~ Q )
				ArrowProof(Q,
					ContraProof(And(Not(P), Not(Q)),
						EAnd(Not(P), Not(Q), false), // ~ Q
						&Reiterate{Formula: Q},         // Q
					), // ~ ( ~ P ^ ~ Q )
				), // Q -> ~ ( ~ P ^ ~ Q )
				EOr(P, Q, Not(And(Not(P), Not(Q)))), // ~ ( ~ P ^ ~ Q )
				ArrowProof(P,
					ContraProof(Q,
						&Reiterate{Formula: P},              // P
						IAnd(P, Q),                       // P ^ Q
						&Reiterate{Formula: Not(And(P, Q))}, // ~ ( P ^ Q )
					), // ~ Q
				), // P -> ~ Q
				ArrowProof(Not(Q),
					ContraProof(Not(P),
						&Reiterate{Formula: Not(Q)},                   // ~ Q
						IAnd(Not(P), Not(Q)),                       // ~ P ^ ~ Q
						&Reiterate{Formula: Not(And(Not(P), Not(Q)))}, // ~ ( ~ P ^ ~ Q )
					), // P
				), // ~ Q -> P
				IDArrow(P, Not(Q)),
			),
		),
	// TODO no distributive
	),
	NewProofsSection("miscellaneous",
		// TODO may be wrong ?  this isn't the implication distributive formula according to wikipedia
		RootProof("( ( P -> Q ) -> R ) -> ( ( P -> Q ) -> ( P -> R ) )",
			ArrowProof(Arrow(Arrow(P, Q), R),
				ArrowProof(Arrow(P, Q),
					&Reiterate{Formula: Arrow(Arrow(P, Q), R)},
					EImply(Arrow(P, Q), R),
					ArrowProof(P,
						&Reiterate{Formula: R},
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
				ArrowConjunctionTheorem(P, Q, true),                                 // P ^ ~ Q
				EAnd(P, Not(Q), false),                                              // ~ Q
				ArrowConjunctionTheorem(Q, R, true),                                 // Q ^ ~ R
				EAnd(Q, Not(R), true),                                               // Q
			), // ( P -> Q ) v ( Q -> R )
		),
		RootProof("( P -> Q ) v ( Q -> P )",
			ContraProof(Not(Or(Arrow(P, Q), Arrow(Q, P))),
				DeMorgansOrToAndTheorem(Arrow(P, Q), Arrow(Q, P), true), // ~ ( P -> Q ) ^ ~ ( Q -> P )
				EAnd(Not(Arrow(P, Q)), Not(Arrow(Q, P)), true),          // ~ ( P -> Q )
				EAnd(Not(Arrow(P, Q)), Not(Arrow(Q, P)), false),         // ~ ( Q -> P )
				ArrowConjunctionTheorem(P, Q, true),                                 // P ^ ~ Q
				EAnd(P, Not(Q), false),                                              // ~ Q
				ArrowConjunctionTheorem(Q, P, true),                                 // Q ^ ~ P
				EAnd(Q, Not(P), true),                                               // Q
			), // ( P -> Q ) v ( Q -> P )
		),
	),
}
