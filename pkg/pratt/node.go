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
			out := []string{"(", v.Open}
			out = append(out, parensHelper(v.Args[0])...)
			return append(out, ")")
		case OpTypeBinary:
			validateArgCount(2, v.Args)
			out := []string{"("}
			out = append(out, parensHelper(v.Args[0])...)
			out = append(out, v.Open)
			out = append(out, parensHelper(v.Args[1])...)
			return append(out, ")")
		case OpTypePostfix:
			validateArgCount(1, v.Args)
			out := []string{"("}
			out = append(out, parensHelper(v.Args[0])...)
			return append(out, v.Open, ")")
		case OpTypeTernary:
			validateArgCount(3, v.Args)
			out := []string{"("}
			out = append(out, parensHelper(v.Args[0])...)
			out = append(out, v.Open)
			out = append(out, parensHelper(v.Args[1])...)
			out = append(out, v.Close)
			out = append(out, parensHelper(v.Args[2])...)
			return append(out, ")")
		case OpTypeGrouping:
			validateArgCount(1, v.Args)
			out := []string{"(" + v.Open}
			out = append(out, parensHelper(v.Args[0])...)
			return append(out, ")"+v.Close)
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
		val = v.Open
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

type OpType string

const (
	OpTypePrefix   OpType = "OpTypePrefix"
	OpTypePostfix  OpType = "OpTypePostfix"
	OpTypeBinary   OpType = "OpTypeBinary"
	OpTypeTernary  OpType = "OpTypeTernary"
	OpTypeGrouping OpType = "OpTypeGrouping"
)

type OpNode struct {
	Open  string
	Close string
	Type  OpType
	Args  []Node
}

func (o *OpNode) NodeString() string {
	return o.Open
}

func Prefix(op string) *OpNode {
	// this is an in-progress node
	return &OpNode{Open: op, Type: OpTypePrefix, Args: []Node{}}
}

func Grouping(op string) *OpNode {
	// this is an in-progress node
	return &OpNode{Open: op, Type: OpTypeGrouping, Args: []Node{}}
}

func Postfix(op string, arg Node) *OpNode {
	// this is a completed node
	return &OpNode{Open: op, Type: OpTypePostfix, Args: []Node{arg}}
}

func Binary(op string, arg Node) *OpNode {
	// this is an in-progress node
	return &OpNode{Open: op, Type: OpTypeBinary, Args: []Node{arg}}
}

func Ternary(open string, arg Node) *OpNode {
	// this is an in-progress node needing *2* more args
	return &OpNode{Open: open, Type: OpTypeTernary, Args: []Node{arg}}
}
