package expr_test

import (
	"github.com/paulgriffiths/goeval/expr"
	"testing"
)

func TestEvaluateComparisonsArithmetic(t *testing.T) {
	a := expr.NewInt(3)
	b := expr.NewInt(4)
	c := expr.NewInt(2)
	x := expr.NewReal(3.0)
	y := expr.NewReal(4.0)
	z := expr.NewReal(2.0)
	bt := expr.NewBool(true)
	bf := expr.NewBool(false)

	testFuncs := []func(expr.Expr, expr.Expr) expr.Expr{
		expr.NewEquality,
		expr.NewNonEquality,
		expr.NewLessThan,
		expr.NewLessThanOrEqual,
		expr.NewGreaterThan,
		expr.NewGreaterThanOrEqual,
	}
	testCases := []struct {
		l, r    expr.Expr
		results [6]expr.Expr
	}{
		{a, b, [6]expr.Expr{bf, bt, bt, bt, bf, bf}},
		{a, c, [6]expr.Expr{bf, bt, bf, bf, bt, bt}},
		{a, a, [6]expr.Expr{bt, bf, bf, bt, bf, bt}},
		{a, y, [6]expr.Expr{bf, bt, bt, bt, bf, bf}},
		{a, z, [6]expr.Expr{bf, bt, bf, bf, bt, bt}},
		{a, x, [6]expr.Expr{bt, bf, bf, bt, bf, bt}},
		{x, b, [6]expr.Expr{bf, bt, bt, bt, bf, bf}},
		{x, c, [6]expr.Expr{bf, bt, bf, bf, bt, bt}},
		{x, a, [6]expr.Expr{bt, bf, bf, bt, bf, bt}},
		{x, y, [6]expr.Expr{bf, bt, bt, bt, bf, bf}},
		{x, z, [6]expr.Expr{bf, bt, bf, bf, bt, bt}},
		{x, x, [6]expr.Expr{bt, bf, bf, bt, bf, bt}},
	}

	for i, tc := range testCases {
		for j := 0; j < 6; j++ {
			result, err := testFuncs[j](tc.l, tc.r).Evaluate(nil)
			if err != nil {
				t.Errorf("case %d.%d, couldn't evaluate expression: %v",
					i+1, j+1, err)
				continue
			}
			if !expr.Equals(result, tc.results[j], nil) {
				t.Errorf("case %d.%d, got %v, want %v", i+1, j+i,
					result, tc.results[j])
			}
		}
	}
}

func TestEvaluateComparisonsNonArithmetic(t *testing.T) {
	sa := expr.NewString("abc")
	sb := expr.NewString("def")
	bt := expr.NewBool(true)
	bf := expr.NewBool(false)

	testFuncs := []func(expr.Expr, expr.Expr) expr.Expr{
		expr.NewEquality,
		expr.NewNonEquality,
	}
	testCases := []struct {
		l, r    expr.Expr
		results [2]expr.Expr
	}{
		{bt, bt, [2]expr.Expr{bt, bf}},
		{bt, bf, [2]expr.Expr{bf, bt}},
		{bf, bt, [2]expr.Expr{bf, bt}},
		{bf, bf, [2]expr.Expr{bt, bf}},
		{sa, sa, [2]expr.Expr{bt, bf}},
		{sa, sb, [2]expr.Expr{bf, bt}},
		{sb, sa, [2]expr.Expr{bf, bt}},
		{sb, sb, [2]expr.Expr{bt, bf}},
	}

	for i, tc := range testCases {
		for j := 0; j < 2; j++ {
			result, err := testFuncs[j](tc.l, tc.r).Evaluate(nil)
			if err != nil {
				t.Errorf("case %d.%d, couldn't evaluate expression: %v",
					i+1, j+1, err)
				continue
			}
			if !expr.Equals(result, tc.results[j], nil) {
				t.Errorf("case %d.%d, got %v, want %v", i+1, j+i,
					result, tc.results[j])
			}
		}
	}
}
