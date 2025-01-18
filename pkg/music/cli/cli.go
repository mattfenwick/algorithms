package cli

import (
	"fmt"
	"strings"

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
	command.AddCommand(SetupGenerateMarkdownCommand())

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
	for _, k := range music.Keys {
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

	for _, k := range music.Keys {
		var rows [][]string
		for _, c := range music.Chords {
			prefix := []string{c.Symbol(k.Start), c.Name}
			row := append(prefix, slice.Map(noteToString, c.Apply(k))...)
			for len(row) < 6 {
				row = append(row, "")
			}
			rows = append(rows, row)
		}
		fmt.Println(utils.NewTable([]string{k.Start.String(), "", "", "", "", ""}, rows...).ToFormattedTable())

		for _, progression := range music.Progressions {
			progressionTable := utils.NewTable([]string{progression.Name, "", "", ""})
			progressionNotes := progression.Apply(k)
			for i, notes := range progressionNotes {
				progressionTable.AddRow(slice.Cons(progression.Chords[i].Name(), slice.Map(noteToString, notes)))
			}
			fmt.Println(progressionTable.ToFormattedTable())
		}
	}
}

func noteToString(n *music.Note) string {
	return n.String()
}

func SetupGenerateMarkdownCommand() *cobra.Command {
	command := &cobra.Command{
		Use:  "markdown",
		Args: cobra.ExactArgs(0),
		Run: func(cmd *cobra.Command, as []string) {
			RunGenerateMarkdown()
		},
	}

	return command
}

func RunGenerateMarkdown() {
	output := strings.Join(slice.Map(generateMarkdownForKey, music.Keys), "\n")
	fmt.Println(output)
}

func generateMarkdownForKey(key *music.Key) string {
	scales := utils.NewTable([]string{"", "1", "2", "3", "4", "5", "6", "7", "8"},
		slice.Cons("Major", slice.Map(noteToString, key.MajorScale())),
		slice.Cons("Minor", slice.Map(noteToString, key.MinorScale())),
	).ToMarkdown()

	var chordRows [][]string
	for _, c := range music.Chords {
		row := slice.Cons(c.Symbol(key.Start), slice.Map(noteToString, c.Apply(key)))
		for len(row) < 5 {
			row = append(row, "")
		}
		row = append(row, c.Name)
		chordRows = append(chordRows, row)
	}
	chords := utils.NewTable([]string{"Name", "", "", "", "", "Description"}, chordRows...).ToMarkdown()

	progressionTables := []string{}
	for _, progression := range music.Progressions {
		progressionTable := utils.NewTable([]string{progression.Name, "", "", ""})
		progressionNotes := progression.Apply(key)
		for i, notes := range progressionNotes {
			progressionTable.AddRow(slice.Cons(progression.Chords[i].Name(), slice.Map(noteToString, notes)))
		}
		progressionTables = append(progressionTables, progressionTable.ToMarkdown())
	}
	progressions := strings.Join(progressionTables, "\n\n")

	return fmt.Sprintf(`# %s

Key signature: %s

## Scales

%s

## Chords

%s

## Progressions

%s`, key.Start.String(), key.Signature(), scales, chords, progressions)
}
