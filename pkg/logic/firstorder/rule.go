package logic

type Rule struct {
	Preconditions []Formula
	Result        Formula
	Name          string
}

func (r *Rule) StepResult() Formula {
	return r.Result
}

func (r *Rule) StepName() string {
	return r.Name
}

func (r *Rule) StandardForm() Formula {
	pres, result := r.Preconditions, r.Result
	left := pres[0]
	for _, l := range pres[1:] {
		left = And(left, l)
	}
	return Implication(left, result)
}

func NewRule(name string, result Formula, preconditions ...Formula) *Rule {
	return &Rule{Preconditions: preconditions, Result: result, Name: name}
}

// I -> -- P, Q => P -> Q
func IImply(iff Formula, then Formula) *Rule {
	return NewRule("I ->", Implication(iff, then), iff, then)
}

// E -> -- (P -> Q), P => Q
func EImply(iff Formula, then Formula) *Rule {
	return NewRule("E ->", then, Implication(iff, then), iff)
}

// I ^ -- P, Q => P ^ Q
func IAnd(left Formula, right Formula) *Rule {
	return NewRule("I ^", And(left, right), left, right)
}

// E ^ -- Left and Right versions
//
//	Left : P ^ Q => P
//	Right: P ^ Q => Q
func EAnd(left Formula, right Formula, isLeft bool) *Rule {
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
//	Left: P -> P v Q
//	Right: Q -> P v Q
func IOr(left Formula, right Formula, isLeft bool) *Rule {
	pre := right
	name := "I v (R)"
	if isLeft {
		pre = left
		name = "I v (L)"
	}
	return NewRule(name, Or(left, right), pre)
}

// E v -- P -> R, Q -> R, P v Q -> R
func EOr(if1 Formula, if2 Formula, then Formula) *Rule {
	return NewRule("E v", then,
		Implication(if1, then),
		Implication(if2, then),
		Or(if1, if2))
}

// I ~ -- P -> ~~P
func INot(formula Formula) *Rule {
	return NewRule("I ~", Not(Not(formula)), formula)
}

// E ~ -- ~~P -> P
func ENot(formula Formula) *Rule {
	return NewRule("E ~", formula, Not(Not(formula)))
}

// I <-> -- (P -> Q), (Q -> P) => P <-> Q
func IBiconditional(l Formula, r Formula) *Rule {
	return NewRule("I <->", Biconditional(l, r), Implication(l, r), Implication(r, l))
}

// E <-> -- Left and Right versions
//
//	Left: (P <-> Q) -> (P -> Q)
//	Right: (P <-> Q) -> (Q -> P)
func EBiconditional(l Formula, r Formula, isLeft bool) *Rule {
	name := "E <-> (R)"
	result := Implication(r, l)
	if isLeft {
		name = "E <-> (L)"
		result = Implication(l, r)
	}
	return NewRule(name, result, Biconditional(l, r))
}

// preconditions: ∀x.( Q(x) )
// result: Q(a) -- from substituting: ∀x.( Q(x) )[x -> a]
func EForall(formula Formula, varName string, arg string) *Rule {
	return NewRule("E ∀",
		substituteVar(formula, varName, arg),
		Forall(varName, formula),
	)
}

// preconditions: Q(a)
// result: ∃x.( Q(x) ) -- from substituting: Q(a)[a -> x]
func IExistential(formula Formula, from string, to string) *Rule {
	return NewRule("I ∃",
		Existential(to, substituteVar(formula, from, to)),
		formula,
	)
}
