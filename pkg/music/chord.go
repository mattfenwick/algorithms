package music

// major
// minor
// augmented
// diminished
// suspended

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
	StepFirst           = &Step{0, 0}
	StepMajorThird      = &Step{4, 2}
	StepMinorThird      = &Step{3, 2}
	StepPerfectFifth    = &Step{7, 4}
	StepAugmentedFifth  = &Step{8, 4}
	StepDiminishedFifth = &Step{6, 4}
)

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
)
