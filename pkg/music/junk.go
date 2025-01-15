package music

import "github.com/mattfenwick/collections/pkg/slice"

var (
	Sharps = []string{"F", "C", "G", "D", "A", "E", "B"}
	Flats  = slice.Reverse(Sharps)
)
