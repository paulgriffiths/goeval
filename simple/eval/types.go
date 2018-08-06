package eval

import (
    "fmt"
)

type expr interface {
    Evaluate() (float64, error)
}

type number struct {
    value float64
}

func (n number) Evaluate() (float64, error) {
    return n.value, nil
}

type add struct {
    left, right expr
}

func (a add) Evaluate() (float64, error) {
    l, err := a.left.Evaluate()
    if err != nil {
        return 1, err
    }
    r, err := a.right.Evaluate()
    if err != nil {
        return 1, err
    }
    return l + r, nil
}

type subtract struct {
    left, right expr
}

func (s subtract) Evaluate() (float64, error) {
    l, err := s.left.Evaluate()
    if err != nil {
        return 1, err
    }
    r, err := s.right.Evaluate()
    if err != nil {
        return 1, err
    }
    return l - r, nil
}

type multiply struct {
    left, right expr
}

func (m multiply) Evaluate() (float64, error) {
    l, err := m.left.Evaluate()
    if err != nil {
        return 1, err
    }
    r, err := m.right.Evaluate()
    if err != nil {
        return 1, err
    }
    return l * r, nil
}

type divide struct {
    left, right expr
}

func (d divide) Evaluate() (float64, error) {
    l, err := d.left.Evaluate()
    if err != nil {
        return 1, err
    }
    r, err := d.right.Evaluate()
    if err != nil {
        return 1, err
    }
    if r == 0 {
        return 1, fmt.Errorf("divide by zero")
    }
    return l / r, nil
}

