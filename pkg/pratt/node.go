package pratt

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

type Node interface {
	NodeString() string
}

func Parens(node Node) string {
	return strings.Join(parensHelper(node), " ")
}

func validateArgCount(expected int, args []Node) {
	if len(args) != expected {
		panic(errors.Errorf("expected %d args, got %d", expected, len(args)))
	}
}

func parensHelper(node Node) []string {
	switch v := node.(type) {
	case *OpNode:
		switch v.Type {
		case OpTypePrefix:
			validateArgCount(1, v.Args)
			out := []string{"(", v.Op}
			out = append(out, parensHelper(v.Args[0])...)
			return append(out, ")")
		case OpTypeBinary:
			validateArgCount(2, v.Args)
			out := []string{"("}
			out = append(out, parensHelper(v.Args[0])...)
			out = append(out, v.Op)
			out = append(out, parensHelper(v.Args[1])...)
			return append(out, ")")
		case OpTypePostfix:
			validateArgCount(1, v.Args)
			out := []string{"("}
			out = append(out, parensHelper(v.Args[0])...)
			return append(out, v.Op, ")")
		default:
			panic(errors.Errorf("invalid op type %s", v.Type))
		}
	case *NumNode:
		return []string{v.Value}
	default:
		panic(errors.Errorf("invalid node type: %+v", node))
	}
}

func NodeString(node Node) string {
	return strings.Join(nodeStringHelper(node, 0), "\n")
}

func nodeStringHelper(node Node, indent int) []string {
	args := []string{}
	val := ""
	switch v := node.(type) {
	case *OpNode:
		val = v.Op
		for _, arg := range v.Args {
			args = append(args, nodeStringHelper(arg, indent+1)...)
		}
	case *NumNode:
		val = v.Value
	default:
		panic(errors.Errorf("invalid node type: %+v", node))
	}
	return append([]string{fmt.Sprintf("%s%s", strings.Repeat("  ", indent), val)}, args...)
}

type NumNode struct {
	Value string
}

func (n *NumNode) NodeString() string {
	return n.Value
}

func Num(val string) *NumNode {
	return &NumNode{Value: val}
}

const (
	OpTypePrefix  = "OpTypePrefix"
	OpTypePostfix = "OpTypePostfix"
	OpTypeBinary  = "OpTypeBinary"
)

type OpNode struct {
	Op   string
	Type string
	Args []Node
}

func (o *OpNode) NodeString() string {
	return o.Op
}

func Prefix(op string) *OpNode {
	// this is an in-progress node
	return &OpNode{Op: op, Type: OpTypePrefix, Args: []Node{}}
}

func Postfix(op string, arg Node) *OpNode {
	// this is a completed node
	return &OpNode{Op: op, Type: OpTypePostfix, Args: []Node{arg}}
}

func Binary(op string, arg Node) *OpNode {
	// this is an in-progress node
	return &OpNode{Op: op, Type: OpTypeBinary, Args: []Node{arg}}
}
