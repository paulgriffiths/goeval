// Package eval provides a simple mathematical expression evaluator.

package vareval

import (
	"fmt"
	"github.com/paulgriffiths/goeval/expr"
	"github.com/paulgriffiths/goeval/tokens"
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

	value, ok := expr.ToValue(result)
	if !ok {
		return nil, UnknownError
	}

	return value, nil
}
