package cli

import (
	"fmt"

	"github.com/mattfenwick/algorithms/pkg/music"
	"github.com/mattfenwick/algorithms/pkg/utils"
	"github.com/mattfenwick/collections/pkg/slice"
	"github.com/pkg/errors"
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
		Use: "music",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return SetUpLogger(flags.Verbosity)
		},
	}

	command.PersistentFlags().StringVarP(&flags.Verbosity, "verbosity", "v", "info", "log level; one of [info, debug, trace, warn, error, fatal, panic]")

	command.AddCommand(SetupScalesCommand())

	return command
}

type ScalesArgs struct {
	Scales []string
}

func SetupScalesCommand() *cobra.Command {
	args := &ScalesArgs{}

	command := &cobra.Command{
		Use:  "scales",
		Args: cobra.ExactArgs(0),
		Run: func(cmd *cobra.Command, as []string) {
			RunScales(args)
		},
	}

	command.Flags().StringSliceVar(&args.Scales, "start", []string{"C", "G"}, "scale starting points")

	return command
}

func RunScales(args *ScalesArgs) {
	for _, start := range args.Scales {
		key, ok := music.StringToKey[start]
		if !ok {
			Die(errors.Errorf("invalid key: %s", start))
		}
		for _, n := range key.MajorScale() {
			fmt.Println(n)
		}
		fmt.Println()
		for _, n := range key.MinorScale() {
			fmt.Println(n)
		}
		fmt.Println()
	}

	var majorRows, minorRows [][]string
	for _, k := range music.KeySignatures {
		majorRow := slice.Cons(fmt.Sprintf("key: %s", k.Start.String()), slice.Map(noteToString, k.MajorScale()))
		majorRows = append(majorRows, majorRow)
		minorRow := slice.Cons(fmt.Sprintf("key: %s", k.Start.String()), slice.Map(noteToString, k.MinorScale()))
		minorRows = append(minorRows, minorRow)
	}
	fmt.Println("major scales:")
	majorTable := utils.NewTable([]string{"", "1", "2", "3", "4", "5", "6", "7", "8"}, majorRows...)
	fmt.Println(majorTable.ToFormattedTable())
	fmt.Println("minor scales:")
	minorTable := utils.NewTable([]string{"", "1", "2", "3", "4", "5", "6", "7", "8"}, minorRows...)
	fmt.Println(minorTable.ToFormattedTable())

	for _, k := range music.KeySignatures {
		var rows [][]string
		for _, c := range []*music.Chord{music.ChordMajorTriad, music.ChordMinorTriad, music.ChordAugmentedTriad, music.ChordDiminishedTriad} {
			rows = append(rows, slice.Cons(c.Name, slice.Map(noteToString, c.Apply(k))))
		}
		fmt.Println(utils.NewTable([]string{k.Start.String(), "1st", "3rd", "5th"}, rows...).ToFormattedTable())
	}
}

func noteToString(n *music.Note) string {
	return n.String()
}
