package pratt

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

var RightAssociative = map[string]bool{
	"**": true,
}

var Precedences = map[string]int{
	"+":  50,
	"-":  50,
	"*":  60,
	"/":  60,
	"%":  60,
	"^":  70,
	"**": 70,
}

func GetPrecedence(op string) int {
	val, ok := Precedences[op]
	if !ok {
		panic(errors.Errorf("invalid op: %s", op))
	}
	return val
}

type Node interface {
	NodeString() string
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
	Op   string
	Args []Node
}

func (o *OpNode) NodeString() string {
	return o.Op
}

func Op(op string, args ...Node) *OpNode {
	return &OpNode{Op: op, Args: args}
}

func Parse(tokens []*Token) Node {
	stack := []Node{}
	for i := 0; ; {
		arg := tokens[i]
		switch arg.Type {
		case TokenTypeNum:
		default:
			panic(errors.Errorf("expected num at %d, found %+v", i, arg))
		}
		i++
		if i >= len(tokens) {
			stack = append(stack, Num(arg.Value))
			// unwind the stack
			for len(stack) > 1 {
				// pop the stack, and add top of stack as arg to next highest
				// this assumes there's only OpNodes on the stack (other than the initial top)
				top := stack[len(stack)-1]
				second := stack[len(stack)-2].(*OpNode)
				second.Args = append(second.Args, top)
				stack = stack[:len(stack)-1]
			}
			break
		}
		op := tokens[i]
		switch op.Type {
		case TokenTypeOp:
			if len(stack) == 0 {
				stack = append(stack, Op(op.Value, Num(arg.Value)))
				break
			}
			var newNode Node = Num(arg.Value)
			for len(stack) > 0 {
				top := stack[len(stack)-1].(*OpNode)
				if GetPrecedence(op.Value) > GetPrecedence(top.Op) {
					// next operator is higher precedence?  stop poppin'
					break
				} else if GetPrecedence(op.Value) == GetPrecedence(top.Op) && RightAssociative[top.Op] {
					// same precedence but right-associative?  stop poppin'
					if !RightAssociative[op.Value] == RightAssociative[top.Op] {
						panic(errors.Errorf("unable to handle same precedence but different associativity: %s vs %s", op.Value, top.Op))
					}
					break
				}
				// time to pop: remove the top, and complete it by adding in previous 'newNode' as its next operand.
				// then set 'newNode' to the popped top
				stack = stack[:len(stack)-1]
				top.Args = append(top.Args, newNode)
				newNode = top
			}
			stack = append(stack, Op(op.Value, newNode))
		default:
			panic(errors.Errorf("expected op at %d, found %+v", i, op))
		}
		i++
	}
	if len(stack) != 1 {
		panic(errors.Errorf("expected stack of size 1, found %d (%+v)", len(stack), stack))
	}
	return stack[0]
}

func ParseString(s string) Node {
	return Parse(Must(Tokenize(s)))
}
