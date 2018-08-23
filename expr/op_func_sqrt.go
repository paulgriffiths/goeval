package expr

import (
	"fmt"
	"math"
)

type sqrtOp struct {
	operand Expr
}

func (op sqrtOp) Evaluate(table *SymTab) (Expr, error) {
	e, err := evalIfArithmetic(table, op.operand)
	if err != nil {
		return nil, err
	}
	result := math.Sqrt(e.floatValue())
	if math.IsNaN(result) {
		return nil, DomainError
	}
	return realValue{result}, nil
}

// NewSqrt creates a new square root function expression.
func NewSqrt(operand Expr) Expr {
	return sqrtOp{operand}
}

func (op sqrtOp) String() string {
	return fmt.Sprintf("sqrt(%v)", op.operand)
}
