package expr

type orOp struct {
	left, right Expr
}

func (op orOp) Evaluate(table *SymTab) (Expr, error) {
	exps, err := evaluateExprs(table, IsBoolean, op.left, op.right)
	if err != nil {
		return nil, err
	}
	cmp := exps[0].(boolValue).or(exps[1].(boolValue))
	return cmp, nil
}

func NewOr(left, right Expr) orOp {
	return orOp{left, right}
}
