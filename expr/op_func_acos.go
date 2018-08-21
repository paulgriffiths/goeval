package expr

import (
	"math"
)

type acosOp struct {
	value Expr
}

func (op acosOp) Evaluate(table *SymTab) (Expr, error) {
	exps, err := evaluateExprs(table, IsNumeric, op.value)
	if err != nil {
		return nil, err
	}
	result := toDegrees(math.Acos(exps[0].(arithmeticValue).floatValue()))
	if math.IsNaN(result) {
		return nil, DomainError
	}
	return realValue{result}, nil
}

func NewAcos(operand Expr) Expr {
	return acosOp{operand}
}
