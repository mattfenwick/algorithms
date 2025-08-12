package logic

import "strings"

type Op string

const (
	AndOp           = "^"
	OrOp            = "v"
	ImplicationOp   = "->"
	BiconditionalOp = "<->"
)

type Term interface {
	TermPrint(isRoot bool) string
}

type VarTerm struct {
	Name string
}

func (v *VarTerm) TermPrint(isRoot bool) string {
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

type OpTerm struct {
	Op       Op
	LeftArg  Term
	RightArg Term
}

func (o *OpTerm) TermPrint(isRoot bool) string {
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

func Var(name string) *VarTerm {
	return &VarTerm{Name: name}
}

func Not(arg Term) *NotTerm {
	return &NotTerm{Arg: arg}
}

func And(l, r Term) *OpTerm {
	return &OpTerm{Op: AndOp, LeftArg: l, RightArg: r}
}

func Or(l, r Term) *OpTerm {
	return &OpTerm{Op: OrOp, LeftArg: l, RightArg: r}
}

func Implication(l, r Term) *OpTerm {
	return &OpTerm{Op: ImplicationOp, LeftArg: l, RightArg: r}
}

func Biconditional(l, r Term) *OpTerm {
	return &OpTerm{Op: BiconditionalOp, LeftArg: l, RightArg: r}
}
