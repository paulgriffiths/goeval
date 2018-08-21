package expr

type divOp struct {
	left, right Expr
}

func (op divOp) Evaluate(table *SymTab) (Expr, error) {
	exps, err := evaluateExprs(table, IsNumeric, op.left, op.right)
	if err != nil {
		return nil, err
	}
	return exps[0].(arithmeticValue).div(exps[1].(arithmeticValue))
}

func NewDiv(left, right Expr) divOp {
	return divOp{left, right}
}
