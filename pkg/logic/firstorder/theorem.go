package logic

import (
	"github.com/pkg/errors"
)

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

// ~ ( P ^ ~ P )
func NonContradictionTheorem(t Formula) *Rule {
	return NewRule("Theorem: non-contradiction", Not(And(t, Not(t))))
}

// ( P   ->   Q ) -> ( ~ Q -> ~ P )
// ( ~ P ->   Q ) -> ( ~ Q ->   P )
// ( P   -> ~ Q ) -> (   Q -> ~ P )
// ( ~ P -> ~ Q ) -> (   Q ->   P )
func ContrapositiveTheorem(p Formula, q Formula) *Rule {
	return NewRule("Theorem: contrapositive",
		Arrow(removeDoubleNegative(Not(q)), removeDoubleNegative(Not(p))),
		Arrow(p, q))
}

// (   P -> Q ) -> ( ~ P v Q )
// ( ~ P -> Q ) -> (   P v Q )
func ArrowDisjunctionTheorem(p Formula, q Formula) *Rule {
	return NewRule("Theorem: -> to v", Or(removeDoubleNegative(Not(p)), q), Arrow(p, q))
}

// (   P v Q ) -> ( ~ P -> Q )
// ( ~ P v Q ) -> (   P -> Q )
func DisjunctionArrowTheorem(p Formula, q Formula) *Rule {
	return NewRule("Theorem: v to ->", Arrow(removeDoubleNegative(Not(p)), q), Or(p, q))
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
		return NewRule("Theorem: (~ ->) to ^", And(p, removeDoubleNegative(Not(q))), Not(Arrow(p, q)))
	} else {
		return NewRule("Theorem: -> to (~ ^)", Not(And(p, removeDoubleNegative(Not(q)))), Arrow(p, q))
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
		return NewRule("Theorem: (~ ^) to ->", Arrow(p, removeDoubleNegative(Not(q))), Not(And(p, q)))
	} else {
		return NewRule("Theorem: ^ to (~ ->)", Not(Arrow(p, removeDoubleNegative(Not(q)))), And(p, q))
	}
}

// ~ ( P <->   Q ) -> ( P <-> ~ Q )
// ~ ( P <-> ~ Q ) -> ( P <->   Q )
func BiconditionalNegationTheorem(p Formula, q Formula) *Rule {
	return NewRule("Theorem: <-> negation",
		DArrow(p, removeDoubleNegative(Not(q))),
		Not(DArrow(p, q)))
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

// DeMorgansExistToForallTheorem handles:
//
//	  ∃x.(   P(x) ) -> ~ ∀x.( ~ P(x) )
//	  ∃x.( ~ P(x) ) -> ~ ∀x.(   P(x) )
//	~ ∃x.(   P(x) ) ->   ∀x.( ~ P(x) )
//	~ ∃x.( ~ P(x) ) ->   ∀x.(   P(x) )
func DeMorgansExistToForallTheorem(exist *QuantifiedFormula, isExistNegated bool) *Rule {
	if exist.Quantifier != ExistentialQuantifier {
		panic(errors.Errorf("expected existential quantifier, got %s", exist.Quantifier))
	}
	if isExistNegated {
		return NewRule("DeMorgan's theorem: (~ ∃) to ∀",
			Forall(exist.Var, removeDoubleNegative(Not(exist.Body))),
			Not(exist),
		)
	} else {
		return NewRule("DeMorgan's theorem: ∃ to (~ ∀)",
			Not(Forall(exist.Var, removeDoubleNegative(Not(exist.Body)))),
			exist,
		)
	}
}

// DeMorgansForallToExistTheorem handles:
//
//	  ∀x.(   P(x) ) -> ~ ∃x.( ~ P(x) )
//	  ∀x.( ~ P(x) ) -> ~ ∃x.(   P(x) )
//	~ ∀x.(   P(x) ) ->   ∃x.( ~ P(x) )
//	~ ∀x.( ~ P(x) ) ->   ∃x.(   P(x) )
func DeMorgansForallToExistTheorem(forall *QuantifiedFormula, isForallNegated bool) *Rule {
	if forall.Quantifier != ForallQuantifier {
		panic(errors.Errorf("expected forall quantifier, got %s", forall.Quantifier))
	}
	if isForallNegated {
		return NewRule("DeMorgan's theorem: (~ ∀) to ∃",
			Exist(forall.Var, removeDoubleNegative(Not(forall.Body))),
			Not(forall),
		)
	} else {
		return NewRule("DeMorgan's theorem: ∀ to (~ ∃)",
			Not(Exist(forall.Var, removeDoubleNegative(Not(forall.Body)))),
			forall,
		)
	}
}

// ∀x.( P(x) ^ Q ) -> ( ∀x.P(x) ^ Q )
// ∀x.( P(x) v Q ) -> ( ∀x.P(x) v Q )
// ∀x.( Q ^ P(x) ) -> ( Q ^ ∀x.P(x) )
// ∀x.( Q v P(x) ) -> ( Q v ∀x.P(x) )
func ForallAndOrTheorem(boundVar string, connective BinOp, left Formula, right Formula, isQuantifierLeft bool) *Rule {
	return quantifierAndOrTheorem(boundVar, left, right, &quantifierAndOrOptions{
		quantifier:       ForallQuantifier,
		connective:       connective,
		isQuantifierLeft: isQuantifierLeft,
		isIntoQuantifier: false,
	})
}

// ( ∀x.P(x) ^ Q ) -> ∀x.( P(x) ^ Q )
// ( ∀x.P(x) v Q ) -> ∀x.( P(x) v Q )
// ( Q ^ ∀x.P(x) ) -> ∀x.( Q ^ P(x) )
// ( Q v ∀x.P(x) ) -> ∀x.( Q v P(x) )
func AndOrForallTheorem(boundVar string, connective BinOp, left Formula, right Formula, isQuantifierLeft bool) *Rule {
	return quantifierAndOrTheorem(boundVar, left, right, &quantifierAndOrOptions{
		quantifier:       ForallQuantifier,
		connective:       connective,
		isQuantifierLeft: isQuantifierLeft,
		isIntoQuantifier: true,
	})
}

// ∃x.( P(x) ^ Q ) -> ( ∃x.P(x) ^ Q )
// ∃x.( P(x) v Q ) -> ( ∃x.P(x) v Q )
// ∃x.( Q ^ P(x) ) -> ( Q ^ ∃x.P(x) )
// ∃x.( Q v P(x) ) -> ( Q v ∃x.P(x) )
func ExistAndOrTheorem(boundVar string, connective BinOp, left Formula, right Formula, isQuantifierLeft bool) *Rule {
	return quantifierAndOrTheorem(boundVar, left, right, &quantifierAndOrOptions{
		quantifier:       ExistentialQuantifier,
		connective:       connective,
		isQuantifierLeft: isQuantifierLeft,
		isIntoQuantifier: false,
	})
}

// ( ∃x.P(x) ^ Q ) -> ∃x.( P(x) ^ Q )
// ( ∃x.P(x) v Q ) -> ∃x.( P(x) v Q )
// ( Q ^ ∃x.P(x) ) -> ∃x.( Q ^ P(x) )
// ( Q v ∃x.P(x) ) -> ∃x.( Q v P(x) )
func AndOrExistTheorem(boundVar string, connective BinOp, left Formula, right Formula, isQuantifierLeft bool) *Rule {
	return quantifierAndOrTheorem(boundVar, left, right, &quantifierAndOrOptions{
		quantifier:       ExistentialQuantifier,
		connective:       connective,
		isQuantifierLeft: isQuantifierLeft,
		isIntoQuantifier: true,
	})
}

type quantifierAndOrOptions struct {
	quantifier       Quantifier
	connective       BinOp
	isQuantifierLeft bool
	isIntoQuantifier bool
}

// example
// ( Q ^ ∀x.P(x) ) -> ∀x.( Q ^ P(x) )
// - quantifier: ∀
// - connective: ^
// - isQuantifierLeft: false
// - isIntoQuantifier: true
//
// ∀x.( P(x) ^ Q ) -> ( ∀x.P(x) ^ Q )
// ∀x.( P(x) v Q ) -> ( ∀x.P(x) v Q )
// ∃x.( P(x) ^ Q ) -> ( ∃x.P(x) ^ Q )
// ∃x.( P(x) v Q ) -> ( ∃x.P(x) v Q )
// ( Q ^ ∀x.P(x) ) -> ∀x.( Q ^ P(x) )
// ( Q v ∀x.P(x) ) -> ∀x.( Q v P(x) )
// ( Q ^ ∃x.P(x) ) -> ∃x.( Q ^ P(x) )
// ( Q v ∃x.P(x) ) -> ∃x.( Q v P(x) )
// ∀x.( Q ^ P(x) ) -> ( Q ^ ∀x.P(x) )
// ∀x.( Q v P(x) ) -> ( Q v ∀x.P(x) )
// ∃x.( Q ^ P(x) ) -> ( Q ^ ∃x.P(x) )
// ∃x.( Q v P(x) ) -> ( Q v ∃x.P(x) )
// ( ∀x.P(x) ^ Q ) -> ∀x.( P(x) ^ Q )
// ( ∀x.P(x) v Q ) -> ∀x.( P(x) v Q )
// ( ∃x.P(x) ^ Q ) -> ∃x.( P(x) ^ Q )
// ( ∃x.P(x) v Q ) -> ∃x.( P(x) v Q )
func quantifierAndOrTheorem(boundVar string, left Formula, right Formula, options *quantifierAndOrOptions) *Rule {
	var connectiveF func(Formula, Formula) *BinOpFormula
	switch options.connective {
	case AndOp:
		connectiveF = And
	case OrOp:
		connectiveF = Or
	default:
		panic(errors.Errorf("invalid binop '%s' for quantifier ^/v theorem", options.connective))
	}
	var l, r Formula
	if options.isQuantifierLeft {
		l = NewQuantifiedFormula(boundVar, left, options.quantifier)
		r = right
	} else {
		l = left
		r = NewQuantifiedFormula(boundVar, right, options.quantifier)
	}
	quantifierFormula := NewQuantifiedFormula(boundVar, connectiveF(left, right), options.quantifier)
	connectiveFormula := connectiveF(l, r)

	var result, precondition Formula
	if options.isIntoQuantifier {
		precondition = connectiveFormula
		result = quantifierFormula
	} else {
		precondition = quantifierFormula
		result = connectiveFormula
	}
	return NewRule("quantifier ^/v theorem",
		result,
		precondition,
		Exist("x", Pred("T")))
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
