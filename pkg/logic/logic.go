package logic

import (
	"strings"

	"github.com/pkg/errors"
)

type Op string

const (
	NotOp           = "~"
	AndOp           = "^"
	OrOp            = "v"
	ImplicationOp   = "->"
	BiconditionalOp = "<->"
)

type Term interface{}

func PrettyPrint(n Term) string {
	var help func(int, Term) []string
	help = func(depth int, node Term) []string {
		out := []string{}
		switch t := node.(type) {
		case *VarTerm:
			out = append(out, t.Name)
		case *OpTerm:
			switch t.Op {
			case NotOp:
				out = append(out, string(t.Op))
				out = append(out, help(depth+1, t.Args[0])...)
			default:
				if depth != 0 {
					out = append(out, "(")
				}
				out = append(out, help(depth+1, t.Args[0])...)
				out = append(out, string(t.Op))
				out = append(out, help(depth+1, t.Args[1])...)
				if depth != 0 {
					out = append(out, ")")
				}
			}
		default:
			panic(errors.Errorf("invalid type: %+V", t))
		}
		return out
	}
	return strings.Join(help(0, n), " ")
}

type VarTerm struct {
	Name string
}

type OpTerm struct {
	Op   Op
	Args []Term
}

func Var(name string) *VarTerm {
	return &VarTerm{Name: name}
}

func Not(arg Term) *OpTerm {
	return &OpTerm{Op: NotOp, Args: []Term{arg}}
}

func And(l, r Term) *OpTerm {
	return &OpTerm{Op: AndOp, Args: []Term{l, r}}
}

func Or(l, r Term) *OpTerm {
	return &OpTerm{Op: OrOp, Args: []Term{l, r}}
}

func Implication(l, r Term) *OpTerm {
	return &OpTerm{Op: ImplicationOp, Args: []Term{l, r}}
}

func Biconditional(l, r Term) *OpTerm {
	return &OpTerm{Op: BiconditionalOp, Args: []Term{l, r}}
}
