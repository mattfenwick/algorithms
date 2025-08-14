package logic

import (
	"fmt"
	"strings"

	"github.com/mattfenwick/collections/pkg/set"
	"github.com/mattfenwick/collections/pkg/slice"
	"github.com/pkg/errors"
)

type EvaluatedStep struct {
	Depth      int
	Step       Step
	ScopeTerms []string
}

type Scope struct {
	Depth     int
	Parent    *Scope
	TrueTerms *set.Set[string]
}

func NewScope(parent *Scope) *Scope {
	depth := 0
	if parent != nil {
		depth = parent.Depth + 1
	}
	return &Scope{
		Depth:     depth,
		Parent:    parent,
		TrueTerms: set.FromSlice([]string{})}
}

func (e *Scope) Find(key string) bool {
	return e.TrueTerms.Contains(key)
}

func (e *Scope) FindInParent(key string) bool {
	if e.Parent == nil {
		return false
	}
	return e.Parent.Find(key) || e.Parent.FindInParent(key)
}

func (e *Scope) Add(key string) error {
	if !e.TrueTerms.Add(key) {
		return errors.Errorf("unable to add key '%s', already present", key)
	}
	return nil
}

func (e *Scope) GetTrueTerms() []string {
	return slice.Sort(e.TrueTerms.ToSlice())
}

func (e *Scope) Print(indent int) {
	fmt.Printf("%s%s\n",
		strings.Join(slice.Replicate(indent*2, " "), ""),
		strings.Join(e.GetTrueTerms(), ","))
	if e.Parent != nil {
		e.Parent.Print(indent + 1)
	}
}

func ApplyProof(proof *Proof, parentScope *Scope) ([]*EvaluatedStep, error) {
	var outSteps []*EvaluatedStep
	scope := NewScope(parentScope)
	if proof.Hypothesis != nil {
		evaledSteps, err := ApplyStep(&Assumption{Term: proof.Hypothesis}, scope)
		if err != nil {
			return nil, err
		}
		outSteps = append(outSteps, evaledSteps...)
	}
	for _, step := range proof.Steps {
		evaledSteps, err := ApplyStep(step, scope)
		if err != nil {
			return nil, err
		}
		outSteps = append(outSteps, evaledSteps...)
	}
	return outSteps, nil
}

func ApplyStep(step Step, scope *Scope) ([]*EvaluatedStep, error) {
	eStep := &EvaluatedStep{Depth: scope.Depth, Step: step, ScopeTerms: scope.GetTrueTerms()}
	outSteps := []*EvaluatedStep{}
	switch t := step.(type) {
	case *Assumption:
		if err := scope.Add(t.StepResult().TermPrint(true)); err != nil {
			return nil, err
		}
	case *Proof:
		evaledSteps, err := ApplyProof(t, scope)
		if err != nil {
			return nil, err
		}
		outSteps = append(outSteps, evaledSteps...)
		if err := scope.Add(t.StepResult().TermPrint(true)); err != nil {
			return nil, err
		}
	case *Reiterate:
		key := t.Term.TermPrint(true)
		if !scope.FindInParent(key) {
			return nil, errors.Errorf("missing or untrue premise '%s' in parent(s)", key)
		}
		if err := scope.Add(t.StepResult().TermPrint(true)); err != nil {
			return nil, errors.Errorf("'%s' aready in scope -- unnecessary step", key)
		}
	case *Repeat:
		key := t.Term.TermPrint(true)
		if !scope.Find(key) {
			return nil, errors.Errorf("missing or untrue premise '%s'", key)
		}
	case *Rule:
		// check preconditions
		for _, n := range t.Preconditions {
			key := n.TermPrint(true)
			if !scope.Find(key) {
				return nil, errors.Errorf("missing or untrue premise '%s'", key)
			}
		}
		// check that result is not in scope and update env
		if err := scope.Add(t.StepResult().TermPrint(true)); err != nil {
			return nil, errors.Errorf("'%s' aready in scope -- unnecessary step", t.StepResult().TermPrint(true))
		}
	default:
		return nil, errors.Errorf("invalid step type %+v", t)
	}
	outSteps = append(outSteps, eStep)
	return outSteps, nil
}

// func printHelper(step *EvaluatedStep, line int, depth int) int {
// 	if step.ChildScope != nil {
// 		for _, s := range step.ChildScope.EvaluatedSteps {
// 			line = step.ChildScope.printHelper(s, line, depth+1)
// 		}
// 	}
// 	indent := strings.Repeat("  ", depth)
// 	fmt.Printf("%s%d: %s\n%s- %s\n",
// 		indent,
// 		line,
// 		step.Step.StepResult().TermPrint(true),
// 		indent,
// 		strings.Join(step.ScopeTerms, ",   "))
// 	line++
// 	return line
// }

// func PrintResult() {
// 	line := 1
// 	for _, step := range e.EvaluatedSteps {
// 		line = e.printHelper(step, line, 0)
// 	}
// }
