// Package eval provides a simple mathematical expression evaluator.

package eval

import (
	"fmt"
	"github.com/paulgriffiths/goeval/tokens"
	"math"
	"strconv"
	"strings"
)

// Evaluates the provided simple mathematical expression.
func Evaluate(expression string) (float64, error) {
	ch, err := NewLexer(strings.NewReader(expression))
	if err != nil {
		return 1, fmt.Errorf("Couldn't create lexer: %v", err)
	}

	ltc := tokens.NewLTChan(ch)
	expr, err := getExpr(ltc)
	if err != nil {
		return 1, err
	}

	if !ltc.IsEmpty() {
		ltc.Flush()
		return 1, TrailingTokensError
	}

	value, err := expr.Evaluate()
	if err != nil {
		return 1, err
	}
	return value, nil
}

func getExpr(ltchan *tokens.LTChan) (expr, error) {
	var left expr
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
			left = add{left, right}
		case ltchan.Match(tokens.OperatorToken("-")):
			right, err := getTerm(ltchan)
			if err != nil {
				return nil, err
			}
			left = subtract{left, right}
		default:
			break loop
		}
	}
	return left, nil
}

func getTerm(ltchan *tokens.LTChan) (expr, error) {
	var left expr
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
			left = multiply{left, right}
		case ltchan.Match(tokens.OperatorToken("/")):
			right, err := getFactor(ltchan)
			if err != nil {
				return nil, err
			}
			left = divide{left, right}
		default:
			break loop
		}
	}
	return left, nil
}

func getSubTerm(ltchan *tokens.LTChan) (expr, error) {
	var left expr
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

			left = power{left, right}
		default:
			break loop
		}
	}
	return left, nil
}

func getFactor(ltchan *tokens.LTChan) (expr, error) {
	neg := false
	if ltchan.Match(tokens.OperatorToken("-")) {
		neg = true
	}
	var result expr

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
		value, err := strconv.ParseFloat(ltchan.Value(), 64)
		if err != nil {
			panic(fmt.Sprintf("Couldn't convert to float: %s", ltchan.Value()))
		}
		result = number{value}
	case ltchan.MatchType(tokens.EmptyWordToken()):
		word := string(ltchan.Value())

		if word == "e" {
			result = number{math.E}
			break
		} else if word == "pi" {
			result = number{math.Pi}
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
			result = cos{ex}
		case "sin":
			result = sin{ex}
		case "tan":
			result = tan{ex}
		case "acos":
			result = acos{ex}
		case "asin":
			result = asin{ex}
		case "atan":
			result = atan{ex}
		case "round":
			result = round{ex}
		case "ceil":
			result = ceil{ex}
		case "floor":
			result = floor{ex}
		case "sqrt":
			result = sqrt{ex}
		case "log":
			result = log{ex}
		case "ln":
			result = ln{ex}
		default:
			return nil, UnknownFunctionError
		}
	default:
		return nil, MissingFactorError
	}

	if neg {
		return negate{result}, nil
	} else {
		return result, nil
	}
}
