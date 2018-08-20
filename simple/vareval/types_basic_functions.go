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

/*
func mustIntegerValue(exp expr) int {
	n, ok := exp.(intValue)
	if !ok {
		panic("integer exp expected")
	}
	return n.value
}

func mustRealValue(exp expr) float64 {
	if !isNumeric(exp) {
		panic("integer exp expected")
	}
	if f, ok := exp.(intValue); ok {
		return float64(f.value)
	}
	f, _ := exp.(realValue)
	return f.value
}

func mustSumIntegerValues(exps ...expr) int {
    var sum int
    for _, e := range exps {
        n, ok := e.(intValue)
        if !ok {
            panic("integer expr expected")
        }
        sum = sum + n.value
    }
	return sum
}

func mustSumRealValues(exps ...expr) float64 {
    var sum float64
    for _, e := range exps {
        if n, ok := e.(intValue); ok {
            sum = sum + float64(n.value)
        } else if n, ok := e.(realValue); ok {
            sum = sum + n.value
        } else {
            panic("real expr expected")
        }
    }
	return sum
}
*/
