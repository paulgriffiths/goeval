package expr

import "fmt"

type lessThanOrEqualOp struct {
	left, right Expr
}

func (op lessThanOrEqualOp) Evaluate(table *SymTab) (Expr, error) {
	exps, err := evaluateExprs(table, IsNumeric, op.left, op.right)
	if err != nil {
		return nil, err
	}
	lt := exps[0].(arithmeticValue).lessThan(exps[1].(arithmeticValue))
	eq := exps[0].(arithmeticValue).equality(exps[1].(arithmeticValue))
	return boolValue{lt || eq}, nil
}

// NewLessThanOrEqual creates a new conditional less-than-or-equal-to
// operator expression.
func NewLessThanOrEqual(left, right Expr) Expr {
	return lessThanOrEqualOp{left, right}
}

func (op lessThanOrEqualOp) String() string {
	return fmt.Sprintf("(%v)<=(%v)", op.left, op.right)
}
