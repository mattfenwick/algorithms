package logic

func removeDoubleNegative(t Formula) Formula {
	switch a := t.(type) {
	case *NotFormula:
		switch b := a.Formula.(type) {
		case *NotFormula:
			return b.Formula
		}
	}
	return t
}

// P v ~ P
func ExcludedMiddleTheorem(t Formula) *Rule {
	return NewRule("Theorem: excluded middle", Or(t, Not(t)))
}

// ( P   ->   Q ) -> ( ~ Q -> ~ P )
// ( ~ P ->   Q ) -> ( ~ Q ->   P )
// ( P   -> ~ Q ) -> (   Q -> ~ P )
// ( ~ P -> ~ Q ) -> (   Q ->   P )
func ContrapositiveTheorem(p Formula, q Formula) *Rule {
	return NewRule("Theorem: contrapositive",
		Implication(removeDoubleNegative(Not(q)), removeDoubleNegative(Not(p))),
		Implication(p, q))
}

// (   P -> Q ) -> ( ~ P v Q )
// ( ~ P -> Q ) -> (   P v Q )
func ArrowDisjunctionTheorem(p Formula, q Formula) *Rule {
	return NewRule("Theorem: -> to v", Or(removeDoubleNegative(Not(p)), q), Implication(p, q))
}

// (   P v Q ) -> ( ~ P -> Q )
// ( ~ P v Q ) -> (   P -> Q )
func DisjunctionArrowTheorem(p Formula, q Formula) *Rule {
	return NewRule("Theorem: v to ->", Implication(removeDoubleNegative(p), q), Or(p, q))
}

// ArrowConjunctionTheorem handles these patterns:
//
//	  ~ (   P ->   Q ) ->   (   P ^ ~ Q )
//	  ~ ( ~ P ->   Q ) ->   ( ~ P ^ ~ Q )
//	  ~ (   P -> ~ Q ) ->   (   P ^   Q )
//	  ~ ( ~ P -> ~ Q ) ->   ( ~ P ^   Q )
//		(   P ->   Q ) -> ~ (   P ^ ~ Q )
//		( ~ P ->   Q ) -> ~ ( ~ P ^ ~ Q )
//		(   P -> ~ Q ) -> ~ (   P ^   Q )
//		( ~ P -> ~ Q ) -> ~ ( ~ P ^   Q )
func ArrowConjunctionTheorem(p Formula, q Formula, isArrowNegated bool) *Rule {
	if isArrowNegated {
		return NewRule("Theorem: (~ ->) to ^", And(p, removeDoubleNegative(Not(q))), Not(Implication(p, q)))
	} else {
		return NewRule("Theorem: -> to (~ ^)", Not(And(p, removeDoubleNegative(Not(q)))), Implication(p, q))
	}
}

// ConjunctionArrowTheorem handles these patterns:
//
//	  (   P ^ ~ Q ) -> ~ (   P ->   Q )
//	  ( ~ P ^ ~ Q ) -> ~ ( ~ P ->   Q )
//	  (   P ^   Q ) -> ~ (   P -> ~ Q )
//	  ( ~ P ^   Q ) -> ~ ( ~ P -> ~ Q )
//	~ (   P ^ ~ Q ) ->   (   P ->   Q )
//	~ ( ~ P ^ ~ Q ) ->   ( ~ P ->   Q )
//	~ (   P ^   Q ) ->   (   P -> ~ Q )
//	~ ( ~ P ^   Q ) ->   ( ~ P -> ~ Q )
func ConjunctionArrowTheorem(p Formula, q Formula, isConjunctionNegated bool) *Rule {
	if isConjunctionNegated {
		return NewRule("Theorem: (~ ^) to ->", Implication(p, removeDoubleNegative(Not(q))), Not(And(p, q)))
	} else {
		return NewRule("Theorem: ^ to (~ ->)", Not(Implication(p, removeDoubleNegative(Not(q)))), And(p, q))
	}
}

// ~ ( P <->   Q ) -> ( P <-> ~ Q )
// ~ ( P <-> ~ Q ) -> ( P <->   Q )
func BiconditionalNegationTheorem(p Formula, q Formula) *Rule {
	return NewRule("Theorem: biconditional negation",
		Biconditional(p, removeDoubleNegative(Not(q))),
		Not(Biconditional(p, q)))
}

// DeMorgansAndToOrTheorem handles:
//
//	  ~ (   P ^   Q ) ->   ( ~ P v ~ Q )
//	  ~ ( ~ P ^   Q ) ->   (   P v ~ Q )
//	  ~ (   P ^ ~ Q ) ->   ( ~ P v   Q )
//	  ~ ( ~ P ^ ~ Q ) ->   (   P v   Q )
//		(   P ^   Q ) -> ~ ( ~ P v ~ Q )
//		( ~ P ^   Q ) -> ~ (   P v ~ Q )
//		(   P ^ ~ Q ) -> ~ ( ~ P v   Q )
//		( ~ P ^ ~ Q ) -> ~ (   P v   Q )
func DeMorgansAndToOrTheorem(p Formula, q Formula, isAndNegated bool) *Rule {
	if isAndNegated {
		return NewRule("DeMorgan's theorem: (~ ^) to v",
			Or(removeDoubleNegative(Not(p)), removeDoubleNegative(Not(q))),
			Not(And(p, q)),
		)
	} else {
		return NewRule("DeMorgan's theorem: ^ to (~ v)",
			Not(Or(removeDoubleNegative(Not(p)), removeDoubleNegative(Not(q)))),
			And(p, q),
		)
	}
}

// DeMorgansOrToAndTheorem handles:
//
//	  ~ (   P v   Q ) ->   ( ~ P ^ ~ Q )
//	  ~ ( ~ P v   Q ) ->   (   P ^ ~ Q )
//	  ~ (   P v ~ Q ) ->   ( ~ P ^   Q )
//	  ~ ( ~ P v ~ Q ) ->   (   P ^   Q )
//		(   P v   Q ) -> ~ ( ~ P ^ ~ Q )
//		( ~ P v   Q ) -> ~ (   P ^ ~ Q )
//		(   P v ~ Q ) -> ~ ( ~ P ^   Q )
//		( ~ P v ~ Q ) -> ~ (   P ^   Q )
func DeMorgansOrToAndTheorem(p Formula, q Formula, isAndNegated bool) *Rule {
	if isAndNegated {
		return NewRule("DeMorgan's theorem: (~ v) to ^",
			And(removeDoubleNegative(Not(p)), removeDoubleNegative(Not(q))),
			Not(Or(p, q)),
		)
	} else {
		return NewRule("DeMorgan's theorem: v to (~ ^)",
			Not(And(removeDoubleNegative(Not(p)), removeDoubleNegative(Not(q)))),
			Or(p, q),
		)
	}
}

// skip -- can be handled instead by 2 step: arrow, eimply
//   ( P -> Q ) ^ ~ Q -> ~ P (modus tollens)
//   ( P v  Q ) ^ ~ P ->   Q
//   ( P v  Q ) ^ ~ Q ->   P
// ~ ( P ^  Q ) ^   P -> ~ Q
// ~ ( P ^  Q ) ^   Q -> ~ P

// used very often ?
// ( P -> Q ) ^ ( Q -> R ) -> ( P -> R )

// TODO not sure -- maybe use subproofs for ENot, INot?
// P <-> ~ ~ P
