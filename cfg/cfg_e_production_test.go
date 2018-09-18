package cfg

import "testing"

func TestCfgEProduction(t *testing.T) {
	for _, tc := range grammarTestCases {
		c, err := getAndParseFile(t, tc.filename)
		if err != nil {
			t.Errorf("couldn't parse grammar file %q: %v", tc.filename, err)
			continue
		}

		if len(c.NonTerminalsWithEProductions()) != len(tc.haveEProds) {
			t.Errorf("case %s, got %d, want %d", tc.filename,
				len(c.NonTerminalsWithEProductions()), len(tc.haveEProds))
			continue
		}

		for n, nt := range c.NonTerminalsWithEProductions() {
			if r := c.NonTerminals[nt]; r != tc.haveEProds[n] {
				t.Errorf("case %s, number %d, got %s, want %s",
					tc.filename, n+1, r, tc.haveEProds[n])
			}
		}
	}
}

func TestCfgHasEProduction(t *testing.T) {
	for _, tc := range grammarTestCases {
		c, err := getAndParseFile(t, tc.filename)
		if err != nil {
			t.Errorf("couldn't parse grammar file %q: %v", tc.filename, err)
			continue
		}

		if r := c.HasEProduction(); r != (len(tc.haveEProds) > 0) {
			t.Errorf("case %s, got %t, want %t", tc.filename,
				r, len(tc.haveEProds) > 0)
		}
	}
}
