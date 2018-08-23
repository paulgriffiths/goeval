package expr

import "fmt"

type powOp struct {
	base, exponent Expr
}

func (op powOp) Evaluate(table *SymTab) (Expr, error) {
	exps, err := evaluateExprs(table, IsNumeric, op.base, op.exponent)
	if err != nil {
		return nil, err
	}
	return exps[0].(arithmeticValue).pow(exps[1].(arithmeticValue))
}

// NewPow creates a new exponentiation operator expression.
func NewPow(base, exponent Expr) Expr {
	return powOp{base, exponent}
}

func (op powOp) String() string {
	return fmt.Sprintf("(%v)^(%v)", op.base, op.exponent)
}
