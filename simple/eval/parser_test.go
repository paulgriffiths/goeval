package eval

import (
	"github.com/paulgriffiths/goeval/expr"
	"math"
	"testing"
)

func almostEqual(a, b, epsilon float64) bool {
	if math.Abs(a-b) <= epsilon {
		return true
	} else {
		return false
	}
}

func TestGoodExpressions(t *testing.T) {
	cases := []struct {
		expr     string
		expected expr.Value
		almost   bool
	}{
		{"0", expr.NewInt(0), false},
		{"7", expr.NewInt(7), false},
		{"-7", expr.NewInt(-7), false},
		{"(1)", expr.NewInt(1), false},
		{"(-1)", expr.NewInt(-1), false},
		{"-(1)", expr.NewInt(-1), false},
		{"((1))", expr.NewInt(1), false},
		{"((-1))", expr.NewInt(-1), false},
		{"(-(1))", expr.NewInt(-1), false},
		{"-((1))", expr.NewInt(-1), false},
		{"1+2", expr.NewInt(3), false},
		{"1+2+3", expr.NewInt(6), false},
		{"1+2+3+4", expr.NewInt(10), false},
		{"1-2", expr.NewInt(-1), false},
		{"1-2-3", expr.NewInt(-4), false},
		{"1-2-3-4", expr.NewInt(-8), false},
		{"-1+6", expr.NewInt(5), false},
		{"-1+-6", expr.NewInt(-7), false},
		{"-1+-6+-8", expr.NewInt(-15), false},
		{"-1-6", expr.NewInt(-7), false},
		{"-1--6", expr.NewInt(5), false},
		{"-1--6--8", expr.NewInt(13), false},
		{"1+2-3", expr.NewInt(0), false},
		{"1-2+3", expr.NewInt(2), false},
		{"1*2", expr.NewInt(2), false},
		{"1*2*3", expr.NewInt(6), false},
		{"1*2*3*4", expr.NewInt(24), false},
		{"1*2*-3*4", expr.NewInt(-24), false},
		{"60/2", expr.NewInt(30), false},
		{"60/2/3", expr.NewInt(10), false},
		{"60/2/-3/5", expr.NewInt(-2), false},
		{"4*6/2", expr.NewInt(12), false},
		{"(4*6)/2", expr.NewInt(12), false},
		{"4*(6/2)", expr.NewInt(12), false},
		{"6/3*3", expr.NewInt(6), false},
		{"(6/3)*3", expr.NewInt(6), false},
		{"8/(4*4)", expr.NewReal(0.5), false},
		{"1+2*3", expr.NewInt(7), false},
		{"1+(2*3)", expr.NewInt(7), false},
		{"(1+2)*3", expr.NewInt(9), false},
		{"1+-(2*3)", expr.NewInt(-5), false},
		{"-(1+2)*3", expr.NewInt(-9), false},
		{"1+2*3+4", expr.NewInt(11), false},
		{"1+6/2+4", expr.NewInt(8), false},
		{"1+(2*3)+4", expr.NewInt(11), false},
		{"(1+2)*(3+4)", expr.NewInt(21), false},
		{"((1+2)*(3+4))", expr.NewInt(21), false},
		{"((-1+2)*(3+-4))", expr.NewInt(-1), false},
		{"((-1+-2)*(-3+-4))", expr.NewInt(21), false},
		{"2^3", expr.NewInt(8), false},
		{"4^3^2", expr.NewInt(4096), false},
		{"(4^3)^2", expr.NewInt(4096), false},
		{"4^(3^2)", expr.NewInt(262144), false},
		{"-3^2", expr.NewInt(9), false},
		{"-(3^2)", expr.NewInt(-9), false},
		{"2^-1", expr.NewReal(0.5), false},
		{"2^-2", expr.NewReal(0.25), false},
		{"-2^-2", expr.NewReal(0.25), false},
		{"-(2^-2)", expr.NewReal(-0.25), false},
		{"4*3^2", expr.NewInt(36), false},
		{"(4*3)^2", expr.NewInt(144), false},
		{"4*(3^2)", expr.NewInt(36), false},
		{"4+3^2", expr.NewInt(13), false},
		{"(4+3)^2", expr.NewInt(49), false},
		{"4+(3^2)", expr.NewInt(13), false},
		{"4-(3^2)", expr.NewInt(-5), false},
		{"pi", expr.NewReal(3.14159265358979323846), true},
		{"e", expr.NewReal(2.71828182845904523536), true},
		{"-pi", expr.NewReal(-3.14159265358979323846), true},
		{"-e", expr.NewReal(-2.71828182845904523536), true},
		{"cos(0)", expr.NewReal(1), true},
		{"cos(60)", expr.NewReal(0.5), true},
		{"cos(90)", expr.NewReal(0), true},
		{"acos(1)", expr.NewReal(0), true},
		{"acos(0.5)", expr.NewReal(60), true},
		{"acos(0)", expr.NewReal(90), true},
		{"-cos(0)", expr.NewReal(-1), true},
		{"-cos(60)", expr.NewReal(-0.5), true},
		{"-cos(90)", expr.NewReal(0), true},
		{"-acos(1)", expr.NewReal(0), true},
		{"-acos(0.5)", expr.NewReal(-60), true},
		{"-acos(0)", expr.NewReal(-90), true},
		{"sin(0)", expr.NewReal(0), true},
		{"sin(30)", expr.NewReal(0.5), true},
		{"sin(90)", expr.NewReal(1), true},
		{"asin(0)", expr.NewReal(0), true},
		{"asin(0.5)", expr.NewReal(30), true},
		{"asin(1)", expr.NewReal(90), true},
		{"tan(0)", expr.NewReal(0), true},
		{"tan(45)", expr.NewReal(1), true},
		{"atan(0)", expr.NewReal(0), true},
		{"atan(1)", expr.NewReal(45), true},
		{"ceil(6.1)", expr.NewReal(7), true},
		{"ceil(6.9)", expr.NewReal(7), true},
		{"floor(6.1)", expr.NewReal(6), true},
		{"floor(6.9)", expr.NewReal(6), true},
		{"round(6.1)", expr.NewReal(6), true},
		{"round(6.9)", expr.NewReal(7), true},
		{"sqrt(16)", expr.NewReal(4), true},
		{"sqrt(144)", expr.NewReal(12), true},
		{"log(100)", expr.NewReal(2), true},
		{"log(1000)", expr.NewReal(3), true},
		{"ln(e^3)", expr.NewReal(3), true},
		{"ln(e^4)", expr.NewReal(4), true},
	}

	for _, c := range cases {
		result, err := Evaluate(c.expr)
		if err != nil {
			t.Errorf("Expr '%s': got error %v, want %v", c.expr, err, nil)
			continue
		}
		if !result.Equals(c.expected) {
			rv, okr := expr.FloatValueIfPossible(result)
			cv, okc := expr.FloatValueIfPossible(c.expected)
			if !okr || !okc || !c.almost || !almostEqual(rv, cv, 0.000001) {
				t.Errorf("Expr '%s': got %v, want %v", c.expr, rv, cv)
			}
		}
	}
}

func TestBadExpressions(t *testing.T) {
	cases := []struct {
		expr string
		err  error
	}{
		{"", MissingFactorError},
		{"()", UnbalancedParenthesesError}, // Should be missing factor?
		{"(", MissingFactorError},
		{"(1", UnbalancedParenthesesError},
		{")", UnbalancedParenthesesError},
		{"4/0", expr.DivideByZeroError},
		{"-", MissingFactorError},
		{"+", MissingFactorError},
		{"--", MissingFactorError},
		{"--2", MissingFactorError},
		{"1+", MissingFactorError},
		{"1*", MissingFactorError},
		{"1^", MissingFactorError},
		{"cos", MissingArgumentError},
		{"cos(", MissingFactorError},
		{"cos()", UnbalancedParenthesesError}, // Should be missing factor?
		{"foobar(4)", UnknownFunctionError},
		{"-1^0.5", expr.DomainError},
		{"sqrt(-1)", expr.DomainError},
		{"log(-1000)", expr.DomainError},
		{"1+2(3+4)", TrailingTokensError},
	}

	for _, c := range cases {
		_, err := Evaluate(c.expr)
		if err != c.err {
			t.Errorf("Expr '%s': got %v, want %v", c.expr, err, c.err)
		}
	}
}
