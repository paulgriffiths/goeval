package eval

import (
    "fmt"
    "io"
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

// Lookahead returns the lookahead character from a lookahead reader.
// Function returns 0 when we're already at the last character and there
// is no lookahead character.
func (r LookaheadReader) Lookahead() byte {
    return r.lookahead
}
