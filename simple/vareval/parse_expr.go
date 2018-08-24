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
			right, err := getTermPower(ltchan)
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
	case ltchan.MatchType(tokens.EmptyStringToken()):
		return expr.NewString(ltchan.Value()), nil
	case ltchan.MatchType(tokens.ZeroNumberToken()):
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
	case ltchan.MatchType(tokens.EmptyKeywordToken()):
		word := string(ltchan.Value())
		switch word {
		case "e":
			return expr.NewReal(math.E), nil
		case "pi":
			return expr.NewReal(math.Pi), nil
		case "true":
			return expr.NewBool(true), nil
		case "false":
			return expr.NewBool(false), nil
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
		}
	case ltchan.MatchType(tokens.EmptyIdentifierToken()):
		word := string(ltchan.Value())
		return expr.NewVariable(word), nil
	case ltchan.MatchType(tokens.EmptyIllegalToken()):
		return nil, SyntaxError
	}
	return nil, MissingFactorError
}
