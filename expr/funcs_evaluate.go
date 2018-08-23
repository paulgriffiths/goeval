package expr

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

func evalIfArithmetic(table *SymTab, e Expr) (arithmeticValue, error) {
	result, err := evaluateExprs(table, IsNumeric, e)
	if err != nil {
		return nil, err
	}
	return result[0].(arithmeticValue), nil
}

func evalIfBoolean(table *SymTab, e Expr) (boolValue, error) {
	result, err := evaluateExprs(table, IsBoolean, e)
	if err != nil {
		return boolValue{false}, err
	}
	return result[0].(boolValue), nil
}

func evalPairIfArithmetic(table *SymTab, a, b Expr) (arithmeticValue,
	arithmeticValue, error) {
	result, err := evaluateExprs(table, IsNumeric, a, b)
	if err != nil {
		return nil, nil, err
	}
	return result[0].(arithmeticValue), result[1].(arithmeticValue), nil
}

func evalPairIfBoolean(table *SymTab, a, b Expr) (boolValue,
	boolValue, error) {
	result, err := evaluateExprs(table, IsBoolean, a, b)
	if err != nil {
		return boolValue{false}, boolValue{false}, err
	}
	return result[0].(boolValue), result[1].(boolValue), nil
}

func evalPairIfString(table *SymTab, a, b Expr) (stringValue,
	stringValue, error) {
	result, err := evaluateExprs(table, IsString, a, b)
	if err != nil {
		return stringValue{""}, stringValue{""}, err
	}
	return result[0].(stringValue), result[1].(stringValue), nil
}
