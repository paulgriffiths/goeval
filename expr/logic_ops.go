package expr

type andOp struct {
	left, right Expr
}

func (op andOp) Evaluate(table *SymTab) (Expr, error) {
	exps, err := evaluateExprs(table, IsBoolean, op.left, op.right)
	if err != nil {
		return nil, err
	}
	cmp := exps[0].(boolValue).and(exps[1].(boolValue))
	return cmp, nil
}

func NewAnd(left, right Expr) andOp {
	return andOp{left, right}
}

type orOp struct {
	left, right Expr
}

func (op orOp) Evaluate(table *SymTab) (Expr, error) {
	exps, err := evaluateExprs(table, IsBoolean, op.left, op.right)
	if err != nil {
		return nil, err
	}
	cmp := exps[0].(boolValue).or(exps[1].(boolValue))
	return cmp, nil
}

func NewOr(left, right Expr) orOp {
	return orOp{left, right}
}

type xorOp struct {
	left, right Expr
}

func (op xorOp) Evaluate(table *SymTab) (Expr, error) {
	exps, err := evaluateExprs(table, IsBoolean, op.left, op.right)
	if err != nil {
		return nil, err
	}
	cmp := exps[0].(boolValue).xor(exps[1].(boolValue))
	return cmp, nil
}

func NewXor(left, right Expr) xorOp {
	return xorOp{left, right}
}

type norOp struct {
	left, right Expr
}

func (op norOp) Evaluate(table *SymTab) (Expr, error) {
	exps, err := evaluateExprs(table, IsBoolean, op.left, op.right)
	if err != nil {
		return nil, err
	}
	cmp := exps[0].(boolValue).nor(exps[1].(boolValue))
	return cmp, nil
}

func NewNor(left, right Expr) norOp {
	return norOp{left, right}
}

type nandOp struct {
	left, right Expr
}

func (op nandOp) Evaluate(table *SymTab) (Expr, error) {
	exps, err := evaluateExprs(table, IsBoolean, op.left, op.right)
	if err != nil {
		return nil, err
	}
	cmp := exps[0].(boolValue).nand(exps[1].(boolValue))
	return cmp, nil
}

func NewNand(left, right Expr) nandOp {
	return nandOp{left, right}
}

type notOp struct {
	operand Expr
}

func (op notOp) Evaluate(table *SymTab) (Expr, error) {
	exps, err := evaluateExprs(table, IsBoolean, op.operand)
	if err != nil {
		return nil, err
	}
	cmp := exps[0].(boolValue).not()
	return cmp, nil
}

func NewNot(operand Expr) notOp {
	return notOp{operand}
}
