package taxes

import (
	"github.com/mattfenwick/collections/pkg/slice"
	"github.com/pkg/errors"
)

func Transpose[A any](as [][]A) [][]A {
	if len(as) == 0 {
		return nil
	}
	firstLength := len(as[0])
	for i, a := range as[1:] {
		if len(a) != firstLength {
			panic(errors.Errorf("unable to transpose slice: not square (%d vs %d at index %d)", firstLength, len(a), i))
		}
	}
	out := [][]A{}
	for i := 0; i < firstLength; i++ {
		out = append(out, slice.Map(func(a []A) A { return a[i] }, as))
	}
	return out
}
