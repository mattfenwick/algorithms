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
	MajorSteps = []int{2, 2, 1, 2, 2, 2, 1}
	MinorSteps = []int{2, 1, 2, 2, 1, 2, 2}
)

func (k *Key) getScale(steps []int) []*Note {
	logrus.Debugf("looking at key of %s", k.Start)
	out := []*Note{k.Start}
	curr := k.Start.BaseNote()
	natural := k.Start.Natural
	for _, step := range steps {
		next := curr.Step(step)
		natural = natural.Next()
		out = append(out, next.KeyNote(natural))
		curr = next
	}
	return out
}

func (k *Key) MajorScale() []*Note {
	return k.getScale(MajorSteps)
}

func (k *Key) MinorScale() []*Note {
	return k.getScale(MinorSteps)
}

func init() {
	for _, k := range KeySignatures {
		StringToKey[k.Start.String()] = k
	}
}
