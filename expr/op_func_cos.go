package expr

import (
	"fmt"
	"math"
)

type cosOp struct {
	operand Expr
}

func (op cosOp) Evaluate(table *SymTab) (Expr, error) {
	exps, err := evaluateExprs(table, IsNumeric, op.operand)
	if err != nil {
		return nil, err
	}
	result := math.Cos(toRadians(exps[0].(arithmeticValue).floatValue()))
	if math.IsNaN(result) {
		return nil, DomainError
	}
	return realValue{result}, nil
}

// NewCos creates a new cosine function expression.
func NewCos(operand Expr) Expr {
	return cosOp{operand}
}

func (op cosOp) String() string {
	return fmt.Sprintf("cos(%v)", op.operand)
}
