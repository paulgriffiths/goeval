package eval

import (
	"github.com/paulgriffiths/goeval/tokens"
	"strings"
	"testing"
)

func TestLexerTokenValues(t *testing.T) {
	cases := []struct {
		input  string
		values []string
	}{
		{"", []string{}},
		{"~", []string{"~"}},
		{"(", []string{"("}},
		{")", []string{")"}},
		{"+", []string{"+"}},
		{"-", []string{"-"}},
		{"*", []string{"*"}},
		{"/", []string{"/"}},
		{"+-", []string{"+", "-"}},
		{"   +  -     ", []string{"+", "-"}},
		{"cos", []string{"cos"}},
		{"  cos   sin   ", []string{"cos", "sin"}},
		{"123", []string{"123"}},
		{"123+", []string{"123", "+"}},
		{"1e6", []string{"1e6"}},
		{"1e6+", []string{"1e6", "+"}},
		{"1.23", []string{"1.23"}},
		{"1.23+", []string{"1.23", "+"}},
		{"1.23e45", []string{"1.23e45"}},
		{"1.23e45+", []string{"1.23e45", "+"}},
		{"1.23e-45", []string{"1.23e-45"}},
		{"1.23e+45", []string{"1.23e+45"}},
		{" (  1.2e-3*(cos(45  ) +sin8.7)  )~?  ",
			[]string{"(", "1.2e-3", "*", "(", "cos", "(", "45",
				")", "+", "sin", "8.7", ")", ")", "~", "?"}},
	}

	for n, c := range cases {
		ch, err := NewLexer(strings.NewReader(c.input))
		if err != nil {
			t.Errorf("Couldn't create lexer: %v", err)
		}

		for _, v := range c.values {
			if token := <-ch; token.Value() != v {
				t.Errorf("Input '%s', got %v, want %v", c.input, c.values, v)
			}
		}

		if _, ok := <-ch; ok != false {
			t.Errorf("Case %d, end of channel not encountered when expected", n)
		}
	}

}

func TestLexerTokenTypes(t *testing.T) {
	cases := []struct {
		input  string
		values []func(*tokens.Token) bool
	}{
		{"0", []func(*tokens.Token) bool{(*tokens.Token).IsNumber}},
		{"1", []func(*tokens.Token) bool{(*tokens.Token).IsNumber}},
		{"123", []func(*tokens.Token) bool{(*tokens.Token).IsNumber}},
		{"1.23", []func(*tokens.Token) bool{(*tokens.Token).IsNumber}},
		{"1e6", []func(*tokens.Token) bool{(*tokens.Token).IsNumber}},
		{"1.23e6", []func(*tokens.Token) bool{(*tokens.Token).IsNumber}},
		{"1e-6", []func(*tokens.Token) bool{(*tokens.Token).IsNumber}},
		{"1.23e-6", []func(*tokens.Token) bool{(*tokens.Token).IsNumber}},
		{"cos", []func(*tokens.Token) bool{(*tokens.Token).IsWord}},
		{"+", []func(*tokens.Token) bool{(*tokens.Token).IsOperator}},
		{"-", []func(*tokens.Token) bool{(*tokens.Token).IsOperator}},
		{"*", []func(*tokens.Token) bool{(*tokens.Token).IsOperator}},
		{"/", []func(*tokens.Token) bool{(*tokens.Token).IsOperator}},
		{"^", []func(*tokens.Token) bool{(*tokens.Token).IsOperator}},
		{"(", []func(*tokens.Token) bool{(*tokens.Token).IsLeftParen}},
		{")", []func(*tokens.Token) bool{(*tokens.Token).IsRightParen}},
		{"~", []func(*tokens.Token) bool{(*tokens.Token).IsIllegal}},
		{"?", []func(*tokens.Token) bool{(*tokens.Token).IsIllegal}},
		{"$", []func(*tokens.Token) bool{(*tokens.Token).IsIllegal}},
		{"&", []func(*tokens.Token) bool{(*tokens.Token).IsIllegal}},
		{"1e", []func(*tokens.Token) bool{(*tokens.Token).IsIllegal}},
		{"1.", []func(*tokens.Token) bool{(*tokens.Token).IsIllegal}},
		{"1e.", []func(*tokens.Token) bool{
			(*tokens.Token).IsIllegal,
			(*tokens.Token).IsIllegal,
		}},
	}

	for n, c := range cases {
		ch, err := NewLexer(strings.NewReader(c.input))
		if err != nil {
			t.Errorf("Couldn't create lexer: %v", err)
		}

		for _, v := range c.values {
			if token := <-ch; !v(&token) {
				t.Errorf("Input '%s', unexpected type", c.input)
			}
		}

		if _, ok := <-ch; ok != false {
			t.Errorf("Case %d, end of channel not encountered when expected", n)
		}
	}
}
