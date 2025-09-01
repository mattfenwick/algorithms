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
	Rx := Pred("R", &TermVar{Name: "x", IsBound: true})
	formulas := [][]Formula{
		{
			Exist("x", P),
			Exist("x", Not(P)),
			Not(Exist("x", P)),
			Not(Exist("x", Not(P))),
			Forall("x", P),
			Forall("x", Not(P)),
			Not(Forall("x", P)),
			Not(Forall("x", Not(P))),
		},
		{
			Exist("x", Qx),
			Exist("x", Not(Qx)),
			Not(Exist("x", Qx)),
			Not(Exist("x", Not(Qx))),
			Forall("x", Qx),
			Forall("x", Not(Qx)),
			Not(Forall("x", Qx)),
			Not(Forall("x", Not(Qx))),
		},
		{
			Exist("x", Rx),
			Exist("x", Not(Rx)),
			Not(Exist("x", Rx)),
			Not(Exist("x", Not(Rx))),
			Forall("x", Rx),
			Forall("x", Not(Rx)),
			Not(Forall("x", Rx)),
			Not(Forall("x", Not(Rx))),
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
	writeOutProofs(quantifierProofs, path.Join(proofDir, "quantifier-proofs.md"))
	writeOutProofs(propositionalProofSections, path.Join(proofDir, "propositional-proofs.md"))
}

func writeOutProofs(proofSections []*ProofsSection, proofPath string) {
	for _, section := range proofSections {
		fmt.Printf("working on section %s\n", section.Name)
		for _, eg := range section.Proofs {
			checked, err := CheckRootProof(eg)
			fmt.Printf("\n\nresult from proof '%s': %s\n", eg.ExpectedResult, err)
			if err == nil {
				checked.PrintSteps()
				fmt.Println(checked.BuildStepTable())
				fmt.Println(checked.BuildStepMarkdownTable())
			}
		}
	}

	if err := file.WriteString(proofPath, GenerateProofsMarkdown(proofSections), 0644); err != nil {
		panic(err)
	}
}

func GenerateProofsMarkdown(proofSections []*ProofsSection) string {
	toc := []string{}
	sections := []string{}
	for i, proofSection := range proofSections {
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
	checked, err := CheckRootProof(proof)
	if err != nil {
		fmt.Println(checked.BuildStepTable())
		panic(err)
	}
	return fmt.Sprintf(`## %s <a name="proof-%d-%d"></a>

%s
`, proof.ExpectedResult, sectionIndex, proofIndex, checked.BuildStepMarkdownTable())
}
