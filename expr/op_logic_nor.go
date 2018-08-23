package expr

import "fmt"

type norOp struct {
	left, right Expr
}

func (op norOp) Evaluate(table *SymTab) (Expr, error) {
	l, r, err := evalPairIfBoolean(table, op.left, op.right)
	if err != nil {
		return nil, err
	}
	return l.nor(r), nil
}

// NewNor creates a new logical nor operator expression.
func NewNor(left, right Expr) Expr {
	return norOp{left, right}
}

func (op norOp) String() string {
	return fmt.Sprintf("nor(%v,%v)", op.left, op.right)
}
