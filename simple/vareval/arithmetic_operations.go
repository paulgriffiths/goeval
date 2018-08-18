package vareval

type addOp struct {
	left, right expr
}

func (op addOp) evaluate(table *symTab) (expr, error) {
	left, err := op.left.evaluate(table)
	if err != nil {
		return nil, err
	}
	right, err := op.right.evaluate(table)
	if err != nil {
		return nil, err
	}

	if !areNumeric(left, right) {
		return nil, TypeError
	}

	if areInteger(left, right) {
		sum := mustIntegerValue(left) + mustIntegerValue(right)
		return intValue{sum}, nil
	}

	sum := mustRealValue(left) + mustRealValue(right)
	return realValue{sum}, nil
}
