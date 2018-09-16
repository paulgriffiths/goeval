package cfg

import (
	"github.com/paulgriffiths/goeval/lar"
	"os"
	"testing"
)

func TestLexerSuccess(t *testing.T) {
	infiles := []string{
		"test_grammars/01_arith_lr.grammar",
		"test_grammars/02_arith_nlr.grammar",
		"test_grammars/03_arith_ambig.grammar",
		"test_grammars/04_bal_parens_1.grammar",
		"test_grammars/05_bal_parens_2.grammar",
		"test_grammars/06_zero_one.grammar",
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
		token{tokenNonTerminal, "E", lar.FilePos{0, 8}},
		token{tokenArrow, ":", lar.FilePos{7, 8}},
		token{tokenNonTerminal, "E", lar.FilePos{9, 8}},
		token{tokenTerminal, "\\+", lar.FilePos{11, 8}},
		token{tokenNonTerminal, "T", lar.FilePos{16, 8}},
		token{tokenAlt, "|", lar.FilePos{18, 8}},
		token{tokenNonTerminal, "T", lar.FilePos{20, 8}},
		token{tokenEndOfLine, "", lar.FilePos{21, 8}},
		token{tokenNonTerminal, "T", lar.FilePos{0, 9}},
		token{tokenArrow, ":", lar.FilePos{7, 9}},
		token{tokenNonTerminal, "T", lar.FilePos{9, 9}},
		token{tokenTerminal, "\\*", lar.FilePos{11, 9}},
		token{tokenNonTerminal, "F", lar.FilePos{16, 9}},
		token{tokenAlt, "|", lar.FilePos{18, 9}},
		token{tokenNonTerminal, "F", lar.FilePos{20, 9}},
		token{tokenEndOfLine, "", lar.FilePos{21, 9}},
		token{tokenNonTerminal, "F", lar.FilePos{0, 10}},
		token{tokenArrow, ":", lar.FilePos{7, 10}},
		token{tokenTerminal, "\\(", lar.FilePos{9, 10}},
		token{tokenNonTerminal, "E", lar.FilePos{14, 10}},
		token{tokenTerminal, "\\)", lar.FilePos{16, 10}},
		token{tokenAlt, "|", lar.FilePos{21, 10}},
		token{tokenNonTerminal, "Digits", lar.FilePos{23, 10}},
		token{tokenEndOfLine, "", lar.FilePos{29, 10}},
		token{tokenNonTerminal, "Digits", lar.FilePos{0, 11}},
		token{tokenArrow, ":", lar.FilePos{7, 11}},
		token{tokenTerminal, "[[:digit:]]+",
			lar.FilePos{9, 11}},
		token{tokenEndOfLine, "", lar.FilePos{23, 11}},
	}

	infileName := "test_grammars/01_arith_lr.grammar"
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
