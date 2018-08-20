package vareval

func isInteger(exp expr) bool {
	if _, ok := exp.(intValue); !ok {
		return false
	}
	return true
}

func areInteger(exps ...expr) bool {
	for _, exp := range exps {
		if !isInteger(exp) {
			return false
		}
	}
	return true
}

func isReal(exp expr) bool {
	if _, ok := exp.(realValue); !ok {
		return false
	}
	return true
}

func areReal(exps ...expr) bool {
	for _, exp := range exps {
		if !isReal(exp) {
			return false
		}
	}
	return true
}

func isNumeric(exp expr) bool {
	return isInteger(exp) || isReal(exp)
}

func areNumeric(exps ...expr) bool {
	for _, exp := range exps {
		if !isNumeric(exp) {
			return false
		}
	}
	return true
}

func isBoolean(exp expr) bool {
	if _, ok := exp.(boolValue); !ok {
		return false
	}
	return true
}

func areBoolean(exps ...expr) bool {
	for _, exp := range exps {
		if !isBoolean(exp) {
			return false
		}
	}
	return true
}

func isString(exp expr) bool {
	if _, ok := exp.(stringValue); !ok {
		return false
	}
	return true
}

func areString(exps ...expr) bool {
	for _, exp := range exps {
		if !isString(exp) {
			return false
		}
	}
	return true
}

func isVariable(exp expr) bool {
	if _, ok := exp.(variableValue); !ok {
		return false
	}
	return true
}

func areVariable(exps ...expr) bool {
	for _, exp := range exps {
		if !isVariable(exp) {
			return false
		}
	}
	return true
}
