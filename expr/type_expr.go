package expr

type Expr interface {
	Evaluate(table *SymTab) (Expr, error)
}
