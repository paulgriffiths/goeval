package cfl

import (
	"os"
	"testing"
)

func stringArraysEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	for n, s := range a {
		if s != b[n] {
			return false
		}
	}

	return true
}

func TestParserParseGoodGrammar(t *testing.T) {
	testCases := []string{
		"test_grammars/arith_lr.grammar",
		"test_grammars/arith_nlr.grammar",
		"test_grammars/arith_ambig.grammar",
		"test_grammars/zero_one.grammar",
		"test_grammars/bal_parens.grammar",
	}

	for _, tc := range testCases {
		infile, fileErr := os.Open(tc)
		if fileErr != nil {
			t.Errorf("couldn't open file %q: %v", tc, fileErr)
			continue
		}

		_, perr := parse(infile)
		if perr != nil {
			t.Errorf("couldn't get tokens for file %q: %v", tc, perr)
		}

		infile.Close()
	}
}

func TestParserFirstPass(t *testing.T) {
	testCases := []struct {
		filename     string
		nonTerminals []string
		terminals    []string
	}{
		{
			"test_grammars/arith_lr.grammar",
			[]string{"E", "T", "F", "Digits"},
			[]string{"+", "*", "(", ")",
				"(0|1|2|3|4|5|6|7|8|9)(0|1|2|3|4|5|6|7|8|9)*"},
		},
		{
			"test_grammars/arith_nlr.grammar",
			[]string{"E", "T", "E'", "F", "T'", "Digits"},
			[]string{"+", "*", "(", ")",
				"(0|1|2|3|4|5|6|7|8|9)(0|1|2|3|4|5|6|7|8|9)*"},
		},
		{
			"test_grammars/arith_ambig.grammar",
			[]string{"E", "Digits"},
			[]string{"+", "*", "(", ")",
				"(0|1|2|3|4|5|6|7|8|9)(0|1|2|3|4|5|6|7|8|9)*"},
		},
		{
			"test_grammars/zero_one.grammar",
			[]string{"S"},
			[]string{"0", "1", "01"},
		},
		{
			"test_grammars/bal_parens.grammar",
			[]string{"S"},
			[]string{"(", ")"},
		},
	}

	for _, tc := range testCases {
		infile, fileErr := os.Open(tc.filename)
		if fileErr != nil {
			t.Errorf("couldn't open file %q: %v", tc.filename, fileErr)
			continue
		}

		tokens, lerr := lex(infile)
		if lerr != nil {
			t.Errorf("couldn't get tokens for file %q: %v",
				tc.filename, lerr)
			infile.Close()
			continue
		}

		infile.Close()

		c := firstPass(tokens)

		if !stringArraysEqual(tc.nonTerminals, c.nonTerminals) {
			t.Errorf("%s, nonterminals, got %v, want %v", tc.filename,
				c.nonTerminals, tc.nonTerminals)
		}
		if !stringArraysEqual(tc.terminals, c.terminals) {
			t.Errorf("%s, terminals, got %v, want %v", tc.filename,
				c.terminals, tc.terminals)
		}

		for n, s := range c.nonTerminals {
			if c.ntTable[s] != n {
				t.Errorf("%s, got %d, want %d", tc.filename,
					c.ntTable[s], n)
			}
		}
		for n, s := range c.terminals {
			if c.tTable[s] != n {
				t.Errorf("%s, got %d, want %d", tc.filename,
					c.tTable[s], n)
			}
		}
	}
}
