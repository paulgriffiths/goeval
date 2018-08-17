package lar

type byteMatch struct {
	result string
	args   []byte
	pos    FilePos
}

var byteMatchGoodCases = []struct {
	input   string
	matches []byteMatch
}{
	{"!", []byteMatch{
		{"!", []byte{'!'}, FilePos{0, 1}},
	}},
	{"?", []byteMatch{
		{"?", []byte{'?'}, FilePos{0, 1}},
	}},
	{"#", []byteMatch{
		{"#", []byte{'#'}, FilePos{0, 1}},
	}},
	{"!", []byteMatch{
		{"!", []byte{'!', '?', '#'}, FilePos{0, 1}},
	}},
	{"?", []byteMatch{
		{"?", []byte{'!', '?', '#'}, FilePos{0, 1}},
	}},
	{"#", []byteMatch{
		{"#", []byte{'!', '?', '#'}, FilePos{0, 1}},
	}},
	{"!?#", []byteMatch{
		{"!", []byte{'!'}, FilePos{0, 1}},
		{"?", []byte{'?'}, FilePos{1, 1}},
		{"#", []byte{'#'}, FilePos{2, 1}},
	}},
	{"!?#", []byteMatch{
		{"!", []byte{'!', '?', '#', '$', '%'}, FilePos{0, 1}},
		{"?", []byte{'!', '?', '#', '$', '%'}, FilePos{1, 1}},
		{"#", []byte{'!', '?', '#', '$', '%'}, FilePos{2, 1}},
	}},
	{"!?#\n%^&\n@()", []byteMatch{
		{"!", []byte{'!'}, FilePos{0, 1}},
		{"?", []byte{'?'}, FilePos{1, 1}},
		{"#", []byte{'#'}, FilePos{2, 1}},
		{"\n", []byte{'\n'}, FilePos{3, 1}},
		{"%", []byte{'%'}, FilePos{0, 2}},
		{"^", []byte{'^'}, FilePos{1, 2}},
		{"&", []byte{'&'}, FilePos{2, 2}},
		{"\n", []byte{'\n'}, FilePos{3, 2}},
		{"@", []byte{'@'}, FilePos{0, 3}},
		{"(", []byte{'('}, FilePos{1, 3}},
		{")", []byte{')'}, FilePos{2, 3}},
	}},
	{"!?#\n%^&\n@()", []byteMatch{
		{"!", []byte{'!', '?', '#', '\n', '%', '^', '&'}, FilePos{0, 1}},
		{"?", []byte{'!', '?', '#', '\n', '%', '^', '&'}, FilePos{1, 1}},
		{"#", []byte{'!', '?', '#', '\n', '%', '^', '&'}, FilePos{2, 1}},
		{"\n", []byte{'!', '?', '#', '\n', '%', '^', '&'}, FilePos{3, 1}},
		{"%", []byte{'\n', '%', '^', '&'}, FilePos{0, 2}},
		{"^", []byte{'\n', '%', '^', '&'}, FilePos{1, 2}},
		{"&", []byte{'\n', '%', '^', '&'}, FilePos{2, 2}},
		{"\n", []byte{'\n', '%', '^', '&'}, FilePos{3, 2}},
		{"@", []byte{'\n', '@', '(', ')'}, FilePos{0, 3}},
		{"(", []byte{'\n', '@', '(', ')'}, FilePos{1, 3}},
		{")", []byte{'\n', '@', '(', ')'}, FilePos{2, 3}},
	}},
}

var byteMatchBadCases = []struct {
	input string
	args  []byte
}{
	{"!", []byte{'@', '#', '$', '%', '^', '&', '\n'}},
}
