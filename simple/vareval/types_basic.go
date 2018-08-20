package vareval

import (
	"fmt"
	"math"
)

type expr interface {
	evaluate(table *symTab) (expr, error)
}

type value interface {
	expr
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
}

type intValue struct {
	value int
}

func (n intValue) equals(other value) bool {
	intOther, ok := other.(intValue)
	if !ok {
		return false
	}
	return n.value == intOther.value
}

func (n intValue) evaluate(table *symTab) (expr, error) {
	return intValue{n.value}, nil
}

func (n intValue) String() string {
	return fmt.Sprintf("%d", n.value)
}

func (n intValue) floatValue() float64 {
	return float64(n.value)
}

func (n intValue) add(other arithmeticValue) arithmeticValue {
	if isReal(other) {
		return realValue{n.floatValue() + other.floatValue()}
	}
	return intValue{n.value + other.(intValue).value}
}

func (n intValue) sub(other arithmeticValue) arithmeticValue {
	if isReal(other) {
		return realValue{n.floatValue() - other.floatValue()}
	}
	return intValue{n.value - other.(intValue).value}
}

func (n intValue) mul(other arithmeticValue) arithmeticValue {
	if isReal(other) {
		return realValue{n.floatValue() * other.floatValue()}
	}
	return intValue{n.value * other.(intValue).value}
}

func (n intValue) div(other arithmeticValue) (arithmeticValue, error) {
	if other.floatValue() == 0.0 || other.floatValue() == -0.0 {
		return nil, DivideByZeroError
	}
	if isReal(other) {
		return realValue{n.floatValue() / other.floatValue()}, nil
	}
	return intValue{n.value / other.(intValue).value}, nil
}

func (n intValue) pow(other arithmeticValue) (arithmeticValue, error) {
	return realValue{n.floatValue()}.pow(other)
}

type realValue struct {
	value float64
}

func (r realValue) equals(other value) bool {
	realOther, ok := other.(realValue)
	if !ok {
		return false
	}
	return r.value == realOther.value
}

func (r realValue) evaluate(table *symTab) (expr, error) {
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
	if other.floatValue() == 0.0 {
		return nil, DivideByZeroError
	}
	prod := math.Pow(r.value, other.floatValue())
	if math.IsNaN(prod) {
		return nil, DomainError
	}
	return realValue{prod}, nil
}

type boolValue struct {
	value bool
}

func (b boolValue) evaluate(table *symTab) (expr, error) {
	return b, nil
}

func (b boolValue) equals(other value) bool {
	boolOther, ok := other.(boolValue)
	if !ok {
		return false
	}
	return b.value == boolOther.value
}

func (b boolValue) String() string {
	return fmt.Sprintf("%t", b.value)
}

type stringValue struct {
	value string
}

func (s stringValue) evaluate(table *symTab) (expr, error) {
	return s, nil
}

func (s stringValue) equals(other value) bool {
	stringOther, ok := other.(stringValue)
	if !ok {
		return false
	}
	return s.value == stringOther.value
}

func (s stringValue) String() string {
	return s.value
}

type variableValue struct {
	key string
}

func (v variableValue) evaluate(table *symTab) (expr, error) {
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
	variableOther, ok := other.(variableValue)
	if !ok {
		return false
	}
	return v.key == variableOther.key
}

func (v variableValue) String() string {
	return "[" + v.key + "]"
}
