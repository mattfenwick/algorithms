package logic

type Rule struct {
	Preconditions []Formula
	Result        Formula
	Name          string
	UseTermVar    *string
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
	return Arrow(left, result)
}

func NewRule(name string, result Formula, preconditions ...Formula) *Rule {
	return &Rule{Preconditions: preconditions, Result: result, Name: name}
}

// I -> -- P, Q => P -> Q
func IImply(iff Formula, then Formula) *Rule {
	return NewRule("I ->", Arrow(iff, then), iff, then)
}

// E -> -- (P -> Q), P => Q
func EImply(iff Formula, then Formula) *Rule {
	return NewRule("E ->", then, Arrow(iff, then), iff)
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
		Arrow(if1, then),
		Arrow(if2, then),
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
func IDArrow(l Formula, r Formula) *Rule {
	return NewRule("I <->", DArrow(l, r), Arrow(l, r), Arrow(r, l))
}

// E <-> -- Left and Right versions
//
//	Left: (P <-> Q) -> (P -> Q)
//	Right: (P <-> Q) -> (Q -> P)
func EDArrow(l Formula, r Formula, isLeft bool) *Rule {
	name := "E <-> (R)"
	result := Arrow(r, l)
	if isLeft {
		name = "E <-> (L)"
		result = Arrow(l, r)
	}
	return NewRule(name, result, DArrow(l, r))
}

// preconditions: ∀x.( Q(x) )
// result: Q(a) -- from substituting: ∀x.( Q(x) )[x -> a]
func EForall(formula Formula, from string, to string) *Rule {
	rule := NewRule("E ∀",
		InstantiateFormula(formula, from, to),
		Forall(from, formula),
	)
	rule.UseTermVar = &to
	return rule
}

// preconditions: Q(a)
// result: ∃x.( Q(x) ) -- from substituting: Q(a)[a -> x]
func IExist(formula Formula, from string, to string) *Rule {
	return NewRule("I ∃",
		Exist(to, GeneralizeFormula(formula, from, to)),
		formula,
	)
}
