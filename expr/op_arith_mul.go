package expr

type mulOp struct {
	left, right Expr
}

func (op mulOp) Evaluate(table *SymTab) (Expr, error) {
	exps, err := evaluateExprs(table, IsNumeric, op.left, op.right)
	if err != nil {
		return nil, err
	}
	return exps[0].(arithmeticValue).mul(exps[1].(arithmeticValue)), nil
}

func NewMul(left, right Expr) Expr {
	return mulOp{left, right}
}
