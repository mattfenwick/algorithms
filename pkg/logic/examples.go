package logic

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

var (
	pToP = NewRootProof(
		"P -> P",
		NewProofImplication(P, &Repeat{P}))

	notPAndNotP = NewRootProof("~ ( P ^ ~ P )",
		NewProofContradiction(
			And(P, Not(P)),
			EAnd(P, Not(P), true),
			EAnd(P, Not(P), false),
		))

	pOrNotP = NewRootProof("P v ~ P",
		NewProofContradiction(
			Not(Or(P, Not(P))),
			NewProofContradiction(
				P,
				IOr(P, Not(P), true),
				&Reiterate{Term: Not(Or(P, Not(P)))},
			),
			IOr(P, Not(P), false),
		),
		ENot(Or(P, Not(P))),
	)

	pAndPToP = NewRootProof("( P ^ P ) -> P",
		NewProofImplication(And(P, P),
			EAnd(P, P, true)),
	)

	pOrPToP = NewRootProof("( P v P ) -> P",
		NewProofImplication(Or(P, P),
			NewProofImplication(P, &Repeat{Term: P}),
			EOr(P, P, P)),
	)
)

var proofSections = []*ProofsSection{
	NewProofsSection("basics",
		pToP,
		notPAndNotP,
		pOrNotP,
		pAndPToP,
		// TODO P -> ( P ^ P )
		// TODO P -> ( P v P )
		pOrPToP,
	),
	NewProofsSection("commutativity",
		NewRootProof("( P ^ Q ) -> ( Q ^ P )",
			NewProofImplication(
				And(P, Q),
				EAnd(P, Q, true),
				EAnd(P, Q, false),
				IAnd(Q, P),
			)),
		NewRootProof("( P v Q ) -> ( Q v P )",
			NewProofImplication(
				Or(P, Q),
				NewProofImplication(P,
					IOr(Q, P, false)),
				NewProofImplication(Q,
					IOr(Q, P, true)),
				EOr(P, Q, Or(Q, P)),
			)),
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
	),
	NewProofsSection("distributivity",
		NewRootProof("( P -> ( Q v R ) ) -> ( ( P -> Q ) v ( P -> R ) )",
			NewProofImplication(Implication(P, Or(Q, R)),
				NewProofImplication(P,
					&Reiterate{Implication(P, Or(Q, R))},
					EImply(P, Or(Q, R)),
					NewProofImplication(Q,
						&Reiterate{P},
						IImply(P, Q),
						IOr(Implication(P, Q), Implication(P, R), true),
					),
					NewProofImplication(R,
						&Reiterate{P},
						IImply(P, R),
						IOr(Implication(P, Q), Implication(P, R), false),
					),
					EOr(Q, R, Or(Implication(P, Q), Implication(P, R))),
				),
				NewProofImplication(Not(P),
					NewProofImplication(Not(Q),
						&Reiterate{Not(P)},
					),
					NewProofImplication(P,
						NewProofContradiction(Not(Q),
							&Reiterate{Implication(Not(Q), Not(P))},
							EImply(Not(Q), Not(P)),
							&Reiterate{P},
						),
						ENot(Q),
					),
					IOr(Implication(P, Q), Implication(P, R), true),
				),
				NewProofContradiction(
					Not(Or(P, Not(P))),
					NewProofContradiction(
						P,
						IOr(P, Not(P), true),
						&Reiterate{Term: Not(Or(P, Not(P)))},
					),
					IOr(P, Not(P), false),
				),
				ENot(Or(P, Not(P))),
				EOr(P, Not(P), Or(Implication(P, Q), Implication(P, R))),
			),
		),
	),
	NewProofsSection("arrows",
		NewRootProof("( P -> Q ) -> ( ~ Q -> ~ P )",
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
		),
		NewRootProof("( ~ Q -> ~ P ) -> ( P -> Q )",
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
					ENot(Q),
				),
			),
		),
		NewRootProof("Q -> ( P -> Q )",
			NewProofImplication(
				Q,
				NewProofImplication(P,
					&Reiterate{Term: Q})),
		),
		// TODO ~ P -> ( P -> Q )
		// TODO (P^Q)->R <-> P->(Q->R)
		// TODO ~PvQ <-> P->Q
		// TODO (P->Q)^(P->~Q)->~P
		// TODO (P->Q)^(~P->Q)->Q
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
	),
}
