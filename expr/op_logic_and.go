package expr

import "fmt"

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

// NewAnd creates a new logical and operator expression.
func NewAnd(left, right Expr) Expr {
	return andOp{left, right}
}

func (op andOp) String() string {
	return fmt.Sprintf("and(%v,%v)", op.left, op.right)
}
