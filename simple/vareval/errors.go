package vareval

// evalError is the internal error value type.
type evalError int

// Error values
const (
	UnbalancedParenthesesError evalError = iota
	DivideByZeroError
	MissingFactorError
	MissingArgumentError
	UnknownFunctionError
	TrailingTokensError
	DomainError
)

// descs contains string descriptions of each error value
var descs = [...]string{
	UnbalancedParenthesesError: "unbalanced parentheses",
	DivideByZeroError:          "divide by zero",
	MissingFactorError:         "missing factor",
	MissingArgumentError:       "missing function argument",
	UnknownFunctionError:       "unknown function",
	TrailingTokensError:        "trailing tokens",
	DomainError:                "domain error",
}

// Error returns a string description of an error
func (e evalError) Error() string {
	return descs[e]
}
