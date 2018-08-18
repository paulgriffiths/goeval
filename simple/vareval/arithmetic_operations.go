package vareval

import "math"

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

type subOp struct {
	left, right expr
}

func (op subOp) evaluate(table *symTab) (expr, error) {
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
		diff := mustIntegerValue(left) - mustIntegerValue(right)
		return intValue{diff}, nil
	}

	diff := mustRealValue(left) - mustRealValue(right)
	return realValue{diff}, nil
}

type mulOp struct {
	left, right expr
}

func (op mulOp) evaluate(table *symTab) (expr, error) {
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
		prod := mustIntegerValue(left) * mustIntegerValue(right)
		return intValue{prod}, nil
	}

	prod := mustRealValue(left) * mustRealValue(right)
	return realValue{prod}, nil
}

type divOp struct {
	left, right expr
}

func (op divOp) evaluate(table *symTab) (expr, error) {
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
		rVal := mustIntegerValue(right)
		if rVal == 0 {
			return nil, DivideByZeroError
		}
		lVal := mustIntegerValue(left)
		if lVal%rVal == 0 {
			return intValue{lVal / rVal}, nil
		}
	}

	rVal := mustRealValue(right)
	if rVal == 0.0 || rVal == -0.0 {
		return nil, DivideByZeroError
	}
	return realValue{mustRealValue(left) / rVal}, nil
}

type powOp struct {
	base, exponent expr
}

func (op powOp) evaluate(table *symTab) (expr, error) {
	base, err := op.base.evaluate(table)
	if err != nil {
		return nil, err
	}
	exponent, err := op.exponent.evaluate(table)
	if err != nil {
		return nil, err
	}

	if !areNumeric(base, exponent) {
		return nil, TypeError
	}

	prod := math.Pow(mustRealValue(base), mustRealValue(exponent))
	if math.IsNaN(prod) {
		return nil, DomainError
	}
	return realValue{prod}, nil
}
