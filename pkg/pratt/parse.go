package pratt

import (
	"github.com/pkg/errors"
)

func (o *Operators) Parse(tokens []*Token) (Node, error) {
	stack := []*OpNode{}
	for i := 0; ; {
		// 1. throw any prefix unary operators on to the stack
		for i < len(tokens) {
			op := tokens[i]
			if op.Type != TokenTypeOp {
				break
			}
			if _, ok := o.Prefix[op.Value]; ok {
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
			if _, ok := o.Postfix[op.Value]; !ok {
				break
			}
			for len(stack) > 0 {
				top := stack[len(stack)-1]
				currPrec, topPrec := o.GetPrecedence(op.Value, OpTypePostfix), o.GetPrecedence(top.Open, top.Type)
				if currPrec > topPrec {
					break
				} else if currPrec == topPrec {
					currIsRight := o.IsRightAssociative(op.Value, OpTypePostfix)
					if currIsRight != o.IsRightAssociative(top.Open, top.Type) {
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

		// 4. process a binary or ternary operator
		op := tokens[i]
		if op.Type != TokenTypeOp {
			return nil, errors.Errorf("expected op at %d, found %+v", i, op)
		}
		var finish func(string, Node) *OpNode
		var opType OpType
		if _, ok := o.Binary[op.Value]; ok {
			finish = Binary
			opType = OpTypeBinary
		} else if _, ok := o.TernaryOpen[op.Value]; ok {
			finish = Ternary
			opType = OpTypeTernary
		} else if v, ok := o.TernaryClose[op.Value]; ok {
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
			topPrec, isTopRightAssociative := o.GetPrecedence(top.Open, top.Type), o.IsRightAssociative(top.Open, top.Type)
			if o.GetPrecedence(op.Value, opType) > topPrec {
				// next operator is higher precedence?  stop poppin'
				break
			}
			if o.GetPrecedence(op.Value, opType) == topPrec {
				if o.IsRightAssociative(op.Value, opType) != isTopRightAssociative {
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

func (o *Operators) ParseString(s string) (Node, error) {
	tokens, err := o.Tokenize(s)
	if err != nil {
		return nil, err
	}
	return o.Parse(tokens)
}
