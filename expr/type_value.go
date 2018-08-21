package expr

// Value contains a numeric, boolean, string or variable value.
type Value interface {
	Expr
	Equals(other Value) bool
	String() string
}
