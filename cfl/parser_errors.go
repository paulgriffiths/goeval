package cfl

import (
	"fmt"
	"github.com/paulgriffiths/goeval/lar"
)

// parseErrorType represents the type of a parser error.
type parseErrorType int

// parser error type values.
const (
	parseErrMissingArrow parseErrorType = iota
	parseErrEmptyBody
	parseErrEmptyNotAlone
	parseErrMissingNonTerminal
)

// parseErr is an interface for parser errors. It is provided so that
// parser functions can return nil errors.
type parseErr interface {
	error
	implementsParseError()
}

// parseError is a concrete parser error type.
type parseError struct {
	t   parseErrorType
	i   string
	pos lar.FilePos
}

// parseErrorNames associate parser error type values with descriptive
// strings.
var parseErrorNames = []string{
	parseErrMissingArrow:       "missng arrow",
	parseErrEmptyBody:          "empty body",
	parseErrEmptyNotAlone:      "empty body not alone",
	parseErrMissingNonTerminal: "missing nonterminal",
}

// implementsParseError is a dummy method to satisfy the interface.
func (e parseError) implementsParseError() {}

// Error returns a string representation of a parser error.
func (e parseError) Error() string {
	return fmt.Sprintf("%s at line %d, char %d",
		parseErrorNames[e.t], e.pos.Line, e.pos.Ch)
}
