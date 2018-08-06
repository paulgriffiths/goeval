package eval

import (
    "strings"
    "testing"
)

func TestLexerEmptyInput(t *testing.T) {
    ch, err := NewLexer(strings.NewReader(""))
    if err != nil {
        t.Errorf("Couldn't create lexer: %v", err)
    }
    if _, ok := <-ch; ok != false {
        t.Errorf("Got %v, want %v", ok, false)
    }
}

func TestLexerLeftParenToken(t *testing.T) {
    ch, err := NewLexer(strings.NewReader("("))
    if err != nil {
        t.Errorf("Couldn't create lexer: %v", err)
    }
    if token := <-ch; token.Value() != "(" {
        t.Errorf("Got value %s, want %s", token.Value(), ")")
    }
    if _, ok := <-ch; ok != false {
        t.Errorf("Got %v, want %v", ok, false)
    }
}

func TestLexerRightParenToken(t *testing.T) {
    ch, err := NewLexer(strings.NewReader(")"))
    if err != nil {
        t.Errorf("Couldn't create lexer: %v", err)
    }
    if token := <-ch; token.Value() != ")" {
        t.Errorf("Got value %s, want %s", token.Value(), ")")
    }
    if _, ok := <-ch; ok != false {
        t.Errorf("Got %v, want %v", ok, false)
    }
}

func TestLexerOperatorToken(t *testing.T) {
    ch, err := NewLexer(strings.NewReader("+"))
    if err != nil {
        t.Errorf("Couldn't create lexer: %v", err)
    }
    if token := <-ch; token.Value() != "+" {
        t.Errorf("Got value %s, want %s", token.Value(), "+")
    }
    if _, ok := <-ch; ok != false {
        t.Errorf("Got %v, want %v", ok, false)
    }
}

func TestLexerTwoOperatorTokens(t *testing.T) {
    ch, err := NewLexer(strings.NewReader("+-"))
    if err != nil {
        t.Errorf("Couldn't create lexer: %v", err)
    }
    if token := <-ch; token.Value() != "+" {
        t.Errorf("Got value %s, want %s", token.Value(), "+")
    }
    if token := <-ch; token.Value() != "-" {
        t.Errorf("Got value %s, want %s", token.Value(), "-")
    }
    if _, ok := <-ch; ok != false {
        t.Errorf("Got %v, want %v", ok, false)
    }
}

func TestLexerTwoOperatorTokensWithWhitespace(t *testing.T) {
    ch, err := NewLexer(strings.NewReader("     +  -        "))
    if err != nil {
        t.Errorf("Couldn't create lexer: %v", err)
    }
    if token := <-ch; token.Value() != "+" {
        t.Errorf("Got value %s, want %s", token.Value(), "+")
    }
    if token := <-ch; token.Value() != "-" {
        t.Errorf("Got value %s, want %s", token.Value(), "-")
    }
    if _, ok := <-ch; ok != false {
        t.Errorf("Got %v, want %v", ok, false)
    }
}

func TestLexerFunctionToken(t *testing.T) {
    ch, err := NewLexer(strings.NewReader("cos"))
    if err != nil {
        t.Errorf("Couldn't create lexer: %v", err)
    }
    if token := <-ch; token.Value() != "cos" {
        t.Errorf("Got value %s, want %s", token.Value(), "cos")
    }
    if _, ok := <-ch; ok != false {
        t.Errorf("Got %v, want %v", ok, false)
    }
}

func TestLexerTwoFunctionTokens(t *testing.T) {
    ch, err := NewLexer(strings.NewReader("  cos  sin  "))
    if err != nil {
        t.Errorf("Couldn't create lexer: %v", err)
    }
    if token := <-ch; token.Value() != "cos" {
        t.Errorf("Got value %s, want %s", token.Value(), "cos")
    }
    if token := <-ch; token.Value() != "sin" {
        t.Errorf("Got value %s, want %s", token.Value(), "sin")
    }
    if _, ok := <-ch; ok != false {
        t.Errorf("Got %v, want %v", ok, false)
    }
}

func TestLexerSimpleIntegerNumberToken(t *testing.T) {
    ch, err := NewLexer(strings.NewReader("123"))
    if err != nil {
        t.Errorf("Couldn't create lexer: %v", err)
    }
    if token := <-ch; token.Value() != "123" {
        t.Errorf("Got value %s, want %s", token.Value(), "123")
    }
    if _, ok := <-ch; ok != false {
        t.Errorf("Got %v, want %v", ok, false)
    }
}

func TestLexerSimpleIntegerNumberTokenWithTrailing(t *testing.T) {
    ch, err := NewLexer(strings.NewReader("123+"))
    if err != nil {
        t.Errorf("Couldn't create lexer: %v", err)
    }
    if token := <-ch; token.Value() != "123" {
        t.Errorf("Got value %s, want %s", token.Value(), "123")
    }
    if token := <-ch; token.Value() != "+" {
        t.Errorf("Got value %s, want %s", token.Value(), "+")
    }
    if _, ok := <-ch; ok != false {
        t.Errorf("Got %v, want %v", ok, false)
    }
}

func TestLexerExponentIntegerNumberToken(t *testing.T) {
    ch, err := NewLexer(strings.NewReader("1e6"))
    if err != nil {
        t.Errorf("Couldn't create lexer: %v", err)
    }
    if token := <-ch; token.Value() != "1e6" {
        t.Errorf("Got value %s, want %s", token.Value(), "1e6")
    }
    if _, ok := <-ch; ok != false {
        t.Errorf("Got %v, want %v", ok, false)
    }
}

func TestLexerExponentIntegerNumberTokenWithTrailing(t *testing.T) {
    ch, err := NewLexer(strings.NewReader("1e6+"))
    if err != nil {
        t.Errorf("Couldn't create lexer: %v", err)
    }
    if token := <-ch; token.Value() != "1e6" {
        t.Errorf("Got value %s, want %s", token.Value(), "1e6")
    }
    if token := <-ch; token.Value() != "+" {
        t.Errorf("Got value %s, want %s", token.Value(), "+")
    }
    if _, ok := <-ch; ok != false {
        t.Errorf("Got %v, want %v", ok, false)
    }
}

func TestLexerBadExponentIntegerNumberToken(t *testing.T) {
    ch, err := NewLexer(strings.NewReader("1e/"))
    if err != nil {
        t.Errorf("Couldn't create lexer: %v", err)
    }
    if token := <-ch; !token.IsIllegal() || token.Value() != "1e" {
        t.Errorf("Didn't get illegal token as expected")
        t.Errorf("Got value %s, want %s", token.Value(), "1e")
    }
    if token := <-ch; token.Value() != "/" {
        t.Errorf("Got value %s, want %s", token.Value(), "/")
    }
    if _, ok := <-ch; ok != false {
        t.Errorf("Got %v, want %v", ok, false)
    }
}

func TestLexerRealNumberToken(t *testing.T) {
    ch, err := NewLexer(strings.NewReader("1.23"))
    if err != nil {
        t.Errorf("Couldn't create lexer: %v", err)
    }
    if token := <-ch; token.Value() != "1.23" {
        t.Errorf("Got value %s, want %s", token.Value(), "1.23")
    }
    if _, ok := <-ch; ok != false {
        t.Errorf("Got %v, want %v", ok, false)
    }
}

func TestLexerRealNumberTokenWithTrailing(t *testing.T) {
    ch, err := NewLexer(strings.NewReader("1.23+"))
    if err != nil {
        t.Errorf("Couldn't create lexer: %v", err)
    }
    if token := <-ch; token.Value() != "1.23" {
        t.Errorf("Got value %s, want %s", token.Value(), "1.23")
    }
    if token := <-ch; token.Value() != "+" {
        t.Errorf("Got value %s, want %s", token.Value(), "+")
    }
    if _, ok := <-ch; ok != false {
        t.Errorf("Got %v, want %v", ok, false)
    }
}

func TestLexerRealExponentNumberToken(t *testing.T) {
    ch, err := NewLexer(strings.NewReader("1.23e72"))
    if err != nil {
        t.Errorf("Couldn't create lexer: %v", err)
    }
    if token := <-ch; token.Value() != "1.23e72" {
        t.Errorf("Got value %s, want %s", token.Value(), "1.23e72")
    }
    if _, ok := <-ch; ok != false {
        t.Errorf("Got %v, want %v", ok, false)
    }
}

func TestLexerRealNegativeExponentNumberToken(t *testing.T) {
    ch, err := NewLexer(strings.NewReader("1.23e-45"))
    if err != nil {
        t.Errorf("Couldn't create lexer: %v", err)
    }
    if token := <-ch; token.Value() != "1.23e-45" {
        t.Errorf("Got value %s, want %s", token.Value(), "1.23e-45")
    }
    if _, ok := <-ch; ok != false {
        t.Errorf("Got %v, want %v", ok, false)
    }
}

func TestLexerRealExplicitPositiveExponentNumberToken(t *testing.T) {
    ch, err := NewLexer(strings.NewReader("1.23e+45"))
    if err != nil {
        t.Errorf("Couldn't create lexer: %v", err)
    }
    if token := <-ch; token.Value() != "1.23e+45" {
        t.Errorf("Got value %s, want %s", token.Value(), "1.23e+45")
    }
    if _, ok := <-ch; ok != false {
        t.Errorf("Got %v, want %v", ok, false)
    }
}

func TestLexerComplexExpression(t *testing.T) {
    ch, err := NewLexer(strings.NewReader("(1.2e-3*(cos45+8.7))"))
    if err != nil {
        t.Errorf("Couldn't create lexer: %v", err)
    }
    if token := <-ch; token.Value() != "(" {
        t.Errorf("Got value %s, want %s", token.Value(), "(")
    }
    if token := <-ch; token.Value() != "1.2e-3" {
        t.Errorf("Got value %s, want %s", token.Value(), "1.2e-3")
    }
    if token := <-ch; token.Value() != "*" {
        t.Errorf("Got value %s, want %s", token.Value(), "*")
    }
    if token := <-ch; token.Value() != "(" {
        t.Errorf("Got value %s, want %s", token.Value(), "(")
    }
    if token := <-ch; token.Value() != "cos" {
        t.Errorf("Got value %s, want %s", token.Value(), "cos")
    }
    if token := <-ch; token.Value() != "45" {
        t.Errorf("Got value %s, want %s", token.Value(), "45")
    }
    if token := <-ch; token.Value() != "+" {
        t.Errorf("Got value %s, want %s", token.Value(), "+")
    }
    if token := <-ch; token.Value() != "8.7" {
        t.Errorf("Got value %s, want %s", token.Value(), "8.7")
    }
    if token := <-ch; token.Value() != ")" {
        t.Errorf("Got value %s, want %s", token.Value(), ")")
    }
    if token := <-ch; token.Value() != ")" {
        t.Errorf("Got value %s, want %s", token.Value(), ")")
    }
    if _, ok := <-ch; ok != false {
        t.Errorf("Got %v, want %v", ok, false)
    }
}

