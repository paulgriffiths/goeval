package expr

import "fmt"

type xorOp struct {
	left, right Expr
}

func (op xorOp) Evaluate(table *SymTab) (Expr, error) {
	l, r, err := evalPairIfBoolean(table, op.left, op.right)
	if err != nil {
		return nil, err
	}
	return l.xor(r), nil
}

// NewXor creates a new logical exclusive-or operator expression.
func NewXor(left, right Expr) Expr {
	return xorOp{left, right}
}

func (op xorOp) String() string {
	return fmt.Sprintf("xor(%v,%v)", op.left, op.right)
}
