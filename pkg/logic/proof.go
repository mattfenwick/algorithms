package logic

import "github.com/pkg/errors"

type Step interface {
	StepResult() Term
}

type Proof struct {
	Steps []Step
}

func NewProof(steps ...Step) *Proof {
	return &Proof{Steps: steps}
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

func NewSubProofContradiction(hypothesis Term, steps ...Step) (*SubProof, error) {
	if len(steps) == 0 {
		return nil, errors.Errorf("expected at least 1 step")
	}
	last := steps[len(steps)-1]
	switch t := last.(type) {
	case Term:
		if t.TermPrint(true) != Not(hypothesis).TermPrint(true) {
			return nil, errors.Errorf("subproof contradiction must end with term negating hypothesis")
		}
		return &SubProof{
			Hypothesis:   hypothesis,
			Steps:        steps,
			Result:       t,
			SubProofType: SubProofTypeContradiction,
		}, nil
	default:
		return nil, errors.Errorf("subproof contradiction must end with term, found %+v", last)
	}
}

func NewSubProofImplication(hypothesis Term, steps ...Step) (*SubProof, error) {
	// last step is the result and must be a term
	if len(steps) == 0 {
		return nil, errors.Errorf("must be at least one step in subproof")
	}
	last := steps[len(steps)-1]
	return &SubProof{
		Hypothesis:   hypothesis,
		Steps:        steps,
		Result:       Implication(hypothesis, last.StepResult()),
		SubProofType: SubProofTypeImplication,
	}, nil
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
