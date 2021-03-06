package eval

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
			case reader.MatchOneOf('+'):
				ch <- tokens.New(tokens.AddOperator, "+")
			case reader.MatchOneOf('-'):
				ch <- tokens.New(tokens.SubOperator, "-")
			case reader.MatchOneOf('*'):
				ch <- tokens.New(tokens.MulOperator, "*")
			case reader.MatchOneOf('/'):
				ch <- tokens.New(tokens.DivOperator, "/")
			case reader.MatchOneOf('^'):
				ch <- tokens.New(tokens.PowOperator, "^")
			case reader.MatchOneOf('('):
				ch <- tokens.New(tokens.LeftParen, "(")
			case reader.MatchOneOf(')'):
				ch <- tokens.New(tokens.RightParen, ")")
			case reader.MatchLetters():
				ch <- tokens.New(tokens.Word, string(reader.Result.Value))
			case reader.MatchDigits():
				value := reader.Result.Value
				if reader.MatchOneOf('.') {
					value = append(value, reader.Result.Value...)
					if !reader.MatchDigits() {
						ch <- tokens.New(tokens.Illegal,
							string(reader.Result.Value))
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
						ch <- tokens.New(tokens.Illegal, string(value))
						continue
					}
					value = append(value, reader.Result.Value...)
				}
				ch <- tokens.New(tokens.Number, string(value))
			default:
				n, _ := reader.Next()
				ch <- tokens.New(tokens.Illegal, string(n))
			}
		}
	}()

	return ch, nil
}
