package expr

import "fmt"

type concatOp struct {
	left, right Expr
}

func (op concatOp) Evaluate(table *SymTab) (Expr, error) {
	l, r, err := evalPairIfString(table, op.left, op.right)
	if err != nil {
		return nil, err
	}
	return l.concat(r), nil
}

// NewConcat creates a new string concatenation operator expression.
func NewConcat(left, right Expr) Expr {
	return concatOp{left, right}
}

func (op concatOp) String() string {
	return fmt.Sprintf("(%v)+(%v)", op.left, op.right)
}
