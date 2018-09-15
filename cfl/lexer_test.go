package cfl

import (
	"github.com/paulgriffiths/goeval/lar"
	"os"
	"testing"
)

func TestLexerSuccess(t *testing.T) {
	infiles := []string{
		"test_grammars/arith_lr.grammar",
		"test_grammars/arith_nlr.grammar",
		"test_grammars/arith_ambig.grammar",
		"test_grammars/bal_parens.grammar",
		"test_grammars/zero_one.grammar",
	}

	for _, f := range infiles {
		infile, fileErr := os.Open(f)
		if fileErr != nil {
			t.Errorf("couldn't open file %q: %v", f, fileErr)
			continue
		}

		_, err := lex(infile)
		if err != nil {
			t.Errorf("couldn't get list of tokens for file %q", f)
		}

		infile.Close()
	}
}

func TestLexerTokenListArithLr(t *testing.T) {
	expected := tokenList{
		token{tokenNonTerminal, "E", lar.FilePos{0, 3}},
		token{tokenArrow, ":", lar.FilePos{7, 3}},
		token{tokenNonTerminal, "E", lar.FilePos{9, 3}},
		token{tokenTerminal, "+", lar.FilePos{11, 3}},
		token{tokenNonTerminal, "T", lar.FilePos{15, 3}},
		token{tokenAlt, "|", lar.FilePos{17, 3}},
		token{tokenNonTerminal, "T", lar.FilePos{19, 3}},
		token{tokenEndOfLine, "", lar.FilePos{20, 3}},
		token{tokenNonTerminal, "T", lar.FilePos{0, 4}},
		token{tokenArrow, ":", lar.FilePos{7, 4}},
		token{tokenNonTerminal, "T", lar.FilePos{9, 4}},
		token{tokenTerminal, "*", lar.FilePos{11, 4}},
		token{tokenNonTerminal, "F", lar.FilePos{15, 4}},
		token{tokenAlt, "|", lar.FilePos{17, 4}},
		token{tokenNonTerminal, "F", lar.FilePos{19, 4}},
		token{tokenEndOfLine, "", lar.FilePos{20, 4}},
		token{tokenNonTerminal, "F", lar.FilePos{0, 5}},
		token{tokenArrow, ":", lar.FilePos{7, 5}},
		token{tokenTerminal, "(", lar.FilePos{9, 5}},
		token{tokenNonTerminal, "E", lar.FilePos{13, 5}},
		token{tokenTerminal, ")", lar.FilePos{15, 5}},
		token{tokenAlt, "|", lar.FilePos{19, 5}},
		token{tokenNonTerminal, "Digits", lar.FilePos{21, 5}},
		token{tokenEndOfLine, "", lar.FilePos{27, 5}},
		token{tokenNonTerminal, "Digits", lar.FilePos{0, 9}},
		token{tokenArrow, ":", lar.FilePos{7, 9}},
		token{tokenTerminal, "(0|1|2|3|4|5|6|7|8|9)(0|1|2|3|4|5|6|7|8|9)*",
			lar.FilePos{9, 9}},
		token{tokenEndOfLine, "", lar.FilePos{54, 9}},
	}

	infileName := "test_grammars/arith_lr.grammar"
	infile, err := os.Open(infileName)
	if err != nil {
		t.Errorf("couldn't open file %q: %v", infileName, err)
		return
	}

	tokens, err := lex(infile)
	if err != nil {
		t.Errorf("couldn't get list of tokens")
		return
	}

	if len(tokens) != len(expected) {
		t.Errorf("Got %d tokens, want %d", len(tokens), len(expected))
		return
	}

	for n, token := range expected {
		if token != tokens[n] {
			t.Errorf("token %d, got %v, want %v", n+1, tokens[n], token)
		}
	}
}

func TestLexerErrors(t *testing.T) {
	testCases := []struct {
		filename string
		err      lexError
	}{
		{
			"test_grammars/bad/unterminated_terminal_1.grammar",
			lexError{lexErrUnterminatedTerminal, "", lar.FilePos{17, 3}},
		},
		{
			"test_grammars/bad/unterminated_terminal_2.grammar",
			lexError{lexErrUnterminatedTerminal, "", lar.FilePos{19, 5}},
		},
		{
			"test_grammars/bad/illegal_character_1.grammar",
			lexError{lexErrIllegalCharacter, "%", lar.FilePos{16, 3}},
		},
		{
			"test_grammars/bad/illegal_character_2.grammar",
			lexError{lexErrIllegalCharacter, "$", lar.FilePos{1, 3}},
		},
	}

	for n, tc := range testCases {
		infile, fileErr := os.Open(tc.filename)
		if fileErr != nil {
			t.Errorf("couldn't open file %q: %v", tc.filename, fileErr)
			continue
		}

		if _, err := lex(infile); err != tc.err {
			t.Errorf("case %d, got %v, want %v", n+1, err, tc.err)
		}

		infile.Close()
	}
}
