package expr_test

import (
    "testing"
    "github.com/paulgriffiths/goeval/expr"
)

func TestEvaluateArithmeticSuccess(t *testing.T) {
    table := expr.NewTable()
    table.Store("i", expr.NewInt(3))
    table.Store("j", expr.NewInt(6))
    table.Store("x", expr.NewReal(16.0))
    table.Store("y", expr.NewReal(4.0))

    i := expr.NewInt(3)
    j := expr.NewInt(6)
    x := expr.NewReal(16.0)
    y := expr.NewReal(4.0)
    vi := expr.NewVariable("i")
    vj := expr.NewVariable("j")
    vx := expr.NewVariable("x")
    vy := expr.NewVariable("y")

	testCases := []struct {
		exp    expr.Expr
		result expr.Expr
	}{
		{expr.NewInt(3), expr.NewInt(3)},
		{expr.NewReal(4.5), expr.NewReal(4.5)},
		{expr.NewAdd(i, j), expr.NewInt(9)},
		{expr.NewAdd(i, y), expr.NewReal(7.0)},
		{expr.NewAdd(x, j), expr.NewReal(22.0)},
		{expr.NewAdd(x, y), expr.NewReal(20.0)},
		{expr.NewAdd(vi, vj), expr.NewInt(9)},
		{expr.NewAdd(i, vy), expr.NewReal(7.0)},
		{expr.NewAdd(vx, j), expr.NewReal(22.0)},
		{expr.NewSub(i, j), expr.NewInt(-3)},
		{expr.NewSub(i, y), expr.NewReal(-1.0)},
		{expr.NewSub(x, j), expr.NewReal(10.0)},
		{expr.NewSub(x, y), expr.NewReal(12.0)},
		{expr.NewSub(vi, vj), expr.NewInt(-3)},
		{expr.NewSub(i, vy), expr.NewReal(-1.0)},
		{expr.NewSub(vx, j), expr.NewReal(10.0)},
		{expr.NewMul(i, j), expr.NewInt(18)},
		{expr.NewMul(i, y), expr.NewReal(12.0)},
		{expr.NewMul(x, j), expr.NewReal(96.0)},
		{expr.NewMul(x, y), expr.NewReal(64.0)},
		{expr.NewMul(vi, vj), expr.NewInt(18)},
		{expr.NewMul(i, vy), expr.NewReal(12.0)},
		{expr.NewMul(vx, j), expr.NewReal(96.0)},
		{expr.NewDiv(i, j), expr.NewReal(0.5)},
		{expr.NewDiv(j, i), expr.NewInt(2)},
		{expr.NewDiv(j, y), expr.NewReal(1.5)},
		{expr.NewDiv(x, y), expr.NewReal(4.0)},
		{expr.NewDiv(vi, vj), expr.NewReal(0.5)},
		{expr.NewDiv(j, vy), expr.NewReal(1.5)},
		{expr.NewDiv(vx, y), expr.NewReal(4.0)},
		{expr.NewPow(i, j), expr.NewInt(729)},
		{expr.NewPow(i, y), expr.NewReal(81.0)},
		{expr.NewPow(x, j), expr.NewReal(16777216.0)},
		{expr.NewPow(x, y), expr.NewReal(65536.0)},
		{expr.NewPow(vi, vj), expr.NewInt(729)},
		{expr.NewPow(i, vy), expr.NewReal(81.0)},
		{expr.NewPow(vx, j), expr.NewReal(16777216.0)},
		{expr.NewNeg(i), expr.NewInt(-3)},
		{expr.NewNeg(expr.NewSub(i, j)), expr.NewInt(3)},
		{expr.NewNeg(x), expr.NewReal(-16.0)},
		{expr.NewNeg(expr.NewSub(y, x)), expr.NewReal(12.0)},
	}

	for n, testCase := range testCases {
        result, err := testCase.exp.Evaluate(table)
        if err != nil {
            t.Errorf("case %d, couldn't evaluate expression: %v", n+1, err)
            continue
        }
		if !expr.Equals(result, testCase.result, table) {
			t.Errorf("case %d, got %v, want %v", n+1,
				result, testCase.result)
		}
	}
}

func TestEvaluateArithmeticNoSymbolTableSuccess(t *testing.T) {
    i := expr.NewInt(3)
    j := expr.NewInt(6)
    x := expr.NewReal(16.0)
    y := expr.NewReal(4.0)

	testCases := []struct {
		exp    expr.Expr
		result expr.Expr
	}{
		{expr.NewAdd(i, j), expr.NewInt(9)},
		{expr.NewAdd(x, y), expr.NewReal(20.0)},
		{expr.NewSub(i, j), expr.NewInt(-3)},
		{expr.NewSub(x, y), expr.NewReal(12.0)},
		{expr.NewMul(i, j), expr.NewInt(18)},
		{expr.NewMul(x, y), expr.NewReal(64.0)},
		{expr.NewDiv(i, j), expr.NewReal(0.5)},
		{expr.NewDiv(x, y), expr.NewReal(4.0)},
		{expr.NewPow(i, j), expr.NewInt(729)},
		{expr.NewPow(x, y), expr.NewReal(65536.0)},
		{expr.NewNeg(i), expr.NewInt(-3)},
		{expr.NewNeg(x), expr.NewReal(-16.0)},
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

func TestEvaluateArithmeticNoSymbolTableFailure(t *testing.T) {
    i := expr.NewInt(3)
    j := expr.NewInt(5)
    vi := expr.NewVariable("i")
    vj := expr.NewVariable("j")

	testCases := []struct {
		exp    expr.Expr
	}{
		{expr.NewAdd(i, vj)},
		{expr.NewAdd(vi, j)},
		{expr.NewAdd(vi, vj)},
		{expr.NewSub(i, vj)},
		{expr.NewSub(vi, j)},
		{expr.NewSub(vi, vj)},
		{expr.NewMul(i, vj)},
		{expr.NewMul(vi, j)},
		{expr.NewMul(vi, vj)},
		{expr.NewDiv(i, vj)},
		{expr.NewDiv(vi, j)},
		{expr.NewDiv(vi, vj)},
		{expr.NewPow(i, vj)},
		{expr.NewPow(vi, j)},
		{expr.NewPow(vi, vj)},
		{expr.NewNeg(vj)},
		{expr.NewNeg(vi)},
	}

	for n, testCase := range testCases {
        _, err := testCase.exp.Evaluate(nil)
        if err != expr.UndefinedVariableError {
            t.Errorf("case %d, got %v, want %v", n+1, err,
                expr.UndefinedVariableError)
            continue
        }
	}
}
