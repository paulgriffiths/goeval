package expr

import "fmt"

type equalityOp struct {
	left, right Expr
}

func (op equalityOp) Evaluate(table *SymTab) (Expr, error) {
	lb, rb, err := evalPairIfBoolean(table, op.left, op.right)
	if err == nil {
		return boolValue{lb.equality(rb)}, nil
	} else if err == UndefinedVariableError {
		return nil, err
	}
	ls, rs, err := evalPairIfString(table, op.left, op.right)
	if err == nil {
		return boolValue{ls.equality(rs)}, nil
	} else if err == UndefinedVariableError {
		return nil, err
	}
	ln, rn, err := evalPairIfArithmetic(table, op.left, op.right)
	if err != nil {
		return nil, err
	}
	return boolValue{ln.equality(rn)}, nil
}

// NewEquality creates a new conditional equality operator expression.
func NewEquality(left, right Expr) Expr {
	return equalityOp{left, right}
}

func (op equalityOp) String() string {
	return fmt.Sprintf("(%v)==(%v)", op.left, op.right)
}
