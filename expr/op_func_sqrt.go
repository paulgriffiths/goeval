package expr

import (
	"math"
)

type sqrtOp struct {
	value Expr
}

func (op sqrtOp) Evaluate(table *SymTab) (Expr, error) {
	exps, err := evaluateExprs(table, IsNumeric, op.value)
	if err != nil {
		return nil, err
	}
	result := math.Sqrt(exps[0].(arithmeticValue).floatValue())
	if math.IsNaN(result) {
		return nil, DomainError
	}
	return realValue{result}, nil
}

func NewSqrt(value Expr) sqrtOp {
	return sqrtOp{value}
}
