package logic

import (
	"fmt"
	"strings"

	"github.com/mattfenwick/collections/pkg/set"
	"github.com/mattfenwick/collections/pkg/slice"
	"github.com/pkg/errors"
)

func CheckProof(proof *Proof) (*Scope, error) {
	env := NewScope(nil)
	err := env.ApplyProof(proof)
	return env, err
}

type Scope struct {
	Parent    *Scope
	TrueTerms *set.Set[string] // TODO string is nice for equality, but should this keep the actual term as well?
}

func NewScope(parent *Scope, trueTerms ...string) *Scope {
	return &Scope{Parent: parent, TrueTerms: set.FromSlice(trueTerms)}
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
	if e.TrueTerms.Contains(key) {
		return errors.Errorf("unable to add key '%s', already present", key)
	}
	e.TrueTerms.Add(key)
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

func (e *Scope) ApplyProof(proof *Proof) error {
	for _, step := range proof.Steps {
		if err := e.ApplyStep(step); err != nil {
			return err
		}
		if err := e.Add(step.StepResult().TermPrint(true)); err != nil {
			return err
		}
	}
	return nil
}

func (e *Scope) ApplyStep(step Step) error {
	switch t := step.(type) {
	case *SubProof:
		childEnv := NewScope(e, t.Hypothesis.TermPrint(true))
		return childEnv.ApplySubProof(t)
		// TODO somehow return childEnv for debugging purposes
	case *Reiterate:
		key := t.Term.TermPrint(true)
		if !e.FindInParent(key) {
			return errors.Errorf("missing or untrue premise '%s' in parent(s)", key)
		}
		if err := e.Add(key); err != nil {
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
		return e.ApplyRule(t)
	default:
		return errors.Errorf("invalid step type %+v", t)
	}
}

func (e *Scope) ApplySubProof(sp *SubProof) error {
	for _, step := range sp.Steps {
		if err := e.ApplyStep(step); err != nil {
			return err
		}
	}
	return nil
}

// ApplyRule implements application of a single rule step
// 1. ensure all args' node strings are in the env.  if not, error -- does not apply
// 2. apply rule to get result
// 3. ensure result is not already in env.  if not, error -- unnecessary step
// 4. add result to env
// example: modus ponens, (P -> Q), P -> Q
//
//	a. env: Q -> (R ^ S), Q
//	   apply modusPonens('Q', '(R ^ S)')
//	   1. p -> yes; q -> yes; ok
//	   2. result: R ^ S
//	   3. not already in env; ok
//	   4. env: Q -> (R ^ S), Q, R ^ S
//
// DON'T check for contradictions, that's not our job right here (TODO: or is it?)
func (e *Scope) ApplyRule(rule *Rule) error {
	// 1. check preconditions
	for _, n := range rule.Preconditions {
		key := n.TermPrint(true)
		if !e.Find(key) {
			return errors.Errorf("missing or untrue premise '%s'", key)
		}
	}
	// 2. result
	result := rule.Result
	// 3. check that result is not in scope
	// 4. update env
	thenString := result.TermPrint(true)
	if err := e.Add(thenString); err != nil {
		return errors.Errorf("'%s' aready in scope -- unnecessary step", thenString)
	}
	return nil
}
