package eval

import (
    "testing"
)

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
    }

    for _, c := range cases {
        result, err := Evaluate(c.expr)
        if err != nil {
            t.Errorf("Expr '%s': got error %v, want %v", c.expr, err, nil)
        } else if err == nil && result != c.expected {
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
        {"()", UnbalancedParenthesesError},
        {"(", MissingFactorError},
        {"(1", UnbalancedParenthesesError},
        {")", UnbalancedParenthesesError},
        {"4/0", DivideByZeroError},
        {"-", MissingFactorError},
        {"+", MissingFactorError},
        {"--", MissingFactorError},
        {"--2", MissingFactorError},
        //{"1+2(3+4)", Trailing information error?},
    }

    for _, c := range cases {
        _, err := Evaluate(c.expr)
        if err != c.err {
            t.Errorf("Expr '%s': got %v, want %v", c.expr, err, c.err)
        }
    }
}
