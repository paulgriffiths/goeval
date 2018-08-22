package vareval

import (
	"fmt"
	"github.com/paulgriffiths/goeval/expr"
	"github.com/paulgriffiths/goeval/tokens"
	"strings"
)

func evaluate(expression string, table *expr.SymTab) (expr.Expr, error) {
	ch, err := NewLexer(strings.NewReader(expression))
	if err != nil {
		return nil, fmt.Errorf("couldn't create lexer: %v", err)
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

	result, err := exp.Evaluate(table)
	if err != nil {
		return nil, err
	}

	return result, nil
}
