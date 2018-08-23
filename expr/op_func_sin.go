package expr

import (
	"fmt"
	"math"
)

type sinOp struct {
	degrees Expr
}

func (op sinOp) Evaluate(table *SymTab) (Expr, error) {
	e, err := evalIfArithmetic(table, op.degrees)
	if err != nil {
		return nil, err
	}
	result := math.Sin(toRadians(e.floatValue()))
	if math.IsNaN(result) {
		return nil, DomainError
	}
	return realValue{result}, nil
}

// NewSin creates a new sine function expression.
func NewSin(degrees Expr) Expr {
	return sinOp{degrees}
}

func (op sinOp) String() string {
	return fmt.Sprintf("sin(%v)", op.degrees)
}
