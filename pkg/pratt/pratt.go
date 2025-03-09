package pratt

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func Run() {
	tokens := []*Token{
		{Type: TokenTypeNum, Value: "3"},
		{Type: TokenTypeOp, Value: "+"},
		{Type: TokenTypeNum, Value: "4"},
		{Type: TokenTypeOp, Value: "*"},
		{Type: TokenTypeNum, Value: "5"},
	}
	node := Parse(tokens)
	fmt.Printf("%s\n", NodeString(node))
}

// TODO associativity
// TODO precedence
// TODO test cases

// type Symbol struct {
// 	Value string
// }

// type Number struct {
// 	Value int
// }

// type Ternary struct {
// 	Condition Node
// 	True Node
// 	False Node
// }

/* Examples
var     x
num     3
un op   - 3
binop   3 + 4
ternop  a ? b : c
paren   ( x )
assoc   a - b - c    ( - ( - a b ) c )
prec    a + b * c    ( + a ( * b c ) )

un/bin  - x - - y      ( - ( - x ) ( - y ))
*/

type TokenType = string

const (
	TokenTypeOp  = "TokenTypeOp"
	TokenTypeNum = "TokenTypeNum"
	// TokenTypeVar = "TokenTypeVar"
)

type Token struct {
	Type  string
	Value string
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

/*
3 + 4 * 5 ^ 6 * 7    answer: 3 + ((4 * (5 ^ 6)) * 7)
[] [3, +, 4, *, 5, ^, 6, *, 7]
[(+, 3)] [4, *, 5, ^, 6, *, 7]
[(+, 3), (*, 4)] [5, ^, 6, *, 7]
[(+, 3), (*, 4), (^, 5)] [6, *, 7]
[(+, 3), (*, (*, 4, (^, 5, 6)))] [7]
- note: multiple things happened in this step
[(+, 3, (*, (*, 4, (^, 5, 6)), 7))]
- note: multiple things happened in this step
*/

type Node interface { /* TODO */
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

type OpNode struct {
	Op   string
	Args []Node
}

type Stack struct {
	Nodes []Node
}

func NewStack() *Stack {
	return &Stack{Nodes: []Node{}}
}

func (s *Stack) PushExpr(node Node) {
	if len(s.Nodes) == 0 {
		s.Nodes = append(s.Nodes, node)
		return
	}
	top := s.Nodes[len(s.Nodes)-1]
	switch topNode := top.(type) {
	case OpNode:
		if len(topNode.Args) != 1 {
			panic(errors.Errorf("expected top node to have 1 arg already, found %d (%+v)", len(topNode.Args)))
		}

	default:
		panic(errors.Errorf("expected top node to be of type OpNode, found %+v", top))
	}
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
			stack = append(stack, &NumNode{Value: arg.Value})
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
				stack = append(stack, &OpNode{Op: op.Value, Args: []Node{&NumNode{Value: arg.Value}}})
				break
			}
			var newNode Node = &NumNode{Value: arg.Value}
			for len(stack) > 0 {
				top := stack[len(stack)-1].(*OpNode)
				if GetPrecedence(op.Value) > GetPrecedence(top.Op) {
					break
				}
				stack = stack[:len(stack)-1]
				top.Args = append(top.Args, newNode)
				newNode = top
			}
			stack = append(stack, &OpNode{Op: op.Value, Args: []Node{newNode}})
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

func Must0(err error) {
	if err != nil {
		logrus.Fatalf("%s", err.Error())
	}
}

func Must[A any](a A, err error) A {
	if err != nil {
		logrus.Fatalf("%s", err.Error())
	}
	return a
}
