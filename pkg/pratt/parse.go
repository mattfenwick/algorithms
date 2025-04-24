package pratt

import (
	"github.com/mattfenwick/algorithms/pkg/utils"
	"github.com/pkg/errors"
)

func unwindE(stack []*OpNode, newNode Node, pred func(top *OpNode) (bool, error)) ([]*OpNode, Node, bool, error) {
	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stop, err := pred(top)
		if err != nil {
			return nil, nil, false, err
		}
		if stop {
			return stack, newNode, true, nil
		}
		top.Args = append(top.Args, newNode)
		newNode = top
		stack = stack[:len(stack)-1]
	}
	return nil, newNode, false, nil
}

func unwind(stack []*OpNode, newNode Node, pred func(top *OpNode) bool) ([]*OpNode, Node, bool) {
	stack, newNode, found, err := unwindE(stack, newNode, func(top *OpNode) (bool, error) {
		return pred(top), nil
	})
	utils.Die0(err)
	return stack, newNode, found
}

func (o *Operators) IsGroupLeft(a string, aType OpType, b string, bType OpType) (bool, error) {
	aPrec, bPrec := o.GetPrecedence(a, aType), o.GetPrecedence(b, bType)
	if aPrec > bPrec {
		return true, nil
	} else if aPrec < bPrec {
		return false, nil
	}
	aAssocRight, bAssocRight := o.IsRightAssociative(a, aType), o.IsRightAssociative(b, bType)
	if aAssocRight != bAssocRight {
		return false, errors.Errorf("same precedence, different associativity: %s vs %s", a, b)
	}
	return !aAssocRight, nil
}

func (o *Operators) Parse(tokens []*Token) (Node, error) {
	var err error
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
			} else if _, ok := o.GroupingOpen[op.Value]; ok {
				stack = append(stack, Grouping(op.Value))
			} else {
				return nil, errors.Errorf("expected prefix op or gouping open at %d, found %+v", i, op)
			}
			i++
		}

		// 2. find a num
		arg := tokens[i]
		if arg.Type != TokenTypeNum {
			return nil, errors.Errorf("expected num at %d, found %+v", i, arg)
		}
		i++
		var newNode Node = Num(arg.Value)

		// 3. process any postfix or grouping-close operators
		for i < len(tokens) {
			op := tokens[i]
			if op.Type != TokenTypeOp {
				break
			}
			if _, ok := o.Postfix[op.Value]; ok {
				stack, newNode, _, err = unwindE(stack, newNode, func(top *OpNode) (bool, error) {
					isLeft, err := o.IsGroupLeft(top.Open, top.Type, op.Value, OpTypePostfix)
					return !isLeft, err
				})
				if err != nil {
					return nil, err
				}
				newNode = Postfix(op.Value, newNode)
			} else if groupingOpen, ok := o.GroupingClose[op.Value]; ok {
				var found bool
				stack, newNode, found = unwind(stack, newNode, func(top *OpNode) bool {
					if top.Type == OpTypeGrouping && top.Open == groupingOpen {
						top.Close = op.Value
						return true
					}
					return false
				})
				if !found {
					return nil, errors.Errorf("found grouping close %s, but no matching open %s at %d", op.Value, groupingOpen, i)
				}
				top := stack[len(stack)-1]
				top.Args = append(top.Args, newNode)
				stack = stack[:len(stack)-1]
				newNode = top
			} else {
				break
			}
			i++
		}

		// no more tokens => done, so it's time to unwind the stack
		if i >= len(tokens) {
			_, newNode, _ = unwind(stack, newNode, func(top *OpNode) bool { return false })
			return newNode, nil
		}

		// 4. find a binary, ternary-open, or ternary-close operator
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
			// 4a. ternary close
			var found bool
			stack, newNode, found = unwind(stack, newNode, func(top *OpNode) bool {
				if top.Open == v.Open {
					// TODO is there a better way to handle this mutation?  or a better place to put it?
					top.Args = append(top.Args, newNode)
					top.Close = op.Value
					return true
				}
				return false
			})
			if !found {
				return nil, errors.Errorf("ternary close '%s' found without corresponding open '%s' at %d", op.Value, v.Open, i)
			}
			i++
			continue
		} else {
			return nil, errors.Errorf("expected binary or ternary operator, found %s at %d", op.Value, i)
		}

		// 5. compare found op to top of stack
		stack, newNode, _, err = unwindE(stack, newNode, func(top *OpNode) (bool, error) {
			isLeft, err := o.IsGroupLeft(top.Open, top.Type, op.Value, opType)
			return !isLeft, err
		})
		if err != nil {
			return nil, err
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
