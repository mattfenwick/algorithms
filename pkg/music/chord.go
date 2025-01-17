package music

type Step struct {
	Base    int
	Natural int
}

func NewStep(b, n int) *Step {
	return &Step{Base: b, Natural: n}
}

type Chord struct {
	Steps []*Step
	Name  string
}

func (c *Chord) Apply(key *Key) []*Note {
	return key.getScale(c.Steps)
}

var (
	StepFirst             = &Step{0, 0}
	StepMinorSecond       = &Step{1, 1}
	StepMajorSecond       = &Step{2, 1}
	StepMinorThird        = &Step{3, 2}
	StepMajorThird        = &Step{4, 2}
	StepFourth            = &Step{5, 3}
	StepDiminishedFifth   = &Step{6, 4}
	StepPerfectFifth      = &Step{7, 4}
	StepAugmentedFifth    = &Step{8, 4}
	StepSixth             = &Step{9, 5}
	StepDiminishedSeventh = &Step{9, 5}
	StepMinorSeventh      = &Step{10, 6}
	StepMajorSeventh      = &Step{11, 6}
)

// see https://en.wikipedia.org/wiki/Chord_(music)#Examples
var (
	ChordMajorTriad = &Chord{
		Steps: []*Step{StepFirst, StepMajorThird, StepPerfectFifth},
		Name:  "Major triad",
	}
	ChordMinorTriad = &Chord{
		Steps: []*Step{StepFirst, StepMinorThird, StepPerfectFifth},
		Name:  "Minor triad",
	}
	ChordAugmentedTriad = &Chord{
		Steps: []*Step{StepFirst, StepMajorThird, StepAugmentedFifth},
		Name:  "Augmented triad",
	}
	ChordDiminishedTriad = &Chord{
		Steps: []*Step{StepFirst, StepMinorThird, StepDiminishedFifth},
		Name:  "Diminished triad",
	}
	ChordSuspendedSecond = &Chord{
		Steps: []*Step{StepFirst, StepMajorSecond, StepPerfectFifth},
		Name:  "Suspended second",
	}
	ChordSuspendedFourth = &Chord{
		Steps: []*Step{StepFirst, StepFourth, StepPerfectFifth},
		Name:  "Suspended fourth",
	}
	ChordFlatFifth = &Chord{
		Steps: []*Step{StepFirst, StepMajorThird, StepDiminishedFifth},
		Name:  "Flat fifth",
	}
	ChordMajorSixth = &Chord{
		Steps: []*Step{StepFirst, StepMajorThird, StepPerfectFifth, StepSixth},
		Name:  "Major sixth",
	}
	ChordMinorSixth = &Chord{
		Steps: []*Step{StepFirst, StepMinorThird, StepPerfectFifth, StepSixth},
		Name:  "Minor sixth",
	}
	ChordMajorSeventh = &Chord{
		Steps: []*Step{StepFirst, StepMajorThird, StepPerfectFifth, StepMajorSeventh},
		Name:  "Major seventh",
	}
	ChordSeventh = &Chord{
		Steps: []*Step{StepFirst, StepMajorThird, StepPerfectFifth, StepMinorSeventh},
		Name:  "Seventh",
	}
	ChordMinorSeventh = &Chord{
		Steps: []*Step{StepFirst, StepMinorThird, StepPerfectFifth, StepMinorSeventh},
		Name:  "Minor seventh",
	}
	ChordAugmentedSeventh = &Chord{
		Steps: []*Step{StepFirst, StepMajorThird, StepAugmentedFifth, StepMinorSeventh},
		Name:  "Augmented seventh",
	}
	ChordMinorMajorSeventh = &Chord{
		Steps: []*Step{StepFirst, StepMinorThird, StepPerfectFifth, StepMajorSeventh},
		Name:  "Minor/major seventh",
	}
	ChordDiminishedSeventh = &Chord{
		Steps: []*Step{StepFirst, StepMinorThird, StepDiminishedFifth, StepDiminishedSeventh},
		Name:  "Diminished seventh",
	}
	ChordHalfDiminishedSeventh = &Chord{
		Steps: []*Step{StepFirst, StepMinorThird, StepDiminishedFifth, StepMinorSeventh},
		Name:  "Half-diminished seventh",
	}
)
