package expr

import (
	"math"
)

type ceilOp struct {
	value Expr
}

func (op ceilOp) Evaluate(table *SymTab) (Expr, error) {
	exps, err := evaluateExprs(table, IsNumeric, op.value)
	if err != nil {
		return nil, err
	}
	result := math.Ceil(exps[0].(arithmeticValue).floatValue())
	return realValue{result}, nil
}

func NewCeil(operand Expr) Expr {
	return ceilOp{operand}
}
