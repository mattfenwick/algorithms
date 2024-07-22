package taxes

import (
	"fmt"

	"github.com/mattfenwick/collections/pkg/slice"
	"github.com/pkg/errors"
	"golang.org/x/exp/constraints"
)

func Transpose[A any](as [][]A) [][]A {
	if len(as) == 0 {
		return nil
	}
	firstLength := len(as[0])
	for i, a := range as[1:] {
		if len(a) != firstLength {
			panic(errors.Errorf("unable to transpose slice: not rectangular (%d vs %d at index %d)", firstLength, len(a), i))
		}
	}
	out := [][]A{}
	for i := 0; i < firstLength; i++ {
		out = append(out, slice.Map(func(a []A) A { return a[i] }, as))
	}
	return out
}

func intToString[A constraints.Integer](a A) string {
	return fmt.Sprintf("%d", a)
}
