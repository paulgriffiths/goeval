package expr

import "fmt"

type lessThanOp struct {
	left, right Expr
}

func (op lessThanOp) Evaluate(table *SymTab) (Expr, error) {
	exps, err := evaluateExprs(table, IsNumeric, op.left, op.right)
	if err != nil {
		return nil, err
	}
	cmp := exps[0].(arithmeticValue).lessThan(exps[1].(arithmeticValue))
	return boolValue{cmp}, nil
}

// NewLessThan creates a new conditional less-than operator expression.
func NewLessThan(left, right Expr) Expr {
	return lessThanOp{left, right}
}

func (op lessThanOp) String() string {
	return fmt.Sprintf("(%v)<(%v)", op.left, op.right)
}
