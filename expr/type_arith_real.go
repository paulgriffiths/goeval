package expr

import (
	"fmt"
	"math"
)

// realValue contains a real value
type realValue struct {
	value float64
}

// NewReal creates a new real value expression.
func NewReal(value float64) Expr {
	return realValue{value}
}

// Equals tests if a real value equals another
func (r realValue) Equals(other value) bool {
	if !IsReal(other) {
		return false
	}
	return r.value == other.(realValue).value
}

// Evaluate evaluates a real value
func (r realValue) Evaluate(_ *SymTab) (Expr, error) {
	return r, nil
}

// String returns a string representation of a real value
func (r realValue) String() string {
	return fmt.Sprintf("%g", r.value)
}

func (r realValue) almostEquals(other value, epsilon float64) bool {
	if !IsReal(other) {
		return false
	}

	// Note that this logic only works for float values sufficiently
	// small so that the epsilon is within its precision. This is ok
	// since the function is only used internally for testing.

	if math.Abs(r.value-other.(realValue).value) <= epsilon {
		return true
	}
	return false
}

func (r realValue) equality(other arithmeticValue) bool {
	return r.value == other.floatValue()
}

func (r realValue) lessThan(other arithmeticValue) bool {
	return r.value < other.floatValue()
}

func (r realValue) floatValue() float64 {
	return r.value
}

func (r realValue) add(other arithmeticValue) arithmeticValue {
	return realValue{r.value + other.floatValue()}
}

func (r realValue) sub(other arithmeticValue) arithmeticValue {
	return realValue{r.value - other.floatValue()}
}

func (r realValue) mul(other arithmeticValue) arithmeticValue {
	return realValue{r.value * other.floatValue()}
}

func (r realValue) div(other arithmeticValue) (arithmeticValue, error) {
	if other.floatValue() == 0.0 {
		return nil, DivideByZeroError
	}
	return realValue{r.value / other.floatValue()}, nil
}

func (r realValue) pow(other arithmeticValue) (arithmeticValue, error) {
	prod := math.Pow(r.value, other.floatValue())
	if math.IsNaN(prod) {
		return nil, DomainError
	}
	return realValue{prod}, nil
}

func (r realValue) negate() arithmeticValue {
	return realValue{-r.value}
}
