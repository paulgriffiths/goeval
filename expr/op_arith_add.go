package expr

type addOp struct {
	left, right Expr
}

func (op addOp) Evaluate(table *SymTab) (Expr, error) {
	exps, err := evaluateExprs(table, IsNumeric, op.left, op.right)
	if err != nil {
		return nil, err
	}
	return exps[0].(arithmeticValue).add(exps[1].(arithmeticValue)), nil
}

func NewAdd(left, right Expr) addOp {
	return addOp{left, right}
}
