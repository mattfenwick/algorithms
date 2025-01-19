package music

import (
	"github.com/mattfenwick/collections/pkg/slice"
)

type RootedChord struct {
	Start int // 1 - 7
	Chord *Chord
}

func (r *RootedChord) Name(key *Key) string {
	return r.Chord.Symbol(key.MajorScale()[r.Start-1])
}

func (r *RootedChord) Apply(key *Key) []*Note {
	return r.Chord.ApplyOn(key.MajorScale()[r.Start-1])
}

// e.g. 1-6-4-5
type Progression struct {
	Chords []*RootedChord
	Name   string
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

	Progression12BarBlues = &Progression{
		Chords: []*RootedChord{
			{1, ChordSeventh},
			{1, ChordSeventh},
			{1, ChordSeventh},
			{1, ChordSeventh},
			{4, ChordSeventh},
			{4, ChordSeventh},
			{1, ChordSeventh},
			{1, ChordSeventh},
			{5, ChordSeventh},
			{4, ChordSeventh},
			{1, ChordSeventh},
			{5, ChordSeventh},
		},
		Name: "12 bar blues",
	}
)

var (
	Progressions = []*Progression{
		Progression1645,
		ProgressionMajorChords,
		Progression12BarBlues,
	}
)
