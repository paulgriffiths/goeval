package lar

type classMatch struct {
	matchFunc func(*LookaheadReader) bool
	result    string
	pos       FilePos
}

var classMatchGoodCases = []struct {
	input   string
	matches []classMatch
}{
	{"", []classMatch{}},
	{"a", []classMatch{
		{(*LookaheadReader).MatchLetter, "a", FilePos{0, 1}},
	}},
	{"ab", []classMatch{
		{(*LookaheadReader).MatchLetter, "a", FilePos{0, 1}},
		{(*LookaheadReader).MatchLetter, "b", FilePos{1, 1}},
	}},
	{"a", []classMatch{
		{(*LookaheadReader).MatchLetters, "a", FilePos{0, 1}},
	}},
	{"ab", []classMatch{
		{(*LookaheadReader).MatchLetters, "ab", FilePos{0, 1}},
	}},
	{"1", []classMatch{
		{(*LookaheadReader).MatchDigit, "1", FilePos{0, 1}},
	}},
	{"12", []classMatch{
		{(*LookaheadReader).MatchDigit, "1", FilePos{0, 1}},
		{(*LookaheadReader).MatchDigit, "2", FilePos{1, 1}},
	}},
	{"1", []classMatch{
		{(*LookaheadReader).MatchDigits, "1", FilePos{0, 1}},
	}},
	{"12", []classMatch{
		{(*LookaheadReader).MatchDigits, "12", FilePos{0, 1}},
	}},
	{" ", []classMatch{
		{(*LookaheadReader).MatchSpace, " ", FilePos{0, 1}},
	}},
	{" \t", []classMatch{
		{(*LookaheadReader).MatchSpace, " ", FilePos{0, 1}},
		{(*LookaheadReader).MatchSpace, "\t", FilePos{1, 1}},
	}},
	{" ", []classMatch{
		{(*LookaheadReader).MatchSpaces, " ", FilePos{0, 1}},
	}},
	{" \t", []classMatch{
		{(*LookaheadReader).MatchSpaces, " \t", FilePos{0, 1}},
	}},
	{"a1", []classMatch{
		{(*LookaheadReader).MatchLetter, "a", FilePos{0, 1}},
		{(*LookaheadReader).MatchDigit, "1", FilePos{1, 1}},
	}},
	{"1 ", []classMatch{
		{(*LookaheadReader).MatchDigit, "1", FilePos{0, 1}},
		{(*LookaheadReader).MatchSpace, " ", FilePos{1, 1}},
	}},
	{" a", []classMatch{
		{(*LookaheadReader).MatchSpace, " ", FilePos{0, 1}},
		{(*LookaheadReader).MatchLetter, "a", FilePos{1, 1}},
	}},
	{"abc123", []classMatch{
		{(*LookaheadReader).MatchLetters, "abc", FilePos{0, 1}},
		{(*LookaheadReader).MatchDigits, "123", FilePos{3, 1}},
	}},
	{"123 \t\r", []classMatch{
		{(*LookaheadReader).MatchDigits, "123", FilePos{0, 1}},
		{(*LookaheadReader).MatchSpaces, " \t\r", FilePos{3, 1}},
	}},
	{" \t\rabc", []classMatch{
		{(*LookaheadReader).MatchSpaces, " \t\r", FilePos{0, 1}},
		{(*LookaheadReader).MatchLetters, "abc", FilePos{3, 1}},
	}},
	{"ab12\n456 \t\r\n def\n", []classMatch{
		{(*LookaheadReader).MatchLetters, "ab", FilePos{0, 1}},
		{(*LookaheadReader).MatchDigits, "12", FilePos{2, 1}},
		{(*LookaheadReader).MatchNewline, "\n", FilePos{4, 1}},
		{(*LookaheadReader).MatchDigits, "456", FilePos{0, 2}},
		{(*LookaheadReader).MatchSpaces, " \t\r", FilePos{3, 2}},
		{(*LookaheadReader).MatchNewline, "\n", FilePos{6, 2}},
		{(*LookaheadReader).MatchSpaces, " ", FilePos{0, 3}},
		{(*LookaheadReader).MatchLetters, "def", FilePos{1, 3}},
		{(*LookaheadReader).MatchNewline, "\n", FilePos{4, 3}},
	}},
}

var classMatchBadCases = []struct {
	input string
	terms []func(*LookaheadReader) bool
}{
	{"a", []func(*LookaheadReader) bool{
		(*LookaheadReader).MatchDigit,
		(*LookaheadReader).MatchDigits,
		(*LookaheadReader).MatchSpace,
		(*LookaheadReader).MatchSpaces,
	}},
	{"1", []func(*LookaheadReader) bool{
		(*LookaheadReader).MatchLetter,
		(*LookaheadReader).MatchLetters,
		(*LookaheadReader).MatchSpace,
		(*LookaheadReader).MatchSpaces,
	}},
	{" ", []func(*LookaheadReader) bool{
		(*LookaheadReader).MatchLetter,
		(*LookaheadReader).MatchLetters,
		(*LookaheadReader).MatchDigit,
		(*LookaheadReader).MatchDigits,
	}},
	{"", []func(*LookaheadReader) bool{
		(*LookaheadReader).MatchLetter,
		(*LookaheadReader).MatchLetters,
		(*LookaheadReader).MatchDigit,
		(*LookaheadReader).MatchDigits,
		(*LookaheadReader).MatchSpace,
		(*LookaheadReader).MatchSpaces,
	}},
}
