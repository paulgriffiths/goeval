package expr

import (
	"math"
)

type sinOp struct {
	value Expr
}

func (op sinOp) Evaluate(table *SymTab) (Expr, error) {
	exps, err := evaluateExprs(table, IsNumeric, op.value)
	if err != nil {
		return nil, err
	}
	result := math.Sin(toRadians(exps[0].(arithmeticValue).floatValue()))
	if math.IsNaN(result) {
		return nil, DomainError
	}
	return realValue{result}, nil
}

func NewSin(operand Expr) Expr {
	return sinOp{operand}
}
