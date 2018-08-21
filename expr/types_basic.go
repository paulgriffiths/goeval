package expr

import (
	"fmt"
	"math"
)

func IntValueIfPossible(expr Expr) (int64, bool) {
	if !IsInteger(expr) {
		return 0, false
	}
	return expr.(intValue).value, true
}

func FloatValueIfPossible(expr Expr) (float64, bool) {
	if !IsNumeric(expr) {
		return 0, false
	}
	return expr.(arithmeticValue).floatValue(), true
}

type Expr interface {
	Evaluate(table *SymTab) (Expr, error)
}

type Value interface {
	Expr
	Equals(other Value) bool
	String() string
}

type arithmeticValue interface {
	Value
	floatValue() float64
	almostEquals(other Value, epsilon float64) bool
	add(other arithmeticValue) arithmeticValue
	sub(other arithmeticValue) arithmeticValue
	mul(other arithmeticValue) arithmeticValue
	div(other arithmeticValue) (arithmeticValue, error)
	pow(other arithmeticValue) (arithmeticValue, error)
	negate() arithmeticValue
	equality(other arithmeticValue) bool
	lessThan(other arithmeticValue) bool
}

type intValue struct {
	value int64
}

func NewInt(value int64) intValue {
	return intValue{value}
}

func (n intValue) Equals(other Value) bool {
	if !IsInteger(other) {
		return false
	}
	return n.value == other.(intValue).value
}

func (n intValue) almostEquals(other Value, _ float64) bool {
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

func (n intValue) Evaluate(_ *SymTab) (Expr, error) {
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

type boolValue struct {
	value bool
}

func NewBool(value bool) boolValue {
	return boolValue{value}
}

func (b boolValue) equality(other boolValue) boolValue {
	return boolValue{b.value == other.value}
}

func (b boolValue) and(other boolValue) boolValue {
	return boolValue{b.value && other.value}
}

func (b boolValue) or(other boolValue) boolValue {
	return boolValue{b.value || other.value}
}

func (b boolValue) xor(other boolValue) boolValue {
	return b.equality(other).not()
}

func (b boolValue) nor(other boolValue) boolValue {
	return b.or(other).not()
}

func (b boolValue) nand(other boolValue) boolValue {
	return b.and(other).not()
}

func (b boolValue) not() boolValue {
	return boolValue{!b.value}
}

func (b boolValue) Evaluate(_ *SymTab) (Expr, error) {
	return b, nil
}

func (b boolValue) Equals(other Value) bool {
	if !IsBoolean(other) {
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

func NewString(value string) stringValue {
	return stringValue{value}
}

func (s stringValue) equality(other stringValue) boolValue {
	return boolValue{s.value == other.value}
}

func (s stringValue) Evaluate(_ *SymTab) (Expr, error) {
	return s, nil
}

func (s stringValue) Equals(other Value) bool {
	if !IsString(other) {
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

func NewVariable(key string) variableValue {
	return variableValue{key}
}

func (v variableValue) Evaluate(table *SymTab) (Expr, error) {
	if table == nil {
		panic("symbol table is nil")
	}

	val, ok := table.Retrieve(v.key)
	if !ok {
		return nil, UnknownIdentifierError
	}
	return val, nil
}

func (v variableValue) Equals(other Value) bool {
	if !IsVariable(other) {
		return false
	}
	return v.key == other.(variableValue).key
}

func (v variableValue) String() string {
	return "[" + v.key + "]"
}
