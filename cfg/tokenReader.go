package cfg

// tokenReader provides lookahead and matching functionality
// for a list of tokens.
type tokenReader struct {
	tokens                       []token
	currentIndex, lookaheadIndex int
}

// newTokenReader creates a new tokenReader.
func newTokenReader(tokens []token) tokenReader {
	return tokenReader{tokens, -1, 0}
}

// reset resets the current and lookahead indices.
func (r *tokenReader) reset() {
	r.currentIndex = -1
	r.lookaheadIndex = 0
}

// atEnd checks if if we've matched the last token.
func (r tokenReader) atEnd() bool {
	return r.lookaheadIndex == len(r.tokens)
}

// match returns true and advances the indices if the next token
// to be read matches the provided type.
func (r *tokenReader) match(t tokenType) bool {
	if r.atEnd() {
		return false
	}
	if r.tokens[r.lookaheadIndex].t == t {
		r.currentIndex++
		r.lookaheadIndex++
		return true
	}
	return false
}

// peek returns true if the next to be read matches the provided type.
func (r tokenReader) peek(t tokenType) bool {
	if r.atEnd() {
		return false
	}
	if r.tokens[r.lookaheadIndex].t == t {
		return true
	}
	return false
}

// current returns the most recently read token. This should only
// be called after a successful match, as it will panic if no tokens
// have yet been read.
func (r tokenReader) current() token {
	if r.currentIndex < 0 || r.currentIndex >= len(r.tokens) {
		panic("no current token")
	}
	return r.tokens[r.currentIndex]
}

// lookahead returns the token that would be read next. Either atEnd,
// peek or match should be called and checked prior to calling
// lookahead, as it will panic if all the tokens have been read.
func (r tokenReader) lookahead() token {
	if r.lookaheadIndex < 0 || r.lookaheadIndex >= len(r.tokens) {
		panic("no current token")
	}
	return r.tokens[r.lookaheadIndex]
}
