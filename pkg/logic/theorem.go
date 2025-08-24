package logic

// P v ~ P
func ExcludedMiddleTheorem(t Term) *Rule {
	return NewRule("Theorem: excluded middle", Or(t, Not(t)))
}

// ( P -> Q ) -> ( ~ Q -> ~ P )
func ContrapositiveTheorem(p Term, q Term) *Rule {
	return NewRule("Theorem: contrapositive", Implication(Not(q), Not(p)), Implication(p, q))
}

// ( ~ Q -> ~ P ) -> ( P -> Q )
func ContrapositiveRTheorem(p Term, q Term) *Rule {
	return NewRule("Theorem: contrapositive (R)", Implication(p, q), Implication(Not(q), Not(p)))
}

// ( ( P v Q ) ^ ~ P ) -> Q
func DisjunctiveSyllogismTheorem(p Term, q Term) *Rule {
	return NewRule("Theorem: disjunctive syllogism", q, Or(p, q), Not(p))
}

// ( P -> Q ) -> ( ~P v Q )
func ArrowDisjunctionTheorem(p Term, q Term) *Rule {
	return NewRule("Theorem: arrow disjunction", Or(Not(p), q), Implication(p, q))
}

// ~ ( P -> Q ) -> ( P ^ ~ Q )
func ArrowNegationTheorem(p Term, q Term) *Rule {
	return NewRule("Theorem: arrow negation", And(p, Not(q)), Not(Implication(p, q)))
}

// ~ ( P <-> Q ) -> ( P <-> ~ Q )
func BiconditionalNegationTheorem(p Term, q Term) *Rule {
	return NewRule("Theorem: biconditional negation", Biconditional(p, Not(q)), Not(Biconditional(p, q)))
}
