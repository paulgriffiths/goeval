package eval

import (
    "testing"
)

func TestGoodExpressionTypes(t *testing.T) {
    cases := []struct {
        t expr
        expected float64
    }{
        {number{4}, 4},
        {number{-4}, -4},
        {number{2.3}, 2.3},
        {number{-2.3}, -2.3},
        {negate{number{4}}, -4},
        {negate{number{-4}}, 4},
        {negate{number{2.3}}, -2.3},
        {negate{number{-2.3}}, 2.3},
        {add{number{3}, number{4}}, 7},
        {add{number{3}, number{-4}}, -1},
        {add{number{-3}, number{4}}, 1},
        {add{number{-3}, number{-4}}, -7},
        {subtract{number{3}, number{4}}, -1},
        {subtract{number{3}, number{-4}}, 7},
        {subtract{number{-3}, number{4}}, -7},
        {subtract{number{-3}, number{-4}}, 1},
        {multiply{number{3}, number{4}}, 12},
        {multiply{number{3}, number{-4}}, -12},
        {multiply{number{-3}, number{4}}, -12},
        {multiply{number{-3}, number{-4}}, 12},
        {multiply{number{3}, add{number{4}, number{5}}}, 27},
        {divide{number{6}, number{2}}, 3},
        {divide{number{6}, number{-2}}, -3},
        {divide{number{-6}, number{2}}, -3},
        {divide{number{-6}, number{-2}}, 3},
        {power{number{16}, number{2}}, 256},
        {power{number{2}, number{-2}}, 0.25},
        {power{number{-16}, number{2}}, 256},
        {power{number{-2}, number{-2}}, 0.25},
        {power{number{4}, number{0.5}}, 2},
        {power{number{16}, number{0.25}}, 2},
        {power{number{4}, number{0}}, 1},
        {power{number{0}, number{0}}, 1},
        {power{number{-4}, number{0}}, 1},
    }

    for n, c := range cases {
        r, err := c.t.Evaluate()
        if err != nil {
            t.Errorf("Case %d, got error %v, want %v", n, err, nil)
        } else if r != c.expected {
            t.Errorf("Case %d, got %v, want %v", n, r, c.expected)
        }
    }
}

func TestBadExpressionTypes(t *testing.T) {
    cases := []struct {
        t expr
        err error
    }{
        {divide{number{4}, number{0}}, DivideByZeroError},
        {power{number{-4}, number{0.5}}, DomainError},
        {power{number{-4}, number{1.5}}, DomainError},
    }

    for n, c := range cases {
        _, err := c.t.Evaluate()
        if err != c.err {
            t.Errorf("Case %d, got error %v, want %v", n, err, c.err)
        }
    }
}
