package vareval

import (
	"github.com/paulgriffiths/goeval/lar"
	"github.com/paulgriffiths/goeval/tokens"
	"io"
)

var keywords = map[string]bool{
	"print": true,
	"true":  true,
	"false": true,
	"and":   true,
	"or":    true,
	"e":     true,
	"pi":    true,
}

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
			case reader.MatchOneOf('='):
				if reader.MatchOneOf('=') {
					ch <- tokens.OperatorToken("==")
				} else {
					ch <- tokens.OperatorToken("=")
				}
			case reader.MatchOneOf('<'):
				if reader.MatchOneOf('=') {
					ch <- tokens.OperatorToken("<=")
				} else {
					ch <- tokens.OperatorToken("<")
				}
			case reader.MatchOneOf('>'):
				if reader.MatchOneOf('=') {
					ch <- tokens.OperatorToken(">=")
				} else {
					ch <- tokens.OperatorToken(">")
				}
			case reader.MatchOneOf('!'):
				if reader.MatchOneOf('=') {
					ch <- tokens.OperatorToken("!=")
				} else {
					ch <- tokens.IllegalToken("!")
				}
			case reader.MatchOneOf('+', '-', '*', '/', '^'):
				ch <- tokens.OperatorToken(string(reader.Result.Value))
			case reader.MatchOneOf('('):
				ch <- tokens.LeftParenToken()
			case reader.MatchOneOf(')'):
				ch <- tokens.RightParenToken()
			case reader.MatchLetters():
				word := string(reader.Result.Value)
				if _, ok := keywords[word]; ok {
					ch <- tokens.KeywordToken(word)
				} else {
					ch <- tokens.WordToken(word)
				}
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
