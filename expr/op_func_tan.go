package expr

import (
	"fmt"
	"math"
)

type tanOp struct {
	degrees Expr
}

func (op tanOp) Evaluate(table *SymTab) (Expr, error) {
	e, err := evalIfArithmetic(table, op.degrees)
	if err != nil {
		return nil, err
	}
	result := math.Tan(toRadians(e.floatValue()))
	if math.IsNaN(result) {
		return nil, DomainError
	}
	return realValue{result}, nil
}

// NewTan creates a new tangent function expression.
func NewTan(degrees Expr) Expr {
	return tanOp{degrees}
}

func (op tanOp) String() string {
	return fmt.Sprintf("tan(%v)", op.degrees)
}
