package expr

import "fmt"

type nonEqualityOp struct {
	left, right Expr
}

func (op nonEqualityOp) Evaluate(table *SymTab) (Expr, error) {
	lb, rb, err := evalPairIfBoolean(table, op.left, op.right)
	if err == nil {
		return boolValue{!lb.equality(rb)}, nil
	}
    ls, rs, err := evalPairIfString(table, op.left, op.right)
	if err == nil {
		return boolValue{!ls.equality(rs)}, nil
	}
    ln, rn, err := evalPairIfArithmetic(table, op.left, op.right)
	if err != nil {
        return nil, TypeError
	}
    return boolValue{!ln.equality(rn)}, nil
}

// NewNonEquality creates a new conditional non-equality
// operator expression.
func NewNonEquality(left, right Expr) Expr {
	return nonEqualityOp{left, right}
}

func (op nonEqualityOp) String() string {
	return fmt.Sprintf("(%v)!=(%v)", op.left, op.right)
}
