package eval

import (
    "fmt"
)

type LTChan struct {
    ch chan Token
    current, lookahead Token
}

func NewLTChan(ch chan Token) *LTChan {
    l := LTChan{ch, NullToken(), NullToken()}
    t, ok := <-ch
    if ok {
        l.lookahead = t
    }
    return &l
}

func (l *LTChan) Next() (Token, error) {
    l.current = l.lookahead
    if l.current == NullToken() {
        return l.current, fmt.Errorf("no more tokens")
    }
    t, ok := <-l.ch
    if !ok {
        l.lookahead = NullToken()
    } else {
        l.lookahead = t
    }
    return l.current, nil
}

func (l *LTChan) Lookahead() Token {
    return l.lookahead
}

func (l *LTChan) MatchToken(t Token) {
    next, err := l.Next()
    if err != nil || next != t {
        panic(fmt.Sprintf("failed to match as expected"))
    }
}

func (l *LTChan) MatchIfEqual(t Token) bool {
    if l.lookahead != t {
        return false
    }
    l.current = l.lookahead
    next, ok := <-l.ch
    if !ok {
        l.lookahead = NullToken()
    } else {
        l.lookahead = next
    }
    return true
}

