package logic

import (
	"fmt"

	"github.com/mattfenwick/collections/pkg/json"
	"github.com/pkg/errors"
)

type Step interface {
	StepResult() Term
}

type Proof struct {
	Name       string
	Hypothesis Term
	Steps      []Step
	Result     Term
	ProofType  ProofType
}

func (p *Proof) StepResult() Term {
	return p.Result
}

type ProofType string

const (
	ProofTypeRoot          ProofType = "Root"
	ProofTypeContradiction ProofType = "Contradiction"
	ProofTypeImplication   ProofType = "Implication"
)

func NewRootProof(name string, steps ...Step) *Proof {
	return &Proof{
		Name:       name,
		Hypothesis: nil,
		Steps:      steps,
		Result:     steps[len(steps)-1].StepResult(),
		ProofType:  ProofTypeRoot,
	}
}

func NewProofContradiction(hypothesis Term, steps ...Step) *Proof {
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
	return &Proof{
		Name:       fmt.Sprintf("subproof contradiction of '%s'", hypothesis.TermPrint(true)),
		Hypothesis: hypothesis,
		Steps:      steps,
		Result:     Not(hypothesis),
		ProofType:  ProofTypeContradiction,
	}
}

func NewProofImplication(hypothesis Term, steps ...Step) *Proof {
	// last step is the result
	if len(steps) == 0 {
		panic(errors.Errorf("must be at least one step in subproof"))
	}
	last := steps[len(steps)-1]
	return &Proof{
		Name:       fmt.Sprintf("subproof implication of '%s'", hypothesis.TermPrint(true)),
		Hypothesis: hypothesis,
		Steps:      steps,
		Result:     Implication(hypothesis, last.StepResult()),
		ProofType:  ProofTypeImplication,
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

// TODO do we really need this?
type Assumption struct {
	Term Term
}

func (a *Assumption) StepResult() Term {
	return a.Term
}
