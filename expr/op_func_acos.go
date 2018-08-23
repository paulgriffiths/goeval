package expr

import (
	"fmt"
	"math"
)

type acosOp struct {
	cosine Expr
}

func (op acosOp) Evaluate(table *SymTab) (Expr, error) {
	e, err := evalIfArithmetic(table, op.cosine)
	if err != nil {
		return nil, err
	}
	result := toDegrees(math.Acos(e.floatValue()))
	if math.IsNaN(result) {
		return nil, DomainError
	}
	return realValue{result}, nil
}

// NewAcos creates a new arc cosine function expression.
func NewAcos(cosine Expr) Expr {
	return acosOp{cosine}
}

func (op acosOp) String() string {
	return fmt.Sprintf("arccos(%v)", op.cosine)
}
