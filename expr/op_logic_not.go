package expr

import "fmt"

type notOp struct {
	operand Expr
}

func (op notOp) Evaluate(table *SymTab) (Expr, error) {
	e, err := evalIfBoolean(table, op.operand)
	if err != nil {
		return nil, err
	}
	return e.not(), nil
}

// NewNot creates a new logical not operator expression.
func NewNot(operand Expr) Expr {
	return notOp{operand}
}

func (op notOp) String() string {
	return fmt.Sprintf("not(%v)", op.operand)
}
