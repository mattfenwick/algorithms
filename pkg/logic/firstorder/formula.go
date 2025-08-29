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

type Formula interface {
	FormulaPrint(isRoot bool) string
}

type PredicateArg struct {
	Var string
}

// PredicateFormula without args is used by 0th-order
// PredicateFormula with args is used by ∀/∃
type PredicateFormula struct {
	Name string
	Args []*PredicateArg
}

func (v *PredicateFormula) FormulaPrint(isRoot bool) string {
	if len(v.Args) == 0 {
		return v.Name
	}
	return fmt.Sprintf("%s(%s)", v.Name, strings.Join(slice.Map(func(p *PredicateArg) string {
		return p.Var
	}, v.Args), ","))
}

type NotFormula struct {
	Formula Formula
}

func (n *NotFormula) FormulaPrint(isRoot bool) string {
	out := []string{"~"}
	out = append(out, n.Formula.FormulaPrint(false))
	return strings.Join(out, " ")
}

type QuantifiedFormula struct {
	Var        string
	Body       Formula
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
//	∀x.(P(y)) -- iff y is bound in an enclosing quantified formula
func NewQuantifiedFormula(varName string, formula Formula, quantifier Quantifier) *QuantifiedFormula {
	// TODO shadowing
	return &QuantifiedFormula{Var: varName, Body: formula, Quantifier: quantifier}
}

func Forall(varName string, formula Formula) *QuantifiedFormula {
	return NewQuantifiedFormula(varName, formula, ForallQuantifier)
}

func Existential(varName string, formula Formula) *QuantifiedFormula {
	return NewQuantifiedFormula(varName, formula, ExistentialQuantifier)
}

func (f *QuantifiedFormula) FormulaPrint(isRoot bool) string {
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
	// intentionally cause subformula to leave off outer parens
	return fmt.Sprintf("%s%s.( %s )", symbol, f.Var, f.Body.FormulaPrint(true))
}

func (f *QuantifiedFormula) Substitute(obj string) Formula {
	return substituteVar(f.Body, f.Var, obj)
}

func substituteVar(formula Formula, from string, to string) Formula {
	switch t := formula.(type) {
	case *NotFormula:
		return &NotFormula{Formula: substituteVar(t.Formula, from, to)}
	case *BinOpFormula:
		return &BinOpFormula{
			Op:    t.Op,
			Left:  substituteVar(t.Left, from, to),
			Right: substituteVar(t.Right, from, to),
		}
	case *PredicateFormula:
		return &PredicateFormula{
			Name: t.Name,
			Args: slice.Map(func(p *PredicateArg) *PredicateArg {
				if p.Var == from {
					return &PredicateArg{Var: to}
				}
				return p
			}, t.Args),
		}
	case *QuantifiedFormula:
		return &QuantifiedFormula{
			Var:        t.Var,
			Body:       substituteVar(t.Body, from, to),
			Quantifier: t.Quantifier,
		}
	default:
		panic(errors.Errorf("invalid formula type: %T", t))
	}
}

type BinOpFormula struct {
	Op    BinOp
	Left  Formula
	Right Formula
}

func (o *BinOpFormula) FormulaPrint(isRoot bool) string {
	var out []string
	if !isRoot {
		out = append(out, "(")
	}
	out = append(out, o.Left.FormulaPrint(false))
	out = append(out, string(o.Op))
	out = append(out, o.Right.FormulaPrint(false))
	if !isRoot {
		out = append(out, ")")
	}
	return strings.Join(out, " ")
}

func Prop(name string, args ...string) *PredicateFormula {
	return &PredicateFormula{Name: name, Args: slice.Map(func(a string) *PredicateArg {
		return &PredicateArg{Var: a}
	}, args)}
}

func Not(arg Formula) *NotFormula {
	return &NotFormula{Formula: arg}
}

func And(l, r Formula) *BinOpFormula {
	return &BinOpFormula{Op: AndOp, Left: l, Right: r}
}

func Or(l, r Formula) *BinOpFormula {
	return &BinOpFormula{Op: OrOp, Left: l, Right: r}
}

func Implication(l, r Formula) *BinOpFormula {
	return &BinOpFormula{Op: ImplicationOp, Left: l, Right: r}
}

func Biconditional(l, r Formula) *BinOpFormula {
	return &BinOpFormula{Op: BiconditionalOp, Left: l, Right: r}
}
