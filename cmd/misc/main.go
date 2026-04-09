package main

import (
	"fmt"
	"strings"

	firstorderproofs "github.com/mattfenwick/algorithms/pkg/logic/firstorder/proofs"
	"github.com/mattfenwick/algorithms/pkg/pratt"
	"github.com/mattfenwick/algorithms/pkg/schema"
	"github.com/mattfenwick/algorithms/pkg/utils"
	"github.com/mattfenwick/collections/pkg/slice"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func main() {
	command := SetupRootCommand()
	utils.Die0(errors.Wrapf(command.Execute(), "run root command"))
}

type RootFlags struct {
	Verbosity string
}

func SetupRootCommand() *cobra.Command {
	flags := &RootFlags{}
	command := &cobra.Command{
		Use: "",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return SetUpLogger(flags.Verbosity)
		},
	}

	command.PersistentFlags().StringVarP(&flags.Verbosity, "verbosity", "v", "info", "log level; one of [info, debug, trace, warn, error, fatal, panic]")

	command.AddCommand(SetupPrattCommand())
	command.AddCommand(SetupSchemaCommand())
	command.AddCommand(SetupLogicCommand())
	command.AddCommand(SetupLinearAlgebraCommand())

	return command
}

type PrattArgs struct {
	Expr string
}

func SetupPrattCommand() *cobra.Command {
	args := &PrattArgs{}

	command := &cobra.Command{
		Use: "pratt",
		Run: func(cmd *cobra.Command, as []string) {
			RunPratt(args.Expr)
		},
	}

	command.Flags().StringVar(&args.Expr, "expr", "", "expression")

	return command
}

func RunPratt(s string) {
	pratt.Run(s)
}

type SchemaArgs struct {
	Path    string
	ReadAll bool
}

func SetupSchemaCommand() *cobra.Command {
	args := &SchemaArgs{}

	command := &cobra.Command{
		Use: "schema",
		Run: func(cmd *cobra.Command, as []string) {
			RunSchema(args.Path, args.ReadAll)
		},
	}

	command.Flags().BoolVar(&args.ReadAll, "all", false, "if true, ignore extension and treat all files as either json or jsonl")
	command.Flags().StringVar(&args.Path, "path", "", "root path to traverse, under which json and jsonl files will be analyzed")
	command.MarkFlagRequired("path")

	return command
}

func RunSchema(path string, readAll bool) {
	schema.Run(path, readAll)
}

type LogicArgs struct {
	ProofPath string
}

func SetupLogicCommand() *cobra.Command {
	args := &LogicArgs{}

	command := &cobra.Command{
		Use: "logic",
		Run: func(cmd *cobra.Command, as []string) {
			RunLogic(args)
		},
	}
	command.Flags().StringVar(&args.ProofPath, "proof-path", "pkg/logic/proofs", "dir path to which markdown file of proofs will be written")

	return command
}

func RunLogic(args *LogicArgs) {
	firstorderproofs.Run()
}

type LinearAlgebraArgs struct {
}

func SetupLinearAlgebraCommand() *cobra.Command {
	args := &LinearAlgebraArgs{}

	command := &cobra.Command{
		Use: "linear",
		Run: func(cmd *cobra.Command, as []string) {
			RunLinearAlgebra(args)
		},
	}

	return command
}

type Matrix struct {
	// From string
	// To string
	Rows [][]string
}

func (m *Matrix) RowCount() int {
	return len(m.Rows)
}

func (m *Matrix) ColumnCount() int {
	return len(m.Rows[0])
}

func (m *Matrix) Print() string {
	header := slice.Replicate(m.ColumnCount(), "")
	return utils.NewTable(header, m.Rows...).ToFormattedTable(nil)
}

// a, b are matrices: a*b
func MatrixMultiply(a, b *Matrix) *Matrix {
	if a.ColumnCount() != b.RowCount() {
		utils.Die0(errors.Errorf("can't multiply -- row/col count mismatch. %d vs %d", a.ColumnCount(), b.RowCount()))
	}
	var rows [][]string
	for _, row := range a.Rows {
		var newRow []string
		for ci := 0; ci < b.ColumnCount(); ci++ {
			var out []string
			for i, ea := range row {
				eb := b.Rows[i][ci]
				out = append(out, fmt.Sprintf("%s*%s", ea, eb))
			}
			newRow = append(newRow, strings.Join(out, " + "))
		}

		rows = append(rows, newRow)
	}
	return NewMatrix(rows)
}

func NewMatrix(rows [][]string) *Matrix {
	length := len(rows[0])
	for i, r := range rows[1:] {
		if len(r) != length {
			utils.Die0(errors.Errorf("matrix must be rectangular: %d vs %d at %d", length, len(r), i))
		}
	}
	return &Matrix{Rows: rows}
}

func RunLinearAlgebra(args *LinearAlgebraArgs) {
	a := NewMatrix([][]string{
		{"a11", "a12", "a13", "a14"},
		{"a21", "a22", "a23", "a24"},
		{"a31", "a32", "a33", "a34"},
	})
	b := NewMatrix([][]string{
		{"b11", "b12"},
		{"b21", "b22"},
		{"b31", "b32"},
		{"b41", "b42"},
	})
	ab := MatrixMultiply(a, b)
	fmt.Println(ab.Print())

	c := NewMatrix([][]string{
		{"1", "2"},
		{"3", "4"},
		{"5", "6"},
	})
	d := NewMatrix([][]string{
		{"6", "5", "4", "3"},
		{"2", "1", "0", "-1"},
	})
	cd := MatrixMultiply(c, d)
	fmt.Println(cd.Print())
}

func SetUpLogger(logLevelStr string) error {
	logLevel, err := logrus.ParseLevel(logLevelStr)
	if err != nil {
		return errors.Wrapf(err, "unable to parse the specified log level: '%s'", logLevel)
	}
	logrus.SetLevel(logLevel)
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	logrus.Infof("log level set to '%s'", logrus.GetLevel())
	return nil
}
