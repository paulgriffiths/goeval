package cfl

import (
	"fmt"
	"github.com/paulgriffiths/goeval/lar"
)

type lexErrorType int

const (
	lexErrIllegalCharacter lexErrorType = iota
	lexErrUnterminatedTerminal
	lexErrBadInput
	lexErrUnknownErr
)

type lexErr interface {
	error
	implementsLexError()
}

type lexError struct {
	t   lexErrorType
	i   string
	pos lar.FilePos
}

var lexErrorNames = []string{
	lexErrIllegalCharacter:     "illegal character",
	lexErrUnterminatedTerminal: "unterminated terminal",
	lexErrBadInput:             "bad input",
	lexErrUnknownErr:           "unknown error",
}

func (e lexError) implementsLexError() {}

func (e lexError) Error() string {
	return fmt.Sprintf("%s at line %d, char %d",
		lexErrorNames[e.t], e.pos.Line, e.pos.Ch)
}
