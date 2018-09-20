package rdp

type testCase struct {
	filename string
	accepts  []string
	rejects  []string
}

var parserTestCases = []testCase{
	{
		tgArithNlr,
		[]string{"3", "3+4", "3+4*5", "(3+4)*5"},
		[]string{"", "+", "3+", "-3", "(3+4*5"},
	},
	{
		tgBalParens2,
		[]string{"", "()", "(())", "((()))", "(()())", "(()()())((()))"},
		[]string{"(", ")", "(()", "())", ")(", "(((((()))))()"},
	},
	{
		tgZeroOne,
		[]string{"01", "0011", "000111"},
		[]string{"", "0", "1", "001", "011", "10", "111000", "001011",
			"0000001111101"},
	},
}
