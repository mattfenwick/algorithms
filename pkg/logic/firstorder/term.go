package logic

import (
	"fmt"
	"strings"

	"github.com/mattfenwick/collections/pkg/slice"
	"github.com/pkg/errors"
)

const (
	ForallSymbol      = "∀"
	ExistentialSymbol = "∃"
)

type BinOp string

const (
	AndOp           BinOp = "^"
	OrOp            BinOp = "v"
	ImplicationOp   BinOp = "->"
	BiconditionalOp BinOp = "<->"
)

type Quantifier string

const (
	ForallQuantifier      Quantifier = "ForallQuantifier"
	ExistentialQuantifier Quantifier = "ExistentialQuantifier"
)

type Term interface {
	TermPrint(isRoot bool) string
}

type PropArg struct {
	Var    string
	Object *string
}

// PropTerm without args is used by 0th-order
// PropTerm with args is used by ∀/∃
type PropTerm struct {
	Name string
	Args []*PropArg
}

func (v *PropTerm) TermPrint(isRoot bool) string {
	if len(v.Args) == 0 {
		return v.Name
	}
	return fmt.Sprintf("%s(%s)", v.Name, strings.Join(slice.Map(func(p *PropArg) string {
		if p.Object != nil {
			return *p.Object
		}
		return p.Var
	}, v.Args), ","))
}

type NotTerm struct {
	Term Term
}

func (n *NotTerm) TermPrint(isRoot bool) string {
	out := []string{"~"}
	out = append(out, n.Term.TermPrint(false))
	return strings.Join(out, " ")
}

type QuantifiedTerm struct {
	Arg        string
	Body       Term
	Quantifier Quantifier
}

// okay:
//
//	∀x.(~P)
//	∀x.( Q(x) )
//	∃x.( Q(x) ^ ∀y.( R(y) ^ S(x) ) )
//	∃x.( ∀y.(R(x,y)) ^ ∀y.(S(y) ) -- two different `y`s, but they aren't shadowing each other
//
// not okay
//
//	∀x.(x) -- x is an object from the domain of discourse; it is not a proposition
//	∃x.( Q(x) ^ ∀x.(R(x)) ) -- the inner forall introduces a new `x` which shadows the outer `x`
//
// maybe okay
//
//	∀x.(P(y)) -- iff y is bound in an enclosing quantified term
func NewQuantifiedTerm(arg string, term Term, quantifier Quantifier) *QuantifiedTerm {
	// TODO shadowing
	return &QuantifiedTerm{Arg: arg, Body: term, Quantifier: quantifier}
}

func Forall(arg string, term Term) *QuantifiedTerm {
	return NewQuantifiedTerm(arg, term, ForallQuantifier)
}

func Existential(arg string, term Term) *QuantifiedTerm {
	return NewQuantifiedTerm(arg, term, ExistentialQuantifier)
}

func (f *QuantifiedTerm) TermPrint(isRoot bool) string {
	// examples
	// ∀x.( ~ P )
	// ∀x.( Q(x) ^ R(y) )
	var symbol string
	switch f.Quantifier {
	case ForallQuantifier:
		symbol = "∀"
	case ExistentialQuantifier:
		symbol = "∃"
	default:
		panic(errors.Errorf("invalid quantifier %s", f.Quantifier))
	}
	// intentionally cause subterm to leave off outer parens
	return fmt.Sprintf("%s%s.( %s )", symbol, f.Arg, f.Body.TermPrint(true))
}

func (f *QuantifiedTerm) Substitute(obj string) Term {
	return substituteVar(f.Body, f.Arg, obj)
}

func substituteVar(term Term, varName string, obj string) Term {
	switch t := term.(type) {
	case *NotTerm:
		return &NotTerm{Term: substituteVar(t.Term, varName, obj)}
	case *BinOpTerm:
		return &BinOpTerm{
			Op:    t.Op,
			Left:  substituteVar(t.Left, varName, obj),
			Right: substituteVar(t.Right, varName, obj),
		}
	case *PropTerm:
		return &PropTerm{
			Name: t.Name,
			Args: slice.Map(func(p *PropArg) *PropArg {
				if p.Var == varName {
					if p.Object != nil {
						panic(errors.Errorf("cannot substitute for proparg %+v: already substituted", p))
					}
					return &PropArg{
						Var:    p.Var,
						Object: &obj,
					}
				}
				return p
			}, t.Args),
		}
	case *QuantifiedTerm:
		return &QuantifiedTerm{
			Arg:        t.Arg,
			Body:       substituteVar(t.Body, varName, obj),
			Quantifier: t.Quantifier,
		}
	default:
		panic(errors.Errorf("invalid term type: %T", t))
	}
}

type BinOpTerm struct {
	Op    BinOp
	Left  Term
	Right Term
}

func (o *BinOpTerm) TermPrint(isRoot bool) string {
	var out []string
	if !isRoot {
		out = append(out, "(")
	}
	out = append(out, o.Left.TermPrint(false))
	out = append(out, string(o.Op))
	out = append(out, o.Right.TermPrint(false))
	if !isRoot {
		out = append(out, ")")
	}
	return strings.Join(out, " ")
}

func Prop(name string, args ...string) *PropTerm {
	return &PropTerm{Name: name, Args: slice.Map(func(a string) *PropArg {
		return &PropArg{Var: a, Object: nil}
	}, args)}
}

func Not(arg Term) *NotTerm {
	return &NotTerm{Term: arg}
}

func And(l, r Term) *BinOpTerm {
	return &BinOpTerm{Op: AndOp, Left: l, Right: r}
}

func Or(l, r Term) *BinOpTerm {
	return &BinOpTerm{Op: OrOp, Left: l, Right: r}
}

func Implication(l, r Term) *BinOpTerm {
	return &BinOpTerm{Op: ImplicationOp, Left: l, Right: r}
}

func Biconditional(l, r Term) *BinOpTerm {
	return &BinOpTerm{Op: BiconditionalOp, Left: l, Right: r}
}
