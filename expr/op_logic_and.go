package expr

import "fmt"

type andOp struct {
	left, right Expr
}

func (op andOp) Evaluate(table *SymTab) (Expr, error) {
	l, r, err := evalPairIfBoolean(table, op.left, op.right)
	if err != nil {
		return nil, err
	}
	return l.and(r), nil
}

// NewAnd creates a new logical and operator expression.
func NewAnd(left, right Expr) Expr {
	return andOp{left, right}
}

func (op andOp) String() string {
	return fmt.Sprintf("and(%v,%v)", op.left, op.right)
}
