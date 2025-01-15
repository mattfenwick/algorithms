package music

type Natural string

const (
	NaturalC Natural = "C"
	NaturalD Natural = "D"
	NaturalE Natural = "E"
	NaturalF Natural = "F"
	NaturalG Natural = "G"
	NaturalA Natural = "A"
	NaturalB Natural = "B"
)

func (n Natural) BaseNote() BaseNote {
	return naturalToBase[n]
}

func (n Natural) Next(size int) Natural {
	index := naturalToIndex[n] + size
	if index >= len(naturals) {
		index -= len(naturals)
	} else if index < 0 {
		index += len(naturals)
	}
	return naturals[index]
}

var naturalToBase = map[Natural]BaseNote{
	NaturalC: Note0,
	NaturalD: Note2,
	NaturalE: Note4,
	NaturalF: Note5,
	NaturalG: Note7,
	NaturalA: Note9,
	NaturalB: Note11,
}

var naturals = []Natural{
	NaturalC,
	NaturalD,
	NaturalE,
	NaturalF,
	NaturalG,
	NaturalA,
	NaturalB,
}

var naturalToIndex = map[Natural]int{}

func init() {
	for i, n := range naturals {
		naturalToIndex[n] = i
	}
}
