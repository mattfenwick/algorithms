package logic

func ExcludedMiddleTheorem(t Term) *Rule {
	return NewRule("P v ~ P", Or(t, Not(t)))
}

func ContrapositiveTheorem(p Term, q Term) *Rule {
	return NewRule("( P -> Q ) -> ( ~ Q -> ~ P )", Implication(Not(q), Not(p)), Implication(p, q))
}

func ContrapositiveRTheorem(p Term, q Term) *Rule {
	return NewRule("( ~ Q -> ~ P ) -> ( P -> Q )", Implication(p, q), Implication(Not(q), Not(p)))
}

func DisjunctiveSyllogismTheorem(p Term, q Term) *Rule {
	return NewRule("( ( P v Q ) ^ ~ P ) -> Q", q, Or(p, q), Not(p))
}

func ArrowDisjunctionTheorem(p Term, q Term) *Rule {
	return NewRule("( P -> Q ) -> ( ~P v Q )", Or(Not(p), q), Implication(p, q))
}

func ArrowConjunctionTheorem(p Term, q Term) *Rule {
	return NewRule("( P -> Q ) -> ( P ^ ~ Q )", And(p, Not(q)), Implication(p, q))
}
