package eval

import (
    "io"
    "github.com/paulgriffiths/goeval/lar"
    "github.com/paulgriffiths/goeval/tokens"
)

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
                ch <- tokens.NewOperatorToken(string(reader.Result()))
            case reader.MatchOneOf('('):
                ch <- tokens.LeftParenToken()
            case reader.MatchOneOf(')'):
                ch <- tokens.RightParenToken()
            case reader.MatchLetters():
                ch <- tokens.NewFunctionToken(string(reader.Result()))
            case reader.MatchDigits():
                value := reader.Result()
                if reader.MatchOneOf('.') {
                    value = append(value, reader.Result()...)
                    if !reader.MatchDigits() {
                        ch <- tokens.NewIllegalToken(string(value))
                        continue
                    }
                    value = append(value, reader.Result()...)
                }
                if reader.MatchOneOf('e', 'E') {
                    value = append(value, reader.Result()...)
                    if reader.MatchOneOf('-', '+')  {
                        value = append(value, reader.Result()...)
                    }
                    if !reader.MatchDigits() {
                        ch <- tokens.NewIllegalToken(string(value))
                        continue
                    }
                    value = append(value, reader.Result()...)
                }
                ch <- tokens.NewNumberToken(string(value))
            }
        }
    } ()

    return ch, nil
}
