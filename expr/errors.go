package expr

// exprError is the internal error value type.
type exprError int

// Error values
const (
	DivideByZeroError exprError = iota
	DomainError
	TypeError
	UnknownIdentifierError
	RangeError
)

// descs contains string descriptions of each error value
var descs = [...]string{
	DivideByZeroError:      "divide by zero",
	DomainError:            "domain error",
	TypeError:              "type error",
	UnknownIdentifierError: "unknown identifier",
	RangeError:             "range error",
}

// Error returns a string description of an error
func (e exprError) Error() string {
	return descs[e]
}
