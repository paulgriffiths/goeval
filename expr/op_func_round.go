package expr

import (
	"fmt"
	"math"
)

type roundOp struct {
	operand Expr
}

func (op roundOp) Evaluate(table *SymTab) (Expr, error) {
	e, err := evalIfArithmetic(table, op.operand)
	if err != nil {
		return nil, err
	}
	result := math.Round(e.floatValue())
	return realValue{result}, nil
}

// NewRound creates a new rounding function expression.
func NewRound(operand Expr) Expr {
	return roundOp{operand}
}

func (op roundOp) String() string {
	return fmt.Sprintf("round(%v)", op.operand)
}
