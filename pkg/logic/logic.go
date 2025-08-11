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

type Node interface{}

func PrettyPrint(n Node) string {
	var help func(int, Node) []string
	help = func(depth int, node Node) []string {
		out := []string{}
		switch t := node.(type) {
		case *VarNode:
			out = append(out, t.Name)
		case *OpNode:
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

type VarNode struct {
	Name string
}

type OpNode struct {
	Op   Op
	Args []Node
}

func Var(name string) *VarNode {
	return &VarNode{Name: name}
}

func Not(arg Node) *OpNode {
	return &OpNode{Op: NotOp, Args: []Node{arg}}
}

func And(l, r Node) *OpNode {
	return &OpNode{Op: AndOp, Args: []Node{l, r}}
}

func Or(l, r Node) *OpNode {
	return &OpNode{Op: OrOp, Args: []Node{l, r}}
}

func Implication(l, r Node) *OpNode {
	return &OpNode{Op: ImplicationOp, Args: []Node{l, r}}
}

func Biconditional(l, r Node) *OpNode {
	return &OpNode{Op: BiconditionalOp, Args: []Node{l, r}}
}
