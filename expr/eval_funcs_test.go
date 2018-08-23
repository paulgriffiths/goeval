package expr_test

import (
	"github.com/paulgriffiths/goeval/expr"
	"math"
	"testing"
)

func TestEvaluateFunctionCos(t *testing.T) {
	testCases := []struct {
		exp    expr.Expr
		result expr.Expr
	}{
		{expr.NewCos(expr.NewReal(0.0)), expr.NewReal(1.0)},
		{expr.NewCos(expr.NewReal(60.0)), expr.NewReal(0.5)},
		{expr.NewCos(expr.NewReal(90.0)), expr.NewReal(0.0)},
	}

	for n, testCase := range testCases {
		result, err := testCase.exp.Evaluate(nil)
		if err != nil {
			t.Errorf("case %d, couldn't evaluate expression: %v", n+1, err)
			continue
		}
		if !expr.AlmostEquals(result, testCase.result, 0.000001, nil) {
			t.Errorf("case %d, got %v, want %v", n+1,
				result, testCase.result)
		}
	}
}

func TestEvaluateFunctionSin(t *testing.T) {
	testCases := []struct {
		exp    expr.Expr
		result expr.Expr
	}{
		{expr.NewSin(expr.NewReal(0.0)), expr.NewReal(0.0)},
		{expr.NewSin(expr.NewReal(30.0)), expr.NewReal(0.5)},
		{expr.NewSin(expr.NewReal(90.0)), expr.NewReal(1.0)},
	}

	for n, testCase := range testCases {
		result, err := testCase.exp.Evaluate(nil)
		if err != nil {
			t.Errorf("case %d, couldn't evaluate expression: %v", n+1, err)
			continue
		}
		if !expr.AlmostEquals(result, testCase.result, 0.000001, nil) {
			t.Errorf("case %d, got %v, want %v", n+1,
				result, testCase.result)
		}
	}
}

func TestEvaluateFunctionTan(t *testing.T) {
	testCases := []struct {
		exp    expr.Expr
		result expr.Expr
	}{
		{expr.NewTan(expr.NewReal(0.0)), expr.NewReal(0.0)},
		{expr.NewTan(expr.NewReal(45.0)), expr.NewReal(1.0)},
	}

	for n, testCase := range testCases {
		result, err := testCase.exp.Evaluate(nil)
		if err != nil {
			t.Errorf("case %d, couldn't evaluate expression: %v", n+1, err)
			continue
		}
		if !expr.AlmostEquals(result, testCase.result, 0.000001, nil) {
			t.Errorf("case %d, got %v, want %v", n+1,
				result, testCase.result)
		}
	}
}

func TestEvaluateFunctionArccos(t *testing.T) {
	testCases := []struct {
		exp    expr.Expr
		result expr.Expr
	}{
		{expr.NewAcos(expr.NewReal(1.0)), expr.NewReal(0.0)},
		{expr.NewAcos(expr.NewReal(0.5)), expr.NewReal(60.0)},
		{expr.NewAcos(expr.NewReal(0.0)), expr.NewReal(90.0)},
	}

	for n, testCase := range testCases {
		result, err := testCase.exp.Evaluate(nil)
		if err != nil {
			t.Errorf("case %d, couldn't evaluate expression: %v", n+1, err)
			continue
		}
		if !expr.AlmostEquals(result, testCase.result, 0.000001, nil) {
			t.Errorf("case %d, got %v, want %v", n+1,
				result, testCase.result)
		}
	}
}

func TestEvaluateFunctionArccosBadDomain(t *testing.T) {
	testCases := []struct {
		exp expr.Expr
	}{
		{expr.NewAcos(expr.NewReal(-1.1))},
		{expr.NewAcos(expr.NewReal(1.1))},
	}

	for n, testCase := range testCases {
		_, err := testCase.exp.Evaluate(nil)
		if err != expr.DomainError {
			t.Errorf("case %d, got %v, want %v", n+1, err,
				expr.DomainError)
		}
	}
}

func TestEvaluateFunctionArcsin(t *testing.T) {
	testCases := []struct {
		exp    expr.Expr
		result expr.Expr
	}{
		{expr.NewAsin(expr.NewReal(0.0)), expr.NewReal(0.0)},
		{expr.NewAsin(expr.NewReal(0.5)), expr.NewReal(30.)},
		{expr.NewAsin(expr.NewReal(1.0)), expr.NewReal(90.0)},
	}

	for n, testCase := range testCases {
		result, err := testCase.exp.Evaluate(nil)
		if err != nil {
			t.Errorf("case %d, couldn't evaluate expression: %v", n+1, err)
			continue
		}
		if !expr.AlmostEquals(result, testCase.result, 0.000001, nil) {
			t.Errorf("case %d, got %v, want %v", n+1,
				result, testCase.result)
		}
	}
}

func TestEvaluateFunctionArcsinBadDomain(t *testing.T) {
	testCases := []struct {
		exp expr.Expr
	}{
		{expr.NewAsin(expr.NewReal(-1.1))},
		{expr.NewAsin(expr.NewReal(1.1))},
	}

	for n, testCase := range testCases {
		_, err := testCase.exp.Evaluate(nil)
		if err != expr.DomainError {
			t.Errorf("case %d, got %v, want %v", n+1, err,
				expr.DomainError)
		}
	}
}

func TestEvaluateFunctionArctan(t *testing.T) {
	testCases := []struct {
		exp    expr.Expr
		result expr.Expr
	}{
		{expr.NewAtan(expr.NewReal(0.0)), expr.NewReal(0.0)},
		{expr.NewAtan(expr.NewReal(1.0)), expr.NewReal(45.0)},
	}

	for n, testCase := range testCases {
		result, err := testCase.exp.Evaluate(nil)
		if err != nil {
			t.Errorf("case %d, couldn't evaluate expression: %v", n+1, err)
			continue
		}
		if !expr.AlmostEquals(result, testCase.result, 0.000001, nil) {
			t.Errorf("case %d, got %v, want %v", n+1,
				result, testCase.result)
		}
	}
}

func TestEvaluateFunctionCeiling(t *testing.T) {
	testCases := []struct {
		exp    expr.Expr
		result expr.Expr
	}{
		{expr.NewCeil(expr.NewReal(5.1)), expr.NewReal(6.0)},
		{expr.NewCeil(expr.NewReal(5.9)), expr.NewReal(6.0)},
		{expr.NewCeil(expr.NewReal(-5.1)), expr.NewReal(-5.0)},
		{expr.NewCeil(expr.NewReal(-5.9)), expr.NewReal(-5.0)},
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

func TestEvaluateFunctionFloor(t *testing.T) {
	testCases := []struct {
		exp    expr.Expr
		result expr.Expr
	}{
		{expr.NewFloor(expr.NewReal(5.1)), expr.NewReal(5.0)},
		{expr.NewFloor(expr.NewReal(5.9)), expr.NewReal(5.0)},
		{expr.NewFloor(expr.NewReal(-5.1)), expr.NewReal(-6.0)},
		{expr.NewFloor(expr.NewReal(-5.9)), expr.NewReal(-6.0)},
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

func TestEvaluateFunctionRound(t *testing.T) {
	testCases := []struct {
		exp    expr.Expr
		result expr.Expr
	}{
		{expr.NewRound(expr.NewReal(5.1)), expr.NewReal(5.0)},
		{expr.NewRound(expr.NewReal(5.9)), expr.NewReal(6.0)},
		{expr.NewRound(expr.NewReal(-5.1)), expr.NewReal(-5.0)},
		{expr.NewRound(expr.NewReal(-5.9)), expr.NewReal(-6.0)},
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

func TestEvaluateFunctionLog(t *testing.T) {
	testCases := []struct {
		exp    expr.Expr
		result expr.Expr
	}{
		{expr.NewLog(expr.NewReal(0.001)), expr.NewReal(-3.0)},
		{expr.NewLog(expr.NewReal(0.01)), expr.NewReal(-2.0)},
		{expr.NewLog(expr.NewReal(0.1)), expr.NewReal(-1.0)},
		{expr.NewLog(expr.NewReal(1.0)), expr.NewReal(0.0)},
		{expr.NewLog(expr.NewReal(10.0)), expr.NewReal(1.0)},
		{expr.NewLog(expr.NewReal(100.0)), expr.NewReal(2.0)},
		{expr.NewLog(expr.NewReal(1000.0)), expr.NewReal(3.0)},
	}

	for n, testCase := range testCases {
		result, err := testCase.exp.Evaluate(nil)
		if err != nil {
			t.Errorf("case %d, couldn't evaluate expression: %v", n+1, err)
			continue
		}
		if !expr.AlmostEquals(result, testCase.result, 0.000001, nil) {
			t.Errorf("case %d, got %v, want %v", n+1,
				result, testCase.result)
		}
	}
}

func TestEvaluateFunctionLogBadDomain(t *testing.T) {
	testCases := []struct {
		exp expr.Expr
	}{
		{expr.NewLog(expr.NewReal(-1.0))},
		{expr.NewLog(expr.NewReal(-10.0))},
		{expr.NewLog(expr.NewReal(-100.0))},
	}

	for n, testCase := range testCases {
		_, err := testCase.exp.Evaluate(nil)
		if err != expr.DomainError {
			t.Errorf("case %d, got %v, want %v", n+1, err,
				expr.DomainError)
		}
	}
}

func TestEvaluateFunctionLn(t *testing.T) {
	testCases := []struct {
		exp    expr.Expr
		result expr.Expr
	}{
		{expr.NewLn(expr.NewReal(math.Pow(math.E, -3.0))),
			expr.NewReal(-3.0)},
		{expr.NewLn(expr.NewReal(math.Pow(math.E, -2.0))),
			expr.NewReal(-2.0)},
		{expr.NewLn(expr.NewReal(math.Pow(math.E, -1.0))),
			expr.NewReal(-1.0)},
		{expr.NewLn(expr.NewReal(math.Pow(math.E, 0.0))),
			expr.NewReal(0.0)},
		{expr.NewLn(expr.NewReal(math.Pow(math.E, 1.0))),
			expr.NewReal(1.0)},
		{expr.NewLn(expr.NewReal(math.Pow(math.E, 2.0))),
			expr.NewReal(2.0)},
		{expr.NewLn(expr.NewReal(math.Pow(math.E, 3.0))),
			expr.NewReal(3.0)},
	}

	for n, testCase := range testCases {
		result, err := testCase.exp.Evaluate(nil)
		if err != nil {
			t.Errorf("case %d, couldn't evaluate expression: %v", n+1, err)
			continue
		}
		if !expr.AlmostEquals(result, testCase.result, 0.000001, nil) {
			t.Errorf("case %d, got %v, want %v", n+1,
				result, testCase.result)
		}
	}
}

func TestEvaluateFunctionLnBadDomain(t *testing.T) {
	testCases := []struct {
		exp expr.Expr
	}{
		{expr.NewLn(expr.NewReal(-1.0))},
		{expr.NewLn(expr.NewReal(-10.0))},
		{expr.NewLn(expr.NewReal(-100.0))},
	}

	for n, testCase := range testCases {
		_, err := testCase.exp.Evaluate(nil)
		if err != expr.DomainError {
			t.Errorf("case %d, got %v, want %v", n+1, err,
				expr.DomainError)
		}
	}
}

func TestEvaluateFunctionSquareRoot(t *testing.T) {
	testCases := []struct {
		exp    expr.Expr
		result expr.Expr
	}{
		{expr.NewSqrt(expr.NewReal(0.25)), expr.NewReal(0.5)},
		{expr.NewSqrt(expr.NewReal(16.0)), expr.NewReal(4.0)},
		{expr.NewSqrt(expr.NewReal(144.0)), expr.NewReal(12.0)},
	}

	for n, testCase := range testCases {
		result, err := testCase.exp.Evaluate(nil)
		if err != nil {
			t.Errorf("case %d, couldn't evaluate expression: %v", n+1, err)
			continue
		}
		if !expr.AlmostEquals(result, testCase.result, 0.000001, nil) {
			t.Errorf("case %d, got %v, want %v", n+1,
				result, testCase.result)
		}
	}
}

func TestEvaluateFunctionSquareRootBadDomain(t *testing.T) {
	testCases := []struct {
		exp expr.Expr
	}{
		{expr.NewSqrt(expr.NewReal(-1.0))},
	}

	for n, testCase := range testCases {
		_, err := testCase.exp.Evaluate(nil)
		if err != expr.DomainError {
			t.Errorf("case %d, got %v, want %v", n+1, err,
				expr.DomainError)
		}
	}
}
