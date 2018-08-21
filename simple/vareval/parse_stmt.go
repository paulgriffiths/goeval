package vareval

import (
	"fmt"
	"github.com/paulgriffiths/goeval/tokens"
	"strings"
)

func ParseStatement(stmt string) (Stmt, error) {
	ch, err := NewLexer(strings.NewReader(stmt))
	if err != nil {
		return nil, fmt.Errorf("Couldn't create lexer: %v", err)
	}

	ltc := tokens.NewLTChan(ch)

	switch {
	case ltc.Match(tokens.KeywordToken("print")):
		exp, err := getExpr(ltc)
		if err != nil {
			return nil, err
		}
		if !ltc.IsEmpty() {
			return nil, TrailingTokensError
		}
		return NewOutputStatement(exp), nil
	case ltc.Match(tokens.KeywordToken("let")):
		if !ltc.MatchType(tokens.EmptyIdentifierToken()) {
			return nil, SyntaxError
		}
		id := string(ltc.Value())
		if _, ok := keywords[id]; ok {
			return nil, IllegalIdentifierError
		}
		if !ltc.Match(tokens.OperatorToken("=")) {
			return nil, SyntaxError
		}
		exp, err := getExpr(ltc)
		if err != nil {
			return nil, err
		}
		if !ltc.IsEmpty() {
			return nil, TrailingTokensError
		}
		return NewAssignStatement(id, exp), nil
	default:
		exp, err := getExpr(ltc)
		if err != nil {
			return nil, err
		}
		if !ltc.IsEmpty() {
			return nil, TrailingTokensError
		}
		return NewOutputExprStatement(exp), nil
	}

	return nil, UnknownError
}
