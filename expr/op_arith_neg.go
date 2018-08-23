package expr

import "fmt"

type negOp struct {
	value Expr
}

func (op negOp) Evaluate(table *SymTab) (Expr, error) {
	e, err := evalIfArithmetic(table, op.value)
	if err != nil {
		return nil, err
	}
	return e.negate(), nil
}

// NewNeg creates a new negation operator expression.
func NewNeg(value Expr) Expr {
	return negOp{value}
}

func (op negOp) String() string {
	return fmt.Sprintf("-(%v)", op.value)
}
