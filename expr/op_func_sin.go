package expr

import (
	"fmt"
	"math"
)

type sinOp struct {
	operand Expr
}

func (op sinOp) Evaluate(table *SymTab) (Expr, error) {
	exps, err := evaluateExprs(table, IsNumeric, op.operand)
	if err != nil {
		return nil, err
	}
	result := math.Sin(toRadians(exps[0].(arithmeticValue).floatValue()))
	if math.IsNaN(result) {
		return nil, DomainError
	}
	return realValue{result}, nil
}

// NewSin creates a new sine function expression.
func NewSin(operand Expr) Expr {
	return sinOp{operand}
}

func (op sinOp) String() string {
	return fmt.Sprintf("sin(%v)", op.operand)
}
