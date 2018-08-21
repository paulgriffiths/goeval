package expr

type equalityOp struct {
	left, right Expr
}

func (op equalityOp) Evaluate(table *SymTab) (Expr, error) {
	exps, err := evaluateExprs(table, IsBoolean, op.left, op.right)
	if err == nil {
		cmp := exps[0].(boolValue).equality(exps[1].(boolValue))
		return cmp, nil
	}
	exps, err = evaluateExprs(table, IsString, op.left, op.right)
	if err == nil {
		cmp := exps[0].(stringValue).equality(exps[1].(stringValue))
		return cmp, nil
	}
	exps, err = evaluateExprs(table, IsNumeric, op.left, op.right)
	if err == nil {
		cmp := exps[0].(arithmeticValue).equality(exps[1].(arithmeticValue))
		return boolValue{cmp}, nil
	}
	return nil, TypeError
}

func NewEquality(left, right Expr) Expr {
	return equalityOp{left, right}
}
