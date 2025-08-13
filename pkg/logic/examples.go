package logic

var (
	P = Var("P")
	Q = Var("Q")
	R = Var("R")
	S = Var("S")
)

var (
	notPAndNotP = NewProof("~ ( P ^ ~ P )",
		NewSubProofContradiction(
			And(P, Not(P)),
			EAnd(P, Not(P), true),
			EAnd(P, Not(P), false),
		))
	andCommutativity = NewProof("P ^ Q -> Q ^ P",
		NewSubProofImplication(
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

	pOrNotP = NewProof("P v ~ P",
		NewSubProofContradiction(
			Not(Or(P, Not(P))),
			NewSubProofContradiction(
				P,
				IOr(P, Not(P), true),
				&Reiterate{Term: Not(Or(P, Not(P)))},
			),
			IOr(P, Not(P), false),
		))
)

var examples = []*Proof{
	notPAndNotP,
	andCommutativity,
	pOrNotP,
}
