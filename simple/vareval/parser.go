package vareval

import (
	"fmt"
	"github.com/paulgriffiths/goeval/expr"
	"github.com/paulgriffiths/goeval/tokens"
	"math"
	"strconv"
)

func getExpr(ltchan *tokens.LTChan) (expr.Expr, error) {
	return getTermEquality(ltchan)
}

func getTermEquality(ltchan *tokens.LTChan) (expr.Expr, error) {
	var left expr.Expr
	left, err := getTermLessThanGreaterThan(ltchan)
	if err != nil {
		return nil, err
	}

loop:
	for {
		switch {
		case ltchan.Match(tokens.OperatorToken("==")):
			right, err := getTermLessThanGreaterThan(ltchan)
			if err != nil {
				return nil, err
			}

			left = expr.NewEquality(left, right)
		case ltchan.Match(tokens.OperatorToken("!=")):
			right, err := getTermLessThanGreaterThan(ltchan)
			if err != nil {
				return nil, err
			}

			left = expr.NewNonEquality(left, right)
		default:
			break loop
		}
	}
	return left, nil
}

func getTermLessThanGreaterThan(ltchan *tokens.LTChan) (expr.Expr, error) {
	var left expr.Expr
	left, err := getTermLogicalOr(ltchan)
	if err != nil {
		return nil, err
	}

loop:
	for {
		switch {
		case ltchan.Match(tokens.OperatorToken("<")):
			right, err := getTermLogicalOr(ltchan)
			if err != nil {
				return nil, err
			}

			left = expr.NewLessThan(left, right)
		case ltchan.Match(tokens.OperatorToken("<=")):
			right, err := getTermLogicalOr(ltchan)
			if err != nil {
				return nil, err
			}

			left = expr.NewLessThanOrEqual(left, right)
		case ltchan.Match(tokens.OperatorToken(">")):
			right, err := getTermLogicalOr(ltchan)
			if err != nil {
				return nil, err
			}

			left = expr.NewGreaterThan(left, right)
		case ltchan.Match(tokens.OperatorToken(">=")):
			right, err := getTermLogicalOr(ltchan)
			if err != nil {
				return nil, err
			}

			left = expr.NewGreaterThanOrEqual(left, right)
		default:
			break loop
		}
	}
	return left, nil
}

func getTermLogicalOr(ltchan *tokens.LTChan) (expr.Expr, error) {
	var left expr.Expr
	left, err := getTermLogicalAnd(ltchan)
	if err != nil {
		return nil, err
	}

loop:
	for {
		switch {
		case ltchan.Match(tokens.KeywordToken("or")):
			right, err := getTermLogicalAnd(ltchan)
			if err != nil {
				return nil, err
			}

			left = expr.NewOr(left, right)
		default:
			break loop
		}
	}
	return left, nil
}

func getTermLogicalAnd(ltchan *tokens.LTChan) (expr.Expr, error) {
	var left expr.Expr
	left, err := getTermPlusMinus(ltchan)
	if err != nil {
		return nil, err
	}

loop:
	for {
		switch {
		case ltchan.Match(tokens.KeywordToken("and")):
			right, err := getTermPlusMinus(ltchan)
			if err != nil {
				return nil, err
			}

			left = expr.NewAnd(left, right)
		default:
			break loop
		}
	}
	return left, nil
}

func getTermPlusMinus(ltchan *tokens.LTChan) (expr.Expr, error) {
	var left expr.Expr
	left, err := getTermMultiplyDivide(ltchan)
	if err != nil {
		return nil, err
	}

loop:
	for {
		switch {
		case ltchan.Match(tokens.OperatorToken("+")):
			right, err := getTermMultiplyDivide(ltchan)
			if err != nil {
				return nil, err
			}
			left = expr.NewAdd(left, right)
		case ltchan.Match(tokens.OperatorToken("-")):
			right, err := getTermMultiplyDivide(ltchan)
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

func getTermMultiplyDivide(ltchan *tokens.LTChan) (expr.Expr, error) {
	var left expr.Expr
	left, err := getTermPower(ltchan)
	if err != nil {
		return nil, err
	}

loop:
	for {
		switch {
		case ltchan.Match(tokens.OperatorToken("*")):
			right, err := getTermPower(ltchan)
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

func getTermPower(ltchan *tokens.LTChan) (expr.Expr, error) {
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
	case ltchan.MatchType(tokens.EmptyKeywordToken()):
		word := string(ltchan.Value())
		switch word {
		case "e":
			result = expr.NewReal(math.E)
		case "pi":
			result = expr.NewReal(math.Pi)
		case "true":
			result = expr.NewBool(true)
		case "false":
			result = expr.NewBool(false)
		}
		if _, ok := functions[word]; ok {
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
		}
	case ltchan.MatchType(tokens.EmptyIdentifierToken()):
		word := string(ltchan.Value())
		result = expr.NewVariable(word)
	default:
		return nil, MissingFactorError
	}

	if neg {
		return expr.NewNeg(result), nil
	} else {
		return result, nil
	}
}
