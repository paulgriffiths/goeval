package tokens

import (
    "fmt"
)

// LTChan implements a single token lookahead channel.
type LTChan struct {
    ch chan Token
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

// Value returns the value of the most recently read token.
func (l *LTChan) Value() string {
    return l.current.Value()
}

// Next reads and returns the next token.
func (l *LTChan) Next() (Token, error) {
    l.current = l.lookahead
    if l.current == nullToken() {
        return l.current, fmt.Errorf("no more tokens")
    }
    if t, ok := <-l.ch; !ok {
        l.lookahead = nullToken()
    } else {
        l.lookahead = t
    }
    return l.current, nil
}

// Match returns true and reads the next token if the next token
// read would match the provided token. Otherwise it returns false
// and doesn't read the next token.
func (l *LTChan) Match(t Token) bool {
    if l.lookahead != t {
        return false
    }
    l.current = l.lookahead
    if next, ok := <-l.ch; !ok {
        l.lookahead = nullToken()
    } else {
        l.lookahead = next
    }
    return true
}

// MatchType returns true and reads the next token if the type of the
// next token read would match the type of the provided token, regardless
// of their values. Otherwise it returns false and doesn't read the next
// token.
func (l *LTChan) MatchType(t Token) bool {
    if l.lookahead.tokenType != t.tokenType {
        return false
    }
    l.current = l.lookahead
    if next, ok := <-l.ch; !ok {
        l.lookahead = nullToken()
    } else {
        l.lookahead = next
    }
    return true
}

// IsEmpty() returns true if there are no more tokens to read in the channel.
func (l *LTChan) IsEmpty() bool {
    return l.lookahead == nullToken()
}

// Flush() reads all remaining tokens in the channel, so the sender can finish.
func (l *LTChan) Flush() {
    for !l.IsEmpty() {
        l.Next()
    }
}
