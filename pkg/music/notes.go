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

func (n Note) Step(size int, useFlats bool) Note {
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
	if len(potentialNotes) == 2 && useFlats {
		return potentialNotes[1]
	}
	return potentialNotes[0]
}

func (n Note) Natural() Note {
	switch n {
	case A, AFlat, ASharp:
		return A
	case B, BFlat:
		return B
	case C, CSharp:
		return C
	case D, DFlat, DSharp:
		return D
	case E, EFlat:
		return E
	case F, FSharp:
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

var MajorSteps = []int{2, 2, 1, 2, 2, 2, 1}

func (k *Key) MajorScale() []Note {
	logrus.Infof("looking at key of %s", k.Start)
	useFlats := k.Flats > 0
	out := []Note{k.Start}
	curr := k.Start
	for _, step := range MajorSteps {
		next := curr.Step(step, useFlats)
		out = append(out, next)
		curr = next
	}
	return out
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
}
