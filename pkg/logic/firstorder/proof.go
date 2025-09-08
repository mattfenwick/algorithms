package logic

import (
	"fmt"

	"github.com/mattfenwick/collections/pkg/json"
	"github.com/mattfenwick/collections/pkg/slice"
	"github.com/pkg/errors"
)

type Step interface {
	StepResult() Formula
	StepName() string
}

type ProofType string

const (
	ProofTypeRoot                     ProofType = "Root"
	ProofTypeContradiction            ProofType = "Contradiction"
	ProofTypeArrow                    ProofType = "Arrow"
	ProofTypeForallIntroduction       ProofType = "Forall Introduction"
	ProofTypeExistentialElimination   ProofType = "Existential Elimination"
	ProofTypeExistentialContradiction ProofType = "Existential Contradiction"
)

func (p ProofType) Name() string {
	switch p {
	case ProofTypeContradiction:
		return "I/E ~"
	case ProofTypeArrow:
		return "I ->"
	case ProofTypeExistentialElimination:
		return "E ∃"
	case ProofTypeExistentialContradiction:
		return "~ ∃"
	case ProofTypeForallIntroduction:
		return "I ∀"
	default:
		panic(errors.Errorf("no Name defined for proof type '%s'", p))
	}

}

type Proof struct {
	ExpectedResult string
	Steps          []Step
	Result         Formula
	ProofType      ProofType
	TermVar        *string
}

func (p *Proof) StepResult() Formula {
	return p.Result
}

func (p *Proof) StepName() string {
	if p.ProofType == ProofTypeForallIntroduction {
		var boundVar string
		switch t := p.Result.(type) {
		case *QuantifiedFormula:
			if t.Quantifier == ForallQuantifier {
				boundVar = t.Var
			}
		}
		if boundVar == "" {
			panic(errors.Errorf("expected forall with bound var, found %T, %+v", p.Result, p.Result))
		}
		return fmt.Sprintf("subproof %s [%s -> %s]", p.ProofType.Name(), *p.TermVar, boundVar)
	}
	return fmt.Sprintf("subproof %s", p.ProofType.Name())
}

func RootProof(expectedResult string, steps ...Step) *Proof {
	return &Proof{
		ExpectedResult: expectedResult,
		Steps:          steps,
		Result:         steps[len(steps)-1].StepResult(),
		ProofType:      ProofTypeRoot,
	}
}

func findNegationInScope(hypothesis Formula, steps []Step) error {
	results := map[string]bool{hypothesis.FormulaPrint(true): true}
	for _, step := range steps[:len(steps)-1] {
		results[step.StepResult().FormulaPrint(true)] = true
	}
	last := steps[len(steps)-1].StepResult()
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
		return errors.Errorf(
			"subproof contradiction must end with negation of some previous formula -- '%s' not found in scope\n  previous formulas: %s",
			last.FormulaPrint(true),
			resultsJson)
	}
	return nil
}

func ContraProof(hypothesis Formula, steps ...Step) *Proof {
	if len(steps) == 0 {
		panic(errors.Errorf("expected at least 1 step"))
	}
	steps = append([]Step{&Assumption{Formula: hypothesis, ProofType: ProofTypeContradiction}}, steps...)
	err := findNegationInScope(hypothesis, steps)
	if err != nil {
		panic(err)
	}
	// avoid introducing double negative. if hypothesis is:
	//    ~ Z, result is Z
	//    Z, result is ~ Z
	// result := removeDoubleNegative(Not(hypothesis)) // TODO wouldn't this be easier?
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

func ArrowProof(hypothesis Formula, steps ...Step) *Proof {
	steps = append([]Step{&Assumption{Formula: hypothesis, ProofType: ProofTypeArrow}}, steps...)
	// last step is the result
	last := steps[len(steps)-1]
	result := Arrow(hypothesis, last.StepResult())
	return &Proof{
		ExpectedResult: result.FormulaPrint(true),
		Steps:          steps,
		Result:         result,
		ProofType:      ProofTypeArrow,
	}
}

func ExistElimProof(termVar string, existential *QuantifiedFormula, steps ...Step) *Proof {
	return existProof(termVar, existential, steps, false)
}

func ExistContraProof(termVar string, existential *QuantifiedFormula, steps ...Step) *Proof {
	return existProof(termVar, existential, steps, true)
}

func existProof(termVar string, existential *QuantifiedFormula, steps []Step, isContra bool) *Proof {
	if existential.Quantifier != ExistentialQuantifier {
		panic(errors.Errorf("existential quantifier required, found '%s'", existential.Quantifier))
	}
	// TODO check that there's no shadowing.  so these must have different var names from hypothesis:
	//   subordinate proofs
	//   existential and forall formulas
	//   forall instantiation
	var proofType ProofType
	if isContra {
		proofType = ProofTypeExistentialContradiction
		// existential contra assumes the existential so DOESN'T require it as as preconditon
		steps = append([]Step{
			&Assumption{Formula: existential, ProofType: proofType},
			&QuantifierAssumption{
				Formula:       InstantiateFormula(existential.Body, existential.Var, termVar),
				ProofType:     proofType,
				Preconditions: []Formula{},
				From:          existential.Var,
				To:            termVar,
			},
		}, steps...)
	} else {
		proofType = ProofTypeExistentialElimination
		// existential elim DOES require the existential as a precondition
		steps = append([]Step{
			&QuantifierAssumption{
				Formula:       InstantiateFormula(existential.Body, existential.Var, termVar),
				ProofType:     proofType,
				Preconditions: []Formula{existential},
				From:          existential.Var,
				To:            termVar,
			},
		}, steps...)
	}
	// verify reiterations don't mention hypothesis
	for _, step := range steps {
		switch t := step.(type) {
		case *Reiterate:
			if containsQuantifierHypothesis(t.Formula, termVar) {
				panic(errors.Errorf("hypothesis '%s' used in reiteration '%s'", termVar, t.Formula.FormulaPrint(true)))
			}
		}
	}
	var result Formula
	last := steps[len(steps)-1]
	if isContra {
		// last step must contradict something else in scope
		err := findNegationInScope(existential, steps)
		if err != nil {
			panic(err)
		}
		// result is negation of existential assumption
		result = removeDoubleNegative(Not(existential))
	} else {
		// last step is the result and must not refer to the hypothesis
		result = last.StepResult()
		// verify hypothesis isn't used in result
		if containsQuantifierHypothesis(result, termVar) {
			panic(errors.Errorf("hypothesis '%s' used in result '%s'", termVar, result.FormulaPrint(true)))
		}
	}
	return &Proof{
		ExpectedResult: result.FormulaPrint(true),
		Steps:          steps,
		Result:         result,
		ProofType:      proofType,
		TermVar:        &termVar,
	}
}

func containsQuantifierHypothesis(formula Formula, hypothesis string) bool {
	switch t := formula.(type) {
	case *NotFormula:
		return containsQuantifierHypothesis(t.Formula, hypothesis)
	case *PredicateFormula:
		return slice.Any(func(p *TermVar) bool { return p.Name == hypothesis }, t.Terms)
	case *BinOpFormula:
		return containsQuantifierHypothesis(t.Left, hypothesis) || containsQuantifierHypothesis(t.Right, hypothesis)
	case *QuantifiedFormula:
		return containsQuantifierHypothesis(t.Body, hypothesis)
	default:
		panic(errors.Errorf("invalid formula type: %T", formula))
	}
}

func ForallIntroProof(boundVar string, hypothesisTermVar string, steps ...Step) *Proof {
	// TODO shadowing
	// verify reiterations don't mention hypothesisTermVar
	for _, step := range steps {
		switch t := step.(type) {
		case *Reiterate:
			if containsQuantifierHypothesis(t.Formula, hypothesisTermVar) {
				panic(errors.Errorf("hypothesis '%s' used in reiteration '%s'", hypothesisTermVar, t.Formula.FormulaPrint(true)))
			}
		}
	}
	// last step is the result
	last := steps[len(steps)-1]
	result := Forall(boundVar, GeneralizeFormula(last.StepResult(), hypothesisTermVar, boundVar))
	return &Proof{
		ExpectedResult: result.FormulaPrint(true),
		Steps:          steps,
		Result:         result,
		ProofType:      ProofTypeForallIntroduction,
		TermVar:        &hypothesisTermVar,
	}
}

func Reit(formula Formula) *Reiterate {
	return &Reiterate{Formula: formula}
}

// Reiterate pulls in a formula from an enclosing scope
type Reiterate struct {
	Formula Formula
}

func (r *Reiterate) StepResult() Formula {
	return r.Formula
}

func (r *Reiterate) StepName() string {
	return "Reiterate"
}

type Assumption struct {
	Formula   Formula
	ProofType ProofType
}

func (a *Assumption) StepResult() Formula {
	return a.Formula
}

func (a *Assumption) StepName() string {
	return fmt.Sprintf("Assume: %s", a.ProofType.Name())
}

type QuantifierAssumption struct {
	Formula       Formula
	ProofType     ProofType
	Preconditions []Formula
	From          string
	To            string
}

func (a *QuantifierAssumption) StepResult() Formula {
	return a.Formula
}

func (a *QuantifierAssumption) StepName() string {
	return fmt.Sprintf("Assume: %s [%s -> %s]", a.ProofType.Name(), a.From, a.To)
}
