package expr

type lessThanOp struct {
	left, right Expr
}

func (op lessThanOp) Evaluate(table *SymTab) (Expr, error) {
	exps, err := evaluateExprs(table, IsNumeric, op.left, op.right)
	if err != nil {
		return nil, err
	}
	cmp := exps[0].(arithmeticValue).lessThan(exps[1].(arithmeticValue))
	return boolValue{cmp}, nil
}

func NewLessThan(left, right Expr) lessThanOp {
	return lessThanOp{left, right}
}
