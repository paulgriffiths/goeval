package expr

import (
	"math"
)

type cosOp struct {
	value Expr
}

func (op cosOp) Evaluate(table *SymTab) (Expr, error) {
	exps, err := evaluateExprs(table, IsNumeric, op.value)
	if err != nil {
		return nil, err
	}
	result := math.Cos(toRadians(exps[0].(arithmeticValue).floatValue()))
	if math.IsNaN(result) {
		return nil, DomainError
	}
	return realValue{result}, nil
}

func NewCos(value Expr) cosOp {
	return cosOp{value}
}
