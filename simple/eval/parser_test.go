package eval

import (
    "testing"
)

func TestVariousExpressions(t *testing.T) {
    cases := []struct{
        expr string
        want float64
        err error
    }{
        // {"", 99, Empty expression error?},
        // {"(", 99, UnbalancedParenthesesError},
        // {"()", 99, Empty expression error?},
        // {"-1", -1, nil},
        // {"-1*-3", -3, nil},
        {"(1", 99, UnbalancedParenthesesError},
        {")", 99, UnbalancedParenthesesError},
        {"0", 0, nil},
        {"1", 1, nil},
        {"(1)", 1, nil},
        {"((1))", 1, nil},
        {"1+2", 3, nil},
        {"1+2+3", 6, nil},
        {"1+2+3+4", 10, nil},
        {"1-2", -1, nil},
        {"1-2-3", -4, nil},
        {"1-2-3-4", -8, nil},
        {"1+2-3", 0, nil},
        {"1-2+3", 2, nil},
        {"1*2", 2, nil},
        {"1*2*3", 6, nil},
        {"1*2*3*4", 24, nil},
        {"60/2", 30, nil},
        {"60/2/3", 10, nil},
        {"60/2/3/5", 2, nil},
        {"4/0", 99, DivideByZeroError},
        {"4*6/2", 12, nil},
        {"6/3*3", 6, nil},
        {"1+2*3", 7, nil},
        {"(1+2)*3", 9, nil},
        {"1+2*3+4", 11, nil},
        {"1+(2*3)+4", 11, nil},
        {"(1+2)*(3+4)", 21, nil},
        {"((1+2)*(3+4))", 21, nil},
    }

    for _, c := range cases {
        result, err := Evaluate(c.expr)
        if err != c.err {
            t.Errorf("Expr '%s': got %v, want %v", c.expr, err, c.err)
        } else if err == nil && result != c.want {
            t.Errorf("Expr '%s': got %v, want %v", c.expr, result, c.want)
        }
    }
}

