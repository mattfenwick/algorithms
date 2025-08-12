package logic

import (
	"github.com/mattfenwick/collections/pkg/set"
	"github.com/pkg/errors"
)

type Environment struct {
	Parent    *Environment
	TrueNodes *set.Set[string]
}

func (e *Environment) Find(key string) bool {
	return e.TrueNodes.Contains(key)
}

func (e *Environment) FindInParent(key string) bool {
	if e.Parent == nil {
		return false
	}
	return e.Parent.Find(key) || e.Parent.FindInParent(key)
}

func (e *Environment) Add(key string) error {
	if e.TrueNodes.Contains(key) {
		return errors.Errorf("unable to add key '%s', already present", key)
	}
	e.TrueNodes.Add(key)
	return nil
}

func (e *Environment) Apply(rule Rule) error {
	// 1. check preconditions
	for _, n := range rule.Preconditions() {
		key := PrettyPrint(n)
		// TODO reiterate -- needs to look up in parent.  how to do this?
		if !e.Find(key) {
			return errors.Errorf("missing or untrue premise '%s'", key)
		}
	}
	// 2. result
	result := rule.Result()
	// 3. check that result is not in environment
	// 4. update env
	thenString := PrettyPrint(result)
	if err := e.Add(thenString); err != nil {
		return errors.Errorf("'%s' aready in environment -- unnecessary step", thenString)
	}
	return nil
}

type Rule interface {
	StandardForm() Rule
	Preconditions() []Node
	Result() Node
}

func StandardForm(rule Rule) Node {
	sf := rule.StandardForm()
	pres, result := sf.Preconditions(), sf.Result()
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
	If   Node
	Then Node
}

func (e *ElimImplicationRule) StandardForm() Rule { //*ElimImplicationRule {
	return &ElimImplicationRule{If: Var("P"), Then: Var("Q")}
}

func (e *ElimImplicationRule) Preconditions() []Node {
	return []Node{
		e.If,
		Implication(e.If, e.Then),
	}
}

func (e *ElimImplicationRule) Result() Node {
	return e.Then
}

type IntroImplicationRule struct {
	If   Node
	Then Node
}

func (e *IntroImplicationRule) StandardForm() Rule { //*ElimImplicationRule {
	return &IntroImplicationRule{If: Var("P"), Then: Var("Q")}
}

func (e *IntroImplicationRule) Preconditions() []Node {
	return []Node{e.If, e.Then}
}

func (e *IntroImplicationRule) Result() Node {
	return Implication(e.If, e.Then)
}

// rule evaluation
// 1. ensure all args' node strings are in the env.  if not, error -- does not apply
// 2. apply rule to get result
// 3. ensure result is not already in env.  if not, error -- unnecessary step
// 4. add result to env
// example: modus ponens, (P -> Q), P -> Q
//   a. env: Q -> (R ^ S), Q
//      apply modusPonens('Q', '(R ^ S)')
//      1. p -> yes; q -> yes; ok
//      2. result: R ^ S
//      3. not already in env; ok
//      4. env: Q -> (R ^ S), Q, R ^ S
// DON'T check for contradictions, that's not our job right now (TODO: or is it?)

// I - ^ -- A, B -> A ^ B
// E - ^ -- A ^ B -> A; A ^ B -> B;

// I - v -- A -> A v B; B -> A v B
// E - v -- P -> R, Q -> R, P v Q -> R

// reiterate
// subproof
//   to implication
//   contradiction to negation of hypothesis
