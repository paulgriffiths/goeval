package expr

func IsInteger(exp Expr) bool {
	if _, ok := exp.(intValue); !ok {
		return false
	}
	return true
}

func AreInteger(exps ...Expr) bool {
	for _, exp := range exps {
		if !IsInteger(exp) {
			return false
		}
	}
	return true
}

func IsReal(exp Expr) bool {
	if _, ok := exp.(realValue); !ok {
		return false
	}
	return true
}

func AreReal(exps ...Expr) bool {
	for _, exp := range exps {
		if !IsReal(exp) {
			return false
		}
	}
	return true
}

func IsNumeric(exp Expr) bool {
	return IsInteger(exp) || IsReal(exp)
}

func AreNumeric(exps ...Expr) bool {
	for _, exp := range exps {
		if !IsNumeric(exp) {
			return false
		}
	}
	return true
}

func IsBoolean(exp Expr) bool {
	if _, ok := exp.(boolValue); !ok {
		return false
	}
	return true
}

func AreBoolean(exps ...Expr) bool {
	for _, exp := range exps {
		if !IsBoolean(exp) {
			return false
		}
	}
	return true
}

func IsString(exp Expr) bool {
	if _, ok := exp.(stringValue); !ok {
		return false
	}
	return true
}

func AreString(exps ...Expr) bool {
	for _, exp := range exps {
		if !IsString(exp) {
			return false
		}
	}
	return true
}

func IsVariable(exp Expr) bool {
	if _, ok := exp.(variableValue); !ok {
		return false
	}
	return true
}

func AreVariable(exps ...Expr) bool {
	for _, exp := range exps {
		if !IsVariable(exp) {
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
