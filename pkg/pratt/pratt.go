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
	{"<-0", 0, false},
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

type TernOp struct {
	Open               string
	Close              string
	Precedence         int
	IsRightAssociative bool
}

var TernarySlice = []*TernOp{
	{"?", ":", 0, true},
	{"if", "else", 0, true},
	{"<l", "<lb", 0, false},
	{"<l1", "<l1b", 1, false},
}

var TernOpsOpen = map[string]*TernOp{}
var TernOpsClose = map[string]*TernOp{}

func init() {
	for _, b := range BinarySlice {
		if _, ok := BinaryOps[b.Symbol]; ok {
			panic(errors.Errorf("dupe binary symbol %s", b.Symbol))
		}
		BinaryOps[b.Symbol] = b
	}
	for _, t := range TernarySlice {
		if _, ok := TernOpsOpen[t.Open]; ok {
			panic(errors.Errorf("dupe ternary symbol %s", t.Open))
		}
		// TODO is it necessary to validate uniqueness of close symbols?
		TernOpsOpen[t.Open] = t
		TernOpsClose[t.Close] = t
	}
}

func GetPrecedence(op string, opType OpType) int {
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
	case OpTypeTernary:
		// this is weird, we only look for precedences of ternary opens.  huh
		if _, ok := TernOpsOpen[op]; !ok {
			panic(errors.Errorf("invalid ternary op %s", op))
		}
		return TernOpsOpen[op].Precedence
	default:
		panic(errors.Errorf("invalid opType for %s: %s", op, opType))
	}
}

func IsRightAssociative(op string, opType OpType) bool {
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
	case OpTypeTernary:
		// this is weird, we only look for associativity of ternary opens.  huh
		if _, ok := TernOpsOpen[op]; !ok {
			panic(errors.Errorf("invalid ternary op %s", op))
		}
		return TernOpsOpen[op].IsRightAssociative
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
				currPrec, topPrec := GetPrecedence(op.Value, OpTypePostfix), GetPrecedence(top.Open, top.Type)
				if currPrec > topPrec {
					break
				} else if currPrec == topPrec {
					currIsRight := IsRightAssociative(op.Value, OpTypePostfix)
					if currIsRight != IsRightAssociative(top.Open, top.Type) {
						return nil, errors.Errorf("unable to handle same precedence but different associativity: %s vs %s", op.Value, top.Open)
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
		if op.Type != TokenTypeOp {
			return nil, errors.Errorf("expected op at %d, found %+v", i, op)
		}
		var finish func(string, Node) *OpNode
		var opType OpType
		if _, ok := BinaryOps[op.Value]; ok {
			finish = Binary
			opType = OpTypeBinary
		} else if _, ok := TernOpsOpen[op.Value]; ok {
			finish = Ternary
			opType = OpTypeTernary
		} else if v, ok := TernOpsClose[op.Value]; ok {
			// pop stack until corresponding open is found
			for len(stack) > 0 {
				top := stack[len(stack)-1]
				top.Args = append(top.Args, newNode)
				if top.Open == v.Open {
					top.Close = op.Value
					break
				}
				newNode = top
				stack = stack[:len(stack)-1]
			}
			if len(stack) == 0 {
				return nil, errors.Errorf("ternary close '%s' found without corresponding open '%s' at %d", op.Value, v.Open, i)
			}
			i++
			continue
		} else {
			return nil, errors.Errorf("expected binary or ternary operator, found %s at %d", op.Value, i)
		}
		for len(stack) > 0 {
			top := stack[len(stack)-1]
			topPrec, isTopRightAssociative := GetPrecedence(top.Open, top.Type), IsRightAssociative(top.Open, top.Type)
			if GetPrecedence(op.Value, opType) > topPrec {
				// next operator is higher precedence?  stop poppin'
				break
			}
			if GetPrecedence(op.Value, opType) == topPrec {
				if IsRightAssociative(op.Value, opType) != isTopRightAssociative {
					return nil, errors.Errorf("unable to handle same precedence but different associativity: %s vs %s", op.Value, top.Open)
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
		stack = append(stack, finish(op.Value, newNode))
		i++
	}
}

func ParseString(s string) (Node, error) {
	return Parse(Must(Tokenize(s)))
}
