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
	ScopeTerms     []string
}

type Scope struct {
	Depth     int
	Parent    *Scope
	TrueTerms map[string]int
}

func NewScope(parent *Scope) *Scope {
	depth := 0
	if parent != nil {
		depth = parent.Depth + 1
	}
	return &Scope{
		Depth:     depth,
		Parent:    parent,
		TrueTerms: map[string]int{}}
}

func (e *Scope) Find(key string) (int, bool) {
	line, ok := e.TrueTerms[key]
	return line, ok
}

func (e *Scope) FindInParent(key string) (int, bool) {
	if e.Parent == nil {
		return 0, false
	}
	line, ok := e.Parent.Find(key)
	if ok {
		return line, true
	}
	return e.Parent.FindInParent(key)
}

func (e *Scope) Add(key string, line int) error {
	if _, ok := e.TrueTerms[key]; ok {
		return errors.Errorf("unable to add key '%s', already present", key)
	}
	e.TrueTerms[key] = line
	return nil
}

func (e *Scope) GetTrueTerms() []string {
	return slice.Sort(dict.Keys(e.TrueTerms))
}

func (e *Scope) Print(indent int) {
	fmt.Printf("%s%s\n",
		strings.Join(slice.Replicate(indent*2, " "), ""),
		strings.Join(e.GetTrueTerms(), ","))
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
		fmt.Printf("%s%d: %s    %s: %s\n",
			indent,
			i+1,
			step.Step.StepResult().TermPrint(true),
			step.Step.StepName(),
			step.LineReferences)
	}
}

func (c *CheckedProof) BuildStepTable() string {
	table := utils.NewTable([]string{"Line", "Formula", "Justification", "Lines used"})
	for i, step := range c.Steps {
		indent := strings.Repeat("  | ", step.Depth)
		table.AddRow([]string{
			fmt.Sprintf("%d", i+1),
			fmt.Sprintf("%s%s", indent, step.Step.StepResult().TermPrint(true)),
			step.Step.StepName(),
			step.LineReferences,
		})
	}
	return table.ToFormattedTable(func(t *tablewriter.Table) {
		t.SetAlignment(tablewriter.ALIGN_LEFT)
	})
}

func (c *CheckedProof) BuildStepMarkdownTable() string {
	table := utils.NewTable([]string{"Line", "Formula", "Justification", "Lines used"})
	for i, step := range c.Steps {
		indent := strings.Repeat(".   ", step.Depth)
		table.AddRow([]string{
			fmt.Sprintf("%d", i+1),
			fmt.Sprintf("<pre>%s%s</pre>", indent, step.Step.StepResult().TermPrint(true)),
			step.Step.StepName(),
			step.LineReferences,
		})
	}
	return table.ToMarkdown()
}

func CheckProof(proof *Proof) (*CheckedProof, error) {
	checked := &CheckedProof{}
	err := CheckProofHelper(proof, nil, checked)
	if err != nil {
		return checked, err
	}
	result := checked.Steps[len(checked.Steps)-1].Step.StepResult().TermPrint(true)
	if result != proof.Name {
		return checked, errors.Errorf("proof name '%s' does not match result '%s'", proof.Name, result)
	}
	return checked, nil
}

func CheckProofHelper(proof *Proof, parentScope *Scope, checked *CheckedProof) error {
	scope := NewScope(parentScope)
	if proof.Hypothesis != nil {
		if err := CheckStep(&Assumption{Term: proof.Hypothesis}, scope, checked); err != nil {
			return err
		}
	}
	for _, step := range proof.Steps {
		if err := CheckStep(step, scope, checked); err != nil {
			return err
		}
	}
	return nil
}

func CheckStep(step Step, scope *Scope, checked *CheckedProof) error {
	trueTerms := scope.GetTrueTerms()
	lineRefs := ""
	shouldAdd := true
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
		key := t.Term.TermPrint(true)
		lineRef, ok := scope.FindInParent(key)
		if !ok {
			return errors.Errorf("missing or untrue premise '%s' in parent(s)", key)
		}
		lineRefs = fmt.Sprintf("%d", lineRef)
	case *Repeat:
		key := t.Term.TermPrint(true)
		lineRef, ok := scope.Find(key)
		if !ok {
			return errors.Errorf("missing or untrue premise '%s'", key)
		}
		lineRefs = fmt.Sprintf("%d", lineRef)
		shouldAdd = false
	case *Rule:
		var linesUsed []int
		for _, n := range t.Preconditions {
			key := n.TermPrint(true)
			lineRef, ok := scope.Find(key)
			if !ok {
				return errors.Errorf("missing or untrue premise '%s'", key)
			}
			linesUsed = append(linesUsed, lineRef)
		}
		// purposely not sorting `linesUsed`, which means they can show up seemingly out of order
		//   for example, '3, 2'
		//   but this matches the actual order they're used in by the rule, so imo it's a good thing
		lineRefs = strings.Join(slice.Map(intToString, linesUsed), ", ")
	default:
		return errors.Errorf("invalid step type %+v", t)
	}
	addedLine := checked.Add(&EvaluatedStep{LineReferences: lineRefs, Depth: scope.Depth, Step: step, ScopeTerms: trueTerms})
	if !shouldAdd {
		return nil
	}
	return scope.Add(step.StepResult().TermPrint(true), addedLine)
}

func intToString(i int) string {
	return fmt.Sprintf("%d", i)
}
