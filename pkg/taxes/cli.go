package taxes

import (
	"github.com/mattfenwick/collections/pkg/json"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func RunRootCommand() {
	command := SetupRootCommand()
	Die(errors.Wrapf(command.Execute(), "run root command"))
}

type RootFlags struct {
	Verbosity string
}

func SetupRootCommand() *cobra.Command {
	flags := &RootFlags{}
	command := &cobra.Command{
		Use:   "taxes",
		Short: "tax estimation",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return SetUpLogger(flags.Verbosity)
		},
	}

	command.PersistentFlags().StringVarP(&flags.Verbosity, "verbosity", "v", "info", "log level; one of [info, debug, trace, warn, error, fatal, panic]")

	command.AddCommand(SetupEstimateCommand())

	return command
}

type EstimateArgs struct {
	ConfigPath string
}

func SetupEstimateCommand() *cobra.Command {
	args := &EstimateArgs{}

	command := &cobra.Command{
		Use:   "estimate",
		Short: "estimate taxes",
		Args:  cobra.ExactArgs(0),
		Run: func(cmd *cobra.Command, as []string) {
			RunEstimate(args)
		},
	}

	command.Flags().StringVar(&args.ConfigPath, "config", "", "path to config file")

	return command
}

func RunEstimate(args *EstimateArgs) {
	var incomes []*Income
	if args.ConfigPath != "" {
		incomesP, err := json.ParseFile[[]*Income](args.ConfigPath)
		Die(err)
		incomes = *incomesP
	} else {
		incomes = defaultIncomes
	}
	RunTaxes(incomes)
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

func Die(err error) {
	if err != nil {
		logrus.Fatalf("%+v", err)
	}
}
