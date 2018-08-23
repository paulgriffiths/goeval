package expr

import "fmt"

type greaterThanOrEqualOp struct {
	left, right Expr
}

func (op greaterThanOrEqualOp) Evaluate(table *SymTab) (Expr, error) {
	exps, err := evaluateExprs(table, IsNumeric, op.left, op.right)
	if err != nil {
		return nil, err
	}
	lt := exps[0].(arithmeticValue).lessThan(exps[1].(arithmeticValue))
	eq := exps[0].(arithmeticValue).equality(exps[1].(arithmeticValue))
	return boolValue{!lt || eq}, nil
}

// NewGreaterThanOrEqual creates a new conditional greater-than-or-equal-to
// operator expression.
func NewGreaterThanOrEqual(left, right Expr) Expr {
	return greaterThanOrEqualOp{left, right}
}

func (op greaterThanOrEqualOp) String() string {
	return fmt.Sprintf("(%v)>=(%v)", op.left, op.right)
}
