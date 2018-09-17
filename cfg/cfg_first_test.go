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
			tgArithNlr,
			[]string{"F"},
			[]string{"\\(", "[[:digit:]]+"},
		},
		{
			tgArithNlr,
			[]string{"T"},
			[]string{"\\(", "[[:digit:]]+"},
		},
		{
			tgArithNlr,
			[]string{"E"},
			[]string{"\\(", "[[:digit:]]+"},
		},
		{
			tgArithNlr,
			[]string{"E'"},
			[]string{"\\+", ""},
		},
		{
			tgArithNlr,
			[]string{"T'"},
			[]string{"\\*", ""},
		},
		{
			tgArithNlr,
			[]string{"F", "T"},
			[]string{"\\(", "[[:digit:]]+"},
		},
		{
			tgArithNlr,
			[]string{"T", "E"},
			[]string{"\\(", "[[:digit:]]+"},
		},
		{
			tgArithNlr,
			[]string{"E", "F"},
			[]string{"\\(", "[[:digit:]]+"},
		},
		{
			tgArithNlr,
			[]string{"E'", "T'"},
			[]string{"\\*", "\\+", ""},
		},
		{
			tgArithNlr,
			[]string{"E'", "F"},
			[]string{"\\+", "\\(", "[[:digit:]]+"},
		},
		{
			tgArithNlr,
			[]string{"F", "E'"},
			[]string{"\\(", "[[:digit:]]+"},
		},
		{
			tgArithNlr,
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
