package logic

import "strings"

type BinOp string

const (
	AndOp           BinOp = "^"
	OrOp            BinOp = "v"
	ImplicationOp   BinOp = "->"
	BiconditionalOp BinOp = "<->"
)

type Term interface {
	TermPrint(isRoot bool) string
}

type PropTerm struct {
	Name string
}

func (v *PropTerm) TermPrint(isRoot bool) string {
	return v.Name
}

type NotTerm struct {
	Arg Term
}

func (n *NotTerm) TermPrint(isRoot bool) string {
	out := []string{"~"}
	out = append(out, n.Arg.TermPrint(false))
	return strings.Join(out, " ")
}

// TODO
// type ForallTerm struct {
//
// }
//
// type ExistentialTerm struct {
//
// }

type BinOpTerm struct {
	Op       BinOp
	LeftArg  Term
	RightArg Term
}

func (o *BinOpTerm) TermPrint(isRoot bool) string {
	var out []string
	if !isRoot {
		out = append(out, "(")
	}
	out = append(out, o.LeftArg.TermPrint(false))
	out = append(out, string(o.Op))
	out = append(out, o.RightArg.TermPrint(false))
	if !isRoot {
		out = append(out, ")")
	}
	return strings.Join(out, " ")
}

func Prop(name string) *PropTerm {
	return &PropTerm{Name: name}
}

func Not(arg Term) *NotTerm {
	return &NotTerm{Arg: arg}
}

func And(l, r Term) *BinOpTerm {
	return &BinOpTerm{Op: AndOp, LeftArg: l, RightArg: r}
}

func Or(l, r Term) *BinOpTerm {
	return &BinOpTerm{Op: OrOp, LeftArg: l, RightArg: r}
}

func Implication(l, r Term) *BinOpTerm {
	return &BinOpTerm{Op: ImplicationOp, LeftArg: l, RightArg: r}
}

func Biconditional(l, r Term) *BinOpTerm {
	return &BinOpTerm{Op: BiconditionalOp, LeftArg: l, RightArg: r}
}
