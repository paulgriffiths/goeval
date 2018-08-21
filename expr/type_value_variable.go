package expr

type variableValue struct {
	key string
}

func NewVariable(key string) variableValue {
	return variableValue{key}
}

func (v variableValue) Evaluate(table *SymTab) (Expr, error) {
	if table == nil {
		return nil, UnknownIdentifierError
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
