package expr

import "fmt"

type notOp struct {
	operand Expr
}

func (op notOp) Evaluate(table *SymTab) (Expr, error) {
	exps, err := evaluateExprs(table, IsBoolean, op.operand)
	if err != nil {
		return nil, err
	}
	cmp := exps[0].(boolValue).not()
	return cmp, nil
}

// NewNot creates a new logical not operator expression.
func NewNot(operand Expr) Expr {
	return notOp{operand}
}

func (op notOp) String() string {
	return fmt.Sprintf("not(%v)", op.operand)
}
