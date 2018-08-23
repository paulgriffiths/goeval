package expr

import (
	"fmt"
	"math"
)

type atanOp struct {
	operand Expr
}

func (op atanOp) Evaluate(table *SymTab) (Expr, error) {
	exps, err := evaluateExprs(table, IsNumeric, op.operand)
	if err != nil {
		return nil, err
	}
	result := toDegrees(math.Atan(exps[0].(arithmeticValue).floatValue()))
	if math.IsNaN(result) {
		return nil, DomainError
	}
	return realValue{result}, nil
}

// NewAtan creates a new arc tangent function expression.
func NewAtan(operand Expr) Expr {
	return atanOp{operand}
}

func (op atanOp) String() string {
	return fmt.Sprintf("arctan(%v)", op.operand)
}
