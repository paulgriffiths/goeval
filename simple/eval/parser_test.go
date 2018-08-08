package eval

import (
    "testing"
    "math"
)

func almostEqual(a, b, epsilon float64) bool {
    if math.Abs(a-b) <= epsilon {
        return true
    } else {
        return false
    }
}

func TestGoodExpressions(t *testing.T) {
    cases := []struct{
        expr string
        expected float64
    }{
        {"0", 0},
        {"7", 7},
        {"-7", -7},
        {"(1)", 1},
        {"(-1)", -1},
        {"-(1)", -1},
        {"((1))", 1},
        {"((-1))", -1},
        {"(-(1))", -1},
        {"-((1))", -1},
        {"1+2", 3},
        {"1+2+3", 6},
        {"1+2+3+4", 10},
        {"1-2", -1},
        {"1-2-3", -4},
        {"1-2-3-4", -8},
        {"-1+6", 5},
        {"-1+-6", -7},
        {"-1+-6+-8", -15},
        {"-1-6", -7},
        {"-1--6", 5},
        {"-1--6--8", 13},
        {"1+2-3", 0},
        {"1-2+3", 2},
        {"1*2", 2},
        {"1*2*3", 6},
        {"1*2*3*4", 24},
        {"1*2*-3*4", -24},
        {"60/2", 30},
        {"60/2/3", 10},
        {"60/2/-3/5", -2},
        {"4*6/2", 12},
        {"(4*6)/2", 12},
        {"4*(6/2)", 12},
        {"6/3*3", 6},
        {"(6/3)*3", 6},
        {"8/(4*4)", 0.5},
        {"1+2*3", 7},
        {"1+(2*3)", 7},
        {"(1+2)*3", 9},
        {"1+-(2*3)", -5},
        {"-(1+2)*3", -9},
        {"1+2*3+4", 11},
        {"1+6/2+4", 8},
        {"1+(2*3)+4", 11},
        {"(1+2)*(3+4)", 21},
        {"((1+2)*(3+4))", 21},
        {"((-1+2)*(3+-4))", -1},
        {"((-1+-2)*(-3+-4))", 21},
        {"2^3", 8},
        {"4^3^2", 4096},
        {"(4^3)^2", 4096},
        {"4^(3^2)", 262144},
        {"-3^2", 9},
        {"-(3^2)", -9},
        {"2^-1", 0.5},
        {"2^-2", 0.25},
        {"-2^-2", 0.25},
        {"-(2^-2)", -0.25},
        {"4*3^2", 36},
        {"(4*3)^2", 144},
        {"4*(3^2)", 36},
        {"4+3^2", 13},
        {"(4+3)^2", 49},
        {"4+(3^2)", 13},
        {"4-(3^2)", -5},
        {"pi", 3.14159265358979323846},
        {"e", 2.71828182845904523536},
        {"-pi", -3.14159265358979323846},
        {"-e", -2.71828182845904523536},
        {"cos(0)", 1},
        {"cos(60)", 0.5},
        {"cos(90)", 0},
        {"acos(1)", 0},
        {"acos(0.5)", 60},
        {"acos(0)", 90},
        {"-cos(0)", -1},
        {"-cos(60)", -0.5},
        {"-cos(90)", 0},
        {"-acos(1)", 0},
        {"-acos(0.5)", -60},
        {"-acos(0)", -90},
        {"sin(0)", 0},
        {"sin(30)", 0.5},
        {"sin(90)", 1},
        {"asin(0)", 0},
        {"asin(0.5)", 30},
        {"asin(1)", 90},
        {"tan(0)", 0},
        {"tan(45)", 1},
        {"atan(0)", 0},
        {"atan(1)", 45},
        {"ceil(6.1)", 7},
        {"ceil(6.9)", 7},
        {"floor(6.1)", 6},
        {"floor(6.9)", 6},
        {"round(6.1)", 6},
        {"round(6.9)", 7},
        {"sqrt(16)", 4},
        {"sqrt(144)", 12},
        {"log(100)", 2},
        {"log(1000)", 3},
        {"ln(e^3)", 3},
        {"ln(e^4)", 4},
    }

    for _, c := range cases {
        result, err := Evaluate(c.expr)
        if err != nil {
            t.Errorf("Expr '%s': got error %v, want %v", c.expr, err, nil)
        } else if err == nil && !almostEqual(result, c.expected, 0.00000001) {
            t.Errorf("Expr '%s': got %v, want %v", c.expr, result, c.expected)
        }
    }
}

func TestBadExpressions(t *testing.T) {
    cases := []struct{
        expr string
        err error
    }{
        {"", MissingFactorError},
        {"()", UnbalancedParenthesesError},     // Should be missing factor?
        {"(", MissingFactorError},
        {"(1", UnbalancedParenthesesError},
        {")", UnbalancedParenthesesError},
        {"4/0", DivideByZeroError},
        {"-", MissingFactorError},
        {"+", MissingFactorError},
        {"--", MissingFactorError},
        {"--2", MissingFactorError},
        {"1+", MissingFactorError},
        {"1*", MissingFactorError},
        {"1^", MissingFactorError},
        {"cos", MissingArgumentError},
        {"cos(", MissingFactorError},
        {"cos()", UnbalancedParenthesesError},  // Should be missing factor?
        {"foobar(4)", UnknownFunctionError},
        {"-1^0.5", DomainError},
        {"sqrt(-1)", DomainError},
        {"log(-1000)", DomainError},
        {"1+2(3+4)", TrailingTokensError},
    }

    for _, c := range cases {
        _, err := Evaluate(c.expr)
        if err != c.err {
            t.Errorf("Expr '%s': got %v, want %v", c.expr, err, c.err)
        }
    }
}
