package music

import (
	"strings"

	"github.com/mattfenwick/collections/pkg/slice"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

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
func (n BaseNote) KeyNote(natural Natural) *KeyNote {
	diff := int(n) - int(natural.BaseNote())
	// which way to go? the shorter way
	if diff > 6 {
		diff -= 12
	} else if diff < -6 {
		diff += 12
	}
	return &KeyNote{Natural: natural, SharpsFlats: diff}
}

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

func (n Natural) Next() Natural {
	index := naturalToIndex[n] + 1
	if index >= len(naturals) {
		index -= len(naturals)
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

// KeyNote TODO is not a good name.  come up with a better one
type KeyNote struct {
	Natural     Natural
	SharpsFlats int
}

func (k *KeyNote) BaseNote() BaseNote {
	return k.Natural.BaseNote().Step(k.SharpsFlats)
}

func (k *KeyNote) Sharps() int {
	if k.SharpsFlats >= 0 {
		return k.SharpsFlats
	}
	return 0
}

func (k *KeyNote) Flats() int {
	if k.SharpsFlats < 0 {
		return -1 * k.SharpsFlats
	}
	return 0
}

func (k *KeyNote) String() string {
	if k.SharpsFlats == 0 {
		return string(k.Natural)
	} else if k.SharpsFlats > 0 {
		return string(k.Natural) + strings.Join(slice.Replicate(k.SharpsFlats, SharpString), "")
	} else {
		return string(k.Natural) + strings.Join(slice.Replicate(-1*k.SharpsFlats, FlatString), "")
	}
}

func NewKeyNote(natural Natural, sharps int, flats int) *KeyNote {
	if sharps != 0 && flats != 0 {
		Die(errors.Errorf("invalid key note: %s, %d, %d", natural, sharps, flats))
	}
	return &KeyNote{
		Natural:     natural,
		SharpsFlats: sharps - flats,
	}
}

func Parse(n string) (*KeyNote, error) {
	var runes []rune
	for _, r := range n {
		runes = append(runes, r)
	}
	natural := Natural(runes[0])
	flats, sharps := 0, 0
	for _, m := range runes[1:] {
		switch m {
		case FlatRune:
			flats++
		case SharpRune:
			sharps++
		case NaturalRune:
		default:
			return nil, errors.Errorf("invalid note modifier: %d (%s)", m, string(m))
		}
	}
	if flats > 0 && sharps > 0 {
		return nil, errors.Errorf("invalid note: %s", n)
	}
	return NewKeyNote(natural, sharps, flats), nil
}

func (n BaseNote) Step(size int) BaseNote {
	end := int(n) + size
	if end < 0 {
		end = end + 12
	} else if end >= 12 {
		end = end - 12
	}
	return BaseNote(end)
}

var (
	Sharps = []string{"F", "C", "G", "D", "A", "E", "B"}
	Flats  = slice.Reverse(Sharps)
)

type Key struct {
	Start  *KeyNote
	Sharps int
	Flats  int
}

var KeySignatures = []*Key{
	{NewKeyNote(NaturalG, 0, 1), 0, 6},
	{NewKeyNote(NaturalD, 0, 1), 0, 5},
	{NewKeyNote(NaturalA, 0, 1), 0, 4},
	{NewKeyNote(NaturalE, 0, 1), 0, 3},
	{NewKeyNote(NaturalB, 0, 1), 0, 2},
	{NewKeyNote(NaturalF, 0, 0), 0, 1},
	{NewKeyNote(NaturalC, 0, 0), 0, 0},
	{NewKeyNote(NaturalG, 0, 0), 1, 0},
	{NewKeyNote(NaturalD, 0, 0), 2, 0},
	{NewKeyNote(NaturalA, 0, 0), 3, 0},
	{NewKeyNote(NaturalE, 0, 0), 4, 0},
	{NewKeyNote(NaturalB, 0, 0), 5, 0},
	{NewKeyNote(NaturalF, 1, 0), 6, 0},
	{NewKeyNote(NaturalC, 2, 0), 7, 0},
}

var StringToKey = map[string]*Key{}

var (
	MajorSteps = []int{2, 2, 1, 2, 2, 2, 1}
	MinorSteps = []int{2, 1, 2, 2, 1, 2, 2}
)

func (k *Key) getScale(steps []int) []*KeyNote {
	logrus.Infof("looking at key of %s", k.Start)
	out := []*KeyNote{k.Start}
	curr := k.Start.BaseNote()
	natural := k.Start.Natural
	for _, step := range steps {
		next := curr.Step(step)
		natural = natural.Next()
		out = append(out, next.KeyNote(natural))
		curr = next
	}
	return out
}

func (k *Key) MajorScale() []*KeyNote {
	return k.getScale(MajorSteps)
}

func (k *Key) MinorScale() []*KeyNote {
	return k.getScale(MinorSteps)
}

func init() {
	for _, k := range KeySignatures {
		StringToKey[k.Start.String()] = k
	}
	for i, n := range naturals {
		naturalToIndex[n] = i
	}
}
