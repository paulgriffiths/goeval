package vareval

import (
	"github.com/paulgriffiths/goeval/lar"
	"github.com/paulgriffiths/goeval/tokens"
	"io"
)

// NewLexer creates a new lexer to read from the provided reader
// and returns the channel to which it writes.
func NewLexer(input io.Reader) (chan tokens.Token, error) {
	reader, err := lar.NewLookaheadReader(input)
	if err != nil {
		return nil, err
	}

	ch := make(chan tokens.Token)

	go func() {
		for {
			switch {
			case reader.EndOfInput():
				close(ch)
				return
			case reader.MatchSpaces():
				continue
			case reader.MatchOneOf('+', '-', '*', '/', '^'):
				ch <- tokens.OperatorToken(string(reader.Result.Value))
			case reader.MatchOneOf('('):
				ch <- tokens.LeftParenToken()
			case reader.MatchOneOf(')'):
				ch <- tokens.RightParenToken()
			case reader.MatchLetters():
				ch <- tokens.WordToken(string(reader.Result.Value))
			case reader.MatchDigits():
				value := reader.Result.Value
				if reader.MatchOneOf('.') {
					value = append(value, reader.Result.Value...)
					if !reader.MatchDigits() {
						ch <- tokens.IllegalToken(string(value))
						continue
					}
					value = append(value, reader.Result.Value...)
				}
				if reader.MatchOneOf('e', 'E') {
					value = append(value, reader.Result.Value...)
					if reader.MatchOneOf('-', '+') {
						value = append(value, reader.Result.Value...)
					}
					if !reader.MatchDigits() {
						ch <- tokens.IllegalToken(string(value))
						continue
					}
					value = append(value, reader.Result.Value...)
				}
				ch <- tokens.NumberToken(string(value))
			default:
				n, _ := reader.Next()
				ch <- tokens.IllegalToken(string(n))
			}
		}
	}()

	return ch, nil
}
