package music

import "fmt"

type Step struct {
	Base    int
	Natural int
	Name    string
}

func NewStep(b, n int, name string) *Step {
	return &Step{Base: b, Natural: n, Name: name}
}

var (
	StepFirst             = NewStep(0, 0, "1")
	StepMinorSecond       = NewStep(1, 1, "♭2")
	StepMajorSecond       = NewStep(2, 1, "2")
	StepMinorThird        = NewStep(3, 2, "♭3")
	StepMajorThird        = NewStep(4, 2, "3")
	StepFourth            = NewStep(5, 3, "4")
	StepSharpFourth       = NewStep(6, 3, "♯4")
	StepDiminishedFifth   = NewStep(6, 4, "♭5")
	StepPerfectFifth      = NewStep(7, 4, "5")
	StepAugmentedFifth    = NewStep(8, 4, "♯5")
	StepFlatSixth         = NewStep(8, 5, "♭6")
	StepSixth             = NewStep(9, 5, "6")
	StepDiminishedSeventh = NewStep(9, 6, "♭♭7")
	StepMinorSeventh      = NewStep(10, 6, "♭7")
	StepMajorSeventh      = NewStep(11, 6, "7")
	StepEighth            = NewStep(12, 7, "8")
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
		BaseSymbol: "<sup>sus2</sup>",
	}
	ChordSuspendedFourth = &Chord{
		Steps:      []*Step{StepFirst, StepFourth, StepPerfectFifth},
		Name:       "Suspended fourth",
		BaseSymbol: "<sup>sus4</sup>",
	}
	ChordFlatFifth = &Chord{
		Steps:      []*Step{StepFirst, StepMajorThird, StepDiminishedFifth},
		Name:       "Flat fifth",
		BaseSymbol: "<sup>♭5</sup>",
	}
	ChordDominantSeventhFlatFifth = &Chord{
		Steps:      []*Step{StepFirst, StepMajorThird, StepDiminishedFifth, StepMinorSeventh},
		Name:       "Dominant seventh flat fifth",
		BaseSymbol: "<sup>7♭5</sup>",
	}
	ChordMajorSixth = &Chord{
		Steps:      []*Step{StepFirst, StepMajorThird, StepPerfectFifth, StepSixth},
		Name:       "Major sixth",
		BaseSymbol: "<sup>6</sup>",
	}
	ChordMinorSixth = &Chord{
		Steps:      []*Step{StepFirst, StepMinorThird, StepPerfectFifth, StepSixth},
		Name:       "Minor sixth",
		BaseSymbol: "m<sup>6</sup>",
	}
	ChordMajorSeventh = &Chord{
		Steps:      []*Step{StepFirst, StepMajorThird, StepPerfectFifth, StepMajorSeventh},
		Name:       "Major seventh",
		BaseSymbol: "M<sup>7</sup>",
	}
	ChordSeventh = &Chord{
		Steps:      []*Step{StepFirst, StepMajorThird, StepPerfectFifth, StepMinorSeventh},
		Name:       "Seventh",
		BaseSymbol: "<sup>7</sup>",
	}
	ChordMinorSeventh = &Chord{
		Steps:      []*Step{StepFirst, StepMinorThird, StepPerfectFifth, StepMinorSeventh},
		Name:       "Minor seventh",
		BaseSymbol: "m<sup>7</sup>",
	}
	ChordAugmentedSeventh = &Chord{
		Steps:      []*Step{StepFirst, StepMajorThird, StepAugmentedFifth, StepMinorSeventh},
		Name:       "Augmented seventh",
		BaseSymbol: "+<sup>7</sup>",
	}
	ChordMinorMajorSeventh = &Chord{
		Steps:      []*Step{StepFirst, StepMinorThird, StepPerfectFifth, StepMajorSeventh},
		Name:       "Minor/major seventh",
		BaseSymbol: "m/M<sup>7</sup>",
	}
	ChordDiminishedSeventh = &Chord{
		Steps:      []*Step{StepFirst, StepMinorThird, StepDiminishedFifth, StepDiminishedSeventh},
		Name:       "Diminished seventh",
		BaseSymbol: "<sup>o7</sup>",
	}
	ChordHalfDiminishedSeventh = &Chord{
		Steps:      []*Step{StepFirst, StepMinorThird, StepDiminishedFifth, StepMinorSeventh},
		Name:       "Half-diminished seventh",
		BaseSymbol: "<sup>ø7</sup>",
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
