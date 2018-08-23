package expr

import (
	"fmt"
	"math"
)

type ceilOp struct {
	operand Expr
}

func (op ceilOp) Evaluate(table *SymTab) (Expr, error) {
	e, err := evalIfArithmetic(table, op.operand)
	if err != nil {
		return nil, err
	}
	result := math.Ceil(e.floatValue())
	return realValue{result}, nil
}

// NewCeil creates a new ceiling function expression.
func NewCeil(operand Expr) Expr {
	return ceilOp{operand}
}

func (op ceilOp) String() string {
	return fmt.Sprintf("ceil(%v)", op.operand)
}
