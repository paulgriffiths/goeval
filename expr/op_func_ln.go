package expr

import (
	"fmt"
	"math"
)

type lnOp struct {
	operand Expr
}

func (op lnOp) Evaluate(table *SymTab) (Expr, error) {
	exps, err := evaluateExprs(table, IsNumeric, op.operand)
	if err != nil {
		return nil, err
	}
	result := math.Log(exps[0].(arithmeticValue).floatValue())
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
