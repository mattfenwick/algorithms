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

var Keys = []*Key{
	{NewNote(NaturalC, 0, 0), 0, 0},
	{NewNote(NaturalG, 0, 0), 1, 0},
	{NewNote(NaturalD, 0, 0), 2, 0},
	{NewNote(NaturalA, 0, 0), 3, 0},
	{NewNote(NaturalE, 0, 0), 4, 0},
	{NewNote(NaturalB, 0, 0), 5, 0},
	{NewNote(NaturalG, 0, 1), 0, 6},
	{NewNote(NaturalF, 1, 0), 6, 0},
	{NewNote(NaturalD, 0, 1), 0, 5},
	{NewNote(NaturalC, 1, 0), 7, 0},
	{NewNote(NaturalA, 0, 1), 0, 4},
	{NewNote(NaturalE, 0, 1), 0, 3},
	{NewNote(NaturalB, 0, 1), 0, 2},
	{NewNote(NaturalF, 0, 0), 0, 1},
}

var StringToKey = map[string]*Key{}

func init() {
	for _, k := range Keys {
		StringToKey[k.Start.String()] = k
	}
}
