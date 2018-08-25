package vareval

import (
	"fmt"
	"github.com/paulgriffiths/goeval/expr"
	"github.com/paulgriffiths/goeval/tokens"
	"math"
	"strconv"
)

type opPair struct {
	tokenType tokens.TokenType
	opType    func(left, right expr.Expr) expr.Expr
}

type binaryTerm struct {
	nextTermFunc func(ltchan *tokens.LTChan) (expr.Expr, error)
	ops          []opPair
}

func getBinaryTerm(ltchan *tokens.LTChan, term binaryTerm) (expr.Expr, error) {
	var left expr.Expr
	left, err := term.nextTermFunc(ltchan)
	if err != nil {
		return nil, err
	}

loop:
	for {
		for _, pair := range term.ops {
			if ltchan.Match(pair.tokenType) {
				right, err := term.nextTermFunc(ltchan)
				if err != nil {
					return nil, err
				}
				left = pair.opType(left, right)
				continue loop
			}
		}
		break
	}
	return left, nil
}

func getExpr(ltchan *tokens.LTChan) (expr.Expr, error) {
	return getTermEquality(ltchan)
}

func getTermEquality(ltchan *tokens.LTChan) (expr.Expr, error) {
	term := binaryTerm{getTermLessThanGreaterThan,
		[]opPair{
			{tokens.EqualityOperator, expr.NewEquality},
			{tokens.NonEqualityOperator, expr.NewNonEquality},
		},
	}
	return getBinaryTerm(ltchan, term)
}

func getTermLessThanGreaterThan(ltchan *tokens.LTChan) (expr.Expr, error) {
	term := binaryTerm{getTermLogicalOr,
		[]opPair{
			{tokens.LessOperator, expr.NewLessThan},
			{tokens.LessEqualOperator, expr.NewLessThanOrEqual},
			{tokens.GreaterOperator, expr.NewGreaterThan},
			{tokens.GreaterEqualOperator, expr.NewGreaterThanOrEqual},
		},
	}
	return getBinaryTerm(ltchan, term)
}

func getTermLogicalOr(ltchan *tokens.LTChan) (expr.Expr, error) {
	term := binaryTerm{getTermLogicalAnd,
		[]opPair{{tokens.OrOperator, expr.NewOr}},
	}
	return getBinaryTerm(ltchan, term)
}

func getTermLogicalAnd(ltchan *tokens.LTChan) (expr.Expr, error) {
	term := binaryTerm{getTermPlusMinus,
		[]opPair{{tokens.AndOperator, expr.NewAnd}},
	}
	return getBinaryTerm(ltchan, term)
}

func getTermPlusMinus(ltchan *tokens.LTChan) (expr.Expr, error) {
	term := binaryTerm{getTermMultiplyDivide,
		[]opPair{
			{tokens.AddOperator, expr.NewAdd},
			{tokens.SubOperator, expr.NewSub},
		},
	}
	return getBinaryTerm(ltchan, term)
}

func getTermMultiplyDivide(ltchan *tokens.LTChan) (expr.Expr, error) {
	term := binaryTerm{getTermPower,
		[]opPair{
			{tokens.MulOperator, expr.NewMul},
			{tokens.DivOperator, expr.NewDiv},
		},
	}
	return getBinaryTerm(ltchan, term)
}

func getTermPower(ltchan *tokens.LTChan) (expr.Expr, error) {
	term := binaryTerm{getTermNegate,
		[]opPair{{tokens.PowOperator, expr.NewPow}},
	}
	return getBinaryTerm(ltchan, term)
}

func getTermNegate(ltchan *tokens.LTChan) (expr.Expr, error) {
	neg := false
	if ltchan.Match(tokens.SubOperator) {
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
	case ltchan.Match(tokens.RightParen):
		return nil, UnbalancedParenthesesError
	case ltchan.Match(tokens.LeftParen):
		ex, err := getExpr(ltchan)
		if err != nil {
			return nil, err
		}
		if !ltchan.Match(tokens.RightParen) {
			return nil, UnbalancedParenthesesError
		}
		return ex, nil
	case ltchan.Match(tokens.String):
		return expr.NewString(ltchan.Value()), nil
	case ltchan.Match(tokens.Number):
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
	case ltchan.Match(tokens.Keyword):
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
			if !ltchan.Match(tokens.LeftParen) {
				return nil, MissingArgumentError
			}
			ex, err := getExpr(ltchan)
			if err != nil {
				return nil, err
			}
			if !ltchan.Match(tokens.RightParen) {
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
	case ltchan.Match(tokens.Identifier):
		word := string(ltchan.Value())
		return expr.NewVariable(word), nil
	case ltchan.Match(tokens.Illegal):
		return nil, SyntaxError
	}
	return nil, MissingFactorError
}
