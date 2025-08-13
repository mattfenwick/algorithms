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
	// there must be at least 1 step
	last := steps[len(steps)-1]
	switch t := last.(type) {
	case *Rule:
		return &SubProof{
			Hypothesis:   hypothesis,
			Steps:        steps,
			Result:       Implication(hypothesis, t.Result),
			SubProofType: SubProofTypeImplication,
		}, nil
	default:
		return nil, errors.Errorf("last step must be a Term, got %+v", last)
	}
}

type Reiterate struct {
	Term Term
}
