package expr

import (
	"fmt"
	"math"
)

type realValue struct {
	value float64
}

func NewReal(value float64) realValue {
	return realValue{value}
}

func (r realValue) Equals(other Value) bool {
	if !IsReal(other) {
		return false
	}
	return r.value == other.(realValue).value
}

func (r realValue) almostEquals(other Value, epsilon float64) bool {
	if !IsReal(other) {
		return false
	}
	if math.Abs(r.value-other.(realValue).value) <= epsilon {
		return true
	} else {
		return false
	}
}

func (r realValue) equality(other arithmeticValue) bool {
	return r.value == other.floatValue()
}

func (r realValue) lessThan(other arithmeticValue) bool {
	return r.value < other.floatValue()
}

func (r realValue) Evaluate(_ *SymTab) (Expr, error) {
	return r, nil
}

func (r realValue) String() string {
	return fmt.Sprintf("%f", r.value)
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
