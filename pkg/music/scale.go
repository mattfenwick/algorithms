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
		NewStep(0, 0),
		NewStep(2, 1),
		NewStep(4, 2),
		NewStep(5, 3),
		NewStep(7, 4),
		NewStep(9, 5),
		NewStep(11, 6),
		NewStep(12, 7),
	)
	ScaleMinor = NewScale("Minor",
		NewStep(0, 0),
		NewStep(2, 1),
		NewStep(3, 2),
		NewStep(5, 3),
		NewStep(7, 4),
		NewStep(8, 5),
		NewStep(10, 6),
		NewStep(12, 7),
	)
	ScaleBlues = NewScale("Blues",
		NewStep(0, 0),
		NewStep(3, 2),
		NewStep(5, 3),
		NewStep(6, 3),
		NewStep(7, 4),
		NewStep(10, 6),
		NewStep(12, 7),
	)
	ScalePentatonic = NewScale("Pentatonic",
		NewStep(0, 0),
		NewStep(2, 1),
		NewStep(4, 2),
		NewStep(7, 4),
		NewStep(9, 5),
		NewStep(12, 7),
	)
	ScalePentatonicMinor = NewScale("Pentatonic minor",
		NewStep(0, 0),
		NewStep(3, 2),
		NewStep(5, 3),
		NewStep(7, 4),
		NewStep(10, 6),
		NewStep(12, 7),
	)

// ChromaticSteps = []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
)

var (
	Scales = []*Scale{
		ScaleMajor,
		ScaleMinor,
		ScaleBlues,
		ScalePentatonic,
		ScalePentatonicMinor,
	}
)
