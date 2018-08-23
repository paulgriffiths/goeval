package expr

import (
	"fmt"
	"math"
)

type roundOp struct {
	operand Expr
}

func (op roundOp) Evaluate(table *SymTab) (Expr, error) {
	exps, err := evaluateExprs(table, IsNumeric, op.operand)
	if err != nil {
		return nil, err
	}
	result := math.Round(exps[0].(arithmeticValue).floatValue())
	return realValue{result}, nil
}

// NewRound creates a new rounding function expression.
func NewRound(operand Expr) Expr {
	return roundOp{operand}
}

func (op roundOp) String() string {
	return fmt.Sprintf("round(%v)", op.operand)
}
