package logic

import (
	"github.com/mattfenwick/collections/pkg/json"
	"github.com/pkg/errors"
)

type Step interface {
	StepResult() Term
}

type Proof struct {
	Name  string
	Steps []Step
}

func NewProof(name string, steps ...Step) *Proof {
	return &Proof{Name: name, Steps: steps}
}

type SubProofType string

const (
	SubProofTypeContradiction SubProofType = "Contradiction"
	SubProofTypeImplication   SubProofType = "Implication"
)

type SubProof struct {
	Hypothesis   Term
	Steps        []Step
	Result       Term
	SubProofType SubProofType
}

func (sp *SubProof) StepResult() Term {
	return sp.Result
}

func NewSubProofContradiction(hypothesis Term, steps ...Step) *SubProof {
	if len(steps) == 0 {
		panic(errors.Errorf("expected at least 1 step"))
	}
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
		panic(errors.Errorf(
			"subproof contradiction must end with negation of some previous term -- '%s' not found in scope\n  previous terms: %s",
			last.TermPrint(true),
			json.MustMarshalToString(results)))
	}
	return &SubProof{
		Hypothesis:   hypothesis,
		Steps:        steps,
		Result:       Not(hypothesis),
		SubProofType: SubProofTypeContradiction,
	}
}

func NewSubProofImplication(hypothesis Term, steps ...Step) *SubProof {
	// last step is the result and must be a term
	if len(steps) == 0 {
		panic(errors.Errorf("must be at least one step in subproof"))
	}
	last := steps[len(steps)-1]
	return &SubProof{
		Hypothesis:   hypothesis,
		Steps:        steps,
		Result:       Implication(hypothesis, last.StepResult()),
		SubProofType: SubProofTypeImplication,
	}
}

// Reiterate pulls in a term from an enclosing scope
type Reiterate struct {
	Term Term
}

func (r *Reiterate) StepResult() Term {
	return r.Term
}

// Repeat asserts that a term is in the current scope.  basically just useful for `P -> P`
type Repeat struct {
	Term Term
}

func (r *Repeat) StepResult() Term {
	return r.Term
}
