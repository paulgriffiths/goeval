package cfg

import "testing"

func TestCfgParseNumNonTerminals(t *testing.T) {
	for _, tc := range grammarTestCases {
		c, err := getAndParseFile(t, tc.filename)
		if err != nil {
			t.Errorf("couldn't parse grammar file %q: %v", tc.filename, err)
			continue
		}

		if len(c.NonTerminals) != tc.numNonTerminals {
			t.Errorf("case %s, got %d, want %d", tc.filename,
				len(c.NonTerminals), tc.numNonTerminals)
		}
	}
}

func TestCfgParseNumTerminals(t *testing.T) {
	for _, tc := range grammarTestCases {
		c, err := getAndParseFile(t, tc.filename)
		if err != nil {
			t.Errorf("couldn't parse grammar file %q: %v", tc.filename, err)
			continue
		}

		if len(c.Terminals) != tc.numTerminals {
			t.Errorf("case %s, got %d, want %d", tc.filename,
				len(c.Terminals), tc.numTerminals)
		}
	}
}

func TestCfgParseNumProductions(t *testing.T) {
	for _, tc := range grammarTestCases {
		c, err := getAndParseFile(t, tc.filename)
		if err != nil {
			t.Errorf("couldn't parse grammar file %q: %v", tc.filename, err)
			continue
		}

		if c.NumProductions() != tc.numProductions {
			t.Errorf("case %s, got %d, want %d", tc.filename,
				c.NumProductions(), tc.numProductions)
		}
	}
}

func TestCfgParseNonTerminalNames(t *testing.T) {
	for _, tc := range grammarTestCases {
		c, err := getAndParseFile(t, tc.filename)
		if err != nil {
			t.Errorf("couldn't parse grammar file %q: %v", tc.filename, err)
			continue
		}

		if len(c.NonTerminals) != len(tc.nonTerminalNames) {
			t.Errorf("case %s, got %d, want %d", tc.filename,
				len(tc.nonTerminalNames), len(c.NonTerminals))
			continue
		}

		for n, ntName := range c.NonTerminals {
			if ntName != tc.nonTerminalNames[n] {
				t.Errorf("case %s, nonterminal %d, got %s, want %s",
					tc.filename, n, tc.nonTerminalNames[n], ntName)
			}
		}
	}
}

func TestCfgParseTerminalNames(t *testing.T) {
	for _, tc := range grammarTestCases {
		c, err := getAndParseFile(t, tc.filename)
		if err != nil {
			t.Errorf("couldn't parse grammar file %q: %v", tc.filename, err)
			continue
		}

		if len(c.Terminals) != len(tc.terminalNames) {
			t.Errorf("case %s, got %d, want %d", tc.filename,
				len(tc.terminalNames), len(c.Terminals))
			continue
		}

		for n, tName := range c.Terminals {
			if tName != tc.terminalNames[n] {
				t.Errorf("case %s, terminal %d, got %s, want %s",
					tc.filename, n, tc.terminalNames[n], tName)
			}
		}
	}
}

func TestCfgParseNonTerminalTable(t *testing.T) {
	for _, tc := range grammarTestCases {
		c, err := getAndParseFile(t, tc.filename)
		if err != nil {
			t.Errorf("couldn't parse grammar file %q: %v", tc.filename, err)
			continue
		}

		for n, ntName := range c.NonTerminals {
			if r := c.NtTable[ntName]; r != n {
				t.Errorf("case %s, nonterminal %s, got %d, want %d",
					tc.filename, ntName, r, n)
			}
		}
	}
}

func TestCfgParseTerminalTable(t *testing.T) {
	for _, tc := range grammarTestCases {
		c, err := getAndParseFile(t, tc.filename)
		if err != nil {
			t.Errorf("couldn't parse grammar file %q: %v", tc.filename, err)
			continue
		}

		for n, tName := range c.Terminals {
			if r := c.TTable[tName]; r != n {
				t.Errorf("case %s, terminal %s, got %d, want %d",
					tc.filename, tName, r, n)
			}
		}
	}
}
