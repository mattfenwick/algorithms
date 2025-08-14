package logic

type Rule struct {
	Preconditions []Term
	Result        Term
	Name          string
}

func (r *Rule) StepResult() Term {
	return r.Result
}

func (r *Rule) StepName() string {
	return r.Name
}

func (r *Rule) StandardForm() Term {
	pres, result := r.Preconditions, r.Result
	left := pres[0]
	for _, l := range pres[1:] {
		left = And(left, l)
	}
	return Implication(left, result)
}

func NewRule(name string, result Term, preconditions ...Term) *Rule {
	return &Rule{Preconditions: preconditions, Result: result, Name: name}
}

// I -> -- P, Q => P -> Q
func IImply(ifTerm Term, then Term) *Rule {
	return NewRule("I ->", Implication(ifTerm, then), ifTerm, then)
}

// E -> -- (P -> Q), P => Q
func EImply(ifTerm Term, then Term) *Rule {
	return NewRule("E ->", then, Implication(ifTerm, then), ifTerm)
}

// I ^ -- P, Q => P ^ Q
func IAnd(left Term, right Term) *Rule {
	return NewRule("I ^", And(left, right), left, right)
}

// E ^ -- Left and Right versions
//
//	Left : P ^ Q => P
//	Right: P ^ Q => Q
func EAnd(left Term, right Term, isLeft bool) *Rule {
	result := right
	name := "E ^ (R)"
	if isLeft {
		result = left
		name = "E ^ (L)"
	}
	return NewRule(name, result, And(left, right))
}

// I v -- Left and Right versions
//
//	Left: A -> A v B
//	Right: B -> A v B
func IOr(left Term, right Term, isLeft bool) *Rule {
	pre := right
	name := "I v (R)"
	if isLeft {
		pre = left
		name = "I v (L)"
	}
	return NewRule(name, Or(left, right), pre)
}

// E v -- P -> R, Q -> R, P v Q -> R
func EOr(if1 Term, if2 Term, then Term) *Rule {
	return NewRule("E v", then,
		Implication(if1, then),
		Implication(if2, then),
		Or(if1, if2))
}

// I ~ -- P -> ~~P
func INot(term Term) *Rule {
	return NewRule("I ~", Not(Not(term)), term)
}

// E ~ -- ~~P -> P
func ENot(term Term) *Rule {
	return NewRule("E ~", term, Not(Not(term)))
}
