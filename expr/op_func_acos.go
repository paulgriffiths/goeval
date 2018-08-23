package expr

import (
	"fmt"
	"math"
)

type acosOp struct {
	operand Expr
}

func (op acosOp) Evaluate(table *SymTab) (Expr, error) {
	exps, err := evaluateExprs(table, IsNumeric, op.operand)
	if err != nil {
		return nil, err
	}
	result := toDegrees(math.Acos(exps[0].(arithmeticValue).floatValue()))
	if math.IsNaN(result) {
		return nil, DomainError
	}
	return realValue{result}, nil
}

// NewAcos creates a new arc cosine function expression.
func NewAcos(operand Expr) Expr {
	return acosOp{operand}
}

func (op acosOp) String() string {
	return fmt.Sprintf("arccos(%v)", op.operand)
}
