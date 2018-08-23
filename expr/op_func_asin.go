package expr

import (
	"fmt"
	"math"
)

type asinOp struct {
	operand Expr
}

func (op asinOp) Evaluate(table *SymTab) (Expr, error) {
	exps, err := evaluateExprs(table, IsNumeric, op.operand)
	if err != nil {
		return nil, err
	}
	result := toDegrees(math.Asin(exps[0].(arithmeticValue).floatValue()))
	if math.IsNaN(result) {
		return nil, DomainError
	}
	return realValue{result}, nil
}

// NewAsin creates a new arc sine function expression.
func NewAsin(operand Expr) Expr {
	return asinOp{operand}
}

func (op asinOp) String() string {
	return fmt.Sprintf("arcsin(%v)", op.operand)
}
