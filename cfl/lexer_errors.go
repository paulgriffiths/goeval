package cfl

import (
	"fmt"
	"github.com/paulgriffiths/goeval/lar"
)

// lexErrorType represents the type of a lexer error.
type lexErrorType int

// lexer error type values.
const (
	lexErrIllegalCharacter lexErrorType = iota
	lexErrUnterminatedTerminal
	lexErrBadInput
	lexErrUnknownErr
)

// lexErr is an interface for lexer errors. It is provided so that
// lexer functions can return nil errors.
type lexErr interface {
	error
	implementsLexError()
}

// lexError is a concrete lexer error type.
type lexError struct {
	t   lexErrorType
	i   string
	pos lar.FilePos
}

// lexErrorNames associate lexer error type values with descriptive
// strings.
var lexErrorNames = []string{
	lexErrIllegalCharacter:     "illegal character",
	lexErrUnterminatedTerminal: "unterminated terminal",
	lexErrBadInput:             "bad input",
	lexErrUnknownErr:           "unknown error",
}

// implementsLexError is a dummy method to satisfy the interface.
func (e lexError) implementsLexError() {}

// Error returns a string representation of a lexer error.
func (e lexError) Error() string {
	return fmt.Sprintf("%s at line %d, char %d",
		lexErrorNames[e.t], e.pos.Line, e.pos.Ch)
}
