package eval

import (
	"fmt"
	"github.com/paulgriffiths/goeval/expr"
	"github.com/paulgriffiths/goeval/tokens"
	"math"
	"strconv"
)

func getExpr(ltchan *tokens.LTChan) (expr.Expr, error) {
	var left expr.Expr
	left, err := getTerm(ltchan)
	if err != nil {
		return nil, err
	}

loop:
	for {
		switch {
		case ltchan.Match(tokens.OperatorToken("+")):
			right, err := getTerm(ltchan)
			if err != nil {
				return nil, err
			}
			left = expr.NewAdd(left, right)
		case ltchan.Match(tokens.OperatorToken("-")):
			right, err := getTerm(ltchan)
			if err != nil {
				return nil, err
			}
			left = expr.NewSub(left, right)
		default:
			break loop
		}
	}
	return left, nil
}

func getTerm(ltchan *tokens.LTChan) (expr.Expr, error) {
	var left expr.Expr
	left, err := getSubTerm(ltchan)
	if err != nil {
		return nil, err
	}

loop:
	for {
		switch {
		case ltchan.Match(tokens.OperatorToken("*")):
			right, err := getSubTerm(ltchan)
			if err != nil {
				return nil, err
			}
			left = expr.NewMul(left, right)
		case ltchan.Match(tokens.OperatorToken("/")):
			right, err := getSubTerm(ltchan)
			if err != nil {
				return nil, err
			}
			left = expr.NewDiv(left, right)
		default:
			break loop
		}
	}
	return left, nil
}

func getSubTerm(ltchan *tokens.LTChan) (expr.Expr, error) {
	var left expr.Expr
	left, err := getTermNegate(ltchan)
	if err != nil {
		return nil, err
	}

loop:
	for {
		switch {
		case ltchan.Match(tokens.OperatorToken("^")):
			right, err := getTermNegate(ltchan)
			if err != nil {
				return nil, err
			}

			// Note that we make the power operator left-associative, here

			left = expr.NewPow(left, right)
		default:
			break loop
		}
	}
	return left, nil
}

func getTermNegate(ltchan *tokens.LTChan) (expr.Expr, error) {
	neg := false
	if ltchan.Match(tokens.OperatorToken("-")) {
		neg = true
	}
	ex, err := getFactor(ltchan)
	if err != nil {
		return nil, err
	}
	if neg {
		return expr.NewNeg(ex), nil
	}
	return ex, nil
}
func getFactor(ltchan *tokens.LTChan) (expr.Expr, error) {
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
	case ltchan.MatchNumber():
		nval, err := strconv.ParseInt(ltchan.Value(), 10, 64)
		if err == nil {
			return expr.NewInt(nval), nil
		}
		rval, err := strconv.ParseFloat(ltchan.Value(), 64)
		if err != nil {
			panic(fmt.Sprintf("Couldn't convert to float: %s",
				ltchan.Value()))
		}
		return expr.NewReal(rval), nil
	case ltchan.MatchWord():
		word := string(ltchan.Value())

		if word == "e" {
			return expr.NewReal(math.E), nil
		} else if word == "pi" {
			return expr.NewReal(math.Pi), nil
		}

		if !ltchan.Match(tokens.LeftParenToken()) {
			return nil, MissingArgumentError
		}
		ex, err := getExpr(ltchan)
		if err != nil {
			return nil, err
		}
		if !ltchan.Match(tokens.RightParenToken()) {
			return nil, UnbalancedParenthesesError
		}
		switch word {
		case "cos":
			return expr.NewCos(ex), nil
		case "sin":
			return expr.NewSin(ex), nil
		case "tan":
			return expr.NewTan(ex), nil
		case "acos":
			return expr.NewAcos(ex), nil
		case "asin":
			return expr.NewAsin(ex), nil
		case "atan":
			return expr.NewAtan(ex), nil
		case "round":
			return expr.NewRound(ex), nil
		case "ceil":
			return expr.NewCeil(ex), nil
		case "floor":
			return expr.NewFloor(ex), nil
		case "sqrt":
			return expr.NewSqrt(ex), nil
		case "log":
			return expr.NewLog(ex), nil
		case "ln":
			return expr.NewLn(ex), nil
		default:
			return nil, UnknownFunctionError
		}
	default:
		return nil, MissingFactorError
	}
	panic("reached end of getFactor")
}
