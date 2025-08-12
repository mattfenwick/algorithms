package logic

import (
	"fmt"
	"strings"

	"github.com/mattfenwick/collections/pkg/set"
	"github.com/mattfenwick/collections/pkg/slice"
	"github.com/pkg/errors"
)

type Environment struct {
	Parent    *Environment
	TrueTerms *set.Set[string] // TODO string is nice for equality, but should this keep the actual term as well?
}

func (e *Environment) Find(key string) bool {
	return e.TrueTerms.Contains(key)
}

func (e *Environment) FindInParent(key string) bool {
	if e.Parent == nil {
		return false
	}
	return e.Parent.Find(key) || e.Parent.FindInParent(key)
}

func (e *Environment) Add(key string) error {
	if e.TrueTerms.Contains(key) {
		return errors.Errorf("unable to add key '%s', already present", key)
	}
	e.TrueTerms.Add(key)
	return nil
}

func (e *Environment) Print(indent int) {
	fmt.Printf("%s%s\n",
		strings.Join(slice.Replicate(indent*2, " "), ""),
		strings.Join(slice.Sort(e.TrueTerms.ToSlice()), ","))
	if e.Parent != nil {
		e.Parent.Print(indent + 1)
	}
}

// Apply implements rule checking
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
func (e *Environment) Apply(rule Rule) error {
	// 1. check preconditions
	for _, n := range rule.Preconditions() {
		key := n.TermPrint(true)
		if rule.FindInParent() {
			if !e.FindInParent(key) {
				return errors.Errorf("missing or untrue premise '%s' in parent", key)
			}
		} else {
			if !e.Find(key) {
				return errors.Errorf("missing or untrue premise '%s'", key)
			}
		}
	}
	// 2. result
	result := rule.Result()
	// 3. check that result is not in environment
	// 4. update env
	thenString := result.TermPrint(true)
	if err := e.Add(thenString); err != nil {
		return errors.Errorf("'%s' aready in environment -- unnecessary step", thenString)
	}
	return nil
}
