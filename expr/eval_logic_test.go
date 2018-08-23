package expr_test

import (
	"github.com/paulgriffiths/goeval/expr"
	"testing"
)

func TestEvaluateLogicalEquality(t *testing.T) {
	a := expr.NewBool(true)
	b := expr.NewBool(false)

	testCases := []struct {
		exp    expr.Expr
		result expr.Expr
	}{
		{expr.NewEquality(a, a), expr.NewBool(true)},
		{expr.NewEquality(a, b), expr.NewBool(false)},
		{expr.NewEquality(b, a), expr.NewBool(false)},
		{expr.NewEquality(b, b), expr.NewBool(true)},
	}

	for n, testCase := range testCases {
		result, err := testCase.exp.Evaluate(nil)
		if err != nil {
			t.Errorf("case %d, couldn't evaluate expression: %v", n+1, err)
			continue
		}
		if !expr.Equals(result, testCase.result, nil) {
			t.Errorf("case %d, got %v, want %v", n+1,
				result, testCase.result)
		}
	}
}

func TestEvaluateLogicalNonEquality(t *testing.T) {
	a := expr.NewBool(true)
	b := expr.NewBool(false)

	testCases := []struct {
		exp    expr.Expr
		result expr.Expr
	}{
		{expr.NewNonEquality(a, a), expr.NewBool(false)},
		{expr.NewNonEquality(a, b), expr.NewBool(true)},
		{expr.NewNonEquality(b, a), expr.NewBool(true)},
		{expr.NewNonEquality(b, b), expr.NewBool(false)},
	}

	for n, testCase := range testCases {
		result, err := testCase.exp.Evaluate(nil)
		if err != nil {
			t.Errorf("case %d, couldn't evaluate expression: %v", n+1, err)
			continue
		}
		if !expr.Equals(result, testCase.result, nil) {
			t.Errorf("case %d, got %v, want %v", n+1,
				result, testCase.result)
		}
	}
}

func TestEvaluateLogicalAnd(t *testing.T) {
	a := expr.NewBool(true)
	b := expr.NewBool(false)

	testCases := []struct {
		exp    expr.Expr
		result expr.Expr
	}{
		{expr.NewAnd(a, a), expr.NewBool(true)},
		{expr.NewAnd(a, b), expr.NewBool(false)},
		{expr.NewAnd(b, a), expr.NewBool(false)},
		{expr.NewAnd(b, b), expr.NewBool(false)},
	}

	for n, testCase := range testCases {
		result, err := testCase.exp.Evaluate(nil)
		if err != nil {
			t.Errorf("case %d, couldn't evaluate expression: %v", n+1, err)
			continue
		}
		if !expr.Equals(result, testCase.result, nil) {
			t.Errorf("case %d, got %v, want %v", n+1,
				result, testCase.result)
		}
	}
}

func TestEvaluateLogicalOr(t *testing.T) {
	a := expr.NewBool(true)
	b := expr.NewBool(false)

	testCases := []struct {
		exp    expr.Expr
		result expr.Expr
	}{
		{expr.NewOr(a, a), expr.NewBool(true)},
		{expr.NewOr(a, b), expr.NewBool(true)},
		{expr.NewOr(b, a), expr.NewBool(true)},
		{expr.NewOr(b, b), expr.NewBool(false)},
	}

	for n, testCase := range testCases {
		result, err := testCase.exp.Evaluate(nil)
		if err != nil {
			t.Errorf("case %d, couldn't evaluate expression: %v", n+1, err)
			continue
		}
		if !expr.Equals(result, testCase.result, nil) {
			t.Errorf("case %d, got %v, want %v", n+1,
				result, testCase.result)
		}
	}
}

func TestEvaluateLogicalXor(t *testing.T) {
	a := expr.NewBool(true)
	b := expr.NewBool(false)

	testCases := []struct {
		exp    expr.Expr
		result expr.Expr
	}{
		{expr.NewXor(a, a), expr.NewBool(false)},
		{expr.NewXor(a, b), expr.NewBool(true)},
		{expr.NewXor(b, a), expr.NewBool(true)},
		{expr.NewXor(b, b), expr.NewBool(false)},
	}

	for n, testCase := range testCases {
		result, err := testCase.exp.Evaluate(nil)
		if err != nil {
			t.Errorf("case %d, couldn't evaluate expression: %v", n+1, err)
			continue
		}
		if !expr.Equals(result, testCase.result, nil) {
			t.Errorf("case %d, got %v, want %v", n+1,
				result, testCase.result)
		}
	}
}

func TestEvaluateLogicalNor(t *testing.T) {
	a := expr.NewBool(true)
	b := expr.NewBool(false)

	testCases := []struct {
		exp    expr.Expr
		result expr.Expr
	}{
		{expr.NewNor(a, a), expr.NewBool(false)},
		{expr.NewNor(a, b), expr.NewBool(false)},
		{expr.NewNor(b, a), expr.NewBool(false)},
		{expr.NewNor(b, b), expr.NewBool(true)},
	}

	for n, testCase := range testCases {
		result, err := testCase.exp.Evaluate(nil)
		if err != nil {
			t.Errorf("case %d, couldn't evaluate expression: %v", n+1, err)
			continue
		}
		if !expr.Equals(result, testCase.result, nil) {
			t.Errorf("case %d, got %v, want %v", n+1,
				result, testCase.result)
		}
	}
}

func TestEvaluateLogicalNand(t *testing.T) {
	a := expr.NewBool(true)
	b := expr.NewBool(false)

	testCases := []struct {
		exp    expr.Expr
		result expr.Expr
	}{
		{expr.NewNand(a, a), expr.NewBool(false)},
		{expr.NewNand(a, b), expr.NewBool(true)},
		{expr.NewNand(b, a), expr.NewBool(true)},
		{expr.NewNand(b, b), expr.NewBool(true)},
	}

	for n, testCase := range testCases {
		result, err := testCase.exp.Evaluate(nil)
		if err != nil {
			t.Errorf("case %d, couldn't evaluate expression: %v", n+1, err)
			continue
		}
		if !expr.Equals(result, testCase.result, nil) {
			t.Errorf("case %d, got %v, want %v", n+1,
				result, testCase.result)
		}
	}
}

func TestEvaluateLogicalNot(t *testing.T) {
	a := expr.NewBool(true)
	b := expr.NewBool(false)

	testCases := []struct {
		exp    expr.Expr
		result expr.Expr
	}{
		{expr.NewNot(a), expr.NewBool(false)},
		{expr.NewNot(b), expr.NewBool(true)},
	}

	for n, testCase := range testCases {
		result, err := testCase.exp.Evaluate(nil)
		if err != nil {
			t.Errorf("case %d, couldn't evaluate expression: %v", n+1, err)
			continue
		}
		if !expr.Equals(result, testCase.result, nil) {
			t.Errorf("case %d, got %v, want %v", n+1,
				result, testCase.result)
		}
	}
}

func TestEvaluateLogicalCompound(t *testing.T) {
	a := expr.NewBool(true)
	b := expr.NewBool(false)

	testCases := []struct {
		exp    expr.Expr
		result expr.Expr
	}{
		{expr.NewAnd(a, expr.NewOr(a, b)), expr.NewBool(true)},
		{expr.NewAnd(b, expr.NewOr(a, b)), expr.NewBool(false)},
		{expr.NewAnd(b, expr.NewNot(expr.NewOr(a, b))),
			expr.NewBool(false)},
		{expr.NewAnd(expr.NewNot(b), expr.NewOr(a, b)),
			expr.NewBool(true)},
	}

	for n, testCase := range testCases {
		result, err := testCase.exp.Evaluate(nil)
		if err != nil {
			t.Errorf("case %d, couldn't evaluate expression: %v", n+1, err)
			continue
		}
		if !expr.Equals(result, testCase.result, nil) {
			t.Errorf("case %d, got %v, want %v", n+1,
				result, testCase.result)
		}
	}
}
