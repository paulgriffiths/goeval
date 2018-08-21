package expr

type nonEqualityOp struct {
	left, right Expr
}

func (op nonEqualityOp) Evaluate(table *SymTab) (Expr, error) {
	exps, err := evaluateExprs(table, IsBoolean, op.left, op.right)
	if err == nil {
		cmp := exps[0].(boolValue).equality(exps[1].(boolValue))
		return boolValue{!cmp.value}, nil
	}
	exps, err = evaluateExprs(table, IsString, op.left, op.right)
	if err == nil {
		cmp := exps[0].(stringValue).equality(exps[1].(stringValue))
		return boolValue{!cmp.value}, nil
	}
	exps, err = evaluateExprs(table, IsNumeric, op.left, op.right)
	if err == nil {
		cmp := !exps[0].(arithmeticValue).equality(exps[1].(arithmeticValue))
		return boolValue{cmp}, nil
	}
	return nil, TypeError
}

func NewNonEquality(left, right Expr) Expr {
	return nonEqualityOp{left, right}
}
