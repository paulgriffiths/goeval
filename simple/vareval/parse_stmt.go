package vareval

import (
	"fmt"
	"github.com/paulgriffiths/goeval/tokens"
	"strings"
)

// ParseStatement parses a statement provided in string form.
func ParseStatement(stmt string) (Stmt, error) {
	ch, err := NewLexer(strings.NewReader(stmt))
	if err != nil {
		return nil, fmt.Errorf("couldn't create lexer: %v", err)
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
		if !ltc.MatchIdentifier() {
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
