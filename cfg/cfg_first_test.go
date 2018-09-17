package cfg

import (
	"testing"
)

func TestCfgFirst(t *testing.T) {
	testCases := []struct {
		filename string
		nt       string
		result   []string
	}{
		{
			"test_grammars/02_arith_nlr.grammar",
			"F",
			[]string{"\\(", "[[:digit:]]+"},
		},
		{
			"test_grammars/02_arith_nlr.grammar",
			"T",
			[]string{"\\(", "[[:digit:]]+"},
		},
		{
			"test_grammars/02_arith_nlr.grammar",
			"E",
			[]string{"\\(", "[[:digit:]]+"},
		},
		{
			"test_grammars/02_arith_nlr.grammar",
			"E'",
			[]string{"\\+", ""},
		},
		{
			"test_grammars/02_arith_nlr.grammar",
			"T'",
			[]string{"\\*", ""},
		},
	}

	for n, tc := range testCases {
		c, err := getAndParseFile(t, tc.filename)
		if err != nil {
			t.Errorf("couldn't parse file %q: %v", tc.filename, err)
			continue
		}

		resultSet := c.First(c.NonTerminalComp(tc.nt))
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
