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
			case reader.MatchOneOf('='):
				if reader.MatchOneOf('=') {
					ch <- tokens.New(tokens.EqualityOperator, "==")
				} else {
					ch <- tokens.New(tokens.AssignmentOperator, "=")
				}
			case reader.MatchOneOf('<'):
				if reader.MatchOneOf('=') {
					ch <- tokens.New(tokens.LessEqualOperator, "<=")
				} else {
					ch <- tokens.New(tokens.LessOperator, "<")
				}
			case reader.MatchOneOf('>'):
				if reader.MatchOneOf('=') {
					ch <- tokens.New(tokens.GreaterEqualOperator, ">=")
				} else {
					ch <- tokens.New(tokens.GreaterOperator, ">")
				}
			case reader.MatchOneOf('!'):
				if reader.MatchOneOf('=') {
					ch <- tokens.New(tokens.NonEqualityOperator, "!=")
				} else {
					ch <- tokens.New(tokens.Illegal,
						string(reader.Result.Value))
				}
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
			case reader.MatchOneOf('"'):
				sval := []rune{}
				for reader.MatchAnyExcept('"', '\n') {
					sval = append(sval, reader.Result.Value...)
				}
				if reader.MatchOneOf('"') {
					ch <- tokens.New(tokens.String, string(sval))
				} else {
					ch <- tokens.New(tokens.Illegal, string(sval))
				}
			case reader.MatchIdentifier():
				word := string(reader.Result.Value)
				if _, ok := keywords[word]; ok {
					switch word {
					case "and":
						ch <- tokens.New(tokens.AndOperator, word)
					case "or":
						ch <- tokens.New(tokens.OrOperator, word)
					case "nand":
						ch <- tokens.New(tokens.NandOperator, word)
					case "nor":
						ch <- tokens.New(tokens.NorOperator, word)
					case "xor":
						ch <- tokens.New(tokens.XorOperator, word)
					case "not":
						ch <- tokens.New(tokens.NotOperator, word)
					default:
						ch <- tokens.New(tokens.Keyword, word)
					}
				} else {
					ch <- tokens.New(tokens.Identifier, word)
				}
			case reader.MatchDigits():
				value := reader.Result.Value
				if reader.MatchOneOf('.') {
					value = append(value, reader.Result.Value...)
					if !reader.MatchDigits() {
						ch <- tokens.New(tokens.Illegal, string(value))
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
