package logic

import (
	"fmt"
	"strings"

	"github.com/mattfenwick/collections/pkg/dict"
	"github.com/mattfenwick/collections/pkg/slice"
	"github.com/pkg/errors"
)

func CheckProof(proof *Proof) (*Scope, error) {
	env := NewScope(nil)
	err := env.ApplyProof(proof)
	return env, err
}

type EvaluatedStep struct {
	Step       Step
	ScopeTerms []string
	ChildScope *Scope
}

type Scope struct {
	Parent         *Scope
	TrueTerms      map[string]*EvaluatedStep
	EvaluatedSteps []*EvaluatedStep
}

func NewScope(parent *Scope) *Scope {
	s := &Scope{
		Parent:    parent,
		TrueTerms: map[string]*EvaluatedStep{}}
	return s
}

func (e *Scope) Find(key string) bool {
	_, ok := e.TrueTerms[key]
	return ok
}

func (e *Scope) FindInParent(key string) bool {
	if e.Parent == nil {
		return false
	}
	return e.Parent.Find(key) || e.Parent.FindInParent(key)
}

func (e *Scope) Add(eStep *EvaluatedStep) error {
	key := eStep.Step.StepResult().TermPrint(true)
	if _, ok := e.TrueTerms[key]; ok {
		return errors.Errorf("unable to add key '%s', already present", key)
	}
	e.TrueTerms[key] = eStep
	return nil
}

func (e *Scope) GetTrueTerms() []string {
	return slice.Sort(dict.Keys(e.TrueTerms))
}

func (e *Scope) Print(indent int) {
	fmt.Printf("%s%s\n",
		strings.Join(slice.Replicate(indent*2, " "), ""),
		strings.Join(e.GetTrueTerms(), ","))
	if e.Parent != nil {
		e.Parent.Print(indent + 1)
	}
}

func (e *Scope) ApplyProof(proof *Proof) error {
	if proof.Hypothesis != nil {
		if err := e.ApplyStep(&Assumption{Term: proof.Hypothesis}); err != nil {
			return err
		}
	}
	for _, step := range proof.Steps {
		if err := e.ApplyStep(step); err != nil {
			return err
		}
	}
	return nil
}

func (e *Scope) ApplyStep(step Step) error {
	eStep := &EvaluatedStep{Step: step, ScopeTerms: e.GetTrueTerms()}
	switch t := step.(type) {
	case *Assumption:
		if err := e.Add(eStep); err != nil {
			return err
		}
	case *Proof:
		childEnv := NewScope(e)
		if err := childEnv.ApplyProof(t); err != nil {
			return err
		}
		if err := e.Add(eStep); err != nil {
			return err
		}
		eStep.ChildScope = childEnv
	case *Reiterate:
		key := t.Term.TermPrint(true)
		if !e.FindInParent(key) {
			return errors.Errorf("missing or untrue premise '%s' in parent(s)", key)
		}
		if err := e.Add(eStep); err != nil {
			return errors.Errorf("'%s' aready in scope -- unnecessary step", key)
		}
		return nil
	case *Repeat:
		key := t.Term.TermPrint(true)
		if !e.Find(key) {
			return errors.Errorf("missing or untrue premise '%s'", key)
		}
		return nil
	case *Rule:
		// check preconditions
		for _, n := range t.Preconditions {
			key := n.TermPrint(true)
			if !e.Find(key) {
				return errors.Errorf("missing or untrue premise '%s'", key)
			}
		}
		// check that result is not in scope and update env
		if err := e.Add(eStep); err != nil {
			return errors.Errorf("'%s' aready in scope -- unnecessary step", eStep.Step.StepResult().TermPrint(true))
		}
		return nil
	default:
		return errors.Errorf("invalid step type %+v", t)
	}
	e.EvaluatedSteps = append(e.EvaluatedSteps, eStep)
	return nil
}

func (e *Scope) printHelper(step *EvaluatedStep, line int, depth int) int {
	if step.ChildScope != nil {
		for _, s := range step.ChildScope.EvaluatedSteps {
			line = step.ChildScope.printHelper(s, line, depth+1)
		}
	}
	indent := strings.Repeat("  ", depth)
	fmt.Printf("%s%d: %s\n%s- %s\n",
		indent,
		line,
		step.Step.StepResult().TermPrint(true),
		indent,
		strings.Join(step.ScopeTerms, ",   "))
	line++
	return line
}

func (e *Scope) PrintResult() {
	line := 1
	for _, step := range e.EvaluatedSteps {
		line = e.printHelper(step, line, 0)
	}
}
