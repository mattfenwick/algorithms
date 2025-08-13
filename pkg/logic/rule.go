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

// E -> -- (P -> Q), P => Q
func ElimImplicationRule(ifTerm Term, then Term) *Rule {
	return NewRule(then, ifTerm, Implication(ifTerm, then))
}

// I -> -- P, Q => P -> Q
func IntroImplicationRule(ifTerm Term, then Term) *Rule {
	return NewRule(Implication(ifTerm, then), ifTerm, then)
}

// I ^ -- A, B => A ^ B
func IntroAndRule(left Term, right Term) *Rule {
	return NewRule(And(left, right), left, right)
}

// E ^ -- A ^ B => A; A ^ B => B;
func ElimAndRule(left Term, right Term, isLeft bool) *Rule {
	result := right
	if isLeft {
		result = left
	}
	return NewRule(result, And(left, right))
}

// I v -- A -> A v B; B -> A v B
func IntroOrRule(left Term, right Term, isLeft bool) *Rule {
	pre := right
	if isLeft {
		pre = left
	}
	return NewRule(Or(left, right), pre)
}

// E v -- P -> R, Q -> R, P v Q -> R
func ElimOrRule(if1 Term, if2 Term, then Term) *Rule {
	return NewRule(then,
		Implication(if1, then),
		Implication(if2, then),
		Or(if1, if2))
}

// I ~ -- P -> ~~P
func IntroNotRule(term Term) *Rule {
	return NewRule(Not(Not(term)), term)
}

// E ~ -- ~~P -> P
func ElimNotRule(term Term) *Rule {
	return NewRule(term, Not(Not(term)))
}
