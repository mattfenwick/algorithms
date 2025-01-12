package music

import "github.com/pkg/errors"

const (
	Natural = '♮'
	Flat    = '♭'
	Sharp   = '♯'
)

type Note struct {
	Names []string
}

func makeNotes(names ...[]string) []*Note {
	var out []*Note
	for _, n := range names {
		out = append(out, &Note{Names: n})
	}
	return out
}

var (
	Notes = makeNotes(
		[]string{"C"},
		[]string{"C♯", "D♭"},
		[]string{"D"},
		[]string{"D♯", "E♭"},
		[]string{"E"},
		[]string{"F"},
		[]string{"F♯", "G♭"},
		[]string{"G"},
		[]string{"G♯", "A♭"},
		[]string{"A"},
		[]string{"A♯", "B♭"},
		[]string{"B"},
	)

	NoteIndex = map[string]int{}
)

func init() {
	for i, n := range Notes {
		for _, name := range n.Names {
			NoteIndex[name] = i
		}
	}
}

func Step(start string, size int) int {
	startIndex, ok := NoteIndex[start]
	if !ok {
		die(errors.Errorf("invalid note %s", start))
	}
	endIndex := startIndex + size
	if endIndex < 0 {
		return endIndex + 12
	} else if endIndex >= 12 {
		return endIndex - 12
	} else {
		return endIndex
	}
}

func MajorScale(start string) []string {
	return []string{
		start,
		Notes[Step(start, 2)].Names[0],  // TODO wrong name
		Notes[Step(start, 4)].Names[0],  // TODO wrong name
		Notes[Step(start, 5)].Names[0],  // TODO wrong name
		Notes[Step(start, 7)].Names[0],  // TODO wrong name
		Notes[Step(start, 9)].Names[0],  // TODO wrong name
		Notes[Step(start, 11)].Names[0], // TODO wrong name
		start,
	}
}
