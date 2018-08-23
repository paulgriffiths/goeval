package expr

import (
	"fmt"
	"math"
)

type floorOp struct {
	operand Expr
}

func (op floorOp) Evaluate(table *SymTab) (Expr, error) {
	exps, err := evaluateExprs(table, IsNumeric, op.operand)
	if err != nil {
		return nil, err
	}
	result := math.Floor(exps[0].(arithmeticValue).floatValue())
	return realValue{result}, nil
}

// NewFloor creates a new floor function expression.
func NewFloor(operand Expr) Expr {
	return floorOp{operand}
}

func (op floorOp) String() string {
	return fmt.Sprintf("floor(%v)", op.operand)
}
