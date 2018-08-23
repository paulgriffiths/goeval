package expr

// Equals tests if two expressions are equal in value and type.
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

// ToInt returns the integral value of an expression if it is an
// integral expression, and returns false otherwise.
func ToInt(exp Expr) (int64, bool) {
	if !IsInteger(exp) {
		return 0, false
	}
	return exp.(intValue).value, true
}

// ToFloat returns the float value of an expression if it is an
// integral or real expression, and returns false otherwise.
func ToFloat(exp Expr) (float64, bool) {
	if !IsNumeric(exp) {
		return 0.0, false
	}
	return exp.(arithmeticValue).floatValue(), true
}

// ToBool returns the boolean value of an expression if it is a
// boolean expression, and returns false otherwise.
func ToBool(exp Expr) (bool, bool) {
	if !IsBoolean(exp) {
		return false, false
	}
	return exp.(boolValue).value, true
}

// ToString returns the string value of an expression if it is a
// string expression, and returns false otherwise.
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

// IsInteger tests if an expression is an integral value.
func IsInteger(exp Expr) bool {
	if _, ok := exp.(intValue); !ok {
		return false
	}
	return true
}

// AreInteger tests if all the provided expressions are integral values.
func AreInteger(exps ...Expr) bool {
	for _, exp := range exps {
		if !IsInteger(exp) {
			return false
		}
	}
	return true
}

// IsReal tests if an expression is a real value.
func IsReal(exp Expr) bool {
	if _, ok := exp.(realValue); !ok {
		return false
	}
	return true
}

// AreReal tests if all the provided expressions are real values.
func AreReal(exps ...Expr) bool {
	for _, exp := range exps {
		if !IsReal(exp) {
			return false
		}
	}
	return true
}

// IsNumeric tests if an expression is an integral or a real value.
func IsNumeric(exp Expr) bool {
	return IsInteger(exp) || IsReal(exp)
}

// AreNumeric tests if all the provided expressions are integral or
// real values.
func AreNumeric(exps ...Expr) bool {
	for _, exp := range exps {
		if !IsNumeric(exp) {
			return false
		}
	}
	return true
}

// IsBoolean tests if an expression is a boolean value.
func IsBoolean(exp Expr) bool {
	if _, ok := exp.(boolValue); !ok {
		return false
	}
	return true
}

// AreBoolean tests if all the provided expressions are boolean values.
func AreBoolean(exps ...Expr) bool {
	for _, exp := range exps {
		if !IsBoolean(exp) {
			return false
		}
	}
	return true
}

// IsString tests if an expression is a string value.
func IsString(exp Expr) bool {
	if _, ok := exp.(stringValue); !ok {
		return false
	}
	return true
}

// AreString tests if all the provided expressions are string values.
func AreString(exps ...Expr) bool {
	for _, exp := range exps {
		if !IsString(exp) {
			return false
		}
	}
	return true
}

// IsVariable tests if an expression is a variable value.
func IsVariable(exp Expr) bool {
	if _, ok := exp.(variableValue); !ok {
		return false
	}
	return true
}

// AreVariable tests if all the provided expressions are variable values.
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
