package vareval

type addOp struct {
	left, right expr
}

func (op addOp) evaluate(table *symTab) (expr, error) {
	exps, err := evaluateExprs(table, isNumeric, op.left, op.right)
	if err != nil {
		return nil, err
	}
	return exps[0].(arithmeticValue).add(exps[1].(arithmeticValue)), nil
}

type subOp struct {
	left, right expr
}

func (op subOp) evaluate(table *symTab) (expr, error) {
	exps, err := evaluateExprs(table, isNumeric, op.left, op.right)
	if err != nil {
		return nil, err
	}
	return exps[0].(arithmeticValue).sub(exps[1].(arithmeticValue)), nil
}

type mulOp struct {
	left, right expr
}

func (op mulOp) evaluate(table *symTab) (expr, error) {
	exps, err := evaluateExprs(table, isNumeric, op.left, op.right)
	if err != nil {
		return nil, err
	}
	return exps[0].(arithmeticValue).mul(exps[1].(arithmeticValue)), nil
}

type divOp struct {
	left, right expr
}

func (op divOp) evaluate(table *symTab) (expr, error) {
	exps, err := evaluateExprs(table, isNumeric, op.left, op.right)
	if err != nil {
		return nil, err
	}
	return exps[0].(arithmeticValue).div(exps[1].(arithmeticValue))
}

type powOp struct {
	base, exponent expr
}

func (op powOp) evaluate(table *symTab) (expr, error) {
	exps, err := evaluateExprs(table, isNumeric, op.base, op.exponent)
	if err != nil {
		return nil, err
	}
	return exps[0].(arithmeticValue).pow(exps[1].(arithmeticValue))
}
