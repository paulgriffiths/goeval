package expr

import (
	"math"
)

type atanOp struct {
	value Expr
}

func (op atanOp) Evaluate(table *SymTab) (Expr, error) {
	exps, err := evaluateExprs(table, IsNumeric, op.value)
	if err != nil {
		return nil, err
	}
	result := toDegrees(math.Atan(exps[0].(arithmeticValue).floatValue()))
	if math.IsNaN(result) {
		return nil, DomainError
	}
	return realValue{result}, nil
}

func NewAtan(value Expr) atanOp {
	return atanOp{value}
}
