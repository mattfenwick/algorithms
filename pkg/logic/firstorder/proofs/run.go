package proofs

import (
	"fmt"
	"os"
	"path"
	"strings"

	. "github.com/mattfenwick/algorithms/pkg/logic/firstorder"
	"github.com/mattfenwick/collections/pkg/file"
	"github.com/olekukonko/tablewriter"
)

func printSomeTruthTables() {
	formulas := [][]Formula{
		{
			Existential("x", Prop("P")),
			Existential("x", Not(Prop("P"))),
			Not(Existential("x", Prop("P"))),
			Not(Existential("x", Not(Prop("P")))),
			Forall("x", Prop("P")),
			Forall("x", Not(Prop("P"))),
			Not(Forall("x", Prop("P"))),
			Not(Forall("x", Not(Prop("P")))),
		},
		{
			Existential("x", Prop("Q", "x")),
			Existential("x", Not(Prop("Q", "x"))),
			Not(Existential("x", Prop("Q", "x"))),
			Not(Existential("x", Not(Prop("Q", "x")))),
			Forall("x", Prop("Q", "x")),
			Forall("x", Not(Prop("Q", "x"))),
			Not(Forall("x", Prop("Q", "x"))),
			Not(Forall("x", Not(Prop("Q", "x")))),
		},
		{
			Existential("x", Prop("R", "x")),
			Existential("x", Not(Prop("R", "x"))),
			Not(Existential("x", Prop("R", "x"))),
			Not(Existential("x", Not(Prop("R", "x")))),
			Forall("x", Prop("R", "x")),
			Forall("x", Not(Prop("R", "x"))),
			Not(Forall("x", Prop("R", "x"))),
			Not(Forall("x", Not(Prop("R", "x")))),
		},
	}
	domains := [][]string{
		{},
		{"a", "b", "c"},
	}
	// for it, formula := range formulas {
	// 	for id, d := range domains {
	// 		env := NewEnv(map[string]bool{
	// 			"Q(a)": true,
	// 			"Q(b)": true,
	// 			"Q(c)": true,
	// 			"R(a)": true,
	// 			"R(b)": false,
	// 			"R(c)": false,
	// 		}, d...)
	// 		fmt.Printf("[%d, %d] for domain %+v, %s:\n", it, id, d, formula.formulaPrint(true))
	// 		fmt.Println(TruthTable(formula, env).ToFormattedTable(func(t *tablewriter.Table) {
	// 			t.SetAlignment(tablewriter.ALIGN_CENTER)
	// 			t.SetAutoFormatHeaders(false)
	// 		}))
	// 	}
	// }
	for id, d := range domains {
		env := NewEnv(map[string]bool{
			"P":    true,
			"Q(a)": true,
			"Q(b)": true,
			"Q(c)": true,
			"R(a)": true,
			"R(b)": false,
			"R(c)": false,
		}, d...)
		fmt.Printf("[%d] for domain %+v\n", id, d)
		fmt.Println(TruthTableFromFormulas(env, formulas...).ToFormattedTable(func(t *tablewriter.Table) {
			t.SetAlignment(tablewriter.ALIGN_CENTER)
			t.SetAutoFormatHeaders(false)
		}))
	}
}

func Run() {
	if false {
		// TODO what to do with this?
		printSomeTruthTables()
		os.Exit(1)
	}

	proofDir := "pkg/logic/firstorder/proofs"
	proofPath := path.Join(proofDir, "proofs.md")
	for _, section := range proofs {
		fmt.Printf("working on section %s\n", section.Name)
		for _, eg := range section.Proofs {
			checked, err := CheckProof(eg)
			fmt.Printf("\n\nresult from proof '%s': %s\n", eg.ExpectedResult, err)
			if err == nil {
				checked.PrintSteps()
				fmt.Println(checked.BuildStepTable())
				fmt.Println(checked.BuildStepMarkdownTable())
			}
		}
	}

	if err := file.WriteString(proofPath, GenerateProofsMarkdown(), 0644); err != nil {
		panic(err)
	}
}

func GenerateProofsMarkdown() string {
	toc := []string{}
	sections := []string{}
	for i, proofSection := range proofs {
		sectionIndex := i + 1
		toc = append(toc, fmt.Sprintf("%d. [%s](#%s)", sectionIndex, proofSection.Name, proofSection.Name))
		sections = append(sections, fmt.Sprintf(`# %s <a name="%s"></a>`+"\n", proofSection.Name, proofSection.Name))
		for j, proof := range proofSection.Proofs {
			name := proof.ExpectedResult
			proofIndex := j + 1
			toc = append(toc, fmt.Sprintf("    %d. [%s](#proof-%d-%d)", proofIndex, name, sectionIndex, proofIndex))
			sections = append(sections, generateMarkdownForProof(sectionIndex, proofIndex, proof))
		}
	}
	return fmt.Sprintf(`
# Table of Contents

%s

%s
`, strings.Join(toc, "\n"), strings.Join(sections, "\n"))
}

func generateMarkdownForProof(sectionIndex int, proofIndex int, proof *Proof) string {
	checked, err := CheckProof(proof)
	if err != nil {
		fmt.Println(checked.BuildStepTable())
		panic(err)
	}
	return fmt.Sprintf(`## %s <a name="proof-%d-%d"></a>

%s
`, proof.ExpectedResult, sectionIndex, proofIndex, checked.BuildStepMarkdownTable())
}
