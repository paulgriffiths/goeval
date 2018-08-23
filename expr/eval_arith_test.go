package expr_test

import (
	"github.com/paulgriffiths/goeval/expr"
	"testing"
)

func TestEvaluateArithmeticAdd(t *testing.T) {
	i := expr.NewInt(4)
	j := expr.NewInt(5)
	x := expr.NewReal(4.0)
	y := expr.NewReal(5.0)

	testCases := []struct {
		exp    expr.Expr
		result expr.Expr
	}{
		{expr.NewAdd(i, j), expr.NewInt(9)},
		{expr.NewAdd(i, y), expr.NewReal(9.0)},
		{expr.NewAdd(x, j), expr.NewReal(9.0)},
		{expr.NewAdd(x, y), expr.NewReal(9.0)},
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

func TestEvaluateArithmeticSubtract(t *testing.T) {
	i := expr.NewInt(4)
	j := expr.NewInt(5)
	x := expr.NewReal(4.0)
	y := expr.NewReal(5.0)

	testCases := []struct {
		exp    expr.Expr
		result expr.Expr
	}{
		{expr.NewSub(i, j), expr.NewInt(-1)},
		{expr.NewSub(i, y), expr.NewReal(-1.0)},
		{expr.NewSub(x, j), expr.NewReal(-1.0)},
		{expr.NewSub(x, y), expr.NewReal(-1.0)},
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

func TestEvaluateArithmeticMultiply(t *testing.T) {
	i := expr.NewInt(4)
	j := expr.NewInt(5)
	x := expr.NewReal(4.0)
	y := expr.NewReal(5.0)

	testCases := []struct {
		exp    expr.Expr
		result expr.Expr
	}{
		{expr.NewMul(i, j), expr.NewInt(20)},
		{expr.NewMul(i, y), expr.NewReal(20.0)},
		{expr.NewMul(x, j), expr.NewReal(20.0)},
		{expr.NewMul(x, y), expr.NewReal(20.0)},
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

func TestEvaluateArithmeticDivide(t *testing.T) {
	i := expr.NewInt(4)
	j := expr.NewInt(2)
	x := expr.NewReal(5.0)
	y := expr.NewReal(2.0)

	testCases := []struct {
		exp    expr.Expr
		result expr.Expr
	}{
		{expr.NewDiv(i, j), expr.NewInt(2)},
		{expr.NewDiv(j, i), expr.NewReal(0.5)},
		{expr.NewDiv(i, y), expr.NewReal(2.0)},
		{expr.NewDiv(x, j), expr.NewReal(2.5)},
		{expr.NewDiv(x, y), expr.NewReal(2.5)},
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

func TestEvaluateArithmeticDivideByZero(t *testing.T) {
	i := expr.NewInt(6)
	j := expr.NewInt(0)
	x := expr.NewReal(5.0)
	y := expr.NewReal(0.0)

	testCases := []struct {
		exp expr.Expr
	}{
		{expr.NewDiv(i, j)},
		{expr.NewDiv(i, y)},
		{expr.NewDiv(x, j)},
		{expr.NewDiv(x, y)},
	}

	for n, testCase := range testCases {
		_, err := testCase.exp.Evaluate(nil)
		if err != expr.DivideByZeroError {
			t.Errorf("case %d, got %v, want %v", n+1, err,
				expr.DivideByZeroError)
			continue
		}
	}
}

func TestEvaluateArithmeticExponentiation(t *testing.T) {
	i := expr.NewInt(4)
	j := expr.NewInt(5)
	k := expr.NewInt(2)
	x := expr.NewReal(4.0)
	y := expr.NewReal(5.0)
	z := expr.NewReal(0.5)
	q := expr.NewReal(-1)
	r := expr.NewReal(-2)

	testCases := []struct {
		exp    expr.Expr
		result expr.Expr
	}{
		{expr.NewPow(i, j), expr.NewInt(1024)},
		{expr.NewPow(i, y), expr.NewReal(1024.0)},
		{expr.NewPow(x, j), expr.NewReal(1024.0)},
		{expr.NewPow(x, y), expr.NewReal(1024.0)},
		{expr.NewPow(i, z), expr.NewReal(2.0)},
		{expr.NewPow(x, z), expr.NewReal(2.0)},
		{expr.NewPow(k, q), expr.NewReal(0.5)},
		{expr.NewPow(k, r), expr.NewReal(0.25)},
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

func TestEvaluateArithmeticExponentiationDomainError(t *testing.T) {
	i := expr.NewInt(-3)
	x := expr.NewReal(-3.0)
	testCases := []struct {
		exp expr.Expr
	}{
		{expr.NewPow(i, expr.NewReal(0.5))},
		{expr.NewPow(i, expr.NewReal(-0.5))},
		{expr.NewPow(x, expr.NewReal(0.5))},
		{expr.NewPow(x, expr.NewReal(-0.5))},
	}

	for n, testCase := range testCases {
		_, err := testCase.exp.Evaluate(nil)
		if err != expr.DomainError {
			t.Errorf("case %d, got %v, want %v", n+1, err,
				expr.DomainError)
			continue
		}
	}
}

func TestEvaluateArithmeticNegation(t *testing.T) {
	i := expr.NewInt(4)
	j := expr.NewInt(-5)
	x := expr.NewReal(4.0)
	y := expr.NewReal(-5.0)

	testCases := []struct {
		exp    expr.Expr
		result expr.Expr
	}{
		{expr.NewNeg(i), expr.NewInt(-4)},
		{expr.NewNeg(j), expr.NewInt(5)},
		{expr.NewNeg(x), expr.NewReal(-4.0)},
		{expr.NewNeg(y), expr.NewReal(5.0)},
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

func TestEvaluateArithmeticCompound(t *testing.T) {
	i := expr.NewInt(2)
	j := expr.NewInt(3)
	k := expr.NewInt(4)
	x := expr.NewReal(2.0)
	y := expr.NewReal(3.0)
	z := expr.NewReal(4.0)

	testCases := []struct {
		exp    expr.Expr
		result expr.Expr
	}{
		{expr.NewAdd(i, expr.NewMul(j, k)), expr.NewInt(14)},
		{expr.NewAdd(x, expr.NewMul(j, k)), expr.NewReal(14.0)},
		{expr.NewAdd(i, expr.NewMul(y, k)), expr.NewReal(14.0)},
		{expr.NewAdd(i, expr.NewMul(j, z)), expr.NewReal(14.0)},
		{expr.NewAdd(x, expr.NewMul(y, z)), expr.NewReal(14.0)},
		{expr.NewMul(expr.NewAdd(i, j), k), expr.NewInt(20)},
		{expr.NewMul(expr.NewAdd(x, j), k), expr.NewReal(20.0)},
		{expr.NewMul(expr.NewAdd(i, y), k), expr.NewReal(20.0)},
		{expr.NewMul(expr.NewAdd(i, j), z), expr.NewReal(20.0)},
		{expr.NewMul(expr.NewAdd(x, y), z), expr.NewReal(20.0)},
		{expr.NewSub(j, expr.NewDiv(k, i)), expr.NewInt(1)},
		{expr.NewSub(j, expr.NewDiv(k, x)), expr.NewReal(1.0)},
		{expr.NewSub(y, expr.NewDiv(k, i)), expr.NewReal(1.0)},
		{expr.NewSub(j, expr.NewDiv(z, i)), expr.NewReal(1.0)},
		{expr.NewSub(y, expr.NewDiv(z, x)), expr.NewReal(1.0)},
		{expr.NewDiv(expr.NewSub(k, expr.NewNeg(i)), j), expr.NewInt(2)},
		{expr.NewDiv(expr.NewSub(k, expr.NewNeg(x)), j), expr.NewReal(2.0)},
		{expr.NewDiv(expr.NewSub(k, expr.NewNeg(i)), y), expr.NewReal(2.0)},
		{expr.NewDiv(expr.NewSub(z, expr.NewNeg(i)), j), expr.NewReal(2.0)},
		{expr.NewDiv(expr.NewSub(z, expr.NewNeg(x)), y), expr.NewReal(2.0)},
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
