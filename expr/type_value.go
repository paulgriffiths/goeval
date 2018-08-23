package expr

// Value contains a numeric, boolean, string or variable value.
// The purpose of this type is to differentiate these from, for
// example, operations, which can evaluate to a value, but which
// do not themselves contain a specific value.
type value interface {
	Expr
	Equals(other value) bool
}
