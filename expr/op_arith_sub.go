package expr

type subOp struct {
	left, right Expr
}

func (op subOp) Evaluate(table *SymTab) (Expr, error) {
	exps, err := evaluateExprs(table, IsNumeric, op.left, op.right)
	if err != nil {
		return nil, err
	}
	return exps[0].(arithmeticValue).sub(exps[1].(arithmeticValue)), nil
}

func NewSub(left, right Expr) subOp {
	return subOp{left, right}
}
