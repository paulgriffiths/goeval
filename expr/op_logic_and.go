package expr

type andOp struct {
	left, right Expr
}

func (op andOp) Evaluate(table *SymTab) (Expr, error) {
	exps, err := evaluateExprs(table, IsBoolean, op.left, op.right)
	if err != nil {
		return nil, err
	}
	cmp := exps[0].(boolValue).and(exps[1].(boolValue))
	return cmp, nil
}

func NewAnd(left, right Expr) andOp {
	return andOp{left, right}
}
