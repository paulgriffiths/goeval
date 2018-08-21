package expr

import (
	"math"
)

type roundOp struct {
	value Expr
}

func (op roundOp) Evaluate(table *SymTab) (Expr, error) {
	exps, err := evaluateExprs(table, IsNumeric, op.value)
	if err != nil {
		return nil, err
	}
	result := math.Round(exps[0].(arithmeticValue).floatValue())
	return realValue{result}, nil
}

func NewRound(value Expr) roundOp {
	return roundOp{value}
}
