package pratt

import "github.com/pkg/errors"

type BinOp struct {
	Symbol             string
	Precedence         int
	IsRightAssociative bool
}

type TernOp struct {
	Open               string
	Close              string
	Precedence         int
	IsRightAssociative bool
}

type Operators struct {
	Prefix        map[string]int
	Postfix       map[string]int
	BinarySlice   []*BinOp
	Binary        map[string]*BinOp
	TernarySlice  []*TernOp
	TernaryOpen   map[string]*TernOp
	TernaryClose  map[string]*TernOp
	GroupingOpen  map[string]string
	GroupingClose map[string]string
}

func NewOperators(
	prefix map[string]int,
	postfix map[string]int,
	binarySlice []*BinOp,
	ternarySlice []*TernOp,
	grouping map[string]string) *Operators {
	binary := map[string]*BinOp{}
	for _, b := range binarySlice {
		if _, ok := binary[b.Symbol]; ok {
			panic(errors.Errorf("dupe binary symbol %s", b.Symbol))
		}
		binary[b.Symbol] = b
	}
	ternaryOpen := map[string]*TernOp{}
	ternaryClose := map[string]*TernOp{}
	for _, t := range ternarySlice {
		if _, ok := ternaryOpen[t.Open]; ok {
			panic(errors.Errorf("dupe ternary open symbol %s", t.Open))
		}
		if _, ok := ternaryClose[t.Close]; ok {
			panic(errors.Errorf("dupe ternary close symbol %s", t.Close))
		}
		ternaryOpen[t.Open] = t
		ternaryClose[t.Close] = t
	}
	groupingClose := map[string]string{}
	for k, v := range grouping {
		if _, ok := groupingClose[v]; ok {
			panic(errors.Errorf("dupe grouping close symbol %s", v))
		}
		groupingClose[v] = k
	}
	return &Operators{
		Prefix:        prefix,
		Postfix:       postfix,
		BinarySlice:   binarySlice,
		Binary:        binary,
		TernarySlice:  ternarySlice,
		TernaryOpen:   ternaryOpen,
		TernaryClose:  ternaryClose,
		GroupingOpen:  grouping,
		GroupingClose: groupingClose,
	}
}

func (o *Operators) GetPrecedence(op string, opType OpType) int {
	switch opType {
	case OpTypePrefix:
		if _, ok := o.Prefix[op]; !ok {
			panic(errors.Errorf("invalid prefix op %s", op))
		}
		return o.Prefix[op]
	case OpTypePostfix:
		if _, ok := o.Postfix[op]; !ok {
			panic(errors.Errorf("invalid postfix op %s", op))
		}
		return o.Postfix[op]
	case OpTypeBinary:
		if _, ok := o.Binary[op]; !ok {
			panic(errors.Errorf("invalid binary op %s", op))
		}
		return o.Binary[op].Precedence
	case OpTypeTernary:
		// this is weird, we only look for precedences of ternary opens.  huh
		if _, ok := o.TernaryOpen[op]; !ok {
			panic(errors.Errorf("invalid ternary op %s", op))
		}
		return o.TernaryOpen[op].Precedence
	case OpTypeGrouping:
		// TODO the intention is that Grouping has the absolute lowest precedence
		return -1000
	default:
		panic(errors.Errorf("invalid opType for %s: %s", op, opType))
	}
}

func (o *Operators) IsRightAssociative(op string, opType OpType) bool {
	switch opType {
	case OpTypePrefix:
		if _, ok := o.Prefix[op]; !ok {
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
		if _, ok := o.Postfix[op]; !ok {
			panic(errors.Errorf("invalid postfix op %s", op))
		}
		// TODO same comment as for the prefix above, but decree right associativity.
		return true
	case OpTypeBinary:
		if _, ok := o.Binary[op]; !ok {
			panic(errors.Errorf("invalid binary op %s", op))
		}
		return o.Binary[op].IsRightAssociative
	case OpTypeTernary:
		// this is weird, we only look for associativity of ternary opens.  huh
		if _, ok := o.TernaryOpen[op]; !ok {
			panic(errors.Errorf("invalid ternary op %s", op))
		}
		return o.TernaryOpen[op].IsRightAssociative
	case OpTypeGrouping:
		panic(errors.Errorf("TODO -- make decision"))
	default:
		panic(errors.Errorf("invalid optype for %s: %s", op, opType))
	}
}
