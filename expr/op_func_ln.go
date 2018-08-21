package expr

import (
	"math"
)

type lnOp struct {
	value Expr
}

func (op lnOp) Evaluate(table *SymTab) (Expr, error) {
	exps, err := evaluateExprs(table, IsNumeric, op.value)
	if err != nil {
		return nil, err
	}
	result := math.Log(exps[0].(arithmeticValue).floatValue())
	if math.IsNaN(result) {
		return nil, DomainError
	}
	return realValue{result}, nil
}

func NewLn(operand Expr) Expr {
	return lnOp{operand}
}
