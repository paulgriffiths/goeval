package expr

import "fmt"

type subOp struct {
	left, right Expr
}

func (op subOp) Evaluate(table *SymTab) (Expr, error) {
	l, r, err := evalPairIfArithmetic(table, op.left, op.right)
	if err != nil {
		return nil, err
	}
	return l.sub(r), nil
}

// NewSub creates a new subtraction operator expression.
func NewSub(left, right Expr) Expr {
	return subOp{left, right}
}

func (op subOp) String() string {
	return fmt.Sprintf("(%v)-(%v)", op.left, op.right)
}
