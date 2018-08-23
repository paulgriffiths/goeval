package expr

import "fmt"

type mulOp struct {
	left, right Expr
}

func (op mulOp) Evaluate(table *SymTab) (Expr, error) {
	l, r, err := evalPairIfArithmetic(table, op.left, op.right)
	if err != nil {
		return nil, err
	}
	return l.mul(r), nil
}

// NewMul creates a new multiplication operator expression.
func NewMul(left, right Expr) Expr {
	return mulOp{left, right}
}

func (op mulOp) String() string {
	return fmt.Sprintf("(%v)*(%v)", op.left, op.right)
}
