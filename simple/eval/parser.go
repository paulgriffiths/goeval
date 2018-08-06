package eval

import (
    "fmt"
    "strconv"
    "strings"
    "github.com/paulgriffiths/goeval/tokens"
)

func evaluate(expression string) (float64, error) {
    ch, err := NewLexer(strings.NewReader(expression))
    if err != nil {
        return 1, fmt.Errorf("Couldn't create lexer: %v", err)
    }

    e, err := getExpr(tokens.NewLTChan(ch))
    if err != nil {
        return 1, fmt.Errorf("Couldn't parse expression: %v", err)
    }

    // TODO: Need to empty channel, in case of error, so that
    // lexing goroutine can finish and return.

    v, err := e.Evaluate()
    if err != nil {
        return 1, fmt.Errorf("Couldn't evaluate expression: %v", err)
    }
    return v, nil
}

func getExpr(ltchan *tokens.LTChan) (expr, error) {
    f, err := getTerm(ltchan)
    if err != nil {
        return nil, err
    }

    switch {
    case ltchan.MatchIfEqual(tokens.NewOperatorToken("+")):
        v, err := getTerm(ltchan)
        if err != nil {
            return nil, err
        }
        return add{f, v}, nil
    case ltchan.MatchIfEqual(tokens.NewOperatorToken("-")):
        v, err := getTerm(ltchan)
        if err != nil {
            return nil, err
        }
        return subtract{f, v}, nil
    default:
        return f, nil
    }
}

func getTerm(ltchan *tokens.LTChan) (expr, error) {
    f, err := getFactor(ltchan)
    if err != nil {
        return nil, err
    }

    switch {
    case ltchan.MatchIfEqual(tokens.NewOperatorToken("*")):
        v, err := getFactor(ltchan)
        if err != nil {
            return nil, err
        }
        return multiply{f, v}, nil
    case ltchan.MatchIfEqual(tokens.NewOperatorToken("/")):
        v, err := getFactor(ltchan)
        if err != nil {
            return nil, err
        }
        return divide{f, v}, nil
    default:
        return f, nil
    }
}

func getFactor(ltchan *tokens.LTChan) (expr, error) {
    t, err := ltchan.Next()
    if err != nil {
        return nil, err
    }

    switch {
    case t.IsRightParen():
        return nil, fmt.Errorf("mismatched parentheses")
    case t.IsLeftParen():
        e, err := getExpr(ltchan)
        if err != nil {
            return nil, err
        }
        if !ltchan.MatchIfEqual(tokens.RightParenToken()) {
            return nil, fmt.Errorf("mismatched parentheses")
        }
        return e, nil
    case t.IsNumber():
        v, err := strconv.ParseFloat(t.Value(), 64)
        if err != nil {
            panic(fmt.Sprintf("Couldn't convert to float: %s", t.Value()))
        }
        return number{v}, nil
    default:
        return nil, fmt.Errorf("bad factor")
    }
    panic(fmt.Errorf("unknown parser error getting factor"))
}
