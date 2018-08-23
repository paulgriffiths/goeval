package vareval

import (
	"github.com/paulgriffiths/goeval/expr"
	"math"
	"testing"
)

func almostEqual(a, b, epsilon float64) bool {
	if math.Abs(a-b) <= epsilon {
		return true
	}
	return false
}

func TestGoodExpressions(t *testing.T) {
	cases := []struct {
		expr     string
		expected expr.Expr
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
		{"-cos(90)", expr.NewReal(0.0), true},
		{"-acos(1)", expr.NewReal(0.0), true},
		{"-acos(0.5)", expr.NewReal(-60.0), true},
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
		{"true", expr.NewBool(true), false},
		{"false", expr.NewBool(false), false},
		{"true or false", expr.NewBool(true), false},
		{"true or true", expr.NewBool(true), false},
		{"false or true", expr.NewBool(true), false},
		{"false or false", expr.NewBool(false), false},
		{"false or false or false", expr.NewBool(false), false},
		{"false or false or true", expr.NewBool(true), false},
		{"false or true or false", expr.NewBool(true), false},
		{"true or false or false", expr.NewBool(true), false},
		{"true and false", expr.NewBool(false), false},
		{"true and true", expr.NewBool(true), false},
		{"false and true", expr.NewBool(false), false},
		{"false and false", expr.NewBool(false), false},
		{"false and false and false", expr.NewBool(false), false},
		{"false and false and true", expr.NewBool(false), false},
		{"false and true and false", expr.NewBool(false), false},
		{"true and false and false", expr.NewBool(false), false},
		{"true and true and true", expr.NewBool(true), false},
		{"true == true", expr.NewBool(true), false},
		{"true == false", expr.NewBool(false), false},
		{"true == true or false", expr.NewBool(true), false},
		{"true == true and false", expr.NewBool(false), false},
		{"true or false == true and false", expr.NewBool(false), false},
		{"true or (false == true) and false", expr.NewBool(true), false},
		{"true != true", expr.NewBool(false), false},
		{"true != false", expr.NewBool(true), false},
		{"true != false == false != false", expr.NewBool(false), false},
		{"1 == 1", expr.NewBool(true), false},
		{"1 == 2", expr.NewBool(false), false},
		{"1 == -1", expr.NewBool(false), false},
		{"1 == -2", expr.NewBool(false), false},
		{"1 == 1.0", expr.NewBool(true), false},
		{"1 == 1.1", expr.NewBool(false), false},
		{"1 != 1", expr.NewBool(false), false},
		{"1 != 2", expr.NewBool(true), false},
		{"1 != -1", expr.NewBool(true), false},
		{"1 != -2", expr.NewBool(true), false},
		{"1 != 1.0", expr.NewBool(false), false},
		{"1 != 1.1", expr.NewBool(true), false},
		{"2 == 1", expr.NewBool(false), false},
		{"2 == 2", expr.NewBool(true), false},
		{"2 == 3", expr.NewBool(false), false},
		{"2 != 1", expr.NewBool(true), false},
		{"2 != 2", expr.NewBool(false), false},
		{"2 != 3", expr.NewBool(true), false},
		{"2 < 1", expr.NewBool(false), false},
		{"2 < 2", expr.NewBool(false), false},
		{"2 < 3", expr.NewBool(true), false},
		{"2 <= 1", expr.NewBool(false), false},
		{"2 <= 2", expr.NewBool(true), false},
		{"2 <= 3", expr.NewBool(true), false},
		{"2 > 1", expr.NewBool(true), false},
		{"2 > 2", expr.NewBool(false), false},
		{"2 > 3", expr.NewBool(false), false},
		{"2 >= 1", expr.NewBool(true), false},
		{"2 >= 2", expr.NewBool(true), false},
		{"2 >= 3", expr.NewBool(false), false},
		{"2 < 3 == 5 > 4", expr.NewBool(true), false},
		{"2 < 3 != 5 < 4", expr.NewBool(true), false},
	}

	for _, c := range cases {
		result, err := evaluate(c.expr, nil)
		if err != nil {
			t.Errorf("Expr '%s': got error %v, want %v", c.expr, err, nil)
			continue
		}
		if !expr.Equals(result, c.expected, nil) {
			rv, okr := expr.ToFloat(result)
			cv, okc := expr.ToFloat(c.expected)
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
		{"foobar(4)", TrailingTokensError},
		{"-1^0.5", expr.DomainError},
		{"sqrt(-1)", expr.DomainError},
		{"log(-1000)", expr.DomainError},
		{"1+2(3+4)", TrailingTokensError},
		{"true or 1", expr.TypeError},
		{"1 or 2", expr.TypeError},
		{"true and 1", expr.TypeError},
		{"1 and 2", expr.TypeError},
		{"1 == true", expr.TypeError},
		{"false == 2", expr.TypeError},
		{"1 != true", expr.TypeError},
		{"false != 2", expr.TypeError},
		{"1 <= true", expr.TypeError},
		{"false <= 2", expr.TypeError},
		{"1 >= true", expr.TypeError},
		{"false >= 2", expr.TypeError},
		{"1 < true", expr.TypeError},
		{"false < 2", expr.TypeError},
		{"1 > true", expr.TypeError},
		{"false > 2", expr.TypeError},
		{"1 < 2 < 3", expr.TypeError},
		{"1 > 2 > 3", expr.TypeError},
		{"1 <= 2 <= 3", expr.TypeError},
		{"1 >= 2 >= 3", expr.TypeError},
		{"frob + 3", expr.UndefinedVariableError},
		{"cos(foobar)", expr.UndefinedVariableError},
	}

	for _, c := range cases {
		_, err := evaluate(c.expr, nil)
		if err != c.err {
			t.Errorf("Expr '%s': got %v, want %v", c.expr, err, c.err)
		}
	}
}

func TestGoodExpressionsWithVariables(t *testing.T) {
	cases := []struct {
		expr     string
		expected expr.Expr
		almost   bool
	}{
		{"tom", expr.NewInt(5), false},
		{"dick", expr.NewReal(10.5), false},
		{"harry", expr.NewBool(true), false},
		{"tom * 5", expr.NewInt(25), false},
		{"(tom * 5)", expr.NewInt(25), false},
		{"tom + tom", expr.NewInt(10), false},
		{"tom + dick", expr.NewReal(15.5), false},
		{"harry == false", expr.NewBool(false), false},
		{"round(dick) + 3", expr.NewReal(14.0), false},
		{"ceil(dick) + 3", expr.NewReal(14.0), false},
		{"floor(dick) + 3", expr.NewReal(13.0), false},
	}

	for _, c := range cases {
		table := expr.NewTable()
		table.Store("tom", expr.NewInt(5))
		table.Store("dick", expr.NewReal(10.5))
		table.Store("harry", expr.NewBool(true))

		result, err := evaluate(c.expr, table)
		if err != nil {
			t.Errorf("Expr '%s': got error %v, want %v", c.expr, err, nil)
			continue
		}
		if !expr.Equals(result, c.expected, table) {
			rv, okr := expr.ToFloat(result)
			cv, okc := expr.ToFloat(c.expected)
			if !okr || !okc || !c.almost || !almostEqual(rv, cv, 0.000001) {
				t.Errorf("Expr '%s': got %v, want %v", c.expr, rv, cv)
			}
		}
	}
}
