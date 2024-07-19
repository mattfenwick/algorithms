package taxes

import (
	"math"

	"github.com/mattfenwick/collections/pkg/builtin"
	"github.com/pkg/errors"
)

type RawBracket struct {
	Rate int
	Max  int
}

type Bracket struct {
	RawBracket *RawBracket
	Start      int
	End        *int
}

type BracketTax struct {
	TaxableAmount int
	Tax           int
}

func (b *Bracket) GetTax(income int) *BracketTax {
	if income <= b.Start {
		return &BracketTax{0, 0}
	}
	amount := income - b.Start
	if b.End != nil {
		end := *b.End
		amount = builtin.Min(end-b.Start, income-b.Start)
	}
	// TODO probably doesn't round right
	return &BracketTax{amount, amount * b.RawBracket.Rate / 100}
}

type StatusBrackets struct {
	RawBrackets []*RawBracket
}

func NewStatusBrackets(rawBrackets []*RawBracket) *StatusBrackets {
	if len(rawBrackets) == 0 {
		panic(errors.Errorf("must have 1+ raw brackets"))
	}
	prev := rawBrackets[0]
	for _, curr := range rawBrackets[1:] {
		if curr.Max-prev.Max <= 0 {
			panic(errors.Errorf("invalid raw brackets: maxes must be increasing (%d to %d)", prev.Max, curr.Max))
		}
	}
	return &StatusBrackets{RawBrackets: rawBrackets}
}

func (b *StatusBrackets) GetBrackets() []*Bracket {
	start := 0
	var out []*Bracket
	for _, r := range b.RawBrackets {
		b := &Bracket{
			RawBracket: r,
			Start:      start,
			End:        nil,
		}
		if r.Max < math.MaxInt {
			end := r.Max
			b.End = &end
			start = end
		}
		out = append(out, b)
	}
	return out
}
