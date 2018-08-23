package expr

type variableValue struct {
	key string
}

// NewVariable creates a new variable value expression.
func NewVariable(name string) Expr {
	return variableValue{name}
}

func (v variableValue) Evaluate(table *SymTab) (Expr, error) {
	if table == nil {
		return nil, UndefinedVariableError
	}

	val, ok := table.Retrieve(v.key)
	if !ok {
		return nil, UndefinedVariableError
	}
	return val, nil
}

func (v variableValue) Equals(other value) bool {
	if !IsVariable(other) {
		return false
	}
	return v.key == other.(variableValue).key
}

func (v variableValue) String() string {
	return v.key
}
