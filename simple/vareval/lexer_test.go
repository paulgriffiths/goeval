package vareval

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
				")", "+", "sin8", ".", "7", ")", ")", "~", "?"}},
	}

	for n, c := range cases {
		ch, err := NewLexer(strings.NewReader(c.input))
		if err != nil {
			t.Errorf("couldn't create lexer: %v", err)
		}

		for _, v := range c.values {
			if token := <-ch; token.Value != v {
				t.Errorf("input '%s', got %v, want %v",
					c.input, c.values, v)
			}
		}

		if _, ok := <-ch; ok != false {
			t.Errorf("case %d, end of channel not encountered", n)
		}
	}

}

func TestLexerTokenTypes(t *testing.T) {
	cases := []struct {
		input string
		value tokens.TokenType
	}{
		{"0", tokens.Number},
		{"1", tokens.Number},
		{"123", tokens.Number},
		{"1.23", tokens.Number},
		{"1e6", tokens.Number},
		{"1.23e6", tokens.Number},
		{"1e-6", tokens.Number},
		{"1.23e-6", tokens.Number},
		{"cos", tokens.Keyword},
		{"foo", tokens.Identifier},
		{"ha12", tokens.Identifier},
		{"+", tokens.AddOperator},
		{"-", tokens.SubOperator},
		{"*", tokens.MulOperator},
		{"/", tokens.DivOperator},
		{"^", tokens.PowOperator},
		{"(", tokens.LeftParen},
		{")", tokens.RightParen},
		{"~", tokens.Illegal},
		{"?", tokens.Illegal},
		{"$", tokens.Illegal},
		{"&", tokens.Illegal},
		{"1e", tokens.Illegal},
		{"1.", tokens.Illegal},
	}

	for n, c := range cases {
		ch, err := NewLexer(strings.NewReader(c.input))
		if err != nil {
			t.Errorf("couldn't create lexer: %v", err)
		}

		if token := <-ch; token.Type != c.value {
			t.Errorf("Input '%s', unexpected type", c.input)
		}

		if _, ok := <-ch; ok != false {
			t.Errorf("case %d, end of channel not encountered", n)
		}
	}
}
