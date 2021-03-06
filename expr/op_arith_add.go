package expr

import "fmt"

type addOp struct {
	left, right Expr
}

func (op addOp) Evaluate(table *SymTab) (Expr, error) {
	ls, rs, err := evalPairIfString(table, op.left, op.right)
	if err == nil {
		return ls.concat(rs), nil
	}
	l, r, err := evalPairIfArithmetic(table, op.left, op.right)
	if err != nil {
		return nil, err
	}
	return l.add(r), nil
}

// NewAdd creates a new addition operator expression.
func NewAdd(left, right Expr) Expr {
	return addOp{left, right}
}

func (op addOp) String() string {
	return fmt.Sprintf("(%v)+(%v)", op.left, op.right)
}
