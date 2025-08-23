package main

import (
	"github.com/mattfenwick/algorithms/pkg/logic/proofs"
	"github.com/mattfenwick/algorithms/pkg/pratt"
	"github.com/mattfenwick/algorithms/pkg/schema"
	"github.com/mattfenwick/algorithms/pkg/utils"
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
	proofs.Run(args.ProofPath)
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
