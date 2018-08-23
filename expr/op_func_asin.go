package expr

import (
	"fmt"
	"math"
)

type asinOp struct {
	sine Expr
}

func (op asinOp) Evaluate(table *SymTab) (Expr, error) {
	e, err := evalIfArithmetic(table, op.sine)
	if err != nil {
		return nil, err
	}
	result := toDegrees(math.Asin(e.floatValue()))
	if math.IsNaN(result) {
		return nil, DomainError
	}
	return realValue{result}, nil
}

// NewAsin creates a new arc sine function expression.
func NewAsin(sine Expr) Expr {
	return asinOp{sine}
}

func (op asinOp) String() string {
	return fmt.Sprintf("arcsin(%v)", op.sine)
}
