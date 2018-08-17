package lar

import (
	"fmt"
	"io"
	"unicode"
)

// LookaheadReader implements a single character lookahead reader.
type LookaheadReader struct {
	reader  io.Reader
	buffer  []byte
	current byte
    pos FilePos
	Result  ReaderResult
}

// NewLookaheadReader returns a single character lookahead reader from
// an io.Reader
func NewLookaheadReader(reader io.Reader) (LookaheadReader, error) {
	r := LookaheadReader{reader, []byte{0}, 0, FilePos{-1, 1},
		ReaderResult{[]byte{}, FilePos{0, 0}}}
	if _, err := r.reader.Read(r.buffer); err != nil && err != io.EOF {
		return r, fmt.Errorf("couldn't create lookahead reader: %v", err)
	}
	return r, nil
}

// Next returns the next character from a lookahead reader.
// If there are no more characters, the function returns 0 and io.EOF.
// On any other error, the function returns 0 and that error.
func (r *LookaheadReader) Next() (byte, error) {
	if r.buffer[0] == 0 {
		return 0, io.EOF
	}

	if r.current == '\n' {
        r.pos.incLine()
	} else {
		r.pos.inc()
	}

	r.current = r.buffer[0]

	if _, err := r.reader.Read(r.buffer); err != nil {
		r.buffer[0] = 0
		if err != io.EOF {
			return 0, err
		}
	}

	return r.current, nil
}

// MatchOneOf returns true if the next character to be read is among
// the characters passed to the function and stores that character in
// the result, otherwise it returns false and clears the result.
func (r *LookaheadReader) MatchOneOf(vals ...byte) bool {
	r.Result.clear()
	for _, b := range vals {
		if r.buffer[0] == b {
			r.Next()
			r.Result.setPos(r.pos)
			r.Result.appendByte(b)
			return true
		}
	}
	return false
}

// MatchNewline returns true if the next character to be read is
// a newline character and stores that character in the result,
// otherwise it returns false and clears the result
func (r *LookaheadReader) MatchNewline() bool {
	return r.MatchOneOf('\n')
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

// EndOfInput returns true if end of input has been reached,
// otherwise it returns false.
func (r LookaheadReader) EndOfInput() bool {
	return r.buffer[0] == 0
}

// matchSingleIsFunc packages up the matching logic for MatchLetter,
// MatchDigit, and MatchSpace, which otherwise differ only in the
// function used to test the byte.
func (r *LookaheadReader) matchSingleIsFunc(isFunc func(rune) bool) bool {
	r.Result.clear()
	if r.buffer[0] != '\n' && isFunc(rune(r.buffer[0])) {
		current, _ := r.Next()
		r.Result.setPos(r.pos)
		r.Result.appendByte(current)
		return true
	}
	return false
}

// matchMultipleIsFunc packages up the matching logic for MatchLetters,
// MatchDigits, and MatchSpaces, which otherwise differ only in the
// function used to test the byte.
func (r *LookaheadReader) matchMultipleIsFunc(isFunc func(rune) bool) bool {
	r.Result.clear()
	found := false
	for r.buffer[0] != '\n' && isFunc(rune(r.buffer[0])) {
		current, _ := r.Next()
		if !found {
			r.Result.setPos(r.pos)
			found = true
		}
		r.Result.appendByte(current)
	}
	return found
}
