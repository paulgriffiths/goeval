package expr

import "fmt"

type greaterThanOrEqualOp struct {
	left, right Expr
}

func (op greaterThanOrEqualOp) Evaluate(table *SymTab) (Expr, error) {
	l, r, err := evalPairIfArithmetic(table, op.left, op.right)
	if err != nil {
		return nil, err
	}
	return boolValue{!l.lessThan(r) || l.equality(r)}, nil
}

// NewGreaterThanOrEqual creates a new conditional greater-than-or-equal-to
// operator expression.
func NewGreaterThanOrEqual(left, right Expr) Expr {
	return greaterThanOrEqualOp{left, right}
}

func (op greaterThanOrEqualOp) String() string {
	return fmt.Sprintf("(%v)>=(%v)", op.left, op.right)
}
