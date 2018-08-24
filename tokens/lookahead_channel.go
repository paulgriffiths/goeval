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

// Value returns the value of the most recently read token.
func (l *LTChan) Value() string {
	return l.current.Value()
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

// Match returns true and reads the next token if the next token
// read would match the provided token. Otherwise it returns false
// and doesn't read the next token.
func (l *LTChan) Match(t Token) bool {
	if l.lookahead == t {
		l.readNext()
		return true
	}
	return false
}

// MatchType returns true and reads the next token if the type of the
// next token read would match the type of the provided token, regardless
// of their values. Otherwise it returns false and doesn't read the next
// token.
func (l *LTChan) MatchType(t Token) bool {
	if l.lookahead.tokenType == t.tokenType {
		l.readNext()
		return true
	}
	return false
}

// MatchString returns true and reads the next token if the next token
// is a string token. Otherwise it returns false and doesn't read the
// next token.
func (l *LTChan) MatchString() bool {
	if l.lookahead.tokenType == stringToken {
		l.readNext()
		return true
	}
	return false
}

// MatchNumber returns true and reads the next token if the next token
// is a number token. Otherwise it returns false and doesn't read the
// next token.
func (l *LTChan) MatchNumber() bool {
	if l.lookahead.tokenType == numberToken {
		l.readNext()
		return true
	}
	return false
}

// MatchWord returns true and reads the next token if the next token
// is a word token. Otherwise it returns false and doesn't read the
// next token.
func (l *LTChan) MatchWord() bool {
	if l.lookahead.tokenType == wordToken {
		l.readNext()
		return true
	}
	return false
}

// MatchKeyword returns true and reads the next token if the next token
// is a keyword token. Otherwise it returns false and doesn't read the
// next token.
func (l *LTChan) MatchKeyword() bool {
	if l.lookahead.tokenType == keywordToken {
		l.readNext()
		return true
	}
	return false
}

// MatchIdentifier returns true and reads the next token if the next token
// is an identifier token. Otherwise it returns false and doesn't read the
// next token.
func (l *LTChan) MatchIdentifier() bool {
	if l.lookahead.tokenType == identifierToken {
		l.readNext()
		return true
	}
	return false
}

// MatchIllegal returns true and reads the next token if the next token
// is an illegal token. Otherwise it returns false and doesn't read the
// next token.
func (l *LTChan) MatchIllegal() bool {
	if l.lookahead.tokenType == illegalToken {
		l.readNext()
		return true
	}
	return false
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
