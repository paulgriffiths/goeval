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
    case mso:
        return "missing factor"
    default:
        return "unspecified error"
    }
}

const (
    ubp evalErrorType = iota
    dbz
    mso
)

var UnbalancedParenthesesError evalError = evalError{ubp}
var DivideByZeroError evalError = evalError{dbz}
var MissingFactorError evalError = evalError{mso}
