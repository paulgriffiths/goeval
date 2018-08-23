package expr

import "fmt"

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

// NewDiv creates a new division operator expression.
func NewDiv(left, right Expr) Expr {
	return divOp{left, right}
}

func (op divOp) String() string {
	return fmt.Sprintf("(%v)/(%v)", op.left, op.right)
}
