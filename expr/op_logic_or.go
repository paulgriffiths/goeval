package expr

import "fmt"

type orOp struct {
	left, right Expr
}

func (op orOp) Evaluate(table *SymTab) (Expr, error) {
	exps, err := evaluateExprs(table, IsBoolean, op.left, op.right)
	if err != nil {
		return nil, err
	}
	cmp := exps[0].(boolValue).or(exps[1].(boolValue))
	return cmp, nil
}

// NewOr creates a new logical or operator expression.
func NewOr(left, right Expr) Expr {
	return orOp{left, right}
}

func (op orOp) String() string {
	return fmt.Sprintf("or(%v,%v)", op.left, op.right)
}
