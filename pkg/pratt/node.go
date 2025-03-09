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

func parensHelper(node Node) []string {
	switch v := node.(type) {
	case *OpNode:
		// TODO uncomment this line to get the number of expected args, can be handy sanity check
//		out := []string{"(", fmt.Sprintf("%s[%d]", v.Op, v.ExpectedArgCount)}
		out := []string{"(", v.Op}
		for _, a := range v.Args {
			out = append(out, parensHelper(a)...)
		}
		return append(out,")")
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

type OpNode struct {
	Op               string
	ExpectedArgCount int
	Args             []Node
}

func (o *OpNode) NodeString() string {
	return o.Op
}

func Op(op string, expectedArgCount int, args ...Node) *OpNode {
	return &OpNode{Op: op, ExpectedArgCount: expectedArgCount, Args: args}
}
