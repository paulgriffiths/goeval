package expr

import "fmt"

type orOp struct {
	left, right Expr
}

func (op orOp) Evaluate(table *SymTab) (Expr, error) {
	l, r, err := evalPairIfBoolean(table, op.left, op.right)
	if err != nil {
		return nil, err
	}
	return l.or(r), nil
}

// NewOr creates a new logical or operator expression.
func NewOr(left, right Expr) Expr {
	return orOp{left, right}
}

func (op orOp) String() string {
	return fmt.Sprintf("or(%v,%v)", op.left, op.right)
}
