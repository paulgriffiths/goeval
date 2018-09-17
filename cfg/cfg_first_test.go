package cfg

import (
	"testing"
)

func TestCfgFirst(t *testing.T) {
	testCases := []struct {
		filename   string
		nt, result []string
	}{
		{
			"test_grammars/02_arith_nlr.grammar",
			[]string{"F"},
			[]string{"\\(", "[[:digit:]]+"},
		},
		{
			"test_grammars/02_arith_nlr.grammar",
			[]string{"T"},
			[]string{"\\(", "[[:digit:]]+"},
		},
		{
			"test_grammars/02_arith_nlr.grammar",
			[]string{"E"},
			[]string{"\\(", "[[:digit:]]+"},
		},
		{
			"test_grammars/02_arith_nlr.grammar",
			[]string{"E'"},
			[]string{"\\+", ""},
		},
		{
			"test_grammars/02_arith_nlr.grammar",
			[]string{"T'"},
			[]string{"\\*", ""},
		},
		{
			"test_grammars/02_arith_nlr.grammar",
			[]string{"F", "T"},
			[]string{"\\(", "[[:digit:]]+"},
		},
		{
			"test_grammars/02_arith_nlr.grammar",
			[]string{"T", "E"},
			[]string{"\\(", "[[:digit:]]+"},
		},
		{
			"test_grammars/02_arith_nlr.grammar",
			[]string{"E", "F"},
			[]string{"\\(", "[[:digit:]]+"},
		},
		{
			"test_grammars/02_arith_nlr.grammar",
			[]string{"E'", "T'"},
			[]string{"\\*", "\\+", ""},
		},
		{
			"test_grammars/02_arith_nlr.grammar",
			[]string{"E'", "F"},
			[]string{"\\+", "\\(", "[[:digit:]]+"},
		},
		{
			"test_grammars/02_arith_nlr.grammar",
			[]string{"F", "E'"},
			[]string{"\\(", "[[:digit:]]+"},
		},
		{
			"test_grammars/02_arith_nlr.grammar",
			[]string{"E'", "T'", "F"},
			[]string{"\\*", "\\+", "\\(", "[[:digit:]]+"},
		},
	}

	for n, tc := range testCases {
		c, err := getAndParseFile(t, tc.filename)
		if err != nil {
			t.Errorf("couldn't parse file %q: %v", tc.filename, err)
			continue
		}

		components := []BodyComp{}
		for _, nt := range tc.nt {
			components = append(components, c.NonTerminalComp(nt))
		}
		resultSet := c.First(components...)
		cmpSet := NewSetBodyComp()

		for _, s := range tc.result {
			if s == "" {
				cmpSet.Insert(BodyComp{BodyEmpty, 0})
				continue
			}
			cmpSet.Insert(c.TerminalComp(s))
		}

		if !resultSet.Equals(cmpSet) {
			t.Errorf("case %d, got %v, want %v", n+1, resultSet, cmpSet)
		}
	}
}
