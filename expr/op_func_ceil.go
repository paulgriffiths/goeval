package expr

import (
	"fmt"
	"math"
)

type ceilOp struct {
	operand Expr
}

func (op ceilOp) Evaluate(table *SymTab) (Expr, error) {
	exps, err := evaluateExprs(table, IsNumeric, op.operand)
	if err != nil {
		return nil, err
	}
	result := math.Ceil(exps[0].(arithmeticValue).floatValue())
	return realValue{result}, nil
}

// NewCeil creates a new ceiling function expression.
func NewCeil(operand Expr) Expr {
	return ceilOp{operand}
}

func (op ceilOp) String() string {
	return fmt.Sprintf("ceil(%v)", op.operand)
}
