package pratt

import (
	"strconv"
	"strings"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

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

var (
	digits = map[byte]bool{
		'0': true,
		'1': true,
		'2': true,
		'3': true,
		'4': true,
		'5': true,
		'6': true,
		'7': true,
		'8': true,
		'9': true,
	}
)

func (o *Operators) Tokenize(s string) ([]*Token, error) {
	var tokens []*Token
	for _, field := range strings.Fields(s) {
		if len(field) == 0 {
			return nil, errors.Errorf("invalid empty token")
		}

		// operator.  we throw away the the operator type information here,
		//   because we don't need it and it might be inaccurate anyway
		//   (example: unary - vs. binary -)
		// prefix op
		if _, ok := o.Prefix[field]; ok {
			tokens = append(tokens, &Token{Type: TokenTypeOp, Value: field})
			continue
		}
		// postfix op
		if _, ok := o.Postfix[field]; ok {
			tokens = append(tokens, &Token{Type: TokenTypeOp, Value: field})
			continue
		}
		// binop
		if _, ok := o.Binary[field]; ok {
			tokens = append(tokens, &Token{Type: TokenTypeOp, Value: field})
			continue
		}
		// ternary op -- open
		if _, ok := o.TernaryOpen[field]; ok {
			tokens = append(tokens, &Token{Type: TokenTypeOp, Value: field})
			continue
		}
		// ternary op -- close
		if _, ok := o.TernaryClose[field]; ok {
			tokens = append(tokens, &Token{Type: TokenTypeOp, Value: field})
			continue
		}

		// number
		if _, ok := digits[field[0]]; !ok {
			return nil, errors.Errorf("invalid num literal: '%s'", field)
		}
		val, err := strconv.ParseInt(field, 10, 32)
		if err != nil {
			return nil, errors.Wrapf(err, "invalid num literal '%s'", field)
		}
		logrus.Warnf("TODO: don't throw away the int token value %d", val)
		tokens = append(tokens, &Token{Type: TokenTypeNum, Value: field})
	}
	return tokens, nil
}
