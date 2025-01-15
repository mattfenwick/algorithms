package music

import (
	"strings"

	"github.com/mattfenwick/collections/pkg/slice"
	"github.com/pkg/errors"
)

// Note starts with a modifier and adds sharps or flats on top.
type Note struct {
	Natural     Natural
	SharpsFlats int
}

func (k *Note) BaseNote() BaseNote {
	return k.Natural.BaseNote().Next(k.SharpsFlats)
}

func (k *Note) Sharps() int {
	if k.SharpsFlats >= 0 {
		return k.SharpsFlats
	}
	return 0
}

func (k *Note) Flats() int {
	if k.SharpsFlats < 0 {
		return -1 * k.SharpsFlats
	}
	return 0
}

func (k *Note) String() string {
	if k.SharpsFlats == 0 {
		return string(k.Natural)
	} else if k.SharpsFlats > 0 {
		return string(k.Natural) + strings.Join(slice.Replicate(k.SharpsFlats, SharpString), "")
	} else {
		return string(k.Natural) + strings.Join(slice.Replicate(-1*k.SharpsFlats, FlatString), "")
	}
}

func NewNote(natural Natural, sharps int, flats int) *Note {
	if sharps != 0 && flats != 0 {
		panic(errors.Errorf("invalid key note: %s, %d, %d", natural, sharps, flats))
	}
	return &Note{
		Natural:     natural,
		SharpsFlats: sharps - flats,
	}
}

func Parse(n string) (*Note, error) {
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
	return NewNote(natural, sharps, flats), nil
}
