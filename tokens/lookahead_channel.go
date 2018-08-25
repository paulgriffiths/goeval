package tokens

import (
	"fmt"
)

// LTChan implements a single token lookahead channel.
type LTChan struct {
	ch                 chan Token
	current, lookahead Token
}

// NewLTChan returns a new single token lookahead channel.
func NewLTChan(ch chan Token) *LTChan {
	l := LTChan{ch, nullToken(), nullToken()}
	if t, ok := <-ch; ok {
		l.lookahead = t
	}
	return &l
}

// readNext reads the next token from the channel.
func (l *LTChan) readNext() {
	l.current = l.lookahead
	if next, ok := <-l.ch; !ok {
		l.lookahead = nullToken()
	} else {
		l.lookahead = next
	}
}

// Next reads and returns the next token.
func (l *LTChan) Next() (Token, error) {
	l.current = l.lookahead
	if l.current == nullToken() {
		return l.current, fmt.Errorf("no more tokens")
	}
	l.readNext()
	return l.current, nil
}

// Match returns true and reads the next token if type of the next token
// read would match the type. Otherwise it returns false and doesn't
// read the next token.
func (l *LTChan) Match(tokenType TokenType) bool {
	if l.lookahead.Type == tokenType {
		l.readNext()
		return true
	}
	return false
}

// MatchValue returns true and reads the next token if type and value
// of the next token read would match the provided type and value.
// Otherwise it returns false and doesn't read the next token.
func (l *LTChan) MatchValue(tokenType TokenType, value string) bool {
	if l.lookahead.Type == tokenType && l.lookahead.Value == value {
		l.readNext()
		return true
	}
	return false
}

// MatchToken returns true and reads the next token if the next token
// read would match the provided token. Otherwise it returns false
// and doesn't read the next token.
func (l *LTChan) MatchToken(token Token) bool {
	if l.lookahead == token {
		l.readNext()
		return true
	}
	return false
}

// Value returns the value of the most recently read token.
func (l *LTChan) Value() string {
	return l.current.Value
}

// IsEmpty returns true if there are no more tokens to read in the channel.
func (l *LTChan) IsEmpty() bool {
	return l.lookahead == nullToken()
}

// Flush reads all remaining tokens in the channel, so the sender
// can finish.
func (l *LTChan) Flush() {
	for !l.IsEmpty() {
		l.Next()
	}
}
