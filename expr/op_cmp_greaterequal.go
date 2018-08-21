package expr

type greaterThanOrEqualOp struct {
	left, right Expr
}

func (op greaterThanOrEqualOp) Evaluate(table *SymTab) (Expr, error) {
	exps, err := evaluateExprs(table, IsNumeric, op.left, op.right)
	if err != nil {
		return nil, err
	}
	lt := exps[0].(arithmeticValue).lessThan(exps[1].(arithmeticValue))
	eq := exps[0].(arithmeticValue).equality(exps[1].(arithmeticValue))
	return boolValue{!lt || eq}, nil
}

func NewGreaterThanOrEqual(left, right Expr) greaterThanOrEqualOp {
	return greaterThanOrEqualOp{left, right}
}