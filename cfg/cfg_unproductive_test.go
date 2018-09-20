package cfg

import "testing"

func TestCfgUnproductive(t *testing.T) {
	for _, tc := range grammarTestCases {
		c, err := getAndParseFile(t, tc.filename)
		if err != nil {
			t.Errorf("couldn't parse grammar file %q: %v", tc.filename, err)
			continue
		}

		if len(c.Unproductive()) != len(tc.unproductive) {
			t.Errorf("case %s, got %d, want %d", tc.filename,
				len(c.Unproductive()), len(tc.unproductive))
			continue
		}

		for n, nt := range c.Unproductive() {
			if r := c.NonTerminals[nt]; r != tc.unproductive[n] {
				t.Errorf("case %s, number %d, got %s, want %s",
					tc.filename, n+1, r, tc.unproductive[n])
			}
		}
	}
}
