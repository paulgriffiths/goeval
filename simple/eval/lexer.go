package eval

import (
    "io"
    "unicode"
    "github.com/paulgriffiths/goeval"
)

type token struct {
    tokenType int
    value string
}

const (
    operatorToken int = iota
    numberToken
    functionToken
    leftParenToken
    rightParenToken
    illegalToken
)

func newToken(tokenType int, value string) token {
    return token{tokenType, value}
}

func NewLexer(input io.Reader) (chan token, error) {
    r, err := eval.NewLookaheadReader(input)
    if err != nil {
        return nil, err
    }

    ch := make(chan token)

    go func() {
        for {
            c, err := r.Next()
            if err != nil {
                close(ch)
                return
            }

            switch {
            case unicode.IsSpace(rune(c)):
                continue
            case c == '+':
                ch <- newToken(operatorToken, "+")
            case c == '-':
                ch <- newToken(operatorToken, "-")
            case c == '*':
                ch <- newToken(operatorToken, "*")
            case c == '/':
                ch <- newToken(operatorToken, "/")
            case c == '^':
                ch <- newToken(operatorToken, "^")
            case c == '%':
                ch <- newToken(operatorToken, "%")
            case c == '%':
                ch <- newToken(operatorToken, "%")
            case c == '(':
                ch <- newToken(leftParenToken, "(")
            case c == ')':
                ch <- newToken(rightParenToken, ")")
            case unicode.IsLetter(rune(c)):
                value := []byte{c}
                for unicode.IsLetter(rune(r.Lookahead())) {
                    n, _ := r.Next()
                    value = append(value, n)
                }
                ch <- newToken(functionToken, string(value))
            case unicode.IsDigit(rune(c)):
                value := []byte{c}
                for unicode.IsDigit(rune(r.Lookahead())) {
                    n, _ := r.Next()
                    value = append(value, n)
                }
                if r.Lookahead() == '.' {
                    n, _ := r.Next()
                    value = append(value, n)
                    if !unicode.IsDigit(rune(r.Lookahead())) {
                        ch <- newToken(illegalToken, string(value))
                        continue
                    }
                    for unicode.IsDigit(rune(r.Lookahead())) {
                        n, _ := r.Next()
                        value = append(value, n)
                    }
                }
                if r.Lookahead() == 'e' || r.Lookahead() == 'E' {
                    n, _ := r.Next()
                    value = append(value, n)
                    if r.Lookahead() == '-' {
                        n, _ := r.Next()
                        value = append(value, n)
                    }
                    if !unicode.IsDigit(rune(r.Lookahead())) {
                        ch <- newToken(illegalToken, string(value))
                        continue
                    }
                    for unicode.IsDigit(rune(r.Lookahead())) {
                        n, _ := r.Next()
                        value = append(value, n)
                    }
                }
                ch <- newToken(numberToken, string(value))
            }
        }
    } ()

    return ch, nil
}
