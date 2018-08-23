package expr

import "fmt"

type nandOp struct {
	left, right Expr
}

func (op nandOp) Evaluate(table *SymTab) (Expr, error) {
	exps, err := evaluateExprs(table, IsBoolean, op.left, op.right)
	if err != nil {
		return nil, err
	}
	cmp := exps[0].(boolValue).nand(exps[1].(boolValue))
	return cmp, nil
}

// NewNand creates a new logical nand operator expression.
func NewNand(left, right Expr) Expr {
	return nandOp{left, right}
}

func (op nandOp) String() string {
	return fmt.Sprintf("nand(%v,%v)", op.left, op.right)
}
