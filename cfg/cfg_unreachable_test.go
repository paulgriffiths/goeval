package cfg

import "testing"

func TestCfgUnreachable(t *testing.T) {
	for _, tc := range grammarTestCases {
		c, err := getAndParseFile(t, tc.filename)
		if err != nil {
			t.Errorf("couldn't parse grammar file %q: %v", tc.filename, err)
			continue
		}

		if len(c.Unreachable()) != len(tc.unreachable) {
			t.Errorf("case %s, got %d, want %d", tc.filename,
				len(c.Unreachable()), len(tc.unreachable))
			continue
		}

		for n, nt := range c.Unreachable() {
			if r := c.NonTerminals[nt]; r != tc.unreachable[n] {
				t.Errorf("case %s, number %d, got %s, want %s",
					tc.filename, n+1, r, tc.unreachable[n])
			}
		}
	}
}
