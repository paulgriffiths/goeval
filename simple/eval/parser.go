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

    expr, err := getExpr(tokens.NewLTChan(ch))
    if err != nil {
        return 1, fmt.Errorf("Couldn't parse expression: %v", err)
    }

    // TODO: Need to empty channel, in case of error, so that
    // lexing goroutine can finish and return.

    value, err := expr.Evaluate()
    if err != nil {
        return 1, fmt.Errorf("Couldn't evaluate expression: %v", err)
    }
    return value, nil
}

func getExpr(ltchan *tokens.LTChan) (expr, error) {
    firstTerm, err := getTerm(ltchan)
    if err != nil {
        return nil, err
    }

    switch {
    case ltchan.Match(tokens.NewOperatorToken("+")):
        secondTerm, err := getTerm(ltchan)
        if err != nil {
            return nil, err
        }
        return add{firstTerm, secondTerm}, nil
    case ltchan.Match(tokens.NewOperatorToken("-")):
        secondTerm, err := getTerm(ltchan)
        if err != nil {
            return nil, err
        }
        return subtract{firstTerm, secondTerm}, nil
    default:
        return firstTerm, nil
    }
}

func getTerm(ltchan *tokens.LTChan) (expr, error) {
    firstFact, err := getFactor(ltchan)
    if err != nil {
        return nil, err
    }

    switch {
    case ltchan.Match(tokens.NewOperatorToken("*")):
        secondFact, err := getFactor(ltchan)
        if err != nil {
            return nil, err
        }
        return multiply{firstFact, secondFact}, nil
    case ltchan.Match(tokens.NewOperatorToken("/")):
        secondFact, err := getFactor(ltchan)
        if err != nil {
            return nil, err
        }
        return divide{firstFact, secondFact}, nil
    default:
        return firstFact, nil
    }
}

func getFactor(ltchan *tokens.LTChan) (expr, error) {
    token, err := ltchan.Next()
    if err != nil {
        return nil, err
    }

    switch {
    case token.IsRightParen():
        return nil, fmt.Errorf("mismatched parentheses")
    case token.IsLeftParen():
        expr, err := getExpr(ltchan)
        if err != nil {
            return nil, err
        }
        if !ltchan.Match(tokens.RightParenToken()) {
            return nil, fmt.Errorf("mismatched parentheses")
        }
        return expr, nil
    case token.IsNumber():
        value, err := strconv.ParseFloat(token.Value(), 64)
        if err != nil {
            panic(fmt.Sprintf("Couldn't convert to float: %s", token.Value()))
        }
        return number{value}, nil
    default:
        return nil, fmt.Errorf("bad factor")
    }
    panic(fmt.Errorf("unknown parser error getting factor"))
}
