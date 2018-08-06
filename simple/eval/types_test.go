package eval

import (
    "testing"
)

func TestInteger(t *testing.T) {
    value := number{4}
    if r, _ := value.Evaluate(); r != 4.0 {
        t.Errorf("Got %v, want %v", r, 4.0)
    }
}

func TestNegativeInteger(t *testing.T) {
    value := number{-5}
    if r, _ := value.Evaluate(); r != -5.0 {
        t.Errorf("Got %v, want %v", r, -5.0)
    }
}

func TestReal(t *testing.T) {
    value := number{2.3}
    if r, _ := value.Evaluate(); r != 2.3 {
        t.Errorf("Got %v, want %v", r, 2.3)
    }
}

func TestNegativeReal(t *testing.T) {
    value := number{-5.6}
    if r, _ := value.Evaluate(); r != -5.6 {
        t.Errorf("Got %v, want %v", r, -5.6)
    }
}

func TestAdd(t *testing.T) {
    value := add{number{3}, number{4}}
    if r, _ := value.Evaluate(); r != 7.0 {
        t.Errorf("Got %v, want %v", r, 7.0)
    }
}

func TestAddNegative(t *testing.T) {
    value := add{number{3}, number{-7}}
    if r, _ := value.Evaluate(); r != -4.0 {
        t.Errorf("Got %v, want %v", r, -4.0)
    }
}

func TestSubtract(t *testing.T) {
    value := subtract{number{3}, number{4}}
    if r, _ := value.Evaluate(); r != -1.0 {
        t.Errorf("Got %v, want %v", r, -1.0)
    }
}

func TestSubtractNegative(t *testing.T) {
    value := subtract{number{3}, number{-7}}
    if r, _ := value.Evaluate(); r != 10.0 {
        t.Errorf("Got %v, want %v", r, 10.0)
    }
}

func TestMultiply(t *testing.T) {
    value := multiply{number{3}, number{4}}
    if r, _ := value.Evaluate(); r != 12.0 {
        t.Errorf("Got %v, want %v", r, 12.0)
    }
}

func TestMultiplyNegative(t *testing.T) {
    value := multiply{number{3}, number{-7}}
    if r, _ := value.Evaluate(); r != -21.0 {
        t.Errorf("Got %v, want %v", r, -21.0)
    }
}

func TestDivide(t *testing.T) {
    value := divide{number{5}, number{2}}
    if r, _ := value.Evaluate(); r != 2.5 {
        t.Errorf("Got %v, want %v", r, 2.5)
    }
}

func TestDivideNegative(t *testing.T) {
    value := divide{number{5}, number{-2}}
    if r, _ := value.Evaluate(); r != -2.5 {
        t.Errorf("Got %v, want %v", r, -2.5)
    }
}

func TestDivideByZero(t *testing.T) {
    value := divide{number{5}, number{0}}
    if _, err := value.Evaluate(); err.Error() != "divide by zero" {
        t.Errorf("Got %v, want %v", err.Error(), "divide by zero")
    }
}

func TestAddAndMultiply(t *testing.T) {
    value := multiply{number{3},add{number{4}, number{5}}}
    if r, _ := value.Evaluate(); r != 27.0 {
        t.Errorf("Got %v, want %v", r, 27.0)
    }
}

func TestAddAndDivide(t *testing.T) {
    value := divide{number{15},add{number{4}, number{2}}}
    if r, _ := value.Evaluate(); r != 2.5 {
        t.Errorf("Got %v, want %v", r, 2.5)
    }
}

func TestAddAndDivideByZero(t *testing.T) {
    value := divide{number{15},add{number{3}, number{-3}}}
    if _, err := value.Evaluate(); err.Error() != "divide by zero" {
        t.Errorf("Got %v, want %v", err.Error(), "divide by zero")
    }
}

