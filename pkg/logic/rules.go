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

type Rule interface {
	StandardForm() Node
	Apply(env *Environment) error
}

type ElimImplicationRule struct {
	If   Node
	Then Node
}

func (e *ElimImplicationRule) StandardForm() Node {
	return Implication(
		And(
			Implication(
				Var("P"),
				Var("Q")),
			Var("P")),
		Var("Q"))
}

func (e *ElimImplicationRule) Apply(env *Environment) error {
	// 1. check preconditions
	for _, n := range []Node{e.If, Implication(e.If, e.Then)} {
		key := PrettyPrint(n)
		if !env.Find(key) {
			return errors.Errorf("missing or untrue premise '%s'", key)
		}
	}
	// 2. result
	result := e.Then
	// 3. check that result is not in environment
	// 4. update env
	thenString := PrettyPrint(result)
	if err := env.Add(thenString); err != nil {
		return errors.Errorf("'%s' aready in environment -- unnecessary step", thenString)
	}
	return nil
}

func AssertTrue(env map[string]bool, node Node) error {
	str := PrettyPrint(node)
	val, ok := env[str]
	if !ok {
		return errors.Errorf("'%s' not found in env", str)
	} else if !val {
		return errors.Errorf("'%s' is false in env", str)
	}
	return nil
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

// I - -> P, Q -> (P -> Q)
func IntroImplication(p Node, q Node) func(map[string]bool) error {
	return func(env map[string]bool) error {
		// 1. check preconditions
		if err := AssertTrue(env, p); err != nil {
			return errors.WithMessagef(err, "unable to apply rule -- missing P")
		} else if err := AssertTrue(env, q); err != nil {
			return errors.WithMessagef(err, "unable to apply rule -- missing Q")
		}
		// 2. result
		result := Implication(p, q)
		// 3. check that result is not in environment
		thenString := PrettyPrint(result)
		if _, ok := env[thenString]; ok {
			return errors.Errorf("'%s' aready in environment -- unnecessary step", thenString)
		}
		// 4. update env
		env[thenString] = true
		return nil
	}
}

// E - -> (P -> Q), P -> Q
func ElimImplication(p Node, q Node) func(map[string]bool) error {
	return func(env map[string]bool) error {
		// 1. check preconditions
		implication := Implication(p, q)
		if err := AssertTrue(env, p); err != nil {
			return errors.WithMessagef(err, "unable to apply rule -- missing P")
		} else if err := AssertTrue(env, implication); err != nil {
			return errors.WithMessagef(err, "unable to apply rule -- missing implication")
		}
		// 2. result
		result := q
		// 3. check that result is not in environment
		thenString := PrettyPrint(result)
		if _, ok := env[thenString]; ok {
			return errors.Errorf("'%s' aready in environment -- unnecessary step", thenString)
		}
		env[thenString] = true
		return nil
	}
}

// I - ^ -- A, B -> A ^ B
// E - ^ -- A ^ B -> A; A ^ B -> B;

// I - v -- A -> A v B; B -> A v B
// E - v -- P -> R, Q -> R, P v Q -> R

// reiterate
// subproof
//   to implication
//   contradiction to negation of hypothesis
