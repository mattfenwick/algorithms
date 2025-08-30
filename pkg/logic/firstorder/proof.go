package logic

import (
	"fmt"

	"github.com/mattfenwick/collections/pkg/json"
	"github.com/mattfenwick/collections/pkg/slice"
	"github.com/pkg/errors"
)

type Step interface {
	StepResult() *Formula
	StepDefineTermVar() *string
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
	Result         Formula
	ProofType      ProofType
}

func (p *Proof) StepResult() *Formula {
	return &p.Result
}

func (p *Proof) StepDefineTermVar() *string {
	return nil
}

func (p *Proof) StepName() string {
	return fmt.Sprintf("subproof %s", p.ProofType.Name())
}

func NewRootProof(expectedResult string, steps ...Step) *Proof {
	return &Proof{
		ExpectedResult: expectedResult,
		Steps:          steps,
		Result:         *steps[len(steps)-1].StepResult(),
		ProofType:      ProofTypeRoot,
	}
}

func NewProofContradiction(hypothesis Formula, steps ...Step) *Proof {
	if len(steps) == 0 {
		panic(errors.Errorf("expected at least 1 step"))
	}
	steps = append([]Step{&Assumption{Formula: hypothesis, ProofType: ProofTypeContradiction}}, steps...)
	results := map[string]bool{hypothesis.FormulaPrint(true): true}
	for _, step := range steps[:len(steps)-1] {
		if result := step.StepResult(); result != nil {
			results[(*result).FormulaPrint(true)] = true
		}
	}
	last := *steps[len(steps)-1].StepResult()
	foundPositive, foundNegative := false, false
	switch t := last.(type) {
	case *NotFormula:
		_, foundPositive = results[t.Formula.FormulaPrint(true)]
	}
	_, foundNegative = results[Not(last).FormulaPrint(true)]
	if !foundNegative && !foundPositive {
		resultsJson, err := json.MarshalWithOptions(results, &json.MarshalOptions{EscapeHTML: false, Indent: true, Sort: true})
		if err != nil {
			fmt.Printf("unable to marshal json for results: %+v\n", err)
		}
		panic(errors.Errorf(
			"subproof contradiction must end with negation of some previous formula -- '%s' not found in scope\n  previous formulas: %s",
			last.FormulaPrint(true),
			resultsJson))
	}
	// avoid introducing double negative. if hypothesis is:
	//    ~ Z, result is Z
	//    Z, result is ~ Z
	var result Formula
	switch a := hypothesis.(type) {
	case *NotFormula:
		result = a.Formula
	default:
		result = Not(hypothesis)
	}
	return &Proof{
		ExpectedResult: result.FormulaPrint(true),
		Steps:          steps,
		Result:         result,
		ProofType:      ProofTypeContradiction,
	}
}

func NewProofImplication(hypothesis Formula, steps ...Step) *Proof {
	steps = append([]Step{&Assumption{Formula: hypothesis, ProofType: ProofTypeImplication}}, steps...)
	// last step is the result
	last := steps[len(steps)-1]
	result := Implication(hypothesis, *last.StepResult())
	return &Proof{
		ExpectedResult: result.FormulaPrint(true),
		Steps:          steps,
		Result:         result,
		ProofType:      ProofTypeImplication,
	}
}

func NewProofExistentialElim(varName string, existential *QuantifiedFormula, steps ...Step) *Proof {
	if existential.Quantifier != ExistentialQuantifier {
		panic(errors.Errorf("existential quantifier required, found '%s'", existential.Quantifier))
	}
	// TODO check that there's no shadowing.  so these must have different var names from hypothesis:
	//   subordinate proofs
	//   existential and forall formulas
	//   forall instantiation
	steps = append([]Step{
		&TermVar{Name: varName},
		&QuantifierAssumption{
			Formula:       substituteVar(existential.Body, existential.Var, varName),
			ProofType:     ProofTypeExistentialElimination,
			Preconditions: []Formula{existential},
		},
	}, steps...)
	// verify reiterations don't mention hypothesis
	for _, step := range steps {
		switch t := step.(type) {
		case *Reiterate:
			if containsQuantifierHypothesis(t.Formula, varName) {
				panic(errors.Errorf("hypothesis '%s' used in reiteration '%s'", varName, t.Formula.FormulaPrint(true)))
			}
		}
	}
	// last step is the result and must not refer to the hypothesis
	last := steps[len(steps)-1]
	result := *last.StepResult()
	// verify hypothesis isn't used in result
	if containsQuantifierHypothesis(result, varName) {
		panic(errors.Errorf("hypothesis '%s' used in result '%s'", varName, result.FormulaPrint(true)))
	}
	return &Proof{
		ExpectedResult: result.FormulaPrint(true),
		Steps:          steps,
		Result:         result,
		ProofType:      ProofTypeExistentialElimination,
	}
}

func containsQuantifierHypothesis(formula Formula, hypothesis string) bool {
	switch t := formula.(type) {
	case *NotFormula:
		return containsQuantifierHypothesis(t.Formula, hypothesis)
	case *PredicateFormula:
		return slice.Any(func(p *PredicateArg) bool { return p.Var == hypothesis }, t.Args)
	case *BinOpFormula:
		return containsQuantifierHypothesis(t.Left, hypothesis) || containsQuantifierHypothesis(t.Right, hypothesis)
	case *QuantifiedFormula:
		return containsQuantifierHypothesis(t.Body, hypothesis)
	default:
		panic(errors.Errorf("invalid formula type: %T", formula))
	}
}

func NewProofForallIntro(varName string, hypothesis string, steps ...Step) *Proof {
	// TODO shadowing
	// TODO verify reiterations don't mention hypothesis
	// last step is the result
	steps = append([]Step{
		&TermVar{Name: hypothesis},
	}, steps...)
	last := steps[len(steps)-1]
	result := Forall(varName, substituteVar(*last.StepResult(), hypothesis, varName))
	return &Proof{
		ExpectedResult: result.FormulaPrint(true),
		Steps:          steps,
		Result:         result,
		ProofType:      ProofTypeForallIntroduction,
	}
}

// Reiterate pulls in a formula from an enclosing scope
type Reiterate struct {
	Formula Formula
}

func (r *Reiterate) StepResult() *Formula {
	return &r.Formula
}

func (t *Reiterate) StepDefineTermVar() *string {
	return nil
}

func (r *Reiterate) StepName() string {
	return "Reiterate"
}

type Assumption struct {
	Formula   Formula
	ProofType ProofType
}

func (a *Assumption) StepResult() *Formula {
	return &a.Formula
}

func (a *Assumption) StepDefineTermVar() *string {
	return nil
}

func (a *Assumption) StepName() string {
	return fmt.Sprintf("Assume: %s", a.ProofType.Name())
}

type QuantifierAssumption struct {
	Formula       Formula
	ProofType     ProofType
	Preconditions []Formula
}

func (a *QuantifierAssumption) StepResult() *Formula {
	return &a.Formula
}

func (a *QuantifierAssumption) StepDefineTermVar() *string {
	return nil
}

func (a *QuantifierAssumption) StepName() string {
	return fmt.Sprintf("Assume: %s", a.ProofType.Name())
}

type TermVar struct {
	Name string
}

func (t *TermVar) StepResult() *Formula {
	return nil
}

func (t *TermVar) StepDefineTermVar() *string {
	return &t.Name
}

func (t *TermVar) StepName() string {
	return "Term var"
}
