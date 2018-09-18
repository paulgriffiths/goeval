package cfg

import "testing"

func TestCfgIsNullable(t *testing.T) {
	for _, tc := range grammarTestCases {
		c, err := getAndParseFile(t, tc.filename)
		if err != nil {
			t.Errorf("couldn't parse grammar file %q: %v", tc.filename, err)
			continue
		}

		if len(c.NonTerminalsNullable()) != len(tc.areNullable) {
			t.Errorf("case %s, got %d, want %d", tc.filename,
				len(c.NonTerminalsNullable()), len(tc.areNullable))
			continue
		}

		for n, nt := range c.NonTerminalsNullable() {
			if r := c.NonTerminals[nt]; r != tc.areNullable[n] {
				t.Errorf("case %s, number %d, got %s, want %s",
					tc.filename, n+1, r, tc.areNullable[n])
			}
		}
	}
}
