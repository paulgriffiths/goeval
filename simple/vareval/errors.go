package vareval

// evalError is the internal error value type.
type evalError int

// Error values
const (
	UnbalancedParenthesesError evalError = iota
	MissingFactorError
	MissingArgumentError
	UnknownFunctionError
	TrailingTokensError
	UnknownIdentifierError
	IllegalIdentifierError
	SyntaxError
	UnknownError
)

// descs contains string descriptions of each error value
var descs = [...]string{
	UnbalancedParenthesesError: "unbalanced parentheses",
	MissingFactorError:         "missing factor",
	MissingArgumentError:       "missing function argument",
	UnknownFunctionError:       "unknown function",
	TrailingTokensError:        "trailing tokens",
	UnknownIdentifierError:     "unknown identifier",
	IllegalIdentifierError:     "illegal identifier",
	SyntaxError:                "syntax error",
	UnknownError:               "unknown error",
}

// Error returns a string description of an error
func (e evalError) Error() string {
	return descs[e]
}
