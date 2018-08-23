package expr

import (
	"fmt"
	"math"
)

type logOp struct {
	operand Expr
}

func (op logOp) Evaluate(table *SymTab) (Expr, error) {
	e, err := evalIfArithmetic(table, op.operand)
	if err != nil {
		return nil, err
	}
	result := math.Log10(e.floatValue())
	if math.IsNaN(result) {
		return nil, DomainError
	}
	return realValue{result}, nil
}

// NewLog creates a new base-10 logarithm function expression.
func NewLog(operand Expr) Expr {
	return logOp{operand}
}

func (op logOp) String() string {
	return fmt.Sprintf("log(%v)", op.operand)
}
