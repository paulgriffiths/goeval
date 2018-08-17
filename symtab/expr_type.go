package symtab

// ExprType represents the type of an expression value
type ExprType int

const (
	ExprTypeInt ExprType = iota
	ExprTypeReal
	ExprTypeBool
	ExprTypeString
)
