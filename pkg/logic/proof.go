package logic

import (
	"fmt"

	"github.com/mattfenwick/collections/pkg/json"
	"github.com/pkg/errors"
)

type Step interface {
	StepResult() Term
	StepName() string
}

type ProofType string

const (
	ProofTypeRoot          ProofType = "Root"
	ProofTypeContradiction ProofType = "Contradiction"
	ProofTypeImplication   ProofType = "Implication"
)

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
	switch p.ProofType {
	case ProofTypeContradiction:
		return "subproof contradiction"
	case ProofTypeImplication:
		return "subproof implication"
	default:
		panic(errors.Errorf("no StepName defined for proof type '%s'", p.ProofType))
	}
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
		_, foundPositive = results[t.Arg.TermPrint(true)]
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
		result = a.Arg
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

// TODO do we really need this?
type Assumption struct {
	Term      Term
	ProofType ProofType
}

func (a *Assumption) StepResult() Term {
	return a.Term
}

func (a *Assumption) StepName() string {
	return fmt.Sprintf("Assume %s", a.ProofType)
}
