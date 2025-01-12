package music

import (
	"github.com/mattfenwick/collections/pkg/slice"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type Note string

const (
	CFlat  Note = "C♭"
	C      Note = "C"
	CSharp Note = "C♯"
	DFlat  Note = "D♭"
	D      Note = "D"
	DSharp Note = "D♯"
	EFlat  Note = "E♭"
	E      Note = "E"
	ESharp Note = "E♯"
	FFlat  Note = "F♭"
	F      Note = "F"
	FSharp Note = "F♯"
	GFlat  Note = "G♭"
	G      Note = "G"
	GSharp Note = "G♯"
	AFlat  Note = "A♭"
	A      Note = "A"
	ASharp Note = "A♯"
	BFlat  Note = "B♭"
	B      Note = "B"
	BSharp Note = "B♯"
)

var naturals = []Note{
	C,
	D,
	E,
	F,
	G,
	A,
	B,
}

var naturalToIndex = map[Note]int{}

var notes = [][]Note{
	{BSharp, C},
	{CSharp, DFlat},
	{D},
	{DSharp, EFlat},
	{E, FFlat},
	{ESharp, F},
	{FSharp, GFlat},
	{G},
	{GSharp, AFlat},
	{A},
	{ASharp, BFlat},
	{B, CFlat},
}

var noteToIndex = map[Note]int{}

func (n Note) Index() int {
	return noteToIndex[n]
}

func (n Note) Step(size int) Note {
	startIndex := noteToIndex[n]
	endIndex := startIndex + size
	var potentialNotes []Note
	if endIndex < 0 {
		potentialNotes = notes[endIndex+12]
	} else if endIndex >= 12 {
		potentialNotes = notes[endIndex-12]
	} else {
		potentialNotes = notes[endIndex]
	}
	nextNatural := n.NextNatural()
	for _, p := range potentialNotes {
		if p.Natural() == nextNatural {
			return p
		}
	}
	panic(errors.Errorf("no natural match found for %s (from %s, %d), looking in %+v", nextNatural, n, size, potentialNotes))
}

func (n Note) NextNatural() Note {
	index := naturalToIndex[n.Natural()] + 1
	return naturals[index%len(naturals)]
}

func (n Note) Natural() Note {
	switch n {
	case A, AFlat, ASharp:
		return A
	case B, BFlat, BSharp:
		return B
	case C, CFlat, CSharp:
		return C
	case D, DFlat, DSharp:
		return D
	case E, EFlat, ESharp:
		return E
	case F, FFlat, FSharp:
		return F
	case G, GFlat, GSharp:
		return G
	default:
		panic(errors.Errorf("invalid note: %s", n))
	}
}

var (
	Sharps = []string{"F", "C", "G", "D", "A", "E", "B"}
	Flats  = slice.Reverse(Sharps)
)

type Key struct {
	Start  Note
	Sharps int
	Flats  int
}

var KeySignatures = []*Key{
	{GFlat, 0, 6},
	{DFlat, 0, 5},
	{AFlat, 0, 4},
	{EFlat, 0, 3},
	{BFlat, 0, 2},
	{F, 0, 1},
	{C, 0, 0},
	{G, 1, 0},
	{D, 2, 0},
	{A, 3, 0},
	{E, 4, 0},
	{B, 5, 0},
	{FSharp, 6, 0},
	{CSharp, 7, 0},
}

var NoteToKey = map[Note]*Key{}

var (
	MajorSteps = []int{2, 2, 1, 2, 2, 2, 1}
	MinorSteps = []int{2, 1, 2, 2, 1, 2, 2}
)

func (k *Key) getScale(steps []int) []Note {
	logrus.Infof("looking at key of %s", k.Start)
	out := []Note{k.Start}
	curr := k.Start
	for _, step := range steps {
		next := curr.Step(step)
		out = append(out, next)
		curr = next
	}
	return out
}

func (k *Key) MajorScale() []Note {
	return k.getScale(MajorSteps)
}

func (k *Key) MinorScale() []Note {
	return k.getScale(MinorSteps)
}

func init() {
	for i, ns := range notes {
		for _, n := range ns {
			noteToIndex[n] = i
		}
	}
	for _, k := range KeySignatures {
		NoteToKey[k.Start] = k
	}
	for i, n := range naturals {
		naturalToIndex[n] = i
	}
}
