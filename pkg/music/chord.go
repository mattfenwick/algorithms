package music

import "fmt"

type Step struct {
	Base    int
	Natural int
}

func NewStep(b, n int) *Step {
	return &Step{Base: b, Natural: n}
}

var (
	StepFirst             = NewStep(0, 0)
	StepMinorSecond       = NewStep(1, 1)
	StepMajorSecond       = NewStep(2, 1)
	StepMinorThird        = NewStep(3, 2)
	StepMajorThird        = NewStep(4, 2)
	StepFourth            = NewStep(5, 3)
	StepSharpFourth       = NewStep(6, 3)
	StepDiminishedFifth   = NewStep(6, 4)
	StepPerfectFifth      = NewStep(7, 4)
	StepAugmentedFifth    = NewStep(8, 4)
	StepFlatSixth         = NewStep(8, 5)
	StepSixth             = NewStep(9, 5)
	StepDiminishedSeventh = NewStep(9, 6)
	StepMinorSeventh      = NewStep(10, 6)
	StepMajorSeventh      = NewStep(11, 6)
	StepEighth            = NewStep(12, 7)
)

type Chord struct {
	Steps      []*Step
	Name       string
	BaseSymbol string
}

func (c *Chord) Apply(key *Key) []*Note {
	return c.ApplyOn(key.Start)
}

func (c *Chord) ApplyOn(note *Note) []*Note {
	return getNotesFrom(note, c.Steps)
}

func (c *Chord) Symbol(note *Note) string {
	return fmt.Sprintf("%s%s", note.String(), c.BaseSymbol)
}

// see https://en.wikipedia.org/wiki/Chord_(music)#Examples
var (
	ChordMajorTriad = &Chord{
		Steps:      []*Step{StepFirst, StepMajorThird, StepPerfectFifth},
		Name:       "Major triad",
		BaseSymbol: "",
	}
	ChordMinorTriad = &Chord{
		Steps:      []*Step{StepFirst, StepMinorThird, StepPerfectFifth},
		Name:       "Minor triad",
		BaseSymbol: "m",
	}
	ChordAugmentedTriad = &Chord{
		Steps:      []*Step{StepFirst, StepMajorThird, StepAugmentedFifth},
		Name:       "Augmented triad",
		BaseSymbol: "+",
	}
	ChordDiminishedTriad = &Chord{
		Steps:      []*Step{StepFirst, StepMinorThird, StepDiminishedFifth},
		Name:       "Diminished triad",
		BaseSymbol: "ᵒ",
	}
	ChordSuspendedSecond = &Chord{
		Steps:      []*Step{StepFirst, StepMajorSecond, StepPerfectFifth},
		Name:       "Suspended second",
		BaseSymbol: "sus2",
	}
	ChordSuspendedFourth = &Chord{
		Steps:      []*Step{StepFirst, StepFourth, StepPerfectFifth},
		Name:       "Suspended fourth",
		BaseSymbol: "sus4",
	}
	ChordFlatFifth = &Chord{
		Steps:      []*Step{StepFirst, StepMajorThird, StepDiminishedFifth},
		Name:       "Flat fifth",
		BaseSymbol: "♭5",
	}
	ChordDominantSeventhFlatFifth = &Chord{
		Steps:      []*Step{StepFirst, StepMajorThird, StepDiminishedFifth, StepMinorSeventh},
		Name:       "Dominant seventh flat fifth",
		BaseSymbol: "7♭5",
	}
	ChordMajorSixth = &Chord{
		Steps:      []*Step{StepFirst, StepMajorThird, StepPerfectFifth, StepSixth},
		Name:       "Major sixth",
		BaseSymbol: "6",
	}
	ChordMinorSixth = &Chord{
		Steps:      []*Step{StepFirst, StepMinorThird, StepPerfectFifth, StepSixth},
		Name:       "Minor sixth",
		BaseSymbol: "m6",
	}
	ChordMajorSeventh = &Chord{
		Steps:      []*Step{StepFirst, StepMajorThird, StepPerfectFifth, StepMajorSeventh},
		Name:       "Major seventh",
		BaseSymbol: "M7",
	}
	ChordSeventh = &Chord{
		Steps:      []*Step{StepFirst, StepMajorThird, StepPerfectFifth, StepMinorSeventh},
		Name:       "Seventh",
		BaseSymbol: "7",
	}
	ChordMinorSeventh = &Chord{
		Steps:      []*Step{StepFirst, StepMinorThird, StepPerfectFifth, StepMinorSeventh},
		Name:       "Minor seventh",
		BaseSymbol: "m7",
	}
	ChordAugmentedSeventh = &Chord{
		Steps:      []*Step{StepFirst, StepMajorThird, StepAugmentedFifth, StepMinorSeventh},
		Name:       "Augmented seventh",
		BaseSymbol: "+7",
	}
	ChordMinorMajorSeventh = &Chord{
		Steps:      []*Step{StepFirst, StepMinorThird, StepPerfectFifth, StepMajorSeventh},
		Name:       "Minor/major seventh",
		BaseSymbol: "m/M7",
	}
	ChordDiminishedSeventh = &Chord{
		Steps:      []*Step{StepFirst, StepMinorThird, StepDiminishedFifth, StepDiminishedSeventh},
		Name:       "Diminished seventh",
		BaseSymbol: "ᵒ7",
	}
	ChordHalfDiminishedSeventh = &Chord{
		Steps:      []*Step{StepFirst, StepMinorThird, StepDiminishedFifth, StepMinorSeventh},
		Name:       "Half-diminished seventh",
		BaseSymbol: "ø7",
	}
)

var (
	Chords = []*Chord{
		ChordMajorTriad,
		ChordMinorTriad,
		ChordAugmentedTriad,
		ChordDiminishedTriad,
		ChordSuspendedSecond,
		ChordSuspendedFourth,
		ChordFlatFifth,
		ChordMajorSixth,
		ChordMinorSixth,
		ChordMajorSeventh,
		ChordSeventh,
		ChordMinorSeventh,
		ChordDominantSeventhFlatFifth,
		ChordAugmentedSeventh,
		ChordMinorMajorSeventh,
		ChordDiminishedSeventh,
		ChordHalfDiminishedSeventh,
	}
)
