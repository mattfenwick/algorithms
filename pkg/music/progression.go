package music

import (
	"fmt"

	"github.com/mattfenwick/collections/pkg/slice"
)

type RootedChord struct {
	Start int // 1 - 7
	Chord *Chord
}

func (r *RootedChord) Name() string {
	return fmt.Sprintf("%d: %s", r.Start, r.Chord.Name)
}

func (r *RootedChord) Apply(key *Key) []*Note {
	return r.Chord.ApplyOn(key.MajorScale()[r.Start-1])
}

// e.g. 1-6-4-5
type Progression struct {
	Chords []*RootedChord
	Name string
}

func (p *Progression) Apply(key *Key) [][]*Note {
	return slice.Map(func(c *RootedChord) []*Note { return c.Apply(key) }, p.Chords)
}

var (
	Progression1645 = &Progression{
		Chords: []*RootedChord{
			{1, ChordMajorTriad},
			{6, ChordMinorTriad},
			{4, ChordMajorTriad},
			{5, ChordMajorTriad},
		},
		Name: "I-vi-IV-V",
	}

	ProgressionMajorChords = &Progression{
		Chords: []*RootedChord{
			{1, ChordMajorTriad},
			{2, ChordMinorTriad},
			{3, ChordMinorTriad},
			{4, ChordMajorTriad},
			{5, ChordMajorTriad},
			{6, ChordMinorTriad},
			{7, ChordDiminishedTriad},
		},
		Name: "Major chords",
	}
)

var (
	Progressions = []*Progression{
		Progression1645,
		ProgressionMajorChords,
	}
)
