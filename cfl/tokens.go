package cfl

import (
	"fmt"
	"github.com/paulgriffiths/goeval/lar"
)

// tokenType represents the type of a token.
type tokenType int

const (
	tokenTerminal tokenType = iota
	tokenNonTerminal
	tokenAlt
	tokenArrow
	tokenEndOfLine
	tokenEmpty
)

// token represents a token.
type token struct {
	t   tokenType
	s   string
	pos lar.FilePos
}

// tokenList represents a list of tokens.
type tokenList []token

// typeNames associates token type values with descriptive strings.
var typeNames = []string{
	tokenTerminal:    "Terminal",
	tokenNonTerminal: "Non-terminal",
	tokenAlt:         "Alternative",
	tokenArrow:       "Arrow",
	tokenEndOfLine:   "End-of-line",
	tokenEmpty:       "Empty",
}

// String returns a string representation of a token.
func (t token) String() string {
	return fmt.Sprintf("%s: %q (line %d, ch %d)",
		typeNames[t.t], t.s, t.pos.Line, t.pos.Ch)
}
