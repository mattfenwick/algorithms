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

	andCommutativity = NewRootProof("( P ^ Q ) -> ( Q ^ P )",
		NewProofImplication(
			And(P, Q),
			EAnd(P, Q, true),
			EAnd(P, Q, false),
			IAnd(Q, P),
		))

	orCommutativity = NewRootProof("( P v Q ) -> ( Q v P )",
		NewProofImplication(
			Or(P, Q),
			NewProofImplication(P,
				IOr(Q, P, false)),
			NewProofImplication(Q,
				IOr(Q, P, true)),
			EOr(P, Q, Or(Q, P)),
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

	pQNotQNotP = NewRootProof("( P -> Q ) -> ( ~ Q -> ~ P )",
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
	)

	qToPToQ = NewRootProof("Q -> ( P -> Q )",
		NewProofImplication(
			Q,
			NewProofImplication(P,
				&Reiterate{Term: Q})),
	)

	notQNotPToPQ = NewRootProof("( ~ Q -> ~ P ) -> ( P -> Q )",
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
	)

	pToQToQToRToPToR = NewRootProof("( ( P -> Q ) ^ ( Q -> R ) ) -> ( P -> R )",
		NewProofImplication(And(Implication(P, Q), Implication(Q, R)),
			NewProofImplication(P,
				&Reiterate{Term: And(Implication(P, Q), Implication(Q, R))},
				EAnd(Implication(P, Q), Implication(Q, R), true),
				EAnd(Implication(P, Q), Implication(Q, R), false),
				EImply(P, Q),
				EImply(Q, R),
			)),
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
		pQNotQNotP,
		notQNotPToPQ,
		qToPToQ,
		pAndPToP,
		pOrPToP,
		pToQToQToRToPToR,
	),
	NewProofsSection("commutativity",
		andCommutativity,
		orCommutativity,
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
	NewProofsSection("distributivity"),
}
