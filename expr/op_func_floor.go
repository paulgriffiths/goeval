package expr

import (
	"fmt"
	"math"
)

type floorOp struct {
	operand Expr
}

func (op floorOp) Evaluate(table *SymTab) (Expr, error) {
	e, err := evalIfArithmetic(table, op.operand)
	if err != nil {
		return nil, err
	}
	result := math.Floor(e.floatValue())
	return realValue{result}, nil
}

// NewFloor creates a new floor function expression.
func NewFloor(operand Expr) Expr {
	return floorOp{operand}
}

func (op floorOp) String() string {
	return fmt.Sprintf("floor(%v)", op.operand)
}
