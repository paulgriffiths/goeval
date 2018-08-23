package expr

import "fmt"

type negOp struct {
	value Expr
}

func (op negOp) Evaluate(table *SymTab) (Expr, error) {
	exps, err := evaluateExprs(table, IsNumeric, op.value)
	if err != nil {
		return nil, err
	}
	return exps[0].(arithmeticValue).negate(), nil
}

// NewNeg creates a new negation operator expression.
func NewNeg(value Expr) Expr {
	return negOp{value}
}

func (op negOp) String() string {
	return fmt.Sprintf("-(%v)", op.value)
}
