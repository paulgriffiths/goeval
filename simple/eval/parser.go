// Package eval provides a simple mathematical expression evaluator.

package eval

import (
	"fmt"
	"github.com/paulgriffiths/goeval/expr"
	"github.com/paulgriffiths/goeval/tokens"
	"math"
	"strconv"
	"strings"
)

// Evaluates the provided simple mathematical expression.
func Evaluate(expression string) (expr.Value, error) {
	ch, err := NewLexer(strings.NewReader(expression))
	if err != nil {
		return nil, fmt.Errorf("Couldn't create lexer: %v", err)
	}

	ltc := tokens.NewLTChan(ch)
	exp, err := getExpr(ltc)
	if err != nil {
		return nil, err
	}

	if !ltc.IsEmpty() {
		ltc.Flush()
		return nil, TrailingTokensError
	}

	result, err := exp.Evaluate(nil)
	if err != nil {
		return nil, err
	}

	value, ok := result.(expr.Value)
	if !ok {
		return nil, UnknownError
	}

	return value, nil
}

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
			right, err := getFactor(ltchan)
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
	left, err := getFactor(ltchan)
	if err != nil {
		return nil, err
	}

loop:
	for {
		switch {
		case ltchan.Match(tokens.OperatorToken("^")):
			right, err := getFactor(ltchan)
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

func getFactor(ltchan *tokens.LTChan) (expr.Expr, error) {
	neg := false
	if ltchan.Match(tokens.OperatorToken("-")) {
		neg = true
	}
	var result expr.Expr

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
		result = ex
	case ltchan.MatchType(tokens.ZeroNumberToken()):
		nval, err := strconv.ParseInt(ltchan.Value(), 10, 64)
		if err == nil {
			result = expr.NewInt(nval)
			break
		}
		rval, err := strconv.ParseFloat(ltchan.Value(), 64)
		if err != nil {
			panic(fmt.Sprintf("Couldn't convert to float: %s",
				ltchan.Value()))
		}
		result = expr.NewReal(rval)
	case ltchan.MatchType(tokens.EmptyWordToken()):
		word := string(ltchan.Value())

		if word == "e" {
			result = expr.NewReal(math.E)
			break
		} else if word == "pi" {
			result = expr.NewReal(math.Pi)
			break
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
			result = expr.NewCos(ex)
		case "sin":
			result = expr.NewSin(ex)
		case "tan":
			result = expr.NewTan(ex)
		case "acos":
			result = expr.NewAcos(ex)
		case "asin":
			result = expr.NewAsin(ex)
		case "atan":
			result = expr.NewAtan(ex)
		case "round":
			result = expr.NewRound(ex)
		case "ceil":
			result = expr.NewCeil(ex)
		case "floor":
			result = expr.NewFloor(ex)
		case "sqrt":
			result = expr.NewSqrt(ex)
		case "log":
			result = expr.NewLog(ex)
		case "ln":
			result = expr.NewLn(ex)
		default:
			return nil, UnknownFunctionError
		}
	default:
		return nil, MissingFactorError
	}

	if neg {
		return expr.NewNeg(result), nil
	} else {
		return result, nil
	}
}
