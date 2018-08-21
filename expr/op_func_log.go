package expr

import (
	"math"
)

type logOp struct {
	value Expr
}

func (op logOp) Evaluate(table *SymTab) (Expr, error) {
	exps, err := evaluateExprs(table, IsNumeric, op.value)
	if err != nil {
		return nil, err
	}
	result := math.Log10(exps[0].(arithmeticValue).floatValue())
	if math.IsNaN(result) {
		return nil, DomainError
	}
	return realValue{result}, nil
}

func NewLog(operand Expr) Expr {
	return logOp{operand}
}
