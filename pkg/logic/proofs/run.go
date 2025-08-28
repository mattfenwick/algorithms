package proofs

import (
	"fmt"
	"path"
	"strings"

	"github.com/mattfenwick/algorithms/pkg/logic"
	. "github.com/mattfenwick/algorithms/pkg/logic"
	"github.com/mattfenwick/collections/pkg/file"
	"github.com/olekukonko/tablewriter"
)

func printSomeTruthTables() {
	terms := []Term{
		Implication(P, Or(Q, R)),
		Or(Implication(P, Q), Implication(P, R)),
		Implication(And(P, Q), R),
		Or(Implication(P, R), Implication(Q, R)),
	}
	for _, t := range terms {
		fmt.Println(TruthTable(t).ToFormattedTable(func(t *tablewriter.Table) {
			t.SetAlignment(tablewriter.ALIGN_CENTER)
		}))
	}
}

func Run(proofDir string) {
	if false {
		// TODO what to do with this?
		printSomeTruthTables()
	}

	proofPath := path.Join(proofDir, "propositional-long-proofs.md")
	for _, section := range propositionalLongProofSections {
		fmt.Printf("working on section %s\n", section.Name)
		for _, eg := range section.Proofs {
			checked, err := logic.CheckProof(eg)
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
	for i, proofSection := range propositionalLongProofSections {
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

func generateMarkdownForProof(sectionIndex int, proofIndex int, proof *logic.Proof) string {
	checked, err := logic.CheckProof(proof)
	if err != nil {
		fmt.Println(checked.BuildStepTable())
		panic(err)
	}
	return fmt.Sprintf(`## %s <a name="proof-%d-%d"></a>

%s
`, proof.ExpectedResult, sectionIndex, proofIndex, checked.BuildStepMarkdownTable())
}
