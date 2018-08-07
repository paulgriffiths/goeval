package eval

import (
    "testing"
)

func TestParseNumber(t *testing.T) {
    v, err := Evaluate("2")
    if err != nil {
        t.Errorf("Got %v, want %v", err, nil)
    }
    if v != 2.0 {
        t.Errorf("Got %v, want %v", v, 2.0)
    }
}

func TestParseParenthesizedNumber(t *testing.T) {
    v, err := Evaluate("(2)")
    if err != nil {
        t.Errorf("Got %v, want %v", err, nil)
    }
    if v != 2.0 {
        t.Errorf("Got %v, want %v", v, 2.0)
    }
}

func TestParseDoubleParenthesizedNumber(t *testing.T) {
    v, err := Evaluate("((2))")
    if err != nil {
        t.Errorf("Got %v, want %v", err, nil)
    }
    if v != 2.0 {
        t.Errorf("Got %v, want %v", v, 2.0)
    }
}

func TestParseNumberMissingRightParenthesis(t *testing.T) {
    _, err := Evaluate("(2")
    if err != UnbalancedParenthesesError {
        t.Errorf("Got %v, want %v", err, UnbalancedParenthesesError)
    }
}

func TestParseRightParenthesis(t *testing.T) {
    _, err := Evaluate(")")
    if err != UnbalancedParenthesesError {
        t.Errorf("Got %v, want %v", err, UnbalancedParenthesesError)
    }
}

func TestParseMultiply(t *testing.T) {
    v, err := Evaluate("2*3")
    if err != nil {
        t.Errorf("Got %v, want %v", err, nil)
    }
    if v != 6.0 {
        t.Errorf("Got %v, want %v", v, 6.0)
    }
}

func TestParseMultiplyThreeNumbers(t *testing.T) {
    v, err := Evaluate("2*3*4")
    if err != nil {
        t.Errorf("Got %v, want %v", err, nil)
    }
    if v != 24.0 {
        t.Errorf("Got %v, want %v", v, 24.0)
    }
}

func TestParseDivide(t *testing.T) {
    v, err := Evaluate("3/2")
    if err != nil {
        t.Errorf("Got %v, want %v", err, nil)
    }
    if v != 1.5 {
        t.Errorf("Got %v, want %v", v, 1.5)
    }
}

func TestParseDivideThreeNumbers(t *testing.T) {
    v, err := Evaluate("60/10/2")
    if err != nil {
        t.Errorf("Got %v, want %v", err, nil)
    }
    if v != 3.0 {
        t.Errorf("Got %v, want %v", v, 3.0)
    }
}

func TestParseDivideByZero(t *testing.T) {
    _, err := Evaluate("3/0")
    if err != DivideByZeroError {
        t.Errorf("Got %v, want %v", err, DivideByZeroError)
    }
}

func TestParseAdd(t *testing.T) {
    v, err := Evaluate("6+9")
    if err != nil {
        t.Errorf("Got %v, want %v", err, nil)
    }
    if v != 15.0 {
        t.Errorf("Got %v, want %v", v, 15.0)
    }
}

func TestParseAddThreeNumbers(t *testing.T) {
    v, err := Evaluate("6+9+12")
    if err != nil {
        t.Errorf("Got %v, want %v", err, nil)
    }
    if v != 27.0 {
        t.Errorf("Got %v, want %v", v, 27.0)
    }
}

func TestParseAddSubtract(t *testing.T) {
    v, err := Evaluate("6+9-12")
    if err != nil {
        t.Errorf("Got %v, want %v", err, nil)
    }
    if v != 3.0 {
        t.Errorf("Got %v, want %v", v, 3.0)
    }
}

func TestParseSubtract(t *testing.T) {
    v, err := Evaluate("11-5")
    if err != nil {
        t.Errorf("Got %v, want %v", err, nil)
    }
    if v != 6.0 {
        t.Errorf("Got %v, want %v", v, 6.0)
    }
}

func TestParseSubtractThreeNumber(t *testing.T) {
    v, err := Evaluate("11-5-18")
    if err != nil {
        t.Errorf("Got %v, want %v", err, nil)
    }
    if v != -12.0 {
        t.Errorf("Got %v, want %v", v, -12.0)
    }
}

func TestParseSubtractAdd(t *testing.T) {
    v, err := Evaluate("6-9+12")
    if err != nil {
        t.Errorf("Got %v, want %v", err, nil)
    }
    if v != 9.0 {
        t.Errorf("Got %v, want %v", v, 9.0)
    }
}

func TestParseAddAndMultiply(t *testing.T) {
    v, err := Evaluate("2+3*4")
    if err != nil {
        t.Errorf("Got %v, want %v", err, nil)
    }
    if v != 14.0 {
        t.Errorf("Got %v, want %v", v, 14.0)
    }
}

func TestParseParenthesizedAddAndMultiply(t *testing.T) {
    v, err := Evaluate("(2+3)*4")
    if err != nil {
        t.Errorf("Got %v, want %v", err, nil)
    }
    if v != 20.0 {
        t.Errorf("Got %v, want %v", v, 20.0)
    }
}

