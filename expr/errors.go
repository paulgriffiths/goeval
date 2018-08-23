package expr

// Error is the error value type for expression evaluation errors.
type Error int

// Error values
const (
	DivideByZeroError Error = iota
	DomainError
	TypeError
	UndefinedVariableError
	RangeError
)

// exprErrorDescs contains string descriptions of each error value
var exprErrorDescs = [...]string{
	DivideByZeroError:      "divide by zero",
	DomainError:            "domain error",
	TypeError:              "type error",
	UndefinedVariableError: "undefined variable",
	RangeError:             "range error",
}

// Error returns a string description of an error
func (e Error) Error() string {
	return exprErrorDescs[e]
}
