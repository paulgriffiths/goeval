package expr

import "fmt"

type norOp struct {
	left, right Expr
}

func (op norOp) Evaluate(table *SymTab) (Expr, error) {
	exps, err := evaluateExprs(table, IsBoolean, op.left, op.right)
	if err != nil {
		return nil, err
	}
	cmp := exps[0].(boolValue).nor(exps[1].(boolValue))
	return cmp, nil
}

// NewNor creates a new logical nor operator expression.
func NewNor(left, right Expr) Expr {
	return norOp{left, right}
}

func (op norOp) String() string {
	return fmt.Sprintf("nor(%v,%v)", op.left, op.right)
}
