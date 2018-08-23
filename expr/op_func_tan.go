package expr

import (
	"fmt"
	"math"
)

type tanOp struct {
	operand Expr
}

func (op tanOp) Evaluate(table *SymTab) (Expr, error) {
	exps, err := evaluateExprs(table, IsNumeric, op.operand)
	if err != nil {
		return nil, err
	}
	result := math.Tan(toRadians(exps[0].(arithmeticValue).floatValue()))
	if math.IsNaN(result) {
		return nil, DomainError
	}
	return realValue{result}, nil
}

// NewTan creates a new tangent function expression.
func NewTan(operand Expr) Expr {
	return tanOp{operand}
}

func (op tanOp) String() string {
	return fmt.Sprintf("tan(%v)", op.operand)
}
