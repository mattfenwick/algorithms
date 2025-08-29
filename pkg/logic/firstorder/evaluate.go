package logic

import (
	"fmt"

	"github.com/mattfenwick/collections/pkg/set"
	"github.com/pkg/errors"
)

type Env struct {
	PropToTruth map[string]bool
	Domain      *set.Set[string]
}

func NewEnv(propToTruth map[string]bool, domain ...string) *Env {
	propToTruth["T"] = true
	propToTruth["F"] = false
	return &Env{
		PropToTruth: propToTruth,
		Domain:      set.FromSlice(domain),
	}
}

func (e *Env) Evaluate(node Term) (bool, error) {
	switch t := node.(type) {
	case *PropTerm:
		propString := t.TermPrint(true)
		val, ok := e.PropToTruth[propString]
		if !ok {
			return false, errors.Errorf("undefined Prop: %s", propString)
		}
		return val, nil
	case *NotTerm:
		val, err := e.Evaluate(t.Term)
		if err != nil {
			return false, err
		}
		return !val, nil
	case *BinOpTerm:
		switch t.Op {
		default:
			l, err := e.Evaluate(t.Left)
			if err != nil {
				return false, err
			}
			r, err := e.Evaluate(t.Right)
			if err != nil {
				return false, err
			}
			switch t.Op {
			case AndOp:
				return l && r, nil
			case OrOp:
				return l || r, nil
			case ImplicationOp:
				return !l || r, nil
			case BiconditionalOp:
				return l == r, nil
			}
			return false, errors.Errorf("invalid bin op: %+v", t)
		}
	case *QuantifiedTerm:
		var trues, falses []string
		for _, obj := range e.Domain.ToSlice() {
			// substitute obj in to body
			term := t.Substitute(obj)
			// evaluate body
			value, err := e.Evaluate(term)
			if err != nil {
				return false, err
			}
			if value {
				trues = append(trues, obj)
			} else {
				falses = append(falses, obj)
			}
		}
		fmt.Printf("quantifier %s results: %+v, %+v\n", t.TermPrint(true), falses, trues)
		switch t.Quantifier {
		case ForallQuantifier:
			return len(falses) == 0, nil
		case ExistentialQuantifier:
			return len(trues) > 0, nil
		default:
			panic(errors.Errorf("invalid quantifer %s", t.Quantifier))
		}

	default:
		panic(errors.Errorf("invalid term type: %+v", t))
	}
}
