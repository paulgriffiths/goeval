package expr

type equalityOp struct {
	left, right Expr
}

func (op equalityOp) Evaluate(table *symTab) (Expr, error) {
	exps, err := evaluateExprs(table, IsBoolean, op.left, op.right)
	if err == nil {
		cmp := exps[0].(boolValue).equality(exps[1].(boolValue))
		return cmp, nil
	}
	exps, err = evaluateExprs(table, IsString, op.left, op.right)
	if err == nil {
		cmp := exps[0].(boolValue).equality(exps[1].(boolValue))
		return cmp, nil
	}
	exps, err = evaluateExprs(table, IsNumeric, op.left, op.right)
	if err == nil {
		cmp := exps[0].(arithmeticValue).equality(exps[1].(arithmeticValue))
		return boolValue{cmp}, nil
	}
	return nil, TypeError
}

/*func (op equalityOp) Evaluate(table *symTab) (Expr, error) {
	exps, err := evaluateExprs(table, IsNumeric, op.left, op.right)
	if err != nil {
		return nil, err
	}
	cmp := exps[0].(arithmeticValue).equality(exps[1].(arithmeticValue))
	return boolValue{cmp}, nil
}*/

func NewEquality(left, right Expr) equalityOp {
	return equalityOp{left, right}
}

type nonEqualityOp struct {
	left, right Expr
}

func (op nonEqualityOp) Evaluate(table *symTab) (Expr, error) {
	exps, err := evaluateExprs(table, IsNumeric, op.left, op.right)
	if err != nil {
		return nil, err
	}
	cmp := !exps[0].(arithmeticValue).equality(exps[1].(arithmeticValue))
	return boolValue{cmp}, nil
}

func NewNonEquality(left, right Expr) nonEqualityOp {
	return nonEqualityOp{left, right}
}

type lessThanOp struct {
	left, right Expr
}

func (op lessThanOp) Evaluate(table *symTab) (Expr, error) {
	exps, err := evaluateExprs(table, IsNumeric, op.left, op.right)
	if err != nil {
		return nil, err
	}
	cmp := exps[0].(arithmeticValue).lessThan(exps[1].(arithmeticValue))
	return boolValue{cmp}, nil
}

func NewLessThan(left, right Expr) lessThanOp {
	return lessThanOp{left, right}
}

type greaterThanOp struct {
	left, right Expr
}

func (op greaterThanOp) Evaluate(table *symTab) (Expr, error) {
	exps, err := evaluateExprs(table, IsNumeric, op.left, op.right)
	if err != nil {
		return nil, err
	}
	lt := exps[0].(arithmeticValue).lessThan(exps[1].(arithmeticValue))
	eq := exps[0].(arithmeticValue).equality(exps[1].(arithmeticValue))
	return boolValue{!lt && !eq}, nil
}

func NewGreaterThan(left, right Expr) greaterThanOp {
	return greaterThanOp{left, right}
}

type lessThanOrEqualOp struct {
	left, right Expr
}

func (op lessThanOrEqualOp) Evaluate(table *symTab) (Expr, error) {
	exps, err := evaluateExprs(table, IsNumeric, op.left, op.right)
	if err != nil {
		return nil, err
	}
	lt := exps[0].(arithmeticValue).lessThan(exps[1].(arithmeticValue))
	eq := exps[0].(arithmeticValue).equality(exps[1].(arithmeticValue))
	return boolValue{lt || eq}, nil
}

func NewLessThanOrEqual(left, right Expr) lessThanOrEqualOp {
	return lessThanOrEqualOp{left, right}
}

type greaterThanOrEqualOp struct {
	left, right Expr
}

func (op greaterThanOrEqualOp) Evaluate(table *symTab) (Expr, error) {
	exps, err := evaluateExprs(table, IsNumeric, op.left, op.right)
	if err != nil {
		return nil, err
	}
	lt := exps[0].(arithmeticValue).lessThan(exps[1].(arithmeticValue))
	eq := exps[0].(arithmeticValue).equality(exps[1].(arithmeticValue))
	return boolValue{!lt || eq}, nil
}

func NewGreaterThanOrEqual(left, right Expr) greaterThanOrEqualOp {
	return greaterThanOrEqualOp{left, right}
}
