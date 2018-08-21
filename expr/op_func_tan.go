package expr

import (
	"math"
)

type tanOp struct {
	value Expr
}

func (op tanOp) Evaluate(table *SymTab) (Expr, error) {
	exps, err := evaluateExprs(table, IsNumeric, op.value)
	if err != nil {
		return nil, err
	}
	result := math.Tan(toRadians(exps[0].(arithmeticValue).floatValue()))
	if math.IsNaN(result) {
		return nil, DomainError
	}
	return realValue{result}, nil
}

func NewTan(value Expr) tanOp {
	return tanOp{value}
}
