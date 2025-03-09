package pratt

import (
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

func Parse(tokens []*Token) (Node, error) {
	stack := []*OpNode{}
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
				return nil, errors.Errorf("expected prefix op at %d, found %+v", i, op)
			}
			i++
		}

		// 2. find a num
		arg := tokens[i]
		switch arg.Type {
		case TokenTypeNum:
		default:
			return nil, errors.Errorf("expected num at %d, found %+v", i, arg)
		}
		i++

		// 3. process any postfix operators
		// TODO
		var newNode Node = Num(arg.Value)

		// no more tokens => done, so it's time to unwind the stack
		if i >= len(tokens) {
			// get the stack down to one completed expression
			for len(stack) > 0 {
				// pop the stack, and add top of stack as arg to next highest
				// this assumes there's only OpNodes on the stack (other than the initial top)
				top := stack[len(stack)-1]
				top.Args = append(top.Args, newNode)
				newNode = top
				stack = stack[:len(stack)-1]
			}
			return newNode, nil
		}

		// 4. process a binary operator
		op := tokens[i]
		switch op.Type {
		case TokenTypeOp:
			if _, ok := Binary[op.Value]; !ok {
				return nil, errors.Errorf("expected binary operator, found %s at %d", op.Value, i)
			}
			if len(stack) == 0 {
				stack = append(stack, Op(op.Value, 2, newNode))
				break
			}
			for len(stack) > 0 {
				top := stack[len(stack)-1]
				topPrec, isTopRightAssociative := GetPrecedence(top.Op, top.ExpectedArgCount), IsRightAssociative(top.Op, top.ExpectedArgCount)
				if GetPrecedence(op.Value, 2) > topPrec {
					// next operator is higher precedence?  stop poppin'
					break
				}
				if GetPrecedence(op.Value, 2) == topPrec {
					logrus.Warnf("assoc: %t vs %t (%s vs %s)", IsRightAssociative(op.Value, 2), isTopRightAssociative, op.Value, top.Op)
					if IsRightAssociative(op.Value, 2) != isTopRightAssociative {
						return nil, errors.Errorf("unable to handle same precedence but different associativity: %s vs %s", op.Value, top.Op)
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
			return nil, errors.Errorf("expected op at %d, found %+v", i, op)
		}
		i++
	}
}

func ParseString(s string) (Node, error) {
	return Parse(Must(Tokenize(s)))
}
