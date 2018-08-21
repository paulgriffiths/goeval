package expr

type Value interface {
	Expr
	Equals(other Value) bool
	String() string
}
