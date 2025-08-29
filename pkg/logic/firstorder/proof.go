package logic

import (
	"fmt"

	"github.com/mattfenwick/collections/pkg/json"
	"github.com/mattfenwick/collections/pkg/slice"
	"github.com/pkg/errors"
)

type Step interface {
	StepResult() Term
	StepName() string
}

type ProofType string

const (
	ProofTypeRoot                   ProofType = "Root"
	ProofTypeContradiction          ProofType = "Contradiction"
	ProofTypeImplication            ProofType = "Implication"
	ProofTypeForallIntroduction     ProofType = "Forall Introduction"
	ProofTypeExistentialElimination ProofType = "Existential Elimination"
)

func (p ProofType) Name() string {
	switch p {
	case ProofTypeContradiction:
		return "contradiction"
	case ProofTypeImplication:
		return "implication"
	case ProofTypeExistentialElimination:
		return "∃ elimination"
	case ProofTypeForallIntroduction:
		return "∀ introduction"
	default:
		panic(errors.Errorf("no Name defined for proof type '%s'", p))
	}

}

type Proof struct {
	ExpectedResult string
	Steps          []Step
	Result         Term
	ProofType      ProofType
}

func (p *Proof) StepResult() Term {
	return p.Result
}

func (p *Proof) StepName() string {
	return fmt.Sprintf("subproof %s", p.ProofType.Name())
}

func NewRootProof(expectedResult string, steps ...Step) *Proof {
	return &Proof{
		ExpectedResult: expectedResult,
		Steps:          steps,
		Result:         steps[len(steps)-1].StepResult(),
		ProofType:      ProofTypeRoot,
	}
}

func NewProofContradiction(hypothesis Term, steps ...Step) *Proof {
	if len(steps) == 0 {
		panic(errors.Errorf("expected at least 1 step"))
	}
	steps = append([]Step{&Assumption{Term: hypothesis, ProofType: ProofTypeContradiction}}, steps...)
	results := map[string]bool{hypothesis.TermPrint(true): true}
	for _, step := range steps[:len(steps)-1] {
		results[step.StepResult().TermPrint(true)] = true
	}
	last := steps[len(steps)-1].StepResult()
	foundPositive, foundNegative := false, false
	switch t := last.(type) {
	case *NotTerm:
		_, foundPositive = results[t.Term.TermPrint(true)]
	}
	_, foundNegative = results[Not(last).TermPrint(true)]
	if !foundNegative && !foundPositive {
		resultsJson, err := json.MarshalWithOptions(results, &json.MarshalOptions{EscapeHTML: false, Indent: true, Sort: true})
		if err != nil {
			fmt.Printf("unable to marshal json for results: %+v\n", err)
		}
		panic(errors.Errorf(
			"subproof contradiction must end with negation of some previous term -- '%s' not found in scope\n  previous terms: %s",
			last.TermPrint(true),
			resultsJson))
	}
	// avoid introducing double negative. if hypothesis is:
	//    ~ Z, result is Z
	//    Z, result is ~ Z
	var result Term
	switch a := hypothesis.(type) {
	case *NotTerm:
		result = a.Term
	default:
		result = Not(hypothesis)
	}
	return &Proof{
		ExpectedResult: result.TermPrint(true),
		Steps:          steps,
		Result:         result,
		ProofType:      ProofTypeContradiction,
	}
}

func NewProofImplication(hypothesis Term, steps ...Step) *Proof {
	steps = append([]Step{&Assumption{Term: hypothesis, ProofType: ProofTypeImplication}}, steps...)
	// last step is the result
	last := steps[len(steps)-1]
	result := Implication(hypothesis, last.StepResult())
	return &Proof{
		ExpectedResult: result.TermPrint(true),
		Steps:          steps,
		Result:         result,
		ProofType:      ProofTypeImplication,
	}
}

func NewProofExistentialElim(varName string, existential *QuantifiedTerm, steps ...Step) *Proof {
	if existential.Quantifier != ExistentialQuantifier {
		panic(errors.Errorf("existential quantifier required, found '%s'", existential.Quantifier))
	}
	// TODO check that there's no shadowing.  so these must have different var names from hypothesis:
	//   subordinate proofs
	//   existential and forall terms
	//   forall instantiation
	steps = append([]Step{
		&QuantifierAssumption{
			Term:          substituteVar(existential.Body, existential.Var, varName),
			ProofType:     ProofTypeExistentialElimination,
			Preconditions: []Term{existential},
		},
	}, steps...)
	// verify reiterations don't mention hypothesis
	for _, step := range steps {
		switch t := step.(type) {
		case *Reiterate:
			if containsQuantifierHypothesis(t.Term, varName) {
				panic(errors.Errorf("hypothesis '%s' used in reiteration '%s'", varName, t.Term.TermPrint(true)))
			}
		}
	}
	// last step is the result and must not refer to the hypothesis
	last := steps[len(steps)-1]
	result := last.StepResult()
	// verify hypothesis isn't used in result
	if containsQuantifierHypothesis(result, varName) {
		panic(errors.Errorf("hypothesis '%s' used in result '%s'", varName, result.TermPrint(true)))
	}
	return &Proof{
		ExpectedResult: result.TermPrint(true),
		Steps:          steps,
		Result:         result,
		ProofType:      ProofTypeExistentialElimination,
	}
}

func containsQuantifierHypothesis(term Term, hypothesis string) bool {
	switch t := term.(type) {
	case *NotTerm:
		return containsQuantifierHypothesis(t.Term, hypothesis)
	case *PropTerm:
		return slice.Any(func(p *PropArg) bool { return p.Var == hypothesis }, t.Args)
	case *BinOpTerm:
		return containsQuantifierHypothesis(t.Left, hypothesis) || containsQuantifierHypothesis(t.Right, hypothesis)
	case *QuantifiedTerm:
		return containsQuantifierHypothesis(t.Body, hypothesis)
	default:
		panic(errors.Errorf("invalid term type: %T", term))
	}
}

func NewProofForallIntro(varName string, hypothesis string, steps ...Step) *Proof {
	// TODO shadowing
	// TODO verify reiterations don't mention hypothesis
	// last step is the result
	last := steps[len(steps)-1]
	result := Forall(varName, substituteVar(last.StepResult(), hypothesis, varName))
	return &Proof{
		ExpectedResult: result.TermPrint(true),
		Steps:          steps,
		Result:         result,
		ProofType:      ProofTypeForallIntroduction,
	}
}

// Reiterate pulls in a term from an enclosing scope
type Reiterate struct {
	Term Term
}

func (r *Reiterate) StepResult() Term {
	return r.Term
}

func (r *Reiterate) StepName() string {
	return "Reiterate"
}

type Assumption struct {
	Term      Term
	ProofType ProofType
}

func (a *Assumption) StepResult() Term {
	return a.Term
}

func (a *Assumption) StepName() string {
	return fmt.Sprintf("Assume: %s", a.ProofType.Name())
}

type QuantifierAssumption struct {
	Term          Term
	ProofType     ProofType
	Preconditions []Term
}

func (a *QuantifierAssumption) StepResult() Term {
	return a.Term
}

func (a *QuantifierAssumption) StepName() string {
	return fmt.Sprintf("Assume: %s", a.ProofType.Name())
}
