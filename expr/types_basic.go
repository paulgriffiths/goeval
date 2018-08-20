package expr

import (
	"fmt"
	"math"
)

func FloatValueIfPossible(expr Expr) (float64, bool) {
	if !isNumeric(expr) {
		return 0, false
	}
	return expr.(arithmeticValue).floatValue(), true
}

type Expr interface {
	Evaluate(table *symTab) (Expr, error)
}

type value interface {
	Expr
	equals(other value) bool
	String() string
}

type arithmeticValue interface {
	value
	floatValue() float64
	add(other arithmeticValue) arithmeticValue
	sub(other arithmeticValue) arithmeticValue
	mul(other arithmeticValue) arithmeticValue
	div(other arithmeticValue) (arithmeticValue, error)
	pow(other arithmeticValue) (arithmeticValue, error)
	negate() arithmeticValue
}

type intValue struct {
	value int64
}

func (n intValue) equals(other value) bool {
	if !isInteger(other) {
		return false
	}
	return n.value == other.(intValue).value
}

func (n intValue) Evaluate(_ *symTab) (Expr, error) {
	return n, nil
}

func (n intValue) String() string {
	return fmt.Sprintf("%d", n.value)
}

func (n intValue) floatValue() float64 {
	return float64(n.value)
}

func (n intValue) toReal() realValue {
	return realValue{float64(n.value)}
}

func (n intValue) add(other arithmeticValue) arithmeticValue {
	if isReal(other) {
		return n.toReal().add(other)
	}
	return intValue{n.value + other.(intValue).value}
}

func (n intValue) sub(other arithmeticValue) arithmeticValue {
	if isReal(other) {
		return n.toReal().sub(other)
	}
	return intValue{n.value - other.(intValue).value}
}

func (n intValue) mul(other arithmeticValue) arithmeticValue {
	if isReal(other) {
		return n.toReal().mul(other)
	}
	return intValue{n.value * other.(intValue).value}
}

func (n intValue) div(other arithmeticValue) (arithmeticValue, error) {
	if isReal(other) {
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
	return n.toReal().pow(other)
}

func (n intValue) negate() arithmeticValue {
	return intValue{-n.value}
}

type realValue struct {
	value float64
}

func NewReal(value float64) realValue {
	return realValue{value}
}

func (r realValue) equals(other value) bool {
	if !isReal(other) {
		return false
	}
	return r.value == other.(realValue).value
}

func (r realValue) Evaluate(_ *symTab) (Expr, error) {
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

type boolValue struct {
	value bool
}

func (b boolValue) Evaluate(_ *symTab) (Expr, error) {
	return b, nil
}

func (b boolValue) equals(other value) bool {
	if !isBoolean(other) {
		return false
	}
	return b.value == other.(boolValue).value
}

func (b boolValue) String() string {
	return fmt.Sprintf("%t", b.value)
}

type stringValue struct {
	value string
}

func (s stringValue) Evaluate(_ *symTab) (Expr, error) {
	return s, nil
}

func (s stringValue) equals(other value) bool {
	if !isString(other) {
		return false
	}
	return s.value == other.(stringValue).value
}

func (s stringValue) String() string {
	return s.value
}

type variableValue struct {
	key string
}

func (v variableValue) Evaluate(table *symTab) (Expr, error) {
	if table == nil {
		panic("symbol table is nil")
	}

	val, ok := table.retrieve(v.key)
	if !ok {
		return nil, UnknownIdentifierError
	}
	return val, nil
}

func (v variableValue) equals(other value) bool {
	if !isVariable(other) {
		return false
	}
	return v.key == other.(variableValue).key
}

func (v variableValue) String() string {
	return "[" + v.key + "]"
}
