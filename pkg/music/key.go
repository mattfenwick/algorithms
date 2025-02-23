package music

import (
	"fmt"
	"strings"
)

type Key struct {
	Start  *Note
	Sharps int
	Flats  int
}

func (k *Key) Signature() string {
	return fmt.Sprintf("%s%s", strings.Repeat(SharpString, k.Sharps), strings.Repeat(FlatString, k.Flats))
}

var (
	KeyC      = &Key{NewNote(NaturalC, 0, 0), 0, 0}
	KeyCSharp = &Key{NewNote(NaturalC, 1, 0), 7, 0}
	KeyDFlat  = &Key{NewNote(NaturalD, 0, 1), 0, 5}
	KeyD      = &Key{NewNote(NaturalD, 0, 0), 2, 0}
	KeyEFlat  = &Key{NewNote(NaturalE, 0, 1), 0, 3}
	KeyE      = &Key{NewNote(NaturalE, 0, 0), 4, 0}
	KeyF      = &Key{NewNote(NaturalF, 0, 0), 0, 1}
	KeyFSharp = &Key{NewNote(NaturalF, 1, 0), 6, 0}
	KeyGFlat  = &Key{NewNote(NaturalG, 0, 1), 0, 6}
	KeyG      = &Key{NewNote(NaturalG, 0, 0), 1, 0}
	KeyAFlat  = &Key{NewNote(NaturalA, 0, 1), 0, 4}
	KeyA      = &Key{NewNote(NaturalA, 0, 0), 3, 0}
	KeyBFlat  = &Key{NewNote(NaturalB, 0, 1), 0, 2}
	KeyB      = &Key{NewNote(NaturalB, 0, 0), 5, 0}
)

var KeysByFifths = []*Key{
	KeyC,
	KeyG,
	KeyD,
	KeyA,
	KeyE,
	KeyB,
	KeyGFlat,
	KeyFSharp,
	KeyDFlat,
	KeyCSharp,
	KeyAFlat,
	KeyEFlat,
	KeyBFlat,
	KeyF,
}

var StringToKey = map[string]*Key{}

func init() {
	for _, k := range KeysByFifths {
		StringToKey[k.Start.String()] = k
	}
}

var KeysChromatic = []*Key{
	{NewNote(NaturalC, 0, 0), 0, 0},
	{NewNote(NaturalC, 1, 0), 7, 0},
	{NewNote(NaturalD, 0, 1), 0, 5},
	{NewNote(NaturalD, 0, 0), 2, 0},
	{NewNote(NaturalD, 1, 0), 9, 0},
	{NewNote(NaturalE, 0, 1), 0, 3},
	{NewNote(NaturalE, 0, 0), 4, 0},
	{NewNote(NaturalF, 0, 0), 0, 1},
	{NewNote(NaturalF, 1, 0), 6, 0},
	{NewNote(NaturalG, 0, 1), 0, 6},
	{NewNote(NaturalG, 0, 0), 1, 0},
	{NewNote(NaturalG, 1, 0), 8, 0},
	{NewNote(NaturalA, 0, 1), 0, 4},
	{NewNote(NaturalA, 0, 0), 3, 0},
	{NewNote(NaturalA, 1, 0), 10, 0},
	{NewNote(NaturalB, 0, 1), 0, 2},
	{NewNote(NaturalB, 0, 0), 5, 0},
}
