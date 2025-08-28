package logic

import (
	"maps"

	"github.com/mattfenwick/algorithms/pkg/utils"
	"github.com/mattfenwick/collections/pkg/set"
	"github.com/mattfenwick/collections/pkg/slice"
	"github.com/pkg/errors"
)

func copyDict[K comparable, V any](d map[K]V) map[K]V {
	out := map[K]V{}
	maps.Copy(out, d)
	return out
}

func mappings(xs []string, vars map[string]bool) []map[string]bool {
	if len(xs) == 0 {
		return []map[string]bool{vars}
	}
	x := xs[0]
	c1 := copyDict(vars)
	c1[x] = true
	out := mappings(xs[1:], c1)
	c2 := copyDict(vars)
	c2[x] = false
	out = append(out, mappings(xs[1:], c2)...)
	return out
}

func TruthTable(term Term, env *Env) *utils.Table {
	rawVars, exprs := findVarsAndExpressions(term)
	vars := slice.Sort(set.FromSlice(rawVars).ToSlice())
	var rows [][]string
	for _, mapping := range mappings(vars, map[string]bool{}) {
		originalMapping := copyDict(env.PropToTruth)
		maps.Copy(originalMapping, mapping)
		env := &Env{PropToTruth: originalMapping, Domain: env.Domain}
		var row []string
		for _, v := range vars {
			value := "F"
			e, ok := env.PropToTruth[v]
			if !ok {
				panic(errors.Errorf("expected mapping for %s", v))
			}
			if e {
				value = "T"
			}
			row = append(row, value)
		}
		for _, expr := range exprs {
			result, err := env.Evaluate(expr)
			if err != nil {
				panic(err)
			}
			// fmt.Printf("result for env %+v, %+v, %+v\n", env, expr, result)
			value := "T"
			if !result {
				value = "F"
			}
			row = append(row, value)
		}
		rows = append(rows, row)
	}
	headers := append(vars, slice.Map(func(t Term) string { return t.TermPrint(true) }, exprs)...)
	return utils.NewTable(headers, rows...)
}

func TruthTableFromTerms(env *Env, termsRow ...[]Term) *utils.Table {
	var rows [][]string
	for _, terms := range termsRow {
		var row []string
		for _, term := range terms {
			result, err := env.Evaluate(term)
			if err != nil {
				panic(err)
			}
			// fmt.Printf("result for env %+v, %+v, %+v\n", env, expr, result)
			value := "T"
			if !result {
				value = "F"
			}
			row = append(row, term.TermPrint(true)+"\n"+value)
		}
		rows = append(rows, row)
	}
	headers := slice.Map(func(t Term) string { return "" }, termsRow[0])
	return utils.NewTable(headers, rows...)
}

func findVarsAndExpressions(term Term) ([]string, []Term) {
	vars := []string{}
	exprs := []Term{}
	switch t := term.(type) {
	case *NotTerm:
		vars, exprs = findVarsAndExpressions(t.Term)
		exprs = append(exprs, t)
	case *PropTerm:
		if len(t.Args) == 0 {
			vars = append(vars, t.Name)
		}
		// only do things like P or Q; skip things like P(x) or Q(x)
	case *BinOpTerm:
		vars, exprs = findVarsAndExpressions(t.Left)
		rightVars, rightExprs := findVarsAndExpressions(t.Right)
		vars = append(vars, rightVars...)
		exprs = append(exprs, rightExprs...)
		exprs = append(exprs, t)
	case *QuantifiedTerm:
		// ignore exprs from inside quantified term since they won't be substituted
		vars, _ = findVarsAndExpressions(t.Body)
		exprs = append(exprs, t)
	default:
		panic(errors.Errorf("invalid term type: %T", term))
	}
	return vars, exprs
}
