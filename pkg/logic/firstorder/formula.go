package logic

import (
	"fmt"
	"strings"

	"github.com/mattfenwick/collections/pkg/set"
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

// GeneralizeFormula: Q(a)[x -> a] => ∃x.( Q(x) )
func GeneralizeFormula(formula Formula, from string, to string) Formula {
	return substituteVar(formula, from, to, true)
}

// InstantiateFormula: ∃x.( Q(x) )[x -> a] => Q(a)
func InstantiateFormula(formula Formula, from string, to string) Formula {
	return substituteVar(formula, from, to, false)
}

// substituteVar can generalize or instantiate
// - generalize: Q(a)[x -> a] => ∃x.( Q(x) )
// - instantiate: ∃x.( Q(x) )[x -> a] => Q(a)
func substituteVar(formula Formula, from string, to string, generalize bool) Formula {
	switch t := formula.(type) {
	case *NotFormula:
		return &NotFormula{Formula: substituteVar(t.Formula, from, to, generalize)}
	case *BinOpFormula:
		return &BinOpFormula{
			Op:    t.Op,
			Left:  substituteVar(t.Left, from, to, generalize),
			Right: substituteVar(t.Right, from, to, generalize),
		}
	case *PredicateFormula:
		return &PredicateFormula{
			Name: t.Name,
			Terms: slice.Map(func(p *TermVar) *TermVar {
				if p.Name == from {
					if generalize {
						if p.IsBound {
							panic(errors.Errorf("unable to generalize term var '%s' to '%s': already bound", from, to))
						}
						return BoundTermVar(to)
					} else {
						if !p.IsBound {
							panic(errors.Errorf("unable to instantiate term var '%s' to '%s': already instantiated", from, to))
						}
						return FreeTermVar(to)
					}
				}
				return p
			}, t.Terms),
		}
	case *QuantifiedFormula:
		return &QuantifiedFormula{
			Var:        t.Var,
			Body:       substituteVar(t.Body, from, to, generalize),
			Quantifier: t.Quantifier,
		}
	default:
		panic(errors.Errorf("invalid formula type: %T", t))
	}
}

func CheckTermVarsInScope(formula Formula, freeTermVars *set.Set[string]) *TermVarUsage {
	t := &TermVarUsage{}
	checkTermVarsInScopeHelper(formula, freeTermVars, set.Empty[string](), t)
	return t
}

type TermVarUsage struct {
	FreeViolations      []string
	BoundViolations     []string
	ShadowingViolations [][2]*TermVar
}

func (tvu *TermVarUsage) ViolationCount() int {
	return len(tvu.FreeViolations) + len(tvu.BoundViolations) + len(tvu.ShadowingViolations)
}

func (tvu *TermVarUsage) AddViolation(t *TermVar) {
	if t.IsBound {
		tvu.BoundViolations = append(tvu.BoundViolations, t.Name)
	} else {
		tvu.FreeViolations = append(tvu.FreeViolations, t.Name)
	}
}

func (tvu *TermVarUsage) AddShadowingViolation(outer *TermVar, inner *TermVar) {
	tvu.ShadowingViolations = append(tvu.ShadowingViolations, [2]*TermVar{outer, inner})
}

func checkTermVarsInScopeHelper(formula Formula, freeTermVars *set.Set[string], boundTermVars *set.Set[string], tvu *TermVarUsage) {
	switch t := formula.(type) {
	case *NotFormula:
		checkTermVarsInScopeHelper(t.Formula, freeTermVars, boundTermVars, tvu)
	case *BinOpFormula:
		checkTermVarsInScopeHelper(t.Left, freeTermVars, boundTermVars, tvu)
		checkTermVarsInScopeHelper(t.Right, freeTermVars, boundTermVars, tvu)
	case *PredicateFormula:
		for _, term := range t.Terms {
			if term.IsBound {
				if !boundTermVars.Contains(term.Name) {
					tvu.AddViolation(term)
				}
			} else {
				if !freeTermVars.Contains(term.Name) {
					tvu.AddViolation(term)
				}
			}
		}
	case *QuantifiedFormula:
		btv := set.FromSlice(boundTermVars.ToSlice())
		if freeTermVars.Contains(t.Var) {
			tvu.AddShadowingViolation(FreeTermVar(t.Var), BoundTermVar(t.Var))
		}
		added := btv.Add(t.Var)
		if !added {
			tvu.AddShadowingViolation(BoundTermVar(t.Var), BoundTermVar(t.Var))
		}
		checkTermVarsInScopeHelper(t.Body, freeTermVars, btv, tvu)
	default:
		panic(errors.Errorf("invalid formula type: %T", t))
	}
}

type TermVar struct {
	Name string
	// IsBound is true for 'x' and false for 'a' in ∃x.( ~ Q(x) ^ P(a) )
	IsBound bool
}

func BoundTermVar(name string) *TermVar {
	return &TermVar{Name: name, IsBound: true}
}

func FreeTermVar(name string) *TermVar {
	return &TermVar{Name: name, IsBound: false}
}

// PredicateFormula without args is used by 0th-order
// PredicateFormula with args is used by ∀/∃
type PredicateFormula struct {
	Name  string
	Terms []*TermVar
}

func (v *PredicateFormula) FormulaPrint(isRoot bool) string {
	if len(v.Terms) == 0 {
		return v.Name
	}
	return fmt.Sprintf("%s(%s)", v.Name, strings.Join(slice.Map(func(p *TermVar) string {
		return p.Name
	}, v.Terms), ","))
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

func Exist(varName string, formula Formula) *QuantifiedFormula {
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
	// TODO this feels wrong, shouldn't it drop the quantifier wrapper?
	return InstantiateFormula(f.Body, f.Var, obj)
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

func Pred(name string, termVars ...*TermVar) *PredicateFormula {
	return &PredicateFormula{Name: name, Terms: termVars}
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

func Arrow(l, r Formula) *BinOpFormula {
	return &BinOpFormula{Op: ImplicationOp, Left: l, Right: r}
}

func DArrow(l, r Formula) *BinOpFormula {
	return &BinOpFormula{Op: BiconditionalOp, Left: l, Right: r}
}
