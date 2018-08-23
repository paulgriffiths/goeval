package expr

import (
	"fmt"
	"math"
)

type cosOp struct {
	degrees Expr
}

func (op cosOp) Evaluate(table *SymTab) (Expr, error) {
	e, err := evalIfArithmetic(table, op.degrees)
	if err != nil {
		return nil, err
	}
	result := math.Cos(toRadians(e.floatValue()))
	if math.IsNaN(result) {
		return nil, DomainError
	}
	return realValue{result}, nil
}

// NewCos creates a new cosine function expression.
func NewCos(degrees Expr) Expr {
	return cosOp{degrees}
}

func (op cosOp) String() string {
	return fmt.Sprintf("cos(%v)", op.degrees)
}
