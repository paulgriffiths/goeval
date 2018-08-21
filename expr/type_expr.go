package expr

// Expr holds an expression or subexpression.
type Expr interface {
	Evaluate(table *SymTab) (Expr, error)
}
