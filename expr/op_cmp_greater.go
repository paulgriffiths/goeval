package expr

import "fmt"

type greaterThanOp struct {
	left, right Expr
}

func (op greaterThanOp) Evaluate(table *SymTab) (Expr, error) {
	l, r, err := evalPairIfArithmetic(table, op.left, op.right)
	if err != nil {
		return nil, err
	}
	return boolValue{!l.lessThan(r) && !l.equality(r)}, nil
}

// NewGreaterThan creates a new conditional greater-than
// operator expression.
func NewGreaterThan(left, right Expr) Expr {
	return greaterThanOp{left, right}
}

func (op greaterThanOp) String() string {
	return fmt.Sprintf("(%v)>(%v)", op.left, op.right)
}
