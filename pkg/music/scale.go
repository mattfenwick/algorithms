package music

type Scale struct {
	Name  string
	Steps []*Step
}

func (s *Scale) Apply(key *Key) []*Note {
	return getNotesFrom(key.Start, s.Steps)
}

func NewScale(name string, steps ...*Step) *Scale {
	return &Scale{Steps: steps, Name: name}
}

var (
	ScaleMajor = NewScale("Major",
		StepFirst,
		StepMajorSecond,
		StepMajorThird,
		StepFourth,
		StepPerfectFifth,
		StepSixth,
		StepMajorSeventh,
		StepEighth,
	)
	ScaleMinor = NewScale("Minor",
		StepFirst,
		StepMajorSecond,
		StepMinorThird,
		StepFourth,
		StepPerfectFifth,
		StepFlatSixth,
		StepMinorSeventh,
		StepEighth,
	)
	ScaleBluesMinor = NewScale("Blues minor",
		StepFirst,
		StepMinorThird,
		StepFourth,
		StepSharpFourth,
		StepPerfectFifth,
		StepMinorSeventh,
		StepEighth,
	)
	ScaleBluesMajor = NewScale("Blues major",
		StepFirst,
		StepMajorSecond,
		StepMinorThird,
		StepMajorThird,
		StepPerfectFifth,
		StepSixth,
		StepEighth,
	)
	ScalePentatonicMajor = NewScale("Pentatonic major",
		StepFirst,
		StepMajorSecond,
		StepMajorThird,
		StepPerfectFifth,
		StepSixth,
		StepEighth,
	)
	ScalePentatonicMinor = NewScale("Pentatonic minor",
		StepFirst,
		StepMinorThird,
		StepFourth,
		StepPerfectFifth,
		StepMinorSeventh,
		StepEighth,
	)

// ChromaticSteps = []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
)

var (
	Scales = []*Scale{
		ScaleMajor,
		ScaleMinor,
		ScaleBluesMajor,
		ScalePentatonicMajor,
		ScaleBluesMinor,
		ScalePentatonicMinor,
	}
)
