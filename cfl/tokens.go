package cfl

import (
	"fmt"
	"github.com/paulgriffiths/goeval/lar"
)

type tokenType int

const (
	tokenTerminal tokenType = iota
	tokenNonTerminal
	tokenAlt
	tokenArrow
	tokenEndOfLine
	tokenEmpty
)

type token struct {
	t   tokenType
	s   string
	pos lar.FilePos
}

type tokenList []token

var typeNames = []string{
	tokenTerminal:    "Terminal",
	tokenNonTerminal: "Non-terminal",
	tokenAlt:         "Alternative",
	tokenArrow:       "Arrow",
	tokenEndOfLine:   "End-of-line",
	tokenEmpty:       "Empty",
}

func (t token) String() string {
	return fmt.Sprintf("%s: %q (line %d, ch %d)",
		typeNames[t.t], t.s, t.pos.Line, t.pos.Ch)
}
