package expr

import (
	"fmt"
)

// intValue contains an integral value
type intValue struct {
	value int64
}

// NewInt creates a new integral value
func NewInt(value int64) Expr {
	return intValue{value}
}

// Equals tests whether an integral value equals another
func (n intValue) Equals(other value) bool {
	if !IsInteger(other) {
		return false
	}
	return n.value == other.(intValue).value
}

// Evaluate evalutes an integral value
func (n intValue) Evaluate(_ *SymTab) (Expr, error) {
	return n, nil
}

// String returns a string representation of an integral value
func (n intValue) String() string {
	return fmt.Sprintf("%d", n.value)
}

func (n intValue) almostEquals(other value, _ float64) bool {
	return n.Equals(other)
}

func (n intValue) equality(other arithmeticValue) bool {
	if !IsInteger(other) {
		return n.toReal().equality(other)
	}
	return n.value == other.(intValue).value
}

func (n intValue) lessThan(other arithmeticValue) bool {
	if !IsInteger(other) {
		return n.toReal().lessThan(other)
	}
	return n.value < other.(intValue).value
}

func (n intValue) floatValue() float64 {
	return float64(n.value)
}

func (n intValue) toReal() realValue {
	return realValue{float64(n.value)}
}

func (n intValue) add(other arithmeticValue) arithmeticValue {
	if IsReal(other) {
		return n.toReal().add(other)
	}
	return intValue{n.value + other.(intValue).value}
}

func (n intValue) sub(other arithmeticValue) arithmeticValue {
	if IsReal(other) {
		return n.toReal().sub(other)
	}
	return intValue{n.value - other.(intValue).value}
}

func (n intValue) mul(other arithmeticValue) arithmeticValue {
	if IsReal(other) {
		return n.toReal().mul(other)
	}
	return intValue{n.value * other.(intValue).value}
}

func (n intValue) div(other arithmeticValue) (arithmeticValue, error) {
	if IsReal(other) {
		return n.toReal().div(other)
	}
	if other.(intValue).value == 0 {
		return nil, DivideByZeroError
	}
	if n.value%other.(intValue).value != 0 {
		return n.toReal().div(other)
	}
	return intValue{n.value / other.(intValue).value}, nil
}

func (n intValue) pow(other arithmeticValue) (arithmeticValue, error) {
	if IsReal(other) || other.(intValue).value < 0 {
		return n.toReal().pow(other)
	}
	if other.(intValue).value == 0 {
		return intValue{1}, nil
	}
	retval := n.value
	for i := int64(1); i < other.(intValue).value; i++ {
		retval *= n.value
	}
	return intValue{retval}, nil
}

func (n intValue) negate() arithmeticValue {
	return intValue{-n.value}
}
