package logic

import (
	"github.com/pkg/errors"
)

func Evaluate(node Term, env map[string]bool) (bool, error) {
	switch t := node.(type) {
	case *VarTerm:
		val, ok := env[t.Name]
		if !ok {
			return false, errors.Errorf("undefined variable: %s", t.Name)
		}
		return val, nil
	case *NotTerm:
		val, err := Evaluate(t.Arg, env)
		if err != nil {
			return false, err
		}
		return !val, nil
	case *BinOpTerm:
		switch t.Op {
		default:
			l, err := Evaluate(t.LeftArg, env)
			if err != nil {
				return false, err
			}
			r, err := Evaluate(t.RightArg, env)
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
	default:
		panic(errors.Errorf("invalid node: %+v", t))
	}
}
