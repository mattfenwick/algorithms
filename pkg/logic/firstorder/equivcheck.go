package logic

import "github.com/mattfenwick/collections/pkg/slice"

type EquivCheck struct {
	Left                    Formula
	Right                   Formula
	Forward                 *TT
	Backward                *TT
	ForwardCounterExamples  []*TTResult
	BackwardCounterExamples []*TTResult
}

func NewEquivCheck(left Formula, right Formula) *EquivCheck {
	forward := NewTT(NewEnv(map[string]bool{}), Arrow(left, right))
	backward := NewTT(NewEnv(map[string]bool{}), Arrow(right, left))
	fces := slice.Filter(func(t *TTResult) bool {
		return !t.FormulaValues[len(t.FormulaValues)-1]
	}, forward.Results)
	bces := slice.Filter(func(t *TTResult) bool {
		return !t.FormulaValues[len(t.FormulaValues)-1]
	}, backward.Results)
	return &EquivCheck{
		Left:                    left,
		Right:                   right,
		Forward:                 forward,
		Backward:                backward,
		ForwardCounterExamples:  fces,
		BackwardCounterExamples: bces,
	}
}

func (e *EquivCheck) IsEquiv() bool {
	return (len(e.ForwardCounterExamples) + len(e.BackwardCounterExamples)) == 0
}

func (e *EquivCheck) Row() []string {
	var leftForward, leftBackward string
	if len(e.ForwardCounterExamples) > 0 {
		leftForward = e.ForwardCounterExamples[0].VarAssignment()
	}
	if len(e.BackwardCounterExamples) > 0 {
		leftBackward = e.BackwardCounterExamples[0].VarAssignment()
	}
	return []string{leftForward, leftBackward, e.Left.FormulaPrint(true), e.Right.FormulaPrint(true)}
}
