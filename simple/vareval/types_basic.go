package vareval

import "fmt"

type expr interface {
	evaluate(table *symTab) (expr, error)
}

type value interface {
	expr
	equals(other value) bool
	String() string
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
