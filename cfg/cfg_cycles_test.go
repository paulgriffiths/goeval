package cfg

import "testing"

func TestCfgCycles(t *testing.T) {
	for _, tc := range grammarTestCases {
		c, err := getAndParseFile(t, tc.filename)
		if err != nil {
			t.Errorf("couldn't parse grammar file %q: %v", tc.filename, err)
			continue
		}

		if len(c.NonTerminalsWithCycles()) != len(tc.haveCycles) {
			t.Errorf("case %s, got %d, want %d", tc.filename,
				len(c.NonTerminalsWithCycles()), len(tc.haveCycles))
			continue
		}

		for n, nt := range c.NonTerminalsWithCycles() {
			if r := c.NonTerminals[nt]; r != tc.haveCycles[n] {
				t.Errorf("case %s, number %d, got %s, want %s",
					tc.filename, n+1, r, tc.haveCycles[n])
			}
		}
	}
}

func TestCfgHasCycles(t *testing.T) {
	for _, tc := range grammarTestCases {
		c, err := getAndParseFile(t, tc.filename)
		if err != nil {
			t.Errorf("couldn't parse grammar file %q: %v", tc.filename, err)
			continue
		}

		if r := c.HasCycle(); r != (len(tc.haveCycles) > 0) {
			t.Errorf("case %s, got %t, want %t", tc.filename,
				r, len(tc.haveCycles) > 0)
		}
	}
}
