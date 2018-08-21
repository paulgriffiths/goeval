package expr

import (
	"math"
)

type asinOp struct {
	value Expr
}

func (op asinOp) Evaluate(table *SymTab) (Expr, error) {
	exps, err := evaluateExprs(table, IsNumeric, op.value)
	if err != nil {
		return nil, err
	}
	result := toDegrees(math.Asin(exps[0].(arithmeticValue).floatValue()))
	if math.IsNaN(result) {
		return nil, DomainError
	}
	return realValue{result}, nil
}

func NewAsin(operand Expr) Expr {
	return asinOp{operand}
}
