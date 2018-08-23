package expr

import "fmt"

type addOp struct {
	left, right Expr
}

func (op addOp) Evaluate(table *SymTab) (Expr, error) {
	exps, err := evaluateExprs(table, IsNumeric, op.left, op.right)
	if err != nil {
		return nil, err
	}
	return exps[0].(arithmeticValue).add(exps[1].(arithmeticValue)), nil
}

// NewAdd creates a new addition operator expression.
func NewAdd(left, right Expr) Expr {
	return addOp{left, right}
}

func (op addOp) String() string {
	return fmt.Sprintf("(%v)+(%v)", op.left, op.right)
}
