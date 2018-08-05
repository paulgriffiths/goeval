package eval

import (
    "io"
    "testing"
    "strings"
)

func TestEmptyStringLookaheadReader(t *testing.T) {
    l, err := NewLookaheadReader(strings.NewReader(""))
    if err != nil {
        t.Errorf("couldn't create lookahead reader: %v", err)
    }

    if l.current != 0 {
        t.Errorf("got %q for l.current, want %q", l.current, 0)
    }

    if l.lookahead != 0 {
        t.Errorf("got %q for l.lookahead, want %q", l.lookahead, 0)
    }

    if c, err := l.Next(); c != 0 || err != io.EOF {
        t.Errorf("got %q, %v, expected %q, %v", c, err, 0, io.EOF)
    }

    if l.current != 0 {
        t.Errorf("got %q for l.current, want %q", l.current, 0)
    }

    if l.lookahead != 0 {
        t.Errorf("got %q for l.lookahead, want %q", l.lookahead, 0)
    }
}

func TestSingleCharacterLookaheadReader(t *testing.T) {
    l, err := NewLookaheadReader(strings.NewReader("a"))
    if err != nil {
        t.Errorf("couldn't create lookahead reader: %v", err)
    }

    if l.current != 0 {
        t.Errorf("got %q for l.current, want %q", l.current, 0)
    }

    if l.lookahead != 'a' {
        t.Errorf("got %q for l.lookahead, want %q", l.lookahead, 'a')
    }

    if c, err := l.Next(); c != 'a' || err != nil {
        t.Errorf("got %q, %v for Next(), want %q, %v", c, err, 'a', nil)
    }

    if l.current != 'a' {
        t.Errorf("got %q for l.current, want %q", l.current, 'a')
    }

    if l.lookahead != 0 {
        t.Errorf("got %q for l.lookahead, want %q", l.lookahead, 0)
    }

    if c, err := l.Next(); c != 0 || err != io.EOF {
        t.Errorf("got %q, %v, expected %q, %v", c, err, 0, io.EOF)
    }

    if l.current != 0 {
        t.Errorf("got %q for l.current, want %q", l.current, 0)
    }

    if l.lookahead != 0 {
        t.Errorf("got %q for l.lookahead, want %q", l.lookahead, 0)
    }
}

func TestTwoCharacterLookaheadReader(t *testing.T) {
    l, err := NewLookaheadReader(strings.NewReader("ab"))
    if err != nil {
        t.Errorf("couldn't create lookahead reader: %v", err)
    }

    if l.current != 0 {
        t.Errorf("got %q for l.current, want %q", l.current, 0)
    }

    if l.lookahead != 'a' {
        t.Errorf("got %q for l.lookahead, want %q", l.lookahead, 'a')
    }

    if c, err := l.Next(); c != 'a' || err != nil {
        t.Errorf("got %q, %v for Next(), want %q, %v", c, err, 'a', nil)
    }

    if l.current != 'a' {
        t.Errorf("got %q for l.current, want %q", l.current, 'a')
    }

    if l.lookahead != 'b' {
        t.Errorf("got %q for l.lookahead, want %q", l.lookahead, 'b')
    }

    if c, err := l.Next(); c != 'b' || err != nil {
        t.Errorf("got %q, %v for Next(), want %q, %v", c, err, 'b', nil)
    }

    if l.current != 'b' {
        t.Errorf("got %q for l.current, want %q", l.current, 'b')
    }

    if l.lookahead != 0 {
        t.Errorf("got %q for l.lookahead, want %q", l.lookahead, 0)
    }

    if c, err := l.Next(); c != 0 || err != io.EOF {
        t.Errorf("got %q, %v, expected %q, %v", c, err, 0, io.EOF)
    }

    if l.current != 0 {
        t.Errorf("got %q for l.current, want %q", l.current, 0)
    }

    if l.lookahead != 0 {
        t.Errorf("got %q for l.lookahead, want %q", l.lookahead, 0)
    }
}

