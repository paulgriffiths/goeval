package eval

import (
    "io"
    "github.com/paulgriffiths/goeval"
)

func NewLexer(input io.Reader) (chan eval.Token, error) {
    reader, err := eval.NewLookaheadReader(input)
    if err != nil {
        return nil, err
    }

    ch := make(chan eval.Token)

    go func() {
        for {
            c, err := reader.Next()
            if err != nil {
                close(ch)
                return
            }

            switch {
            case reader.IsSpace():
                continue
            case c == '+':
                ch <- eval.NewOperatorToken("+")
            case c == '-':
                ch <- eval.NewOperatorToken("-")
            case c == '*':
                ch <- eval.NewOperatorToken("*")
            case c == '/':
                ch <- eval.NewOperatorToken("/")
            case c == '^':
                ch <- eval.NewOperatorToken("^")
            case c == '%':
                ch <- eval.NewOperatorToken("%")
            case c == '(':
                ch <- eval.LeftParenToken()
            case c == ')':
                ch <- eval.RightParenToken()
            case reader.IsLetter():
                ch <- eval.NewFunctionToken(string(reader.GetLetters()))
            case reader.IsDigit():
                value := reader.GetDigits()
                if reader.LookaheadIs('.') {
                    value = append(value, reader.NextSafe())
                    if !reader.LookaheadIsDigit() {
                        ch <- eval.NewIllegalToken(string(value))
                        continue
                    }
                    reader.Next()
                    value = append(value, reader.GetDigits()...)
                }
                if reader.LookaheadIs('e') || reader.LookaheadIs('E') {
                    value = append(value, reader.NextSafe())
                    if reader.LookaheadIs('-')  {
                        value = append(value, reader.NextSafe())
                    }
                    if !reader.LookaheadIsDigit() {
                        ch <- eval.NewIllegalToken(string(value))
                        continue
                    }
                    reader.Next()
                    value = append(value, reader.GetDigits()...)
                }
                ch <- eval.NewNumberToken(string(value))
            }
        }
    } ()

    return ch, nil
}
