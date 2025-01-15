package music

type BaseNote int

const (
	Note0  BaseNote = 0
	Note1  BaseNote = 1
	Note2  BaseNote = 2
	Note3  BaseNote = 3
	Note4  BaseNote = 4
	Note5  BaseNote = 5
	Note6  BaseNote = 6
	Note7  BaseNote = 7
	Note8  BaseNote = 8
	Note9  BaseNote = 9
	Note10 BaseNote = 10
	Note11 BaseNote = 11
)

// KeyNote is how to get from natural, to base note
// example: D -> C , result D♭♭
// example: D -> E , result D♯♯
// example: D -> G , result D♯♯♯♯♯
func (n BaseNote) KeyNote(natural Natural) *Note {
	diff := int(n) - int(natural.BaseNote())
	// which way to go? the shorter way
	if diff > 6 {
		diff -= 12
	} else if diff < -6 {
		diff += 12
	}
	return &Note{Natural: natural, SharpsFlats: diff}
}

func (n BaseNote) Next(size int) BaseNote {
	end := int(n) + size
	if end < 0 {
		end = end + 12
	} else if end >= 12 {
		end = end - 12
	}
	return BaseNote(end)
}
