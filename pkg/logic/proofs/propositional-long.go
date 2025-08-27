package proofs

import . "github.com/mattfenwick/algorithms/pkg/logic"

var (
	P = Var("P")
	Q = Var("Q")
	R = Var("R")
	S = Var("S")
)

type ProofsSection struct {
	Name   string
	Proofs []*Proof
}

func NewProofsSection(name string, proofs ...*Proof) *ProofsSection {
	return &ProofsSection{Name: name, Proofs: proofs}
}

var propositionalLongProofSections = []*ProofsSection{
	NewProofsSection("theorems",
		NewRootProof("P v ~ P",
			NewProofContradiction(
				Not(Or(P, Not(P))),
				NewProofContradiction(
					P,
					IOr(P, Not(P), true),
					&Reiterate{Term: Not(Or(P, Not(P)))},
				),
				IOr(P, Not(P), false),
			),
		),
		NewRootProof("( P -> Q ) <-> ( ~ Q -> ~ P )",
			NewProofImplication(
				Implication(P, Q),
				NewProofImplication(
					Not(Q),
					NewProofContradiction(
						P,
						&Reiterate{Term: Implication(P, Q)},
						EImply(P, Q),
						&Reiterate{Term: Not(Q)},
					),
				),
			),
			NewProofImplication(
				Implication(Not(Q), Not(P)),
				NewProofImplication(
					P,
					NewProofContradiction(
						Not(Q),
						&Reiterate{Term: Implication(Not(Q), Not(P))},
						EImply(Not(Q), Not(P)),
						&Reiterate{Term: P},
					),
				),
			),
			IBiconditional(Implication(P, Q), Implication(Not(Q), Not(P))),
		),
		NewRootProof("~ ( P -> Q ) -> ( P ^ ~ Q )",
			NewProofImplication(Not(Implication(P, Q)),
				NewProofContradiction(Not(P),
					NewProofImplication(Not(Q),
						&Reiterate{Term: Not(P)},
					), // ~ Q -> ~ P
					ContrapositiveTheorem(Not(Q), Not(P)), // P -> Q
					&Reiterate{Term: Not(Implication(P, Q))},
				), // P
				NewProofContradiction(Q,
					&Reiterate{Term: P},
					IImply(P, Q),
					&Reiterate{Term: Not(Implication(P, Q))},
				),
				IAnd(P, Not(Q)),
			),
		),
		NewRootProof("( P ^ ~ Q ) -> ~ ( P -> Q )",
			NewProofImplication(And(P, Not(Q)),
				NewProofContradiction(Implication(P, Q),
					&Reiterate{Term: And(P, Not(Q))},
					EAnd(P, Not(Q), true),
					EAnd(P, Not(Q), false),
					EImply(P, Q),
				),
			),
		),
		// NewRootProof("( P -> Q ) -> ( ~ P v Q )",
		// 	NewProofImplication(Implication(P, Q),
		// 		NewProofImplication(P,
		// 			&Reiterate{Term: Implication(P, Q)},
		// 			EImply(P, Q),
		// 			IOr(Not(P), Q, false),
		// 		),
		// 		NewProofImplication(Not(P),
		// 			IOr(Not(P), Q, true),
		// 		),
		// 		// P v ~ P
		// 		EOr(P, Not(P), Or(Not(P), Q)),
		// 	),
		// ),
		// NewRootProof("( ~ P v Q ) -> ( P -> Q )"),
		NewRootProof("( P v Q ) -> ( ~ P -> Q )",
			NewProofImplication(Or(P, Q),
				NewProofImplication(Q,
					NewProofImplication(Not(P),
						&Reiterate{Term: Q},
					),
				), // Q -> ( ~ P -> Q )
				NewProofImplication(P,
					NewProofImplication(Not(Q),
						&Reiterate{Term: P},
					), // ~ Q -> P
					ContrapositiveTheorem(Not(Q), P), // ~ P -> Q
				), // P -> ( ~ P -> Q )
				EOr(P, Q, Implication(Not(P), Q)),
			),
		),
		NewRootProof("( ~ P -> Q ) -> ( P v Q )",
			NewProofImplication(Implication(Not(P), Q),
				NewProofImplication(P,
					IOr(P, Q, true),
				),
				NewProofImplication(Not(P),
					&Reiterate{Term: Implication(Not(P), Q)},
					EImply(Not(P), Q),
					IOr(P, Q, false),
				),
				ExcludedMiddleTheorem(P),
				EOr(P, Not(P), Or(P, Q)),
			),
		),
		NewRootProof("~ ( P v Q ) -> ( ~ P ^ ~ Q )",
			NewProofImplication(Not(Or(P, Q)),
				NewProofContradiction(P,
					IOr(P, Q, true),
					&Reiterate{Term: Not(Or(P, Q))},
				),
				NewProofContradiction(Q,
					IOr(P, Q, false),
					&Reiterate{Term: Not(Or(P, Q))},
				),
				IAnd(Not(P), Not(Q)),
			),
		),
		NewRootProof("( ~ P ^ ~ Q ) -> ~ ( P v Q )",
			NewProofImplication(And(Not(P), Not(Q)),
				NewProofImplication(P,
					NewProofContradiction(And(Not(P), Not(Q)),
						EAnd(Not(P), Not(Q), true),
						&Reiterate{Term: P},
					),
				),
				NewProofImplication(Q,
					NewProofContradiction(And(Not(P), Not(Q)),
						EAnd(Not(P), Not(Q), false),
						&Reiterate{Term: Q},
					),
				),
				NewProofContradiction(Or(P, Q),
					&Reiterate{Term: Implication(P, Not(And(Not(P), Not(Q))))},
					&Reiterate{Term: Implication(Q, Not(And(Not(P), Not(Q))))},
					EOr(P, Q, Not(And(Not(P), Not(Q)))),
					&Reiterate{Term: And(Not(P), Not(Q))},
				),
			),
		),
		NewRootProof("~ ( P ^ Q ) -> ( ~ P v ~ Q )",
			NewProofImplication(Not(And(P, Q)),
				NewProofContradiction(Not(Or(Not(P), Not(Q))),
					NewProofContradiction(Not(P),
						IOr(Not(P), Not(Q), true),
						&Reiterate{Term: Not(Or(Not(P), Not(Q)))},
					),
					NewProofContradiction(Not(Q),
						IOr(Not(P), Not(Q), false),
						&Reiterate{Term: Not(Or(Not(P), Not(Q)))},
					),
					IAnd(P, Q),
					&Reiterate{Term: Not(And(P, Q))},
				),
			),
		),
		NewRootProof("( ~ P v ~ Q ) -> ~ ( P ^ Q )",
			NewProofImplication(Or(Not(P), Not(Q)),
				NewProofImplication(Not(P),
					NewProofContradiction(And(P, Q),
						EAnd(P, Q, true),
						&Reiterate{Term: Not(P)},
					),
				),
				NewProofImplication(Not(Q),
					NewProofContradiction(And(P, Q),
						EAnd(P, Q, false),
						&Reiterate{Term: Not(Q)},
					),
				),
				EOr(Not(P), Not(Q), Not(And(P, Q))),
			),
		),
	),
	NewProofsSection("basics",
		NewRootProof(
			"P -> P",
			NewProofImplication(P)),
		NewRootProof("~ ( P ^ ~ P )",
			NewProofContradiction(
				And(P, Not(P)),
				EAnd(P, Not(P), true),
				EAnd(P, Not(P), false),
			)),
		NewRootProof("( P ^ P ) -> P",
			NewProofImplication(And(P, P),
				EAnd(P, P, true)),
		),
		NewRootProof("P -> ( P ^ P )",
			NewProofImplication(P,
				IAnd(P, P),
			),
		),
		NewRootProof("( P v P ) -> P",
			NewProofImplication(Or(P, P),
				NewProofImplication(P),
				EOr(P, P, P)),
		),
		NewRootProof("P -> ( P v P )",
			NewProofImplication(P,
				IOr(P, P, true),
			),
		),
		NewRootProof("( P ^ Q ) -> ( P v Q )",
			NewProofImplication(And(P, Q),
				EAnd(P, Q, true),
				IOr(P, Q, true),
			),
		),
		NewRootProof("( P -> ~ P ) -> ~ P",
			NewProofImplication(Implication(P, Not(P)),
				NewProofContradiction(P,
					&Reiterate{Term: Implication(P, Not(P))},
					EImply(P, Not(P)),
				),
			),
		),
	),
	NewProofsSection("arrows",
		NewRootProof("Q -> ( P -> Q )",
			NewProofImplication(
				Q,
				NewProofImplication(P,
					&Reiterate{Term: Q})),
		),
		NewRootProof("~ P -> ( P -> Q )",
			NewProofImplication(Not(P),
				NewProofImplication(Not(Q),
					&Reiterate{Term: Not(P)},
				),
				ContrapositiveTheorem(Not(Q), Not(P)),
			),
		),
		NewRootProof("( ( P -> Q ) ^ ( Q -> R ) ) -> ( P -> R )",
			NewProofImplication(And(Implication(P, Q), Implication(Q, R)),
				NewProofImplication(P,
					&Reiterate{Term: And(Implication(P, Q), Implication(Q, R))},
					EAnd(Implication(P, Q), Implication(Q, R), true),
					EAnd(Implication(P, Q), Implication(Q, R), false),
					EImply(P, Q),
					EImply(Q, R),
				)),
		),
		NewRootProof("( ( P ^ Q ) -> R ) -> ( P -> ( Q -> R ) )",
			NewProofImplication(Implication(And(P, Q), R),
				NewProofImplication(P,
					NewProofImplication(Q,
						&Reiterate{Term: P},
						IAnd(P, Q),
						&Reiterate{Term: Implication(And(P, Q), R)},
						EImply(And(P, Q), R),
					),
				),
			),
		),
		NewRootProof("( P -> ( Q -> R ) ) -> ( ( P ^ Q ) -> R )",
			NewProofImplication(Implication(P, Implication(Q, R)),
				NewProofImplication(And(P, Q),
					EAnd(P, Q, true),
					EAnd(P, Q, false),
					&Reiterate{Term: Implication(P, Implication(Q, R))},
					EImply(P, Implication(Q, R)),
					EImply(Q, R),
				),
			),
		),
		NewRootProof("( ( P -> Q ) ^ ( P -> ~ Q ) ) -> ~ P",
			NewProofImplication(And(Implication(P, Q), Implication(P, Not(Q))),
				NewProofContradiction(P,
					&Reiterate{Term: And(Implication(P, Q), Implication(P, Not(Q)))},
					EAnd(Implication(P, Q), Implication(P, Not(Q)), true),
					EAnd(Implication(P, Q), Implication(P, Not(Q)), false),
					EImply(P, Q),
					EImply(P, Not(Q)),
				),
			),
		),
		NewRootProof("( ( P -> Q ) ^ ( ~ P -> Q ) ) -> Q",
			NewProofImplication(And(Implication(P, Q), Implication(Not(P), Q)),
				EAnd(Implication(P, Q), Implication(Not(P), Q), true),
				EAnd(Implication(P, Q), Implication(Not(P), Q), false),
				ExcludedMiddleTheorem(P),
				EOr(P, Not(P), Q),
			),
		),
		NewRootProof("( P -> ( Q -> R ) ) -> ( Q -> ( P -> R ) )",
			NewProofImplication(Implication(P, Implication(Q, R)),
				NewProofImplication(Q,
					NewProofImplication(P,
						&Reiterate{Term: Implication(P, Implication(Q, R))},
						EImply(P, Implication(Q, R)),
						&Reiterate{Term: Q},
						EImply(Q, R),
					),
				),
			),
		),
		NewRootProof("( P -> Q ) -> ( P -> ( Q v R ) )",
			NewProofImplication(Implication(P, Q),
				NewProofImplication(P,
					&Reiterate{Term: Implication(P, Q)},
					EImply(P, Q),
					IOr(Q, R, true),
				),
			),
		),
		NewRootProof("( ( P v Q ) -> R ) -> ( P -> R )",
			NewProofImplication(Implication(Or(P, Q), R),
				NewProofImplication(P,
					IOr(P, Q, true),
					&Reiterate{Term: Implication(Or(P, Q), R)},
					EImply(Or(P, Q), R),
				),
			),
		),
		NewRootProof("( P -> R ) -> ( ( P ^ Q ) -> R )",
			NewProofImplication(Implication(P, R),
				NewProofImplication(And(P, Q),
					&Reiterate{Term: Implication(P, R)},
					EAnd(P, Q, true),
					EImply(P, R),
				),
			),
		),
		NewRootProof("( P -> ( Q ^ R ) ) -> ( P -> Q )",
			NewProofImplication(Implication(P, And(Q, R)),
				NewProofImplication(P,
					&Reiterate{Term: Implication(P, And(Q, R))},
					EImply(P, And(Q, R)),
					EAnd(Q, R, true),
				),
			),
		),
	),
	NewProofsSection("commutativity",
		NewRootProof("( P ^ Q ) -> ( Q ^ P )",
			NewProofImplication(
				And(P, Q),
				EAnd(P, Q, true),
				EAnd(P, Q, false),
				IAnd(Q, P),
			),
		),
		NewRootProof("( P v Q ) -> ( Q v P )",
			NewProofImplication(
				Or(P, Q),
				NewProofImplication(P,
					IOr(Q, P, false)),
				NewProofImplication(Q,
					IOr(Q, P, true)),
				EOr(P, Q, Or(Q, P)),
			),
		),
		NewRootProof("( P <-> Q ) -> ( Q <-> P )",
			NewProofImplication(Biconditional(P, Q),
				EBiconditional(P, Q, true),
				EBiconditional(P, Q, false),
				IBiconditional(Q, P),
			),
		),
	),
	NewProofsSection("associativity",
		NewRootProof("( ( P ^ Q ) ^ R ) -> ( P ^ ( Q ^ R ) )",
			NewProofImplication(And(And(P, Q), R),
				EAnd(And(P, Q), R, true),
				EAnd(And(P, Q), R, false),
				EAnd(P, Q, true),
				EAnd(P, Q, false),
				IAnd(Q, R),
				IAnd(P, And(Q, R)),
			),
		),
		NewRootProof("( P ^ ( Q ^ R ) ) -> ( ( P ^ Q ) ^ R )",
			NewProofImplication(And(P, And(Q, R)),
				EAnd(P, And(Q, R), true),
				EAnd(P, And(Q, R), false),
				EAnd(Q, R, true),
				EAnd(Q, R, false),
				IAnd(P, Q),
				IAnd(And(P, Q), R),
			),
		),
		NewRootProof("( ( P v Q ) v R ) -> ( P v ( Q v R ) )",
			NewProofImplication(Or(Or(P, Q), R),
				NewProofImplication(Or(P, Q),
					NewProofImplication(P,
						IOr(P, Or(Q, R), true),
					),
					NewProofImplication(Q,
						IOr(Q, R, true),
						IOr(P, Or(Q, R), false),
					),
					EOr(P, Q, Or(P, Or(Q, R))),
				),
				NewProofImplication(R,
					IOr(Q, R, false),
					IOr(P, Or(Q, R), false),
				),
				EOr(Or(P, Q), R, Or(P, Or(Q, R))),
			),
		),
		NewRootProof("( P v ( Q v R ) ) -> ( ( P v Q ) v R )",
			NewProofImplication(Or(P, Or(Q, R)),
				NewProofImplication(P,
					IOr(P, Q, true),
					IOr(Or(P, Q), R, true),
				),
				NewProofImplication(Or(Q, R),
					NewProofImplication(Q,
						IOr(P, Q, false),
						IOr(Or(P, Q), R, true),
					),
					NewProofImplication(R,
						IOr(Or(P, Q), R, false),
					),
					EOr(Q, R, Or(Or(P, Q), R)),
				),
				EOr(P, Or(Q, R), Or(Or(P, Q), R)),
			),
		),
		NewRootProof("( ( P <-> Q ) <-> R ) -> ( P <-> ( Q <-> R ) )",
			NewProofImplication(Biconditional(Biconditional(P, Q), R),
				EBiconditional(Biconditional(P, Q), R, true),  // ( P <-> Q ) -> R
				EBiconditional(Biconditional(P, Q), R, false), // R -> ( P <-> Q )
				NewProofContradiction(Not(Implication(P, Biconditional(Q, R))),
					ArrowConjunctionTheorem(P, Biconditional(Q, R), true), // P ^ ~ ( Q <-> R )
					EAnd(P, Not(Biconditional(Q, R)), true),               // P
					EAnd(P, Not(Biconditional(Q, R)), false),              // ~ ( Q <-> R )
					BiconditionalNegationTheorem(Q, R),                    // Q <-> ~ R
					NewProofContradiction(R,
						&Reiterate{Term: Implication(R, Biconditional(P, Q))}, // R -> ( P <-> Q )
						EImply(R, Biconditional(P, Q)),                        // P <-> Q
						EBiconditional(P, Q, true),                            // P -> Q
						&Reiterate{Term: P},                                   // P
						EImply(P, Q),                                          // Q
						&Reiterate{Term: Biconditional(Q, Not(R))},            // Q <-> ~ R
						EBiconditional(Q, Not(R), true),                       // Q -> ~ R
						EImply(Q, Not(R)),                                     // ~ R
					), // ~ R
					EBiconditional(Q, Not(R), false),                      // ~ R -> Q
					EImply(Not(R), Q),                                     // Q
					&Reiterate{Term: Implication(Biconditional(P, Q), R)}, // ( P <-> Q ) -> R
					ContrapositiveTheorem(Biconditional(P, Q), R),         // ~ R -> ~ ( P <-> Q )
					EImply(Not(R), Not(Biconditional(P, Q))),              // ~ ( P <-> Q )
					BiconditionalNegationTheorem(P, Q),                    // P <-> ~ Q
					EBiconditional(P, Not(Q), true),                       // P -> ~ Q
					EImply(P, Not(Q)),                                     // ~ Q
				), // ( P -> ( Q <-> R ) )
				NewProofContradiction(Not(Implication(Biconditional(Q, R), P)),
					ArrowConjunctionTheorem(Biconditional(Q, R), P, true), // ( Q <-> R ) ^ ~ P
					EAnd(Biconditional(Q, R), Not(P), true),               // Q <-> R
					EAnd(Biconditional(Q, R), Not(P), false),              // ~ P
					NewProofContradiction(R,
						&Reiterate{Term: Implication(R, Biconditional(P, Q))}, // R -> ( P <-> Q )
						&Reiterate{Term: Biconditional(Q, R)},                 // Q <-> R
						EBiconditional(Q, R, false),                           // R -> Q
						EImply(R, Q),                                          // Q
						EImply(R, Biconditional(P, Q)),                        // P <-> Q
						EBiconditional(P, Q, false),                           // Q -> P
						EImply(Q, P),                                          // P
						&Reiterate{Term: Not(P)},                              // ~ P
					), // ~ R
					EBiconditional(Q, R, true),                            // Q -> R
					ContrapositiveTheorem(Q, R),                           // ~ R -> ~ Q
					EImply(Not(R), Not(Q)),                                // ~ Q
					&Reiterate{Term: Implication(Biconditional(P, Q), R)}, // ( P <-> Q ) -> R
					ContrapositiveTheorem(Biconditional(P, Q), R),         // ~ R -> ~ ( P <-> Q)
					EImply(Not(R), Not(Biconditional(P, Q))),              // ~ ( P <-> Q )
					BiconditionalNegationTheorem(P, Q),                    // P <-> ~ Q
					EBiconditional(P, Not(Q), false),                      // ~ Q -> P
					EImply(Not(Q), P),                                     // P
				), // ( ( Q <-> R ) -> P )
				IBiconditional(P, Biconditional(Q, R)), // P <-> ( Q <-> R )
			),
		),
		NewRootProof("( P <-> ( Q <-> R ) ) -> ( ( P <-> Q ) <-> R )",
			NewProofImplication(Biconditional(P, Biconditional(Q, R)),
				EBiconditional(P, Biconditional(Q, R), true),  // P -> ( Q <-> R )
				EBiconditional(P, Biconditional(Q, R), false), // ( Q <-> R ) -> P
				// prove: ( P <-> Q ) -> R
				NewProofContradiction(Not(Implication(Biconditional(P, Q), R)),
					ArrowConjunctionTheorem(Biconditional(P, Q), R, true), // ( P <-> Q ) ^ ~ R
					EAnd(Biconditional(P, Q), Not(R), true),               // P <-> Q
					EAnd(Biconditional(P, Q), Not(R), false),              // ~ R
					EBiconditional(P, Q, true),                            // P -> Q
					NewProofContradiction(P,
						&Reiterate{Term: Implication(P, Q)}, // P -> Q
						EImply(P, Q),                        // Q
						&Reiterate{Term: Implication(P, Biconditional(Q, R))}, // P -> ( Q <-> R )
						EImply(P, Biconditional(Q, R)),                        // Q <-> R
						EBiconditional(Q, R, true),                            // Q -> R
						EImply(Q, R),                                          // R
						&Reiterate{Term: Not(R)},                              // ~ R
					), // ~ P
					&Reiterate{Term: Implication(Biconditional(Q, R), P)}, // ( Q <-> R ) -> P
					ContrapositiveTheorem(Biconditional(Q, R), P),         // ~ P -> ~ ( Q <-> R )
					EImply(Not(P), Not(Biconditional(Q, R))),              // ~ ( Q <-> R )
					BiconditionalNegationTheorem(Q, R),                    // Q <-> ~ R
					EBiconditional(Q, Not(R), false),                      // ~ R -> Q
					EImply(Not(R), Q),                                     // Q
					EBiconditional(P, Q, false),                           // Q -> P
					EImply(Q, P),                                          // P
				), // ( ( P <-> Q ) -> R )
				NewProofContradiction(Not(Implication(R, Biconditional(P, Q))),
					ArrowConjunctionTheorem(R, Biconditional(P, Q), true), // R ^ ~ ( P <-> Q )
					EAnd(R, Not(Biconditional(P, Q)), true),               // R
					EAnd(R, Not(Biconditional(P, Q)), false),              // ~ ( P <-> Q )
					BiconditionalNegationTheorem(P, Q),                    // P <-> ~ Q
					EBiconditional(P, Not(Q), true),                       // P -> ~ Q
					NewProofContradiction(P,
						&Reiterate{Term: Implication(P, Not(Q))},              // P -> ~ Q
						EImply(P, Not(Q)),                                     // ~ Q
						&Reiterate{Term: Implication(P, Biconditional(Q, R))}, // P -> ( Q <-> R )
						EImply(P, Biconditional(Q, R)),                        // Q <-> R
						EBiconditional(Q, R, false),                           // R -> Q
						&Reiterate{Term: R},                                   // R
						EImply(R, Q),                                          // Q
					), // ~ P
					EBiconditional(P, Not(Q), false),                      // ~ Q -> P
					ContrapositiveTheorem(Not(Q), P),                      // ~ P -> Q
					EImply(Not(P), Q),                                     // Q
					&Reiterate{Term: Implication(Biconditional(Q, R), P)}, // ( Q <-> R ) -> P
					ContrapositiveTheorem(Biconditional(Q, R), P),         // ~ P -> ~ ( Q <-> R )
					EImply(Not(P), Not(Biconditional(Q, R))),              // ~ ( Q <-> R )
					BiconditionalNegationTheorem(Q, R),                    // Q <-> ~ R
					EBiconditional(Q, Not(R), true),                       // Q -> ~ R
					EImply(Q, Not(R)),                                     // ~ R
				), // ( R -> ( P <-> Q ) )
				IBiconditional(Biconditional(P, Q), R),
			),
		),
	),
	NewProofsSection("distributivity",
		NewRootProof("( P -> ( Q -> R ) ) -> ( ( P -> Q ) -> ( P -> R ) )",
			NewProofImplication(Implication(P, Implication(Q, R)),
				NewProofImplication(Implication(P, Q),
					NewProofImplication(P,
						&Reiterate{Term: Implication(P, Implication(Q, R))},
						EImply(P, Implication(Q, R)),
						&Reiterate{Term: Implication(P, Q)},
						EImply(P, Q),
						EImply(Q, R),
					),
				),
			),
		),
		NewRootProof("( ( P -> Q ) -> ( P -> R ) ) -> ( P -> ( Q -> R ) )",
			NewProofImplication(Implication(Implication(P, Q), Implication(P, R)),
				NewProofImplication(P,
					NewProofImplication(Q,
						&Reiterate{Term: Implication(Implication(P, Q), Implication(P, R))},
						&Reiterate{Term: P},
						IImply(P, Q),
						EImply(Implication(P, Q), Implication(P, R)),
						EImply(P, R),
					),
				),
			),
		),
		NewRootProof("( P -> ( Q v R ) ) -> ( ( P -> Q ) v ( P -> R ) )",
			NewProofImplication(Implication(P, Or(Q, R)),
				NewProofImplication(P,
					&Reiterate{Term: Implication(P, Or(Q, R))},
					EImply(P, Or(Q, R)),
					NewProofImplication(Q,
						&Reiterate{Term: P},
						IImply(P, Q),
						IOr(Implication(P, Q), Implication(P, R), true),
					),
					NewProofImplication(R,
						&Reiterate{Term: P},
						IImply(P, R),
						IOr(Implication(P, Q), Implication(P, R), false),
					),
					EOr(Q, R, Or(Implication(P, Q), Implication(P, R))),
				),
				NewProofImplication(Not(P),
					NewProofImplication(Not(Q),
						&Reiterate{Term: Not(P)},
					),
					ContrapositiveTheorem(Not(Q), Not(P)),
					IOr(Implication(P, Q), Implication(P, R), true),
				),
				ExcludedMiddleTheorem(P),
				EOr(P, Not(P), Or(Implication(P, Q), Implication(P, R))),
			),
		),
		NewRootProof("( ( P -> Q ) v ( P -> R ) ) -> ( P -> ( Q v R ) )",
			NewProofImplication(Or(Implication(P, Q), Implication(P, R)),
				NewProofImplication(P,
					NewProofImplication(Implication(P, Q),
						&Reiterate{Term: P},
						EImply(P, Q),
						IOr(Q, R, true),
					),
					NewProofImplication(Implication(P, R),
						&Reiterate{Term: P},
						EImply(P, R),
						IOr(Q, R, false),
					),
					&Reiterate{Term: Or(Implication(P, Q), Implication(P, R))},
					EOr(Implication(P, Q), Implication(P, R), Or(Q, R)),
				),
			),
		),
		NewRootProof("( P -> ( Q ^ R ) ) -> ( ( P -> Q ) ^ ( P -> R ) )",
			NewProofImplication(Implication(P, And(Q, R)),
				NewProofImplication(P,
					&Reiterate{Term: Implication(P, And(Q, R))},
					EImply(P, And(Q, R)),
					EAnd(Q, R, true),
				),
				NewProofImplication(P,
					&Reiterate{Term: Implication(P, And(Q, R))},
					EImply(P, And(Q, R)),
					EAnd(Q, R, false),
				),
				IAnd(Implication(P, Q), Implication(P, R)),
			),
		),
		NewRootProof("( ( P -> Q ) ^ ( P -> R ) ) -> ( P -> ( Q ^ R ) )",
			NewProofImplication(And(Implication(P, Q), Implication(P, R)),
				NewProofImplication(P,
					&Reiterate{Term: And(Implication(P, Q), Implication(P, R))},
					EAnd(Implication(P, Q), Implication(P, R), true),
					EAnd(Implication(P, Q), Implication(P, R), false),
					EImply(P, Q),
					EImply(P, R),
					IAnd(Q, R),
				),
			),
		),
		NewRootProof("( P ^ ( Q v R ) ) -> ( ( P ^ Q ) v ( P ^ R ) )",
			NewProofImplication(And(P, Or(Q, R)),
				EAnd(P, Or(Q, R), true),
				EAnd(P, Or(Q, R), false),
				NewProofImplication(Q,
					&Reiterate{Term: P},
					IAnd(P, Q),
					IOr(And(P, Q), And(P, R), true),
				),
				NewProofImplication(R,
					&Reiterate{Term: P},
					IAnd(P, R),
					IOr(And(P, Q), And(P, R), false),
				),
				EOr(Q, R, Or(And(P, Q), And(P, R))),
			),
		),
		NewRootProof("( ( P ^ Q ) v ( P ^ R ) ) -> ( P ^ ( Q v R ) )",
			NewProofImplication(Or(And(P, Q), And(P, R)),
				NewProofImplication(And(P, Q),
					EAnd(P, Q, true),
					EAnd(P, Q, false),
					IOr(Q, R, true),
					IAnd(P, Or(Q, R)),
				),
				NewProofImplication(And(P, R),
					EAnd(P, R, true),
					EAnd(P, R, false),
					IOr(Q, R, false),
					IAnd(P, Or(Q, R)),
				),
				EOr(And(P, Q), And(P, R), And(P, Or(Q, R))),
			),
		),
		NewRootProof("( P v ( Q ^ R ) ) -> ( ( P v Q ) ^ ( P v R ) )",
			NewProofImplication(Or(P, And(Q, R)),
				NewProofImplication(P,
					IOr(P, Q, true),
					IOr(P, R, true),
					IAnd(Or(P, Q), Or(P, R)),
				),
				NewProofImplication(And(Q, R),
					EAnd(Q, R, true),
					EAnd(Q, R, false),
					IOr(P, Q, false),
					IOr(P, R, false),
					IAnd(Or(P, Q), Or(P, R)),
				),
				EOr(P, And(Q, R), And(Or(P, Q), Or(P, R))),
			),
		),
		NewRootProof("( ( P v Q ) ^ ( P v R ) ) -> ( P v ( Q ^ R ) )",
			NewProofImplication(And(Or(P, Q), Or(P, R)),
				EAnd(Or(P, Q), Or(P, R), true),
				EAnd(Or(P, Q), Or(P, R), false),
				NewProofImplication(P,
					IOr(P, And(Q, R), true),
				),
				NewProofImplication(P,
					NewProofImplication(Not(Q),
						&Reiterate{Term: P},
					),
					ContrapositiveTheorem(Not(Q), P),
				), // P -> ( ~ P -> Q )
				NewProofImplication(P,
					NewProofImplication(Not(R),
						&Reiterate{Term: P},
					),
					ContrapositiveTheorem(Not(R), P),
				), // P -> ( ~ P -> R )
				NewProofImplication(Not(P),
					&Reiterate{Term: Or(P, Q)},
					&Reiterate{Term: Or(P, R)},
					NewProofImplication(Q,
						&Reiterate{Term: Not(P)},
						IImply(Not(P), Q),
					), // Q -> ( ~ P -> Q )
					&Reiterate{Term: Implication(P, Implication(Not(P), Q))},
					EOr(P, Q, Implication(Not(P), Q)),
					EImply(Not(P), Q),
					NewProofImplication(R,
						&Reiterate{Term: Not(P)},
						IImply(Not(P), R),
					), // R -> ( ~ P -> R )
					&Reiterate{Term: Implication(P, Implication(Not(P), R))},
					EOr(P, R, Implication(Not(P), R)),
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
		NewRootProof("( ( ( P -> R ) ^ ( Q -> S ) ) ^ ( P v Q ) ) -> ( R v S )",
			NewProofImplication(And(And(Implication(P, R), Implication(Q, S)), Or(P, Q)),
				EAnd(And(Implication(P, R), Implication(Q, S)), Or(P, Q), true),
				EAnd(And(Implication(P, R), Implication(Q, S)), Or(P, Q), false),
				EAnd(Implication(P, R), Implication(Q, S), true),
				EAnd(Implication(P, R), Implication(Q, S), false),
				NewProofImplication(P,
					&Reiterate{Term: Implication(P, R)},
					EImply(P, R),
					IOr(R, S, true),
				),
				NewProofImplication(Q,
					&Reiterate{Term: Implication(Q, S)},
					EImply(Q, S),
					IOr(R, S, false),
				),
				EOr(P, Q, Or(R, S)),
			),
		),
		NewRootProof("( ( ( P -> R ) ^ ( Q -> S ) ) ^ ( ~ R v Q ) ) -> ( ~ P v S )",
			NewProofImplication(And(And(Implication(P, R), Implication(Q, S)), Or(Not(R), Q)),
				EAnd(And(Implication(P, R), Implication(Q, S)), Or(Not(R), Q), true),
				EAnd(And(Implication(P, R), Implication(Q, S)), Or(Not(R), Q), false), // ~ R v Q
				EAnd(Implication(P, R), Implication(Q, S), true),                      // P -> R
				EAnd(Implication(P, R), Implication(Q, S), false),                     // Q -> S
				NewProofImplication(Not(R),
					&Reiterate{Term: Implication(P, R)},
					ContrapositiveTheorem(P, R), // ~ R -> ~ P
					EImply(Not(R), Not(P)),      // ~ P
					IOr(Not(P), S, true),
				), // ~ R -> ( ~ P v S )
				NewProofImplication(Q,
					&Reiterate{Term: Implication(Q, S)},
					EImply(Q, S),
					IOr(Not(P), S, false),
				), // Q -> ( ~ P v S )
				EOr(Not(R), Q, Or(Not(P), S)),
			),
		),
		NewRootProof("( ( ( P -> R ) ^ ( Q -> S ) ) ^ ( ~ R v ~ S ) ) -> ( ~ P v ~ Q )",
			NewProofImplication(And(And(Implication(P, R), Implication(Q, S)), Or(Not(R), Not(S))),
				EAnd(And(Implication(P, R), Implication(Q, S)), Or(Not(R), Not(S)), true),
				EAnd(And(Implication(P, R), Implication(Q, S)), Or(Not(R), Not(S)), false), // ~ R v ~ S
				EAnd(Implication(P, R), Implication(Q, S), true),                           // P -> R
				EAnd(Implication(P, R), Implication(Q, S), false),                          // Q -> S
				NewProofImplication(Not(R),
					&Reiterate{Term: Implication(P, R)},
					ContrapositiveTheorem(P, R), // ~ R -> ~ P
					EImply(Not(R), Not(P)),      // ~ P
					IOr(Not(P), Not(Q), true),
				), // ~ R -> ( ~ P v ~ Q )
				NewProofImplication(Not(S),
					&Reiterate{Term: Implication(Q, S)},
					ContrapositiveTheorem(Q, S), // ~ S -> ~ Q
					EImply(Not(S), Not(Q)),      // ~ Q
					IOr(Not(P), Not(Q), false),
				), // ~ S -> ( ~ P v ~ Q )
				EOr(Not(R), Not(S), Or(Not(P), Not(Q))),
			),
		),
		NewRootProof("( ( ( P -> R ) v ( Q -> R ) ) ^ ( P ^ Q ) ) -> R",
			NewProofImplication(And(Or(Implication(P, R), Implication(Q, R)), And(P, Q)),
				EAnd(Or(Implication(P, R), Implication(Q, R)), And(P, Q), true),  // ( P -> R ) v ( Q -> R)
				EAnd(Or(Implication(P, R), Implication(Q, R)), And(P, Q), false), // P ^ Q
				EAnd(P, Q, true),  // P
				EAnd(P, Q, false), // Q
				NewProofImplication(Implication(P, R),
					&Reiterate{Term: P},
					EImply(P, R),
				), // ( P -> R ) -> R
				NewProofImplication(Implication(Q, R),
					&Reiterate{Term: Q},
					EImply(Q, R),
				), // ( Q -> R ) -> R
				EOr(Implication(P, R), Implication(Q, R), R),
			),
		),
		// TODO ( ( ( P -> R ) v ( Q -> S ) ) ^ ( P ^ Q ) ) -> ( R v S )
		NewRootProof("( ( ( P -> R ) v ( Q -> R ) ) ^ ~ R ) -> ( ~ P v ~ Q )",
			NewProofImplication(And(Or(Implication(P, R), Implication(Q, R)), Not(R)),
				EAnd(Or(Implication(P, R), Implication(Q, R)), Not(R), true),  // ( P -> R ) v ( Q -> R)
				EAnd(Or(Implication(P, R), Implication(Q, R)), Not(R), false), // ~ R
				NewProofImplication(Implication(P, R),
					&Reiterate{Term: Not(R)},
					ContrapositiveTheorem(P, R),
					EImply(Not(R), Not(P)),
					IOr(Not(P), Not(Q), true),
				), // ( P -> R ) -> ( ~ P v ~ Q )
				NewProofImplication(Implication(Q, R),
					&Reiterate{Term: Not(R)},
					ContrapositiveTheorem(Q, R),
					EImply(Not(R), Not(Q)),
					IOr(Not(P), Not(Q), false),
				), // ( Q -> R ) -> ( ~ P v ~ Q )
				EOr(Implication(P, R), Implication(Q, R), Or(Not(P), Not(Q))),
			),
		),
		// TODO ( ( ( P -> R ) v ( Q -> S ) ) ^ ( P ^ ~ S ) ) -> ( R v ~ Q )
		// TODO ( ( ( P -> R ) v ( Q -> S ) ) ^ ( ~ R ^ ~ S ) ) -> ( ~ P v ~ Q )
	),

	NewProofsSection("biconditional",
		NewRootProof("( ( P <-> Q ) ^ P ) -> Q",
			NewProofImplication(And(Biconditional(P, Q), P),
				EAnd(Biconditional(P, Q), P, true),
				EAnd(Biconditional(P, Q), P, false),
				EBiconditional(P, Q, true),
				EImply(P, Q),
			),
		),
		NewRootProof("( ( P <-> Q ) ^ ~ P ) -> ~ Q",
			NewProofImplication(And(Biconditional(P, Q), Not(P)),
				EAnd(Biconditional(P, Q), Not(P), true),
				EAnd(Biconditional(P, Q), Not(P), false),
				EBiconditional(P, Q, false), // Q -> P
				NewProofContradiction(Q,
					&Reiterate{Term: Not(P)},
					&Reiterate{Term: Implication(Q, P)},
					EImply(Q, P),
				),
			),
		),
		NewRootProof("( P <-> ~ Q ) -> ( ~ P <-> Q )",
			NewProofImplication(Biconditional(P, Not(Q)),
				EBiconditional(P, Not(Q), true),  // P -> ~ Q
				ContrapositiveTheorem(P, Not(Q)), // Q -> ~ P
				EBiconditional(P, Not(Q), false), // ~ Q -> P
				ContrapositiveTheorem(Not(Q), P), // ~ P -> Q
				IBiconditional(Not(P), Q),
			),
		),
		NewRootProof("( P <-> ~ Q ) -> ~ ( P <-> Q )",
			NewProofImplication(Biconditional(P, Not(Q)),
				EBiconditional(P, Not(Q), true),  // P -> ~ Q
				EBiconditional(P, Not(Q), false), // ~ Q -> P
				NewProofImplication(P,
					NewProofContradiction(Biconditional(P, Q),
						&Reiterate{Term: Implication(P, Not(Q))}, // P -> ~ Q
						EBiconditional(P, Q, true),               // P -> Q
						&Reiterate{Term: P},                      // P
						EImply(P, Q),                             // Q
						EImply(P, Not(Q)),                        // ~ Q
					), // ~ ( P <-> Q )
				), // P -> ~ ( P <-> Q )
				NewProofImplication(Not(P),
					NewProofContradiction(Biconditional(P, Q),
						EBiconditional(P, Q, false),              // Q -> P
						ContrapositiveTheorem(Q, P),              // ~ P -> ~ Q
						&Reiterate{Term: Not(P)},                 // ~ P
						EImply(Not(P), Not(Q)),                   // ~ Q
						&Reiterate{Term: Implication(Not(Q), P)}, // ~ Q -> P
						EImply(Not(Q), P),                        // P
					), // ~ ( P <-> Q )
				), // ~ P -> ~ ( P <-> Q )
				ExcludedMiddleTheorem(P), // P v ~ P
				EOr(P, Not(P), Not(Biconditional(P, Q))),
			),
		),
		NewRootProof("~ ( P <-> Q ) -> ( P <-> ~ Q )",
			NewProofImplication(Not(Biconditional(P, Q)),
				NewProofContradiction(Not(Implication(P, Not(Q))),
					ArrowConjunctionTheorem(P, Not(Q), true),
					EAnd(P, Q, true),
					EAnd(P, Q, false),
					IImply(P, Q),
					IImply(Q, P),
					IBiconditional(P, Q),
					&Reiterate{Term: Not(Biconditional(P, Q))},
				),
				NewProofContradiction(Not(Implication(Not(Q), P)),
					ArrowConjunctionTheorem(Not(Q), P, true),
					EAnd(Not(Q), Not(P), true),
					EAnd(Not(Q), Not(P), false),
					IImply(Not(P), Not(Q)),
					ContrapositiveTheorem(Not(P), Not(Q)),
					IImply(Not(Q), Not(P)),
					ContrapositiveTheorem(Not(Q), Not(P)),
					IBiconditional(P, Q),
					&Reiterate{Term: Not(Biconditional(P, Q))},
				),
				IBiconditional(P, Not(Q)),
			),
		),
		NewRootProof("( P <-> Q ) -> ( ~ P <-> ~ Q )",
			NewProofImplication(Biconditional(P, Q),
				EBiconditional(P, Q, true),
				ContrapositiveTheorem(P, Q),
				EBiconditional(P, Q, false),
				ContrapositiveTheorem(Q, P),
				IBiconditional(Not(P), Not(Q)),
			),
		),
		NewRootProof("( ~ P <-> ~ Q ) -> ( P <-> Q )",
			NewProofImplication(Biconditional(Not(P), Not(Q)),
				EBiconditional(Not(P), Not(Q), true),
				ContrapositiveTheorem(Not(P), Not(Q)),
				EBiconditional(Not(P), Not(Q), false),
				ContrapositiveTheorem(Not(Q), Not(P)),
				IBiconditional(P, Q),
			),
		),
		NewRootProof("( P <-> Q ) v ~ ( P <-> Q )",
			NewProofContradiction(Not(Or(Biconditional(P, Q), Not(Biconditional(P, Q)))),
				NewProofContradiction(Biconditional(P, Q),
					IOr(Biconditional(P, Q), Not(Biconditional(P, Q)), true),
					&Reiterate{Term: Not(Or(Biconditional(P, Q), Not(Biconditional(P, Q))))},
				),
				IOr(Biconditional(P, Q), Not(Biconditional(P, Q)), false),
			),
		),
		NewRootProof("( P <-> ~ Q ) -> ( ( P v Q ) ^ ( ~ P v ~ Q ) )",
			NewProofImplication(Biconditional(P, Not(Q)),
				EBiconditional(P, Not(Q), true),  // P -> ~ Q
				EBiconditional(P, Not(Q), false), // ~ Q -> P
				ContrapositiveTheorem(Not(Q), P), // ~ P -> Q
				NewProofImplication(P,
					IOr(P, Q, true),                          // P v Q
					&Reiterate{Term: Implication(P, Not(Q))}, // P -> ~ Q
					EImply(P, Not(Q)),                        // ~ Q
					IOr(Not(P), Not(Q), false),               // ~ P v ~ Q
					IAnd(Or(P, Q), Or(Not(P), Not(Q))),       // ( P v Q ) ^ ( ~ P v ~ Q )
				), // P -> ( ( P v Q ) ^ ( ~ P v ~ Q ) )
				NewProofImplication(Not(P),
					IOr(Not(P), Not(Q), true),                // ~ P v ~ Q
					&Reiterate{Term: Implication(Not(P), Q)}, // ~ P -> Q
					EImply(Not(P), Q),                        // Q
					IOr(P, Q, false),                         // P v Q
					IAnd(Or(P, Q), Or(Not(P), Not(Q))),       // ( P v Q ) ^ ( ~ P v ~ Q )
				), // ~ P -> ( ( P v Q ) ^ ( ~ P v ~ Q ) )
				ExcludedMiddleTheorem(P), // P v ~ P
				EOr(P, Not(P), And(Or(P, Q), Or(Not(P), Not(Q)))),
			),
		),
		NewRootProof("( ( P v Q ) ^ ( ~ P v ~ Q ) ) -> ( P <-> ~ Q )",
			NewProofImplication(And(Or(P, Q), Or(Not(P), Not(Q))),
				EAnd(Or(P, Q), Or(Not(P), Not(Q)), true),  // P v Q
				EAnd(Or(P, Q), Or(Not(P), Not(Q)), false), // ~ P v ~ Q
				NewProofImplication(Not(P),
					NewProofContradiction(And(P, Q),
						EAnd(P, Q, true),         // P
						&Reiterate{Term: Not(P)}, // ~ P
					), // ~ ( P ^ Q )
				), // ~ P -> ~ ( P ^ Q )
				NewProofImplication(Not(Q),
					NewProofContradiction(And(P, Q),
						EAnd(P, Q, false),        // Q
						&Reiterate{Term: Not(Q)}, // ~ Q
					), // ~ ( P ^ Q )
				), // ~ Q -> ~ ( P ^ Q )
				EOr(Not(P), Not(Q), Not(And(P, Q))), // ~ ( P ^ Q )
				NewProofImplication(P,
					NewProofContradiction(And(Not(P), Not(Q)),
						EAnd(Not(P), Not(Q), true), // ~ P
						&Reiterate{Term: P},        // P
					), // ~ ( ~ P ^ ~ Q )
				), // P -> ~ ( ~ P ^ ~ Q )
				NewProofImplication(Q,
					NewProofContradiction(And(Not(P), Not(Q)),
						EAnd(Not(P), Not(Q), false), // ~ Q
						&Reiterate{Term: Q},         // Q
					), // ~ ( ~ P ^ ~ Q )
				), // Q -> ~ ( ~ P ^ ~ Q )
				EOr(P, Q, Not(And(Not(P), Not(Q)))), // ~ ( ~ P ^ ~ Q )
				NewProofImplication(P,
					NewProofContradiction(Q,
						&Reiterate{Term: P},              // P
						IAnd(P, Q),                       // P ^ Q
						&Reiterate{Term: Not(And(P, Q))}, // ~ ( P ^ Q )
					), // ~ Q
				), // P -> ~ Q
				NewProofImplication(Not(Q),
					NewProofContradiction(Not(P),
						&Reiterate{Term: Not(Q)},                   // ~ Q
						IAnd(Not(P), Not(Q)),                       // ~ P ^ ~ Q
						&Reiterate{Term: Not(And(Not(P), Not(Q)))}, // ~ ( ~ P ^ ~ Q )
					), // P
				), // ~ Q -> P
				IBiconditional(P, Not(Q)),
			),
		),
	// TODO no distributive
	),
	NewProofsSection("miscellaneous",
		// TODO may be wrong ?  this isn't the implication distributive formula according to wikipedia
		NewRootProof("( ( P -> Q ) -> R ) -> ( ( P -> Q ) -> ( P -> R ) )",
			NewProofImplication(Implication(Implication(P, Q), R),
				NewProofImplication(Implication(P, Q),
					&Reiterate{Term: Implication(Implication(P, Q), R)},
					EImply(Implication(P, Q), R),
					NewProofImplication(P,
						&Reiterate{Term: R},
					),
				),
			),
		),
		// TODO may be wrong ?  this isn't the implication distributive formula according to wikipedia
		// NewRootProof("( ( P -> Q ) -> ( P -> R ) ) -> ( ( P -> Q ) -> R )",
		// 	NewProofImplication(Implication(Implication(P, Q), Implication(P, R)),
		// 		NewProofImplication(Implication(P, Q)),
		// 	),
		// ),
		NewRootProof("( P -> Q ) v ( Q -> R )",
			NewProofContradiction(Not(Or(Implication(P, Q), Implication(Q, R))),
				DeMorgansOrToAndTheorem(Implication(P, Q), Implication(Q, R), true), // ~ ( P -> Q ) ^ ~ ( Q -> R )
				EAnd(Not(Implication(P, Q)), Not(Implication(Q, R)), true),          // ~ ( P -> Q )
				EAnd(Not(Implication(P, Q)), Not(Implication(Q, R)), false),         // ~ ( Q -> R )
				ArrowConjunctionTheorem(P, Q, true),                                 // P ^ ~ Q
				EAnd(P, Not(Q), false),                                              // ~ Q
				ArrowConjunctionTheorem(Q, R, true),                                 // Q ^ ~ R
				EAnd(Q, Not(R), true),                                               // Q
			), // ( P -> Q ) v ( Q -> R )
		),
		NewRootProof("( P -> Q ) v ( Q -> P )",
			NewProofContradiction(Not(Or(Implication(P, Q), Implication(Q, P))),
				DeMorgansOrToAndTheorem(Implication(P, Q), Implication(Q, P), true), // ~ ( P -> Q ) ^ ~ ( Q -> P )
				EAnd(Not(Implication(P, Q)), Not(Implication(Q, P)), true),          // ~ ( P -> Q )
				EAnd(Not(Implication(P, Q)), Not(Implication(Q, P)), false),         // ~ ( Q -> P )
				ArrowConjunctionTheorem(P, Q, true),                                 // P ^ ~ Q
				EAnd(P, Not(Q), false),                                              // ~ Q
				ArrowConjunctionTheorem(Q, P, true),                                 // Q ^ ~ P
				EAnd(Q, Not(P), true),                                               // Q
			), // ( P -> Q ) v ( Q -> P )
		),
	),
}
