package cfg

import "testing"

func TestCfgLeftRecursive(t *testing.T) {
	for _, tc := range grammarTestCases {
		c, err := getAndParseFile(t, tc.filename)
		if err != nil {
			t.Errorf("couldn't parse grammar file %q: %v", tc.filename, err)
			continue
		}

		if len(c.NonTerminalsLeftRecursive()) != len(tc.leftRecursive) {
			t.Errorf("case %s, got %d, want %d", tc.filename,
				len(c.NonTerminalsLeftRecursive()), len(tc.leftRecursive))
			continue
		}

		for n, nt := range c.NonTerminalsLeftRecursive() {
			if r := c.NonTerminals[nt]; r != tc.leftRecursive[n] {
				t.Errorf("case %s, number %d, got %s, want %s",
					tc.filename, n+1, r, tc.leftRecursive[n])
			}
		}
	}
}

func TestCfgImmediatelyLeftRecursive(t *testing.T) {
	for _, tc := range grammarTestCases {
		c, err := getAndParseFile(t, tc.filename)
		if err != nil {
			t.Errorf("couldn't parse grammar file %q: %v", tc.filename, err)
			continue
		}

		nilr := len(c.NonTerminalsImmediatelyLeftRecursive())
		if nilr != len(tc.immediatelyLeftRecursive) {
			t.Errorf("case %s, got %d, want %d", tc.filename,
				nilr, len(tc.immediatelyLeftRecursive))
			continue
		}

		for n, nt := range c.NonTerminalsImmediatelyLeftRecursive() {
			if r := c.NonTerminals[nt]; r != tc.immediatelyLeftRecursive[n] {
				t.Errorf("case %s, number %d, got %s, want %s",
					tc.filename, n+1, r, tc.immediatelyLeftRecursive[n])
			}
		}
	}
}

func TestCfgIsLeftRecursive(t *testing.T) {
	for _, tc := range grammarTestCases {
		c, err := getAndParseFile(t, tc.filename)
		if err != nil {
			t.Errorf("couldn't parse grammar file %q: %v", tc.filename, err)
			continue
		}

		if r := c.IsLeftRecursive(); r != tc.isLeftRecursive {
			t.Errorf("case %s, got %t, want %t", tc.filename,
				r, tc.isLeftRecursive)
		}
	}
}
