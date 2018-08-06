package eval

import (
    "testing"
)

func TestParseNumber(t *testing.T) {
    v, err := evaluate("2")
    if err != nil {
        t.Errorf("Got %v, want %v", err, nil)
    }
    if v != 2.0 {
        t.Errorf("Got %v, want %v", v, 2.0)
    }
}

func TestParseParenthesizedNumber(t *testing.T) {
    v, err := evaluate("(2)")
    if err != nil {
        t.Errorf("Got %v, want %v", err, nil)
    }
    if v != 2.0 {
        t.Errorf("Got %v, want %v", v, 2.0)
    }
}

func TestParseDoubleParenthesizedNumber(t *testing.T) {
    v, err := evaluate("((2))")
    if err != nil {
        t.Errorf("Got %v, want %v", err, nil)
    }
    if v != 2.0 {
        t.Errorf("Got %v, want %v", v, 2.0)
    }
}

func TestParseNumberMissingRightParenthesis(t *testing.T) {
    _, err := evaluate("(2")
    if err == nil || err.Error() == "mismatched parentheses" {
        t.Errorf("Got %v, want %v", err, "mismatched parentheses")
    }
}

func TestParseRightParenthesis(t *testing.T) {
    _, err := evaluate(")")
    if err == nil || err.Error() == "mismatched parentheses" {
        t.Errorf("Got %v, want %v", err, "mismatched parentheses")
    }
}

func TestParseMultiply(t *testing.T) {
    v, err := evaluate("2*3")
    if err != nil {
        t.Errorf("Got %v, want %v", err, nil)
    }
    if v != 6.0 {
        t.Errorf("Got %v, want %v", v, 6.0)
    }
}

func TestParseDivide(t *testing.T) {
    v, err := evaluate("3/2")
    if err != nil {
        t.Errorf("Got %v, want %v", err, nil)
    }
    if v != 1.5 {
        t.Errorf("Got %v, want %v", v, 1.5)
    }
}

func TestParseAdd(t *testing.T) {
    v, err := evaluate("6+9")
    if err != nil {
        t.Errorf("Got %v, want %v", err, nil)
    }
    if v != 15.0 {
        t.Errorf("Got %v, want %v", v, 15.0)
    }
}

func TestParseSubtract(t *testing.T) {
    v, err := evaluate("11-5")
    if err != nil {
        t.Errorf("Got %v, want %v", err, nil)
    }
    if v != 6.0 {
        t.Errorf("Got %v, want %v", v, 6.0)
    }
}

func TestParseAddAndMultiply(t *testing.T) {
    v, err := evaluate("2+3*4")
    if err != nil {
        t.Errorf("Got %v, want %v", err, nil)
    }
    if v != 14.0 {
        t.Errorf("Got %v, want %v", v, 14.0)
    }
}

func TestParseParenthesizedAddAndMultiply(t *testing.T) {
    v, err := evaluate("(2+3)*4")
    if err != nil {
        t.Errorf("Got %v, want %v", err, nil)
    }
    if v != 20.0 {
        t.Errorf("Got %v, want %v", v, 20.0)
    }
}

