package eval

import (
    "fmt"
    "strconv"
    "strings"
    "github.com/paulgriffiths/goeval/tokens"
)

func Evaluate(expression string) (float64, error) {
    ch, err := NewLexer(strings.NewReader(expression))
    if err != nil {
        return 1, fmt.Errorf("Couldn't create lexer: %v", err)
    }

    expr, err := getExpr(tokens.NewLTChan(ch))
    if err != nil {
        return 1, err
    }

    // TODO: Need to empty channel, in case of error, so that
    // lexing goroutine can finish and return.

    value, err := expr.Evaluate()
    if err != nil {
        return 1, err
    }
    return value, nil
}

func getExpr(ltchan *tokens.LTChan) (expr, error) {
    firstTerm, err := getTerm(ltchan)
    if err != nil {
        return nil, err
    }

    switch {
    case ltchan.Match(tokens.OperatorToken("+")):
        secondTerm, err := getTerm(ltchan)
        if err != nil {
            return nil, err
        }
        return add{firstTerm, secondTerm}, nil
    case ltchan.Match(tokens.OperatorToken("-")):
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
    case ltchan.Match(tokens.OperatorToken("*")):
        secondFact, err := getFactor(ltchan)
        if err != nil {
            return nil, err
        }
        return multiply{firstFact, secondFact}, nil
    case ltchan.Match(tokens.OperatorToken("/")):
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
    switch {
    case ltchan.Match(tokens.RightParenToken()):
        return nil, UnbalancedParenthesesError
    case ltchan.Match(tokens.LeftParenToken()):
        ex, err := getExpr(ltchan)
        if err != nil {
            return nil, err
        }
        if !ltchan.Match(tokens.RightParenToken()) {
            return nil, UnbalancedParenthesesError
        }
        return ex, nil
    case ltchan.MatchType(tokens.ZeroNumberToken()):
        value, err := strconv.ParseFloat(ltchan.Value(), 64)
        if err != nil {
            panic(fmt.Sprintf("Couldn't convert to float: %s", ltchan.Value()))
        }
        return number{value}, nil
    default:
        return nil, fmt.Errorf("bad factor")
    }
    panic(fmt.Errorf("unknown parser error getting factor"))
}
