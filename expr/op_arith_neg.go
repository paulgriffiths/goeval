package expr

type negOp struct {
	value Expr
}

func (op negOp) Evaluate(table *SymTab) (Expr, error) {
	exps, err := evaluateExprs(table, IsNumeric, op.value)
	if err != nil {
		return nil, err
	}
	return exps[0].(arithmeticValue).negate(), nil
}

func NewNeg(value Expr) negOp {
	return negOp{value}
}
