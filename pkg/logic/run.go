package logic

import (
	"fmt"
	"strings"

	"github.com/mattfenwick/collections/pkg/file"
)

func Run(proofPath string) {
	for _, eg := range examples {
		checked, err := CheckProof(eg)
		fmt.Printf("\n\nresult from proof '%s': %s\n", eg.ExpectedResult, err)
		if err == nil {
			checked.PrintSteps()
			fmt.Println(checked.BuildStepTable())
			fmt.Println(checked.BuildStepMarkdownTable())
		}
	}

	if err := file.WriteString(proofPath, GenerateProofsMarkdown(), 0644); err != nil {
		panic(err)
	}
}

func GenerateProofsMarkdown() string {
	toc := []string{}
	sections := []string{}
	for i, proof := range examples {
		name := proof.ExpectedResult
		toc = append(toc, fmt.Sprintf("%d. [%s](#proof-%d)", i+1, name, i))
		sections = append(sections, generateMarkdownForProof(i, proof))
	}
	return fmt.Sprintf(`
# Table of Contents

%s

%s
`, strings.Join(toc, "\n"), strings.Join(sections, "\n"))
}

func generateMarkdownForProof(index int, proof *Proof) string {
	checked := Must(CheckProof(proof))
	return fmt.Sprintf(`

# %s <a name="proof-%d"></a>

%s`, proof.ExpectedResult, index, checked.BuildStepMarkdownTable())
}

func Must[A any](a A, err error) A {
	if err != nil {
		panic(err)
	}
	return a
}
