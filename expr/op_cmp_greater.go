package expr

import "fmt"

type greaterThanOp struct {
	left, right Expr
}

func (op greaterThanOp) Evaluate(table *SymTab) (Expr, error) {
	exps, err := evaluateExprs(table, IsNumeric, op.left, op.right)
	if err != nil {
		return nil, err
	}
	lt := exps[0].(arithmeticValue).lessThan(exps[1].(arithmeticValue))
	eq := exps[0].(arithmeticValue).equality(exps[1].(arithmeticValue))
	return boolValue{!lt && !eq}, nil
}

// NewGreaterThan creates a new conditional greater-than
// operator expression.
func NewGreaterThan(left, right Expr) Expr {
	return greaterThanOp{left, right}
}

func (op greaterThanOp) String() string {
	return fmt.Sprintf("(%v)>(%v)", op.left, op.right)
}
