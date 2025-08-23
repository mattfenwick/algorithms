package logic

// P v ~ P
func ExcludedMiddleTheorem(t Term) *Rule {
	return NewRule("P v ~ P", Or(t, Not(t)))
}

// ( P -> Q ) -> ( ~ Q -> ~ P )
func ContrapositiveTheorem(p Term, q Term) *Rule {
	return NewRule("( P -> Q ) -> ( ~ Q -> ~ P )", Implication(Not(q), Not(p)), Implication(p, q))
}

// ( ~ Q -> ~ P ) -> ( P -> Q )
func ContrapositiveRTheorem(p Term, q Term) *Rule {
	return NewRule("( ~ Q -> ~ P ) -> ( P -> Q )", Implication(p, q), Implication(Not(q), Not(p)))
}

// ( ( P v Q ) ^ ~ P ) -> Q
func DisjunctiveSyllogismTheorem(p Term, q Term) *Rule {
	return NewRule("( ( P v Q ) ^ ~ P ) -> Q", q, Or(p, q), Not(p))
}

// ( P -> Q ) -> ( ~P v Q )
func ArrowDisjunctionTheorem(p Term, q Term) *Rule {
	return NewRule("( P -> Q ) -> ( ~P v Q )", Or(Not(p), q), Implication(p, q))
}

// ~ ( P -> Q ) -> ( P ^ ~ Q )
func ArrowNegationTheorem(p Term, q Term) *Rule {
	return NewRule("~ ( P -> Q ) -> ( P ^ ~ Q )", And(p, Not(q)), Not(Implication(p, q)))
}
