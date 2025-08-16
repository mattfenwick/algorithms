package logic

var (
	P = Var("P")
	Q = Var("Q")
	R = Var("R")
	S = Var("S")
)

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

var examples = []*Proof{
	pToP,
	notPAndNotP,
	andCommutativity,
	orCommutativity,
	pOrNotP,
	pQNotQNotP,
	notQNotPToPQ,
	qToPToQ,
	pToQToQToRToPToR,
	pAndPToP,
	pOrPToP,
}
