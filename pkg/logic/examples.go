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
		NewProofImplication(Var("P"), &Repeat{Var("P")}))

	pToPAndPOrNotPTodoDelete = NewRootProof(
		"P -> P ^ P v ~P",
		NewProofImplication(Var("P"),
			IAnd(Var("P"), Var("P")),
			IOr(And(Var("P"), Var("P")), Not(Var("P")), true),
		),
	)

	notPAndNotP = NewRootProof("~ ( P ^ ~ P )",
		NewProofContradiction(
			And(P, Not(P)),
			EAnd(P, Not(P), true),
			EAnd(P, Not(P), false),
		))
	andCommutativity = NewRootProof("P ^ Q -> Q ^ P",
		NewProofImplication(
			And(P, Q),
			EAnd(P, Q, true),
			EAnd(P, Q, false),
			IAnd(Q, P),
		))
	// tempDelete = NewSubProofContradiction(
	// 	P,
	// 	IOr(P, Not(P), true),
	// 	&Reiterate{Term: Not(Or(P, Not(P)))},
	// )

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
)

var examples = []*Proof{
	pToP,
	pToPAndPOrNotPTodoDelete,
	notPAndNotP,
	andCommutativity,
	pOrNotP,
}
