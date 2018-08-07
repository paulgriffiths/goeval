package eval

type evalErrorType int

type evalError struct {
    id evalErrorType
}

func (e evalError) Error() string {
    switch e.id {
    case ubp:
        return "unbalanced parentheses"
    case dbz:
        return "divide by zero"
    default:
        return "unspecified error"
    }
}

const (
    ubp evalErrorType = iota
    dbz
)

var UnbalancedParenthesesError evalError = evalError{ubp}
var DivideByZeroError evalError = evalError{dbz}
