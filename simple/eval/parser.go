package eval

import (
    "fmt"
    "strconv"
    "strings"
    "github.com/paulgriffiths/goeval"
)

func evaluate(expression string) (float64, error) {
    ch, err := NewLexer(strings.NewReader(expression))
    if err != nil {
        return 1, fmt.Errorf("Couldn't create lexer: %v", err)
    }

    e, err := getExpr(eval.NewLTChan(ch))
    if err != nil {
        return 1, fmt.Errorf("Couldn't parse expression: %v", err)
    }

    v, err := e.Evaluate()
    if err != nil {
        return 1, fmt.Errorf("Couldn't evaluate expression: %v", err)
    }
    return v, nil
}

func getExpr(ltchan *eval.LTChan) (expr, error) {
    f, err := getTerm(ltchan)
    if err != nil {
        return nil, err
    }

    switch {
    case ltchan.MatchIfEqual(eval.NewOperatorToken("+")):
        v, err := getTerm(ltchan)
        if err != nil {
            return nil, err
        }
        return add{f, v}, nil
    case ltchan.MatchIfEqual(eval.NewOperatorToken("-")):
        v, err := getTerm(ltchan)
        if err != nil {
            return nil, err
        }
        return subtract{f, v}, nil
    default:
        return f, nil
    }
}

func getTerm(ltchan *eval.LTChan) (expr, error) {
    f, err := getFactor(ltchan)
    if err != nil {
        return nil, err
    }

    switch {
    case ltchan.MatchIfEqual(eval.NewOperatorToken("*")):
        v, err := getFactor(ltchan)
        if err != nil {
            return nil, err
        }
        return multiply{f, v}, nil
    case ltchan.MatchIfEqual(eval.NewOperatorToken("/")):
        v, err := getFactor(ltchan)
        if err != nil {
            return nil, err
        }
        return divide{f, v}, nil
    default:
        return f, nil
    }
}

func getFactor(ltchan *eval.LTChan) (expr, error) {
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
        if !ltchan.MatchIfEqual(eval.RightParenToken()) {
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
