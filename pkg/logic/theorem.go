package logic

func ExcludedMiddleTheorem(t Term) *Rule {
	return NewRule("P v ~ P", Or(t, Not(t)))
}
