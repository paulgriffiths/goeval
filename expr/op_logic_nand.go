package expr

type nandOp struct {
	left, right Expr
}

func (op nandOp) Evaluate(table *SymTab) (Expr, error) {
	exps, err := evaluateExprs(table, IsBoolean, op.left, op.right)
	if err != nil {
		return nil, err
	}
	cmp := exps[0].(boolValue).nand(exps[1].(boolValue))
	return cmp, nil
}

func NewNand(left, right Expr) nandOp {
	return nandOp{left, right}
}
