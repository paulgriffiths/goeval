// Package lar implements a single character lookahead reader with
// basic matching functions.
package lar

import (
    "fmt"
    "io"
    "unicode"
)

// LookaheadReader implements a single character lookahead reader.
type LookaheadReader struct {
    reader io.Reader
    buffer []byte
    result []byte
    lookahead byte
}

// NewLookaheadReader returns a single character lookahead reader from
// an io.Reader
func NewLookaheadReader(reader io.Reader) (LookaheadReader, error) {
    r := LookaheadReader{reader, []byte{0}, []byte{}, 0}
    _, err := r.reader.Read(r.buffer)
    if err != nil {
        if err == io.EOF {
            return r, nil
        }
        return r, fmt.Errorf("couldn't create lookahead reader: %v", err)
    }
    r.lookahead = r.buffer[0]
    return r, nil
}

// Next returns the next character from a lookahead reader.
// If there are no more characters, the function returns 0 and io.EOF.
// On any other error, the function returns 0 and that error.
func (r *LookaheadReader) Next() (byte, error) {
    if r.lookahead == 0 {
        return r.lookahead, io.EOF
    }
    current := r.lookahead

    if _, err := r.reader.Read(r.buffer); err == nil {
        r.lookahead = r.buffer[0]
    } else if err == io.EOF {
        r.lookahead = 0
    } else {
        r.lookahead = 0
        return 0, err
    }

    return current, nil
}

// MatchOneOf returns true if the next character to be read is among
// the characters passed to the function and stores that character in
// the result, otherwise it returns false and clears the result.
func (r *LookaheadReader) MatchOneOf(vals ...byte) bool {
    r.result = []byte{}
    for _, b := range vals {
        if r.lookahead == b {
            r.Next()
            r.result = append(r.result, b)
            return true
        }
    }
    return false
}

// matchSingleIsFunc packages up the matching logic for MatchLetter,
// MatchDigit, and MatchSpace, which otherwise differ only in the
// function used to test the byte.
func (r *LookaheadReader) matchSingleIsFunc(isFunc func (rune) bool) bool {
    r.result = []byte{}
    if isFunc(rune(r.lookahead)) {
        current, _ := r.Next()
        r.result = append(r.result, current)
        return true
    }
    return false
}

// matchMultipleIsFunc packages up the matching logic for MatchLetters,
// MatchDigits, and MatchSpaces, which otherwise differ only in the
// function used to test the byte.
func (r *LookaheadReader) matchMultipleIsFunc(isFunc func (rune) bool) bool {
    r.result = []byte{}
    found := false
    for isFunc(rune(r.lookahead)) {
        found = true
        current, _ := r.Next()
        r.result = append(r.result, current)
    }
    return found
}

// MatchLetter returns true if the next character to be read is a letter
// and stores that character in the result, otherwise it returns false
// and clears the result.
func (r *LookaheadReader) MatchLetter() bool {
    return r.matchSingleIsFunc(unicode.IsLetter)
}

// MatchSpace returns true if the next character to be read is whitespace
// and stores that character in the result, otherwise it returns false
// and clears the result.
func (r *LookaheadReader) MatchSpace() bool {
    return r.matchSingleIsFunc(unicode.IsSpace)
}

// MatchDigit returns true if the next character to be read is a digit
// and stores that character in the result, otherwise it returns false
// and clears the result.
func (r *LookaheadReader) MatchDigit() bool {
    return r.matchSingleIsFunc(unicode.IsDigit)
}

// MatchLetters returns true if the next character to be read is a letter
// and stores that and all immediately following letter characters in
// the result, otherwise it returns false and clears the result.
func (r *LookaheadReader) MatchLetters() bool {
    return r.matchMultipleIsFunc(unicode.IsLetter)
}

// MatchSpaces returns true if the next character to be read is whitespace
// and stores that and all immediately following whitespace characters in
// the result, otherwise it returns false and clears the result.
func (r *LookaheadReader) MatchSpaces() bool {
    return r.matchMultipleIsFunc(unicode.IsSpace)
}

// MatchDigits returns true if the next character to be read is a digit
// and stores that and all immediately following digit characters in
// the result, otherwise it returns false and clears the result.
func (r *LookaheadReader) MatchDigits() bool {
    return r.matchMultipleIsFunc(unicode.IsDigit)
}

// Result returns the result of the most recent matching test.
func (r LookaheadReader) Result() []byte {
    return r.result[:]
}

// EndOfInput returns true if end of input has been reached, otherwise false.
func (r LookaheadReader) EndOfInput() bool {
    return r.lookahead == 0
}
