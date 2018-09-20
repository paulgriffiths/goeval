package cfg

type grammarTestCase struct {
	filename                 string
	isLeftRecursive          bool
	numNonTerminals          int
	numTerminals             int
	numProductions           int
	nonTerminalNames         []string
	terminalNames            []string
	leftRecursive            []string
	immediatelyLeftRecursive []string
	haveCycles               []string
	haveEProds               []string
	areNullable              []string
	follows                  map[string][]string
}

var grammarTestCases = []grammarTestCase{
	{
		tgArithLr, true, 4, 5, 7,
		[]string{"E", "T", "F", "Digits"},
		[]string{"\\+", "\\*", "\\(", "\\)", "[[:digit:]]+"},
		[]string{"E", "T"}, []string{"E", "T"},
		[]string{}, []string{}, []string{},
		map[string][]string{
			"F":      []string{"\\+", "\\*", "\\)", "$"},
			"T":      []string{"\\+", "\\*", "\\)", "$"},
			"E":      []string{"\\+", "\\)", "$"},
			"Digits": []string{"\\*", "\\+", "\\)", "$"},
		},
	},
	{
		tgArithNlr, false, 6, 5, 9,
		[]string{"E", "T", "E'", "F", "T'", "Digits"},
		[]string{"\\+", "\\*", "\\(", "\\)", "[[:digit:]]+"},
		[]string{}, []string{},
		[]string{}, []string{"E'", "T'"}, []string{"E'", "T'"},
		map[string][]string{
			"F":      []string{"\\+", "\\*", "\\)", "$"},
			"T":      []string{"\\+", "\\)", "$"},
			"E":      []string{"\\)", "$"},
			"E'":     []string{"\\)", "$"},
			"T'":     []string{"\\+", "\\)", "$"},
			"Digits": []string{"\\*", "\\+", "\\)", "$"},
		},
	},
	{
		tgArithAmbig, true, 2, 5, 5,
		[]string{"E", "Digits"},
		[]string{"\\+", "\\*", "\\(", "\\)", "[[:digit:]]+"},
		[]string{"E"}, []string{"E"},
		[]string{}, []string{}, []string{},
		map[string][]string{
			"E":      []string{"\\*", "\\+", "\\)", "$"},
			"Digits": []string{"\\*", "\\+", "\\)", "$"},
		},
	},
	{
		tgBalParens1, true, 1, 2, 2,
		[]string{"S"},
		[]string{"\\(", "\\)"},
		[]string{"S"}, []string{"S"},
		[]string{}, []string{"S"}, []string{"S"},
		map[string][]string{
			"S": []string{"\\(", "\\)", "$"},
		},
	},
	{
		tgBalParens2, false, 1, 2, 2,
		[]string{"S"},
		[]string{"\\(", "\\)"},
		[]string{}, []string{},
		[]string{}, []string{"S"}, []string{"S"},
		map[string][]string{
			"S": []string{"\\)", "$"},
		},
	},
	{
		tgZeroOne, false, 1, 3, 2,
		[]string{"S"},
		[]string{"0", "1", "01"},
		[]string{}, []string{},
		[]string{}, []string{}, []string{},
		map[string][]string{
			"S": []string{"1", "$"},
		},
	},
	{
		tgIndirectLr1, true, 2, 4, 5,
		[]string{"S", "A"},
		[]string{"a", "b", "c", "d"},
		[]string{"S", "A"}, []string{"A"},
		[]string{}, []string{"A"}, []string{"A"},
		map[string][]string{
			"S": []string{"d", "$"},
			"A": []string{"a", "c"},
		},
	},
	{
		tgIndirectLr2, true, 4, 5, 8,
		[]string{"S", "A", "B", "C"},
		[]string{"a", "b", "c", "d", "e"},
		[]string{"S", "A", "B", "C"}, []string{},
		[]string{}, []string{"A", "B", "C"}, []string{"A", "B", "C"},
		map[string][]string{
			"S": []string{"e", "$"},
			"A": []string{"a"},
			"B": []string{"c"},
			"C": []string{"d"},
		},
	},
	{
		tgIndirectLr3, true, 5, 6, 10,
		[]string{"S", "A", "B", "C", "D"},
		[]string{"a", "b", "c", "d", "e", "f"},
		[]string{"A", "B", "C", "D"}, []string{},
		[]string{},
		[]string{"A", "B", "C", "D"}, []string{"A", "B", "C", "D"},
		map[string][]string{
			"S": []string{"$"},
			"A": []string{"a", "f"},
			"B": []string{"c"},
			"C": []string{"d"},
			"D": []string{"e"},
		},
	},
	{
		tgCycle1, true, 1, 2, 3,
		[]string{"S"},
		[]string{"a", "b"},
		[]string{"S"}, []string{"S"},
		[]string{"S"}, []string{}, []string{},
		map[string][]string{
			"S": []string{"$"},
		},
	},
	{
		tgCycle2, true, 2, 4, 6,
		[]string{"S", "A"},
		[]string{"a", "b", "c", "d"},
		[]string{"A"}, []string{"A"},
		[]string{"A"}, []string{}, []string{},
		map[string][]string{
			"S": []string{"$"},
			"A": []string{"$"},
		},
	},
	{
		tgCycle3, true, 2, 4, 6,
		[]string{"S", "A"},
		[]string{"a", "b", "c", "d"},
		[]string{"S", "A"}, []string{},
		[]string{"S", "A"}, []string{}, []string{},
		map[string][]string{
			"S": []string{"$"},
			"A": []string{"$"},
		},
	},
	{
		tgCycle4, true, 3, 6, 9,
		[]string{"S", "A", "B"},
		[]string{"a", "b", "c", "d", "e", "f"},
		[]string{"S", "A", "B"}, []string{},
		[]string{"S", "A", "B"}, []string{}, []string{},
		map[string][]string{
			"S": []string{"$"},
			"A": []string{"$"},
			"B": []string{"$"},
		},
	},
	{
		tgNullable1, false, 7, 2, 10,
		[]string{"S", "A", "B", "C", "D", "E", "F"},
		[]string{"a", "b"},
		[]string{}, []string{},
		[]string{}, []string{"C", "D"}, []string{"S", "C", "D"},
		map[string][]string{
			"S": []string{"$"},
			"A": []string{"b"},
			"B": []string{"$"},
			"C": []string{"b", "$"},
			"D": []string{"$"},
			"E": []string{"b", "$"},
			"F": []string{"$"},
		},
	},
	{
		tgNullable2, false, 7, 2, 12,
		[]string{"S", "A", "B", "C", "D", "F", "E"},
		[]string{"a", "b"},
		[]string{}, []string{},
		[]string{}, []string{"B", "D"}, []string{"S", "B", "C", "D"},
		map[string][]string{
			"S": []string{"$"},
			"A": []string{"a", "b", "$"},
			"B": []string{"b", "$"},
			"C": []string{"b", "$"},
			"D": []string{"$"},
			"E": []string{"b", "$"},
			"F": []string{"a", "b", "$"},
		},
	},
	{
		tgNullable3, true, 8, 2, 16,
		[]string{"S", "A", "B", "G", "C", "D", "F", "E"},
		[]string{"a", "b"},
		[]string{"S", "G"}, []string{},
		[]string{"S", "G"}, []string{"B", "D"},
		[]string{"S", "B", "G", "C", "D"},
		map[string][]string{
			"S": []string{"$"},
			"A": []string{"a", "b", "$"},
			"B": []string{"b", "$"},
			"C": []string{"b", "$"},
			"D": []string{"$"},
			"E": []string{"b", "$"},
			"F": []string{"a", "b", "$"},
			"G": []string{"$"},
		},
	},
}
