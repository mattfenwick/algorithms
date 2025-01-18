package music

import "github.com/sirupsen/logrus"

type Key struct {
	Start  *Note
	Sharps int
	Flats  int
}

var KeySignatures = []*Key{
	{NewNote(NaturalG, 0, 1), 0, 6},
	{NewNote(NaturalD, 0, 1), 0, 5},
	{NewNote(NaturalA, 0, 1), 0, 4},
	{NewNote(NaturalE, 0, 1), 0, 3},
	{NewNote(NaturalB, 0, 1), 0, 2},
	{NewNote(NaturalF, 0, 0), 0, 1},
	{NewNote(NaturalC, 0, 0), 0, 0},
	{NewNote(NaturalG, 0, 0), 1, 0},
	{NewNote(NaturalD, 0, 0), 2, 0},
	{NewNote(NaturalA, 0, 0), 3, 0},
	{NewNote(NaturalE, 0, 0), 4, 0},
	{NewNote(NaturalB, 0, 0), 5, 0},
	{NewNote(NaturalF, 1, 0), 6, 0},
	{NewNote(NaturalC, 2, 0), 7, 0},
}

var StringToKey = map[string]*Key{}

var (
	MajorSteps = []*Step{
		NewStep(0, 0),
		NewStep(2, 1),
		NewStep(4, 2),
		NewStep(5, 3),
		NewStep(7, 4),
		NewStep(9, 5),
		NewStep(11, 6),
		NewStep(12, 7),
	}
	MinorSteps = []*Step{
		{0, 0},
		{2, 1},
		{3, 2},
		{5, 3},
		{7, 4},
		{8, 5},
		{10, 6},
		{12, 7},
	}

// ChromaticSteps = []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
)

func (k *Key) getScale(steps []*Step) []*Note {
	logrus.Debugf("looking at key of %s", k.Start)
	return getNotesFrom(k.Start, steps)
}

func (k *Key) ChromaticScale() []*Note {
	// TODO what naturals to use?
	panic("TODO")
}

func (k *Key) MajorScale() []*Note {
	return k.getScale(MajorSteps)
}

func (k *Key) MinorScale() []*Note {
	return k.getScale(MinorSteps)
}

func (k *Key) Chords() []*Chord {
	return []*Chord{
		// TODO figure out how to raise these chords
		ChordMajorTriad,      // 0
		ChordMinorTriad,      // 1
		ChordMinorTriad,      // 2
		ChordMajorTriad,      // 3
		ChordMajorTriad,      // 4
		ChordMinorTriad,      // 5
		ChordDiminishedTriad, // 6
	}
}

func init() {
	for _, k := range KeySignatures {
		StringToKey[k.Start.String()] = k
	}
}
