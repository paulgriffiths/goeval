package expr

import (
	"fmt"
	"math"
)

type atanOp struct {
	tangent Expr
}

func (op atanOp) Evaluate(table *SymTab) (Expr, error) {
	e, err := evalIfArithmetic(table, op.tangent)
	if err != nil {
		return nil, err
	}
	result := toDegrees(math.Atan(e.floatValue()))
	if math.IsNaN(result) {
		return nil, DomainError
	}
	return realValue{result}, nil
}

// NewAtan creates a new arc tangent function expression.
func NewAtan(tangent Expr) Expr {
	return atanOp{tangent}
}

func (op atanOp) String() string {
	return fmt.Sprintf("arctan(%v)", op.tangent)
}
