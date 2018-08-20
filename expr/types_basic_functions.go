package expr

func isInteger(exp Expr) bool {
	if _, ok := exp.(intValue); !ok {
		return false
	}
	return true
}

func areInteger(exps ...Expr) bool {
	for _, exp := range exps {
		if !isInteger(exp) {
			return false
		}
	}
	return true
}

func isReal(exp Expr) bool {
	if _, ok := exp.(realValue); !ok {
		return false
	}
	return true
}

func areReal(exps ...Expr) bool {
	for _, exp := range exps {
		if !isReal(exp) {
			return false
		}
	}
	return true
}

func isNumeric(exp Expr) bool {
	return isInteger(exp) || isReal(exp)
}

func areNumeric(exps ...Expr) bool {
	for _, exp := range exps {
		if !isNumeric(exp) {
			return false
		}
	}
	return true
}

func isBoolean(exp Expr) bool {
	if _, ok := exp.(boolValue); !ok {
		return false
	}
	return true
}

func areBoolean(exps ...Expr) bool {
	for _, exp := range exps {
		if !isBoolean(exp) {
			return false
		}
	}
	return true
}

func isString(exp Expr) bool {
	if _, ok := exp.(stringValue); !ok {
		return false
	}
	return true
}

func areString(exps ...Expr) bool {
	for _, exp := range exps {
		if !isString(exp) {
			return false
		}
	}
	return true
}

func isVariable(exp Expr) bool {
	if _, ok := exp.(variableValue); !ok {
		return false
	}
	return true
}

func areVariable(exps ...Expr) bool {
	for _, exp := range exps {
		if !isVariable(exp) {
			return false
		}
	}
	return true
}

func evaluateExprs(table *symTab, testFunc func(Expr) bool,
	exps ...Expr) ([]Expr, error) {
	result := []Expr{}
	for _, val := range exps {
		v, err := val.Evaluate(table)
		if err != nil {
			return nil, err
		}
		if testFunc != nil && !testFunc(v) {
			return nil, TypeError
		}
		result = append(result, v)
	}
	return result, nil
}
