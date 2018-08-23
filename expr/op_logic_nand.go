package expr

import "fmt"

type nandOp struct {
	left, right Expr
}

func (op nandOp) Evaluate(table *SymTab) (Expr, error) {
	l, r, err := evalPairIfBoolean(table, op.left, op.right)
	if err != nil {
		return nil, err
	}
	return l.nand(r), nil
}

// NewNand creates a new logical nand operator expression.
func NewNand(left, right Expr) Expr {
	return nandOp{left, right}
}

func (op nandOp) String() string {
	return fmt.Sprintf("nand(%v,%v)", op.left, op.right)
}
