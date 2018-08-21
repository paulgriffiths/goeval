package expr

type xorOp struct {
	left, right Expr
}

func (op xorOp) Evaluate(table *SymTab) (Expr, error) {
	exps, err := evaluateExprs(table, IsBoolean, op.left, op.right)
	if err != nil {
		return nil, err
	}
	cmp := exps[0].(boolValue).xor(exps[1].(boolValue))
	return cmp, nil
}

func NewXor(left, right Expr) Expr {
	return xorOp{left, right}
}
