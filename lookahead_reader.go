package eval

import (
    "fmt"
    "io"
    "unicode"
)

// Lookahead reader implements a single character lookahead reader.
type LookaheadReader struct {
    reader io.Reader
    buffer []byte
    current byte
    lookahead byte
}

// NewLookaheadReader returns a single character lookahead reader from
// an io.Reader
func NewLookaheadReader(reader io.Reader) (LookaheadReader, error) {
    r := LookaheadReader{reader, []byte{0}, 0, 0}
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
    r.current = r.lookahead
    if r.current == 0 {
        return r.current, io.EOF
    }

    if _, err := r.reader.Read(r.buffer); err == nil {
        r.lookahead = r.buffer[0]
    } else if err == io.EOF {
        r.lookahead = 0
    } else {
        return 0, err
    }

    return r.current, nil
}

func (r *LookaheadReader) NextSafe() byte {
    n, err := r.Next()
    if err != nil {
        panic(fmt.Sprintf("NextSafe was not safe"))
    }
    return n
}

// Lookahead returns the lookahead character from a lookahead reader.
// Function returns 0 when we're already at the last character and there
// is no lookahead character.
func (r LookaheadReader) Lookahead() byte {
    return r.lookahead
}

// LookaheadIs returns true if the lookahead character matches the
// argument, otherwise it returns false.
func (r LookaheadReader) LookaheadIs(b byte) bool {
    return r.lookahead == b
}

func (r LookaheadReader) IsSpace() bool {
    return unicode.IsSpace(rune(r.current))
}

func (r LookaheadReader) IsDigit() bool {
    return unicode.IsDigit(rune(r.current))
}

func (r LookaheadReader) IsLetter() bool {
    return unicode.IsLetter(rune(r.current))
}

func (r LookaheadReader) LookaheadIsSpace() bool {
    return unicode.IsSpace(rune(r.lookahead))
}

func (r LookaheadReader) LookaheadIsDigit() bool {
    return unicode.IsDigit(rune(r.lookahead))
}

func (r LookaheadReader) LookaheadIsLetter() bool {
    return unicode.IsLetter(rune(r.lookahead))
}

func (r *LookaheadReader) GetLetters() []byte {
    if !r.IsLetter() {
        panic(fmt.Sprintf("current is not letter in GetLetters()"))
    }

    result := []byte{r.current}
    for r.LookaheadIsLetter() {
        r.Next()
        result = append(result, r.current)
    }
    return result
}

func (r *LookaheadReader) GetDigits() []byte {
    if !r.IsDigit() {
        panic(fmt.Sprintf("current is not digit in GetDigits()"))
    }

    result := []byte{r.current}
    for r.LookaheadIsDigit() {
        r.Next()
        result = append(result, r.current)
    }
    return result
}

