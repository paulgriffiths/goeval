package expr

import (
	"fmt"
)

type boolValue struct {
	value bool
}

func NewBool(value bool) Expr {
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

func (b boolValue) Equals(other value) bool {
	if !IsBoolean(other) {
		return false
	}
	return b.value == other.(boolValue).value
}

func (b boolValue) String() string {
	return fmt.Sprintf("%t", b.value)
}
