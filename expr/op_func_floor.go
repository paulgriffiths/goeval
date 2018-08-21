package expr

import (
	"math"
)

type floorOp struct {
	value Expr
}

func (op floorOp) Evaluate(table *SymTab) (Expr, error) {
	exps, err := evaluateExprs(table, IsNumeric, op.value)
	if err != nil {
		return nil, err
	}
	result := math.Floor(exps[0].(arithmeticValue).floatValue())
	return realValue{result}, nil
}

func NewFloor(operand Expr) Expr {
	return floorOp{operand}
}
