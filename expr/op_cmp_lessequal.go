package expr

import "fmt"

type lessThanOrEqualOp struct {
	left, right Expr
}

func (op lessThanOrEqualOp) Evaluate(table *SymTab) (Expr, error) {
	l, r, err := evalPairIfArithmetic(table, op.left, op.right)
	if err != nil {
		return nil, err
	}
	return boolValue{l.lessThan(r) || l.equality(r)}, nil
}

// NewLessThanOrEqual creates a new conditional less-than-or-equal-to
// operator expression.
func NewLessThanOrEqual(left, right Expr) Expr {
	return lessThanOrEqualOp{left, right}
}

func (op lessThanOrEqualOp) String() string {
	return fmt.Sprintf("(%v)<=(%v)", op.left, op.right)
}
