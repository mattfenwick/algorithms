package logic

type Rule struct {
	Preconditions []Term
	Result        Term
}

func (r *Rule) StepResult() Term {
	return r.Result
}

func (r *Rule) StandardForm() Term {
	pres, result := r.Preconditions, r.Result
	left := pres[0]
	for _, l := range pres[1:] {
		left = And(left, l)
	}
	return Implication(left, result)
}

func NewRule(result Term, preconditions ...Term) *Rule {
	return &Rule{Preconditions: preconditions, Result: result}
}

// I -> -- P, Q => P -> Q
func IImply(ifTerm Term, then Term) *Rule {
	return NewRule(Implication(ifTerm, then), ifTerm, then)
}

// E -> -- (P -> Q), P => Q
func EImply(ifTerm Term, then Term) *Rule {
	return NewRule(then, ifTerm, Implication(ifTerm, then))
}

// I ^ -- P, Q => P ^ Q
func IAnd(left Term, right Term) *Rule {
	return NewRule(And(left, right), left, right)
}

// E ^ -- P ^ Q => P; P ^ Q => Q;
func EAnd(left Term, right Term, isLeft bool) *Rule {
	result := right
	if isLeft {
		result = left
	}
	return NewRule(result, And(left, right))
}

// I v -- A -> A v B; B -> A v B
func IOr(left Term, right Term, isLeft bool) *Rule {
	pre := right
	if isLeft {
		pre = left
	}
	return NewRule(Or(left, right), pre)
}

// E v -- P -> R, Q -> R, P v Q -> R
func EOr(if1 Term, if2 Term, then Term) *Rule {
	return NewRule(then,
		Implication(if1, then),
		Implication(if2, then),
		Or(if1, if2))
}

// I ~ -- P -> ~~P
func INot(term Term) *Rule {
	return NewRule(Not(Not(term)), term)
}

// E ~ -- ~~P -> P
func ENot(term Term) *Rule {
	return NewRule(term, Not(Not(term)))
}
