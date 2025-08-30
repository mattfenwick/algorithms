package logic

import (
	"fmt"
	"strings"

	"github.com/mattfenwick/algorithms/pkg/utils"
	"github.com/mattfenwick/collections/pkg/dict"
	"github.com/mattfenwick/collections/pkg/slice"
	"github.com/olekukonko/tablewriter"
	"github.com/pkg/errors"
)

type EvaluatedStep struct {
	Depth          int
	LineReferences string
	Step           Step
	ScopeFormulas  []string
	TermVars       []string
}

type Scope struct {
	Depth        int
	Parent       *Scope
	TrueFormulas map[string]int
	TermVars     map[string]int
}

func NewScope(parent *Scope) *Scope {
	depth := 0
	if parent != nil {
		depth = parent.Depth + 1
	}
	return &Scope{
		Depth:        depth,
		Parent:       parent,
		TrueFormulas: map[string]int{},
		TermVars:     map[string]int{},
	}
}

func (e *Scope) AddTermVar(varName string, line int) error {
	if _, ok := e.TermVars[varName]; ok {
		return errors.Errorf("unable to add term var '%s', already present", varName)
	}
	e.TermVars[varName] = line
	return nil
}

func (e *Scope) FindTermVar(key string) (int, bool) {
	line, ok := e.TermVars[key]
	if ok {
		return line, ok
	}
	if e.Parent == nil {
		return 0, false
	}
	return e.Parent.FindTermVar(key)
}

func (e *Scope) GetTermVars() []string {
	var parentTermVars []string
	if e.Parent != nil {
		parentTermVars = append(parentTermVars, e.Parent.GetTermVars()...)
	}
	return append(parentTermVars, slice.Sort(dict.Keys(e.TermVars))...)
}

func (e *Scope) FindFormula(key string) (int, bool) {
	line, ok := e.TrueFormulas[key]
	return line, ok
}

func (e *Scope) FindFormulaInParent(key string) (int, bool) {
	if e.Parent == nil {
		return 0, false
	}
	line, ok := e.Parent.FindFormula(key)
	if ok {
		return line, true
	}
	return e.Parent.FindFormulaInParent(key)
}

func (e *Scope) AddFormula(formula string, line int) error {
	if _, ok := e.TrueFormulas[formula]; ok {
		return errors.Errorf("unable to add formula '%s', already present", formula)
	}
	e.TrueFormulas[formula] = line
	return nil
}

func (e *Scope) GetTrueFormulas() []string {
	return slice.Sort(dict.Keys(e.TrueFormulas))
}

func (e *Scope) Print(indent int) {
	fmt.Printf("%s%s\n",
		strings.Join(slice.Replicate(indent*2, " "), ""),
		strings.Join(e.GetTrueFormulas(), ","))
	if e.Parent != nil {
		e.Parent.Print(indent + 1)
	}
}

type CheckedProof struct {
	Steps []*EvaluatedStep
}

func (c *CheckedProof) Add(step *EvaluatedStep) int {
	c.Steps = append(c.Steps, step)
	return len(c.Steps)
}

func (c *CheckedProof) PrintSteps() {
	for i, step := range c.Steps {
		indent := strings.Repeat("  ", step.Depth)
		result := ""
		stepResult := step.Step.StepResult()
		if stepResult != nil {
			result = (*stepResult).FormulaPrint(true)
		}
		fmt.Printf("%s%d: %s    %s: %s\n",
			indent,
			i+1,
			result,
			step.Step.StepName(),
			step.LineReferences)
	}
}

func (c *CheckedProof) BuildStepTable() string {
	table := utils.NewTable([]string{"Line", "Formula", "Term var", "Justification", "Lines used"})
	for i, step := range c.Steps {
		indent := strings.Repeat("  | ", step.Depth)
		result := ""
		termVar := ""
		stepResult := step.Step.StepResult()
		if stepResult != nil {
			result = (*stepResult).FormulaPrint(true)
		}
		stepTermVar := step.Step.StepDefineTermVar()
		if stepTermVar != nil {
			termVar = *stepTermVar
		}
		table.AddRow([]string{
			fmt.Sprintf("%d", i+1),
			fmt.Sprintf("%s%s", indent, result),
			termVar,
			step.Step.StepName(),
			step.LineReferences,
		})
	}
	return table.ToFormattedTable(func(t *tablewriter.Table) {
		t.SetAlignment(tablewriter.ALIGN_LEFT)
	})
}

func (c *CheckedProof) BuildStepMarkdownTable() string {
	table := utils.NewTable([]string{"Line", "Formula", "Term var", "Justification", "Lines used"})
	for i, step := range c.Steps {
		indent := strings.Repeat(".   ", step.Depth)
		result := ""
		termVar := ""
		stepResult := step.Step.StepResult()
		if stepResult != nil {
			result = (*stepResult).FormulaPrint(true)
		}
		stepTermVar := step.Step.StepDefineTermVar()
		if stepTermVar != nil {
			termVar = *stepTermVar
		}

		table.AddRow([]string{
			fmt.Sprintf("%d", i+1),
			fmt.Sprintf("<pre>%s%s</pre>", indent, result),
			termVar,
			step.Step.StepName(),
			step.LineReferences,
		})
	}
	return table.ToMarkdown()
}

func CheckRootProof(proof *Proof) (*CheckedProof, error) {
	checked := &CheckedProof{}
	rootScope := NewScope(nil)
	rootScope.AddFormula("T", 0)
	err := CheckProofHelper(proof, rootScope, checked)
	if err != nil {
		return checked, err
	}
	result := *checked.Steps[len(checked.Steps)-1].Step.StepResult()
	resultString := result.FormulaPrint(true)
	if resultString != proof.ExpectedResult {
		return checked, errors.Errorf("proof expected result '%s' does not match actual result '%s'", proof.ExpectedResult, resultString)
	}
	return checked, nil
}

func CheckProofHelper(proof *Proof, parentScope *Scope, checked *CheckedProof) error {
	scope := NewScope(parentScope)
	for _, step := range proof.Steps {
		if err := CheckStep(step, scope, checked); err != nil {
			return err
		}
	}
	return nil
}

func CheckStep(step Step, scope *Scope, checked *CheckedProof) error {
	trueFormulas := scope.GetTrueFormulas()
	termVars := scope.GetTermVars()
	lineRefs := ""
	shouldAddFormula := true
	switch t := step.(type) {
	case *Assumption:
	case *Proof:
		startLine := len(checked.Steps) + 1
		if err := CheckProofHelper(t, scope, checked); err != nil {
			return err
		}
		endLine := len(checked.Steps)
		lineRefs = fmt.Sprintf("%d - %d", startLine, endLine)
	case *Reiterate:
		key := t.Formula.FormulaPrint(true)
		lineRef, ok := scope.FindFormulaInParent(key)
		if !ok {
			return errors.Errorf("missing or untrue premise '%s' in parent(s)", key)
		}
		lineRefs = fmt.Sprintf("%d", lineRef)
	case *QuantifierAssumption:
		var err error
		lineRefs, err = findLineRefs(scope, t.Preconditions, true)
		if err != nil {
			return err
		}
	case *Rule:
		var err error
		lineRefs, err = findLineRefs(scope, t.Preconditions, false)
		if err != nil {
			return err
		}
		if t.UseTermVar != nil {
			lineRef, ok := scope.FindTermVar(*t.UseTermVar)
			if !ok {
				return errors.Errorf("unable to find term var '%s' in scope", *t.UseTermVar)
			}
			if lineRefs == "" {
				lineRefs = intToString(lineRef)
			} else {
				// TODO hmm the line refs now may not be in order because the term var is looked up last
				lineRefs = fmt.Sprintf("%s, %s", lineRefs, intToString(lineRef))
			}
		}
	case *TermVar:
		addedLine := checked.Add(&EvaluatedStep{
			LineReferences: lineRefs,
			Depth:          scope.Depth,
			Step:           step,
			ScopeFormulas:  trueFormulas,
			TermVars:       termVars,
		})
		return scope.AddTermVar(t.Name, addedLine)
	default:
		return errors.Errorf("invalid step type %+v", t)
	}
	addedLine := checked.Add(&EvaluatedStep{
		LineReferences: lineRefs,
		Depth:          scope.Depth,
		Step:           step,
		ScopeFormulas:  trueFormulas,
		TermVars:       termVars,
	})
	if !shouldAddFormula {
		return nil
	}
	return scope.AddFormula((*step.StepResult()).FormulaPrint(true), addedLine)
}

func findLineRefs(scope *Scope, preconditions []Formula, findInParent bool) (string, error) {
	var linesUsed []int
	for _, n := range preconditions {
		key := n.FormulaPrint(true)
		var lineRef int
		var ok bool
		if findInParent {
			lineRef, ok = scope.FindFormulaInParent(key)
		} else {
			lineRef, ok = scope.FindFormula(key)
		}
		if !ok {
			return "", errors.Errorf("missing or untrue premise '%s'", key)
		}
		linesUsed = append(linesUsed, lineRef)
	}
	// purposely not sorting `linesUsed`, which means they can show up seemingly out of order
	//   for example, '3, 2'
	//   but this matches the actual order they're used in by the rule, so imo it's a good thing
	return strings.Join(slice.Map(intToString, linesUsed), ", "), nil
}

func intToString(i int) string {
	return fmt.Sprintf("%d", i)
}
