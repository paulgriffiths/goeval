package expr_test

import (
	"github.com/paulgriffiths/goeval/expr"
	"testing"
)

func TestEvaluateIsFunctions(t *testing.T) {
	n := expr.NewInt(3)
	r := expr.NewReal(3.0)
	b := expr.NewBool(true)
	s := expr.NewString("teeth")
	v := expr.NewVariable("vampire")

	testFuncs := []func(expr.Expr) bool{
		expr.IsInteger,
		expr.IsReal,
		expr.IsNumeric,
		expr.IsBoolean,
		expr.IsString,
		expr.IsVariable,
	}
	testCases := []struct {
		e       expr.Expr
		results [6]bool
	}{
		{n, [6]bool{true, false, true, false, false, false}},
		{r, [6]bool{false, true, true, false, false, false}},
		{b, [6]bool{false, false, false, true, false, false}},
		{s, [6]bool{false, false, false, false, true, false}},
		{v, [6]bool{false, false, false, false, false, true}},
	}

	for i, tc := range testCases {
		for j := 0; j < 6; j++ {
			result := testFuncs[j](tc.e)
			if result != tc.results[j] {
				t.Errorf("case %d.%d, got %v, want %v", i+1, j+i,
					result, tc.results[j])
			}
		}
	}
}

func TestEvaluateAreFunctions(t *testing.T) {
	n := expr.NewInt(3)
	r := expr.NewReal(3.0)
	b := expr.NewBool(true)
	s := expr.NewString("teeth")
	v := expr.NewVariable("vampire")

	testFuncs := []func(...expr.Expr) bool{
		expr.AreInteger,
		expr.AreReal,
		expr.AreNumeric,
		expr.AreBoolean,
		expr.AreString,
		expr.AreVariable,
	}
	testCases := []struct {
		e, f    expr.Expr
		results [6]bool
	}{
		{n, n, [6]bool{true, false, true, false, false, false}},
		{n, r, [6]bool{false, false, true, false, false, false}},
		{n, b, [6]bool{false, false, false, false, false, false}},
		{n, s, [6]bool{false, false, false, false, false, false}},
		{n, v, [6]bool{false, false, false, false, false, false}},
		{r, r, [6]bool{false, true, true, false, false, false}},
		{r, n, [6]bool{false, false, true, false, false, false}},
		{r, b, [6]bool{false, false, false, false, false, false}},
		{r, s, [6]bool{false, false, false, false, false, false}},
		{r, v, [6]bool{false, false, false, false, false, false}},
		{b, b, [6]bool{false, false, false, true, false, false}},
		{b, n, [6]bool{false, false, false, false, false, false}},
		{b, r, [6]bool{false, false, false, false, false, false}},
		{b, s, [6]bool{false, false, false, false, false, false}},
		{b, v, [6]bool{false, false, false, false, false, false}},
		{s, s, [6]bool{false, false, false, false, true, false}},
		{s, n, [6]bool{false, false, false, false, false, false}},
		{s, r, [6]bool{false, false, false, false, false, false}},
		{s, b, [6]bool{false, false, false, false, false, false}},
		{s, v, [6]bool{false, false, false, false, false, false}},
		{v, v, [6]bool{false, false, false, false, false, true}},
		{v, n, [6]bool{false, false, false, false, false, false}},
		{v, r, [6]bool{false, false, false, false, false, false}},
		{v, b, [6]bool{false, false, false, false, false, false}},
		{v, s, [6]bool{false, false, false, false, false, false}},
	}

	for i, tc := range testCases {
		for j := 0; j < 6; j++ {
			result := testFuncs[j](tc.e, tc.f)
			if result != tc.results[j] {
				t.Errorf("case %d.%d, got %v, want %v", i+1, j+i,
					result, tc.results[j])
			}
		}
	}
}
