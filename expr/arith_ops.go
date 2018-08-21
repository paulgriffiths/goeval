package expr

type addOp struct {
	left, right Expr
}

func (op addOp) Evaluate(table *SymTab) (Expr, error) {
	exps, err := evaluateExprs(table, IsNumeric, op.left, op.right)
	if err != nil {
		return nil, err
	}
	return exps[0].(arithmeticValue).add(exps[1].(arithmeticValue)), nil
}

func NewAdd(left, right Expr) addOp {
	return addOp{left, right}
}

type subOp struct {
	left, right Expr
}

func (op subOp) Evaluate(table *SymTab) (Expr, error) {
	exps, err := evaluateExprs(table, IsNumeric, op.left, op.right)
	if err != nil {
		return nil, err
	}
	return exps[0].(arithmeticValue).sub(exps[1].(arithmeticValue)), nil
}

func NewSub(left, right Expr) subOp {
	return subOp{left, right}
}

type mulOp struct {
	left, right Expr
}

func (op mulOp) Evaluate(table *SymTab) (Expr, error) {
	exps, err := evaluateExprs(table, IsNumeric, op.left, op.right)
	if err != nil {
		return nil, err
	}
	return exps[0].(arithmeticValue).mul(exps[1].(arithmeticValue)), nil
}

func NewMul(left, right Expr) mulOp {
	return mulOp{left, right}
}

type divOp struct {
	left, right Expr
}

func (op divOp) Evaluate(table *SymTab) (Expr, error) {
	exps, err := evaluateExprs(table, IsNumeric, op.left, op.right)
	if err != nil {
		return nil, err
	}
	return exps[0].(arithmeticValue).div(exps[1].(arithmeticValue))
}

func NewDiv(left, right Expr) divOp {
	return divOp{left, right}
}

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

func NewPow(base, exponent Expr) powOp {
	return powOp{base, exponent}
}

type negOp struct {
	value Expr
}

func (op negOp) Evaluate(table *SymTab) (Expr, error) {
	exps, err := evaluateExprs(table, IsNumeric, op.value)
	if err != nil {
		return nil, err
	}
	return exps[0].(arithmeticValue).negate(), nil
}

func NewNeg(value Expr) negOp {
	return negOp{value}
}
