package expr

import (
	"fmt"
	"math"
)

type lnOp struct {
	operand Expr
}

func (op lnOp) Evaluate(table *SymTab) (Expr, error) {
	e, err := evalIfArithmetic(table, op.operand)
	if err != nil {
		return nil, err
	}
	result := math.Log(e.floatValue())
	if math.IsNaN(result) {
		return nil, DomainError
	}
	return realValue{result}, nil
}

// NewLn creates a new natural logarithm function expression.
func NewLn(operand Expr) Expr {
	return lnOp{operand}
}

func (op lnOp) String() string {
	return fmt.Sprintf("ln(%v)", op.operand)
}
