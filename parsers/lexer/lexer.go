package lexer

import (
	"fmt"
	"github.com/paulgriffiths/goeval/cfg"
	"regexp"
	"sort"
	"unicode"
)

type lexer struct {
	terminals TerminalList   // List of terminals to search for
	r         *regexp.Regexp // Compiled regular expression
	input     string         // The input string
	n         int            // Current position in the input
}

// Lex returns a list of terminals of the provided grammar extracted
// from the provided input.
func Lex(grammar *cfg.Cfg, input string) (TerminalList, error) {
	list := TerminalList{}
	lexer, err := newLexer(grammar, input)
	if err != nil {
		return nil, err
	}

	for {
		terminal, err := lexer.getNextTerminal()
		if err != nil {
			if err.(lexError).t == lexErrEndOfInput {
				break
			}
			return nil, err
		}
		list = append(list, terminal)
	}

	return list, nil
}

func newLexer(grammar *cfg.Cfg, input string) (*lexer, lexErr) {

	// Retrieve terminals from grammar and sort in reverse order.
	// Sorting in reverse order is done in an attempt to avoid
	// conflicts and match longest strings. It's not guaranteed
	// to avoid problems with amgibuous terminals, and differs
	// from the usual Lex approach of matching the earliest-defined
	// pattern first in the event of multiple matches, but this is
	// not a general-purpose library and it's likely sufficient
	// for our current purposes.

	terminals := TerminalList{}
	for n, t := range grammar.Terminals {
		terminals = append(terminals, Terminal{n, t})
	}
	sort.Sort(sort.Reverse(terminals))

	// Build a single regular expression of the form
	// (^term1)|(^term2)|...|(^termn). Go's regular expression package
	// will prefer leftmost submatches, and this is why we sorted them
	// in reverse order. The parenthesized expressions will enable us
	// to identify which terminal was matched.

	rstring := ""
	for i, t := range terminals {
		if i != 0 {
			rstring += "|"
		}
		rstring += fmt.Sprintf("(^%s)", t.S)
	}

	r, err := regexp.Compile(rstring)
	if err != nil {
		fmt.Printf("failed to compile regex\n")
		return nil, lexError{lexErrBadRegexp}
	}

	// Build lexer and return it.

	l := lexer{terminals, r, input, 0}
	return &l, nil
}

func (l *lexer) skipWhitespace() {
	for l.n < len(l.input) && unicode.IsSpace(rune(l.input[l.n])) {
		l.n++
	}
}

func (l *lexer) endOfInput() bool {
	return l.n >= len(l.input)
}

func (l *lexer) getNextTerminal() (Terminal, lexErr) {
	l.skipWhitespace()
	if l.endOfInput() {
		return Terminal{}, lexError{lexErrEndOfInput}
	}

	result := l.r.FindAllStringSubmatchIndex(l.input[l.n:], 1)
	if len(result) == 0 {
		return Terminal{}, lexError{lexErrMatchFailed}
	}
	matches := result[0]
	if matches[0] == -1 {
		return Terminal{}, lexError{lexErrMatchFailed}
	}

	// Find out which subexpression matched, and build and
	// return a terminal accordingly, and advance the input index.

	for i := range l.terminals {
		beg, end := matches[2*(i+1)], matches[2*(i+1)+1]
		if beg == -1 {
			continue
		}
		t := Terminal{l.terminals[i].N, l.input[l.n : l.n+end-beg]}
		l.n += end - beg
		return t, nil
	}

	panic("failed to find regexp match index")
}
