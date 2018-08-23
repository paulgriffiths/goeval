package expr

import "fmt"

type lessThanOp struct {
	left, right Expr
}

func (op lessThanOp) Evaluate(table *SymTab) (Expr, error) {
	l, r, err := evalPairIfArithmetic(table, op.left, op.right)
	if err != nil {
		return nil, err
	}
	return boolValue{l.lessThan(r)}, nil
}

// NewLessThan creates a new conditional less-than operator expression.
func NewLessThan(left, right Expr) Expr {
	return lessThanOp{left, right}
}

func (op lessThanOp) String() string {
	return fmt.Sprintf("(%v)<(%v)", op.left, op.right)
}
