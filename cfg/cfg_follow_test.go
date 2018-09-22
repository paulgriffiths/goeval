package cfg

import "testing"

func TestCfgFollows(t *testing.T) {
	for n, tc := range grammarTestCases {
		c, err := getAndParseFile(t, tc.filename)
		if err != nil {
			t.Errorf("couldn't parse file %q: %v", tc.filename, err)
			continue
		}

		cmpSet := make([]SetBodyComp, len(c.NonTerminals))
		for i := 0; i < len(c.NonTerminals); i++ {
			cmpSet[i] = NewSetBodyComp()
		}

		for nonTerm, terminals := range tc.follows {
			n := c.NtTable[nonTerm]
			for _, term := range terminals {
				if term == "$" {
					cmpSet[n].Insert(NewBodyInputEnd())
					continue
				}
				cmpSet[n].Insert(c.TerminalComp(term))
			}
		}

		for i := 0; i < len(c.NonTerminals); i++ {
			if !c.Follow(i).Equals(cmpSet[i]) {
				t.Errorf("case %d, nonterminal %s, got %v, want %v",
					n+1, c.NonTerminals[i], c.Follow(i).Elements(),
					cmpSet[i].Elements())
			}
		}
	}
}
