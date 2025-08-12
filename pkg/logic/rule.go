package logic

type Rule interface {
	StandardForm() Rule
	Preconditions() []Term
	Result() Term
	FindInParent() bool
}

func StandardForm(rule Rule) Term {
	// sf := rule.StandardForm()
	pres, result := rule.Preconditions(), rule.Result()
	if len(pres) == 0 {
		return result
	}
	left := pres[0]
	for _, l := range pres[1:] {
		left = And(left, l)
	}
	return Implication(left, result)
}

type ElimImplicationRule struct {
	If   Term
	Then Term
}

func (e *ElimImplicationRule) FindInParent() bool {
	return false
}

func (e *ElimImplicationRule) StandardForm() Rule { //*ElimImplicationRule {
	return &ElimImplicationRule{If: Var("P"), Then: Var("Q")}
}

func (e *ElimImplicationRule) Preconditions() []Term {
	return []Term{
		e.If,
		Implication(e.If, e.Then),
	}
}

func (e *ElimImplicationRule) Result() Term {
	return e.Then
}

type IntroImplicationRule struct {
	If   Term
	Then Term
}

func (e *IntroImplicationRule) FindInParent() bool {
	return false
}

func (e *IntroImplicationRule) StandardForm() Rule { //*ElimImplicationRule {
	return &IntroImplicationRule{If: Var("P"), Then: Var("Q")}
}

func (e *IntroImplicationRule) Preconditions() []Term {
	return []Term{e.If, e.Then}
}

func (e *IntroImplicationRule) Result() Term {
	return Implication(e.If, e.Then)
}

// I - ^ -- A, B -> A ^ B
type IntroAndRule struct {
	Left  Term
	Right Term
}

func (e *IntroAndRule) FindInParent() bool {
	return false
}

func (e *IntroAndRule) StandardForm() Rule {
	return &IntroAndRule{Left: Var("P"), Right: Var("Q")}
}

func (e *IntroAndRule) Preconditions() []Term {
	return []Term{e.Left, e.Right}
}

func (e *IntroAndRule) Result() Term {
	return And(e.Left, e.Right)
}

// E - ^ -- A ^ B -> A; A ^ B -> B;
type ElimAndRule struct {
	Left   Term
	Right  Term
	IsLeft bool
}

func (e *ElimAndRule) FindInParent() bool {
	return false
}

func (e *ElimAndRule) StandardForm() Rule {
	return &ElimAndRule{Left: Var("P"), Right: Var("Q"), IsLeft: true}
}

func (e *ElimAndRule) Preconditions() []Term {
	return []Term{And(e.Left, e.Right)}
}

func (e *ElimAndRule) Result() Term {
	if e.IsLeft {
		return e.Left
	}
	return e.Right
}

// I - v -- A -> A v B; B -> A v B
type IntroOrRule struct {
	Left   Term
	Right  Term
	IsLeft bool
}

func (e *IntroOrRule) FindInParent() bool {
	return false
}

func (e *IntroOrRule) StandardForm() Rule {
	return &IntroOrRule{Left: Var("P"), Right: Var("Q"), IsLeft: true}
}

func (e *IntroOrRule) Preconditions() []Term {
	if e.IsLeft {
		return []Term{e.Left}
	} else {
		return []Term{e.Right}
	}
}

func (e *IntroOrRule) Result() Term {
	return Or(e.Left, e.Right)
}

// E - v -- P -> R, Q -> R, P v Q -> R
type ElimOrRule struct {
	If1  Term
	If2  Term
	Then Term
}

func (e *ElimOrRule) FindInParent() bool {
	return false
}

func (e *ElimOrRule) StandardForm() Rule {
	return &ElimOrRule{If1: Var("P"), If2: Var("Q"), Then: Var("R")}
}

func (e *ElimOrRule) Preconditions() []Term {
	return []Term{
		Implication(e.If1, e.Then),
		Implication(e.If2, e.Then),
		Or(e.If1, e.If2)}
}

func (e *ElimOrRule) Result() Term {
	return e.Then
}

// I - ~ -- TODO
// E - ~ -- TODO

// reiterate
type ReiterateRule struct {
	Term Term
}

func (e *ReiterateRule) FindInParent() bool {
	return true
}

func (e *ReiterateRule) StandardForm() Rule {
	return &ReiterateRule{Term: Var("P")}
}

func (e *ReiterateRule) Preconditions() []Term {
	return []Term{e.Term}
}

func (e *ReiterateRule) Result() Term {
	return e.Term
}

// subproof
//   to implication
//   contradiction to negation of hypothesis
