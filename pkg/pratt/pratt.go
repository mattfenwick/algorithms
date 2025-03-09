package pratt

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

var Prefix = map[string]int{
	"!":       100,
	"+":       100,
	"-":       100,
	"^":       100,
	"*":       100,
	"&":       100,
	"<-":      100,
	"~":       100,
	"pre-low": 0,
	"pre-mid": 50,
}

type BinOp struct {
	Symbol             string
	Precedence         int
	IsRightAssociative bool
}

var BinarySlice = []*BinOp{
	{"+", 40, false},
	{"-", 40, false},
	{"bin-mid-left", 50, false},
	{"bin-mid-right", 50, true},
	{"*", 60, false},
	{"/", 60, false},
	{"%", 60, false},
	{"^", 70, false},
	{"**", 70, true},
}

var Binary = map[string]*BinOp{}

func init() {
	for _, b := range BinarySlice {
		if _, ok := Binary[b.Symbol]; ok {
			panic(errors.Errorf("dupe binary symbol %s", b.Symbol))
		}
		Binary[b.Symbol] = b
	}
}

func GetPrecedence(op string, argCount int) int {
	switch argCount {
	case 1:
		if _, ok := Prefix[op]; !ok {
			panic(errors.Errorf("invalid prefix op %s", op))
		}
		return Prefix[op]
	case 2:
		if _, ok := Binary[op]; !ok {
			panic(errors.Errorf("invalid binary op %s", op))
		}
		return Binary[op].Precedence
	default:
		panic(errors.Errorf("invalid argCount for %s: %d", op, argCount))
	}
}

func IsRightAssociative(op string, argCount int) bool {
	switch argCount {
	case 1:
		if _, ok := Prefix[op]; !ok {
			panic(errors.Errorf("invalid prefix op %s", op))
		}
		// TODO should we allow specifying associativity of prefix ops?
		// for example, does 'pre 3 bin 4' parse as '(pre 3) bin 4' or 'pre (3 bin 4)',
		// when 'pre' and 'bin' have the same precedence?
		// for now, punt on this and decree that all prefix ops have left associativity,
		// thereby parsing above example as '(pre 3) bin 4'.
		// This also means that a prefix operator can't be used with a right-associative
		// binary operator of the same precedence.
		return false
	case 2:
		if _, ok := Binary[op]; !ok {
			panic(errors.Errorf("invalid binary op %s", op))
		}
		return Binary[op].IsRightAssociative
	default:
		panic(errors.Errorf("invalid argCount for %s: %d", op, argCount))
	}
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

func Parse(tokens []*Token) Node {
	stack := []Node{}
	for i := 0; ; {
		// 1. throw any prefix unary operators on to the stack
		for i < len(tokens) {
			op := tokens[i]
			if op.Type != TokenTypeOp {
				break
			}
			if _, ok := Prefix[op.Value]; ok {
				stack = append(stack, Op(op.Value, 1))
			} else {
				panic(errors.Errorf("expected prefix op at %d, found %+v", i, op))
			}
			i++
		}
		// 2. find a num
		arg := tokens[i]
		switch arg.Type {
		case TokenTypeNum:
		default:
			panic(errors.Errorf("expected num at %d, found %+v", i, arg))
		}
		i++
		// no more tokens => done, so it's time to unwind the stack
		if i >= len(tokens) {
			// throw the last number on top of the stack
			stack = append(stack, Num(arg.Value))
			// get the stack down to one completed expression
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
		// 3. process a binary operator
		op := tokens[i]
		switch op.Type {
		case TokenTypeOp:
			if _, ok := Binary[op.Value]; !ok {
				panic(errors.Errorf("expected binary operator, found %s at %d", op.Value, i))
			}
			if len(stack) == 0 {
				stack = append(stack, Op(op.Value, 2, Num(arg.Value)))
				break
			}
			var newNode Node = Num(arg.Value)
			for len(stack) > 0 {
				top := stack[len(stack)-1].(*OpNode)
				topPrec, isTopRightAssociative := GetPrecedence(top.Op, top.ExpectedArgCount), IsRightAssociative(top.Op, top.ExpectedArgCount)
				if GetPrecedence(op.Value, 2) > topPrec {
					// next operator is higher precedence?  stop poppin'
					break
				}
				if GetPrecedence(op.Value, 2) == topPrec {
					logrus.Warnf("assoc: %t vs %t (%s vs %s)", IsRightAssociative(op.Value, 2), isTopRightAssociative, op.Value, top.Op)
					if IsRightAssociative(op.Value, 2) != isTopRightAssociative {
						panic(errors.Errorf("unable to handle same precedence but different associativity: %s vs %s", op.Value, top.Op))
					}
					// same precedence but right-associative?  stop poppin'
					if isTopRightAssociative {
						break
					}
				}
				// time to pop: remove the top, and complete it by adding in previous 'newNode' as its next operand.
				// then set 'newNode' to the popped top
				stack = stack[:len(stack)-1]
				top.Args = append(top.Args, newNode)
				newNode = top
			}
			stack = append(stack, Op(op.Value, 2, newNode))
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
