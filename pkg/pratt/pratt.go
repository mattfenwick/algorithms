package pratt

import (
	"github.com/pkg/errors"
)

var PrefixOps = map[string]int{
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

var PostfixOps = map[string]int{
	";":        100,
	"post-low": 0,
	"post-mid": 50,
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

var BinaryOps = map[string]*BinOp{}

func init() {
	for _, b := range BinarySlice {
		if _, ok := BinaryOps[b.Symbol]; ok {
			panic(errors.Errorf("dupe binary symbol %s", b.Symbol))
		}
		BinaryOps[b.Symbol] = b
	}
}

func GetPrecedence(op string, opType string) int {
	switch opType {
	case OpTypePrefix:
		if _, ok := PrefixOps[op]; !ok {
			panic(errors.Errorf("invalid prefix op %s", op))
		}
		return PrefixOps[op]
	case OpTypePostfix:
		if _, ok := PostfixOps[op]; !ok {
			panic(errors.Errorf("invalid postfix op %s", op))
		}
		return PostfixOps[op]
	case OpTypeBinary:
		if _, ok := BinaryOps[op]; !ok {
			panic(errors.Errorf("invalid binary op %s", op))
		}
		return BinaryOps[op].Precedence
	default:
		panic(errors.Errorf("invalid opType for %s: %s", op, opType))
	}
}

func IsRightAssociative(op string, opType string) bool {
	switch opType {
	case OpTypePrefix:
		if _, ok := PrefixOps[op]; !ok {
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
	case OpTypePostfix:
		if _, ok := PostfixOps[op]; !ok {
			panic(errors.Errorf("invalid postfix op %s", op))
		}
		// TODO same comment as for the prefix above, but decree right associativity.
		return true
	case OpTypeBinary:
		if _, ok := BinaryOps[op]; !ok {
			panic(errors.Errorf("invalid binary op %s", op))
		}
		return BinaryOps[op].IsRightAssociative
	default:
		panic(errors.Errorf("invalid optype for %s: %s", op, opType))
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
			if _, ok := PrefixOps[op.Value]; ok {
				stack = append(stack, Prefix(op.Value))
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
		var newNode Node = Num(arg.Value)

		// 3. process any postfix operators
		for i < len(tokens) {
			op := tokens[i]
			if op.Type != TokenTypeOp {
				break
			}
			if _, ok := PostfixOps[op.Value]; !ok {
				break
			}
			for len(stack) > 0 {
				top := stack[len(stack)-1]
				currPrec, topPrec := GetPrecedence(op.Value, OpTypePostfix), GetPrecedence(top.Op, top.Type)
				if currPrec > topPrec {
					break
				} else if currPrec == topPrec {
					currIsRight := IsRightAssociative(op.Value, OpTypePostfix)
					if currIsRight != IsRightAssociative(top.Op, top.Type) {
						return nil, errors.Errorf("unable to handle same precedence but different associativity: %s vs %s", op.Value, top.Op)
					}
					if currIsRight {
						break
					}
					// left associative => continue loop, b/c top of stack wins
				}
				// top of stack wins
				// keep poppin' until top of stack doesn't win
				top.Args = append(top.Args, newNode)
				newNode = top
				stack = stack[:len(stack)-1]
			}
			newNode = Postfix(op.Value, newNode)
			i++
		}

		// no more tokens => done, so it's time to unwind the stack
		if i >= len(tokens) {
			// get the stack down to one completed expression
			for len(stack) > 0 {
				// pop the stack, add newNode as arg to popped top
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
			if _, ok := BinaryOps[op.Value]; !ok {
				return nil, errors.Errorf("expected binary operator, found %s at %d", op.Value, i)
			}
			for len(stack) > 0 {
				top := stack[len(stack)-1]
				topPrec, isTopRightAssociative := GetPrecedence(top.Op, top.Type), IsRightAssociative(top.Op, top.Type)
				if GetPrecedence(op.Value, OpTypeBinary) > topPrec {
					// next operator is higher precedence?  stop poppin'
					break
				}
				if GetPrecedence(op.Value, OpTypeBinary) == topPrec {
					if IsRightAssociative(op.Value, OpTypeBinary) != isTopRightAssociative {
						return nil, errors.Errorf("unable to handle same precedence but different associativity: %s vs %s", op.Value, top.Op)
					}
					// same precedence but right-associative?  stop poppin'
					if isTopRightAssociative {
						break
					}
					// same precedence but left-associative -- continue loop
				}
				// time to pop: remove the top, and complete it by adding in previous 'newNode' as its next operand.
				// then set 'newNode' to the popped top
				stack = stack[:len(stack)-1]
				top.Args = append(top.Args, newNode)
				newNode = top
			}
			stack = append(stack, Binary(op.Value, newNode))
		default:
			return nil, errors.Errorf("expected op at %d, found %+v", i, op)
		}
		i++
	}
}

func ParseString(s string) (Node, error) {
	return Parse(Must(Tokenize(s)))
}
