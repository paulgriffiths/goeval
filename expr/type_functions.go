package expr

func Equals(left, right Expr, table *SymTab) bool {
	exps, err := evaluateExprs(nil, isValue, left, right)
	if err != nil {
		return false
	}
	if exps[0].(value).Equals(exps[1].(value)) {
		return true
	}
	return false
}

func ToInt(exp Expr) (int64, bool) {
	if !IsInteger(exp) {
		return 0, false
	}
	return exp.(intValue).value, true
}

func ToFloat(exp Expr) (float64, bool) {
	if !IsNumeric(exp) {
		return 0.0, false
	}
	return exp.(arithmeticValue).floatValue(), true
}

func ToBool(exp Expr) (bool, bool) {
	if !IsBoolean(exp) {
		return false, false
	}
	return exp.(boolValue).value, true
}

func ToString(exp Expr) (string, bool) {
	if !IsString(exp) {
		return "", false
	}
	return exp.(stringValue).value, true
}

func isValue(exp Expr) bool {
	if _, ok := exp.(value); !ok {
		return false
	}
	return true
}

func areValue(exps ...Expr) bool {
	for _, exp := range exps {
		if !isValue(exp) {
			return false
		}
	}
	return true
}

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

func evaluateExprs(table *SymTab, testFunc func(Expr) bool,
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
