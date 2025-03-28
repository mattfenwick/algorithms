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
	command.AddCommand(SetupGenerateScalesAndChordsMarkdownCommand())
	command.AddCommand(SetupGenerateKeysMarkdownCommand())

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
			RunScalesDeprecated(args)
		},
	}

	command.Flags().StringSliceVar(&args.Scales, "start", []string{"C", "G"}, "scale starting points")

	return command
}

func RunScalesDeprecated(args *ScalesArgs) {
	for _, start := range args.Scales {
		key, ok := music.StringToKey[start]
		if !ok {
			Die(errors.Errorf("invalid key: %s", start))
		}
		for _, n := range music.ScaleMajor.Apply(key) {
			fmt.Println(n)
		}
		fmt.Println()
		for _, n := range music.ScaleMinor.Apply(key) {
			fmt.Println(n)
		}
		fmt.Println()
	}

	var majorRows, minorRows [][]string
	for _, k := range music.KeysByFifths {
		majorRow := slice.Cons(fmt.Sprintf("key: %s", k.Start.String()), slice.Map(noteToString, music.ScaleMajor.Apply(k)))
		majorRows = append(majorRows, majorRow)
		minorRow := slice.Cons(fmt.Sprintf("key: %s", k.Start.String()), slice.Map(noteToString, music.ScaleMinor.Apply(k)))
		minorRows = append(minorRows, minorRow)
	}
	fmt.Println("major scales:")
	majorTable := utils.NewTable([]string{"", "1", "2", "3", "4", "5", "6", "7", "8"}, majorRows...)
	fmt.Println(majorTable.ToFormattedTable())
	fmt.Println("minor scales:")
	minorTable := utils.NewTable([]string{"", "1", "2", "3", "4", "5", "6", "7", "8"}, minorRows...)
	fmt.Println(minorTable.ToFormattedTable())

	for _, k := range music.KeysByFifths {
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
				progressionTable.AddRow(slice.Cons(progression.Chords[i].Name(k), slice.Map(noteToString, notes)))
			}
			fmt.Println(progressionTable.ToFormattedTable())
		}
	}
}

func noteToString(n *music.Note) string {
	return n.String()
}

func SetupGenerateScalesAndChordsMarkdownCommand() *cobra.Command {
	command := &cobra.Command{
		Use:  "scalesandchords",
		Args: cobra.ExactArgs(0),
		Run: func(cmd *cobra.Command, as []string) {
			RunGenerateScalesAndChordsMarkdown()
		},
	}

	return command
}

func RunGenerateScalesAndChordsMarkdown() {
	header := []string{"Chord", "", "", "", "", "Description"}
	var rows [][]string
	for _, chord := range music.Chords {
		row := slice.Cons("X"+chord.BaseSymbol, slice.Map(func(s *music.Step) string { return s.Name }, chord.Steps))
		for len(row) < 5 {
			row = append(row, "")
		}
		row = append(row, chord.Name)
		rows = append(rows, row)
	}
	chordTable := utils.NewTable(header, rows...)
	fmt.Printf("\n\n%s\n", chordTable.ToMarkdown())

	scaleHeader := slice.Cons("Scale", slice.Replicate(8, ""))
	var scaleRows [][]string
	for _, scale := range music.Scales {
		row := slice.Cons(scale.Name, slice.Map(func(s *music.Step) string { return s.Name }, scale.Steps))
		for len(row) < len(scaleHeader) {
			row = append(row, "")
		}
		scaleRows = append(scaleRows, row)
	}
	scaleTable := utils.NewTable(scaleHeader, scaleRows...)
	fmt.Printf("\n\n")
	fmt.Println(scaleTable.ToMarkdown())
}

func SetupGenerateKeysMarkdownCommand() *cobra.Command {
	var order string
	command := &cobra.Command{
		Use:  "keys",
		Args: cobra.ExactArgs(0),
		Run: func(cmd *cobra.Command, as []string) {
			RunGenerateKeysMarkdown(order)
		},
	}

	command.Flags().StringVar(&order, "order", "fifths", "fifths|chromatic")

	return command
}

func RunGenerateKeysMarkdown(order string) {
	var keys []*music.Key
	switch order {
	case "fifths":
		keys = music.KeysByFifths
	case "chromatic":
		keys = music.KeysChromatic
	default:
		Die(errors.Errorf("invalid order: %s", order))
	}
	toc := []string{}
	for i, key := range keys {
		name := key.Start.String()
		toc = append(toc, fmt.Sprintf("%d. [%s](#%s)", i+1, name, strings.ToLower(name)))
	}
	sections := strings.Join(slice.Map(generateMarkdownForKey, keys), "\n")
	fmt.Printf(`
# Table of Contents

%s

%s
`, strings.Join(toc, "\n"), sections)
}

func generateMarkdownForKey(key *music.Key) string {
	scaleHeader := []string{"", "1", "2", "3", "4", "5", "6", "7", "8"}
	var scaleRows [][]string
	for _, scale := range music.Scales {
		row := slice.Cons(scale.Name, slice.Map(noteToString, scale.Apply(key)))
		for len(row) < len(scaleHeader) {
			row = append(row, "")
		}
		scaleRows = append(scaleRows, row)
	}
	scales := utils.NewTable(scaleHeader, scaleRows...).ToMarkdown()

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
		header := []string{progression.Name, "", "", "", ""}
		progressionTable := utils.NewTable(header)
		progressionNotes := progression.Apply(key)
		for i, notes := range progressionNotes {
			row := slice.Cons(progression.Chords[i].Name(key), slice.Map(noteToString, notes))
			for len(row) < len(header) {
				row = append(row, "")
			}
			progressionTable.AddRow(row)
		}
		progressionTables = append(progressionTables, progressionTable.ToMarkdown())
	}
	progressions := strings.Join(progressionTables, "\n\n")

	return fmt.Sprintf(`

# %s <a name="%s"></a>

Key signature: %s

## Scales

%s

## Chords

%s

## Progressions

%s`, key.Start.String(), strings.ToLower(key.Start.String()), key.Signature(), scales, chords, progressions)
}
