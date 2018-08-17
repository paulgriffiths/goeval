package lar

import (
	"strings"
	"testing"
)

type match struct {
	matchFunc func(*LookaheadReader) bool
	result    string
}

func TestMatching(t *testing.T) {
	cases := []struct {
		input string
		terms []match
	}{
		{"", []match{}},
		{"a", []match{{(*LookaheadReader).MatchLetter, "a"}}},
		{"ab", []match{
			{(*LookaheadReader).MatchLetter, "a"},
			{(*LookaheadReader).MatchLetter, "b"},
		}},
		{"a", []match{{(*LookaheadReader).MatchLetters, "a"}}},
		{"ab", []match{{(*LookaheadReader).MatchLetters, "ab"}}},
		{"1", []match{{(*LookaheadReader).MatchDigit, "1"}}},
		{"12", []match{
			{(*LookaheadReader).MatchDigit, "1"},
			{(*LookaheadReader).MatchDigit, "2"},
		}},
		{"1", []match{{(*LookaheadReader).MatchDigits, "1"}}},
		{"12", []match{{(*LookaheadReader).MatchDigits, "12"}}},
		{" ", []match{{(*LookaheadReader).MatchSpace, " "}}},
		{" \t", []match{
			{(*LookaheadReader).MatchSpace, " "},
			{(*LookaheadReader).MatchSpace, "\t"},
		}},
		{" ", []match{{(*LookaheadReader).MatchSpaces, " "}}},
		{" \t", []match{{(*LookaheadReader).MatchSpaces, " \t"}}},
		{"a1", []match{
			{(*LookaheadReader).MatchLetter, "a"},
			{(*LookaheadReader).MatchDigit, "1"},
		}},
		{"1 ", []match{
			{(*LookaheadReader).MatchDigit, "1"},
			{(*LookaheadReader).MatchSpace, " "},
		}},
		{" a", []match{
			{(*LookaheadReader).MatchSpace, " "},
			{(*LookaheadReader).MatchLetter, "a"},
		}},
		{"abc123", []match{
			{(*LookaheadReader).MatchLetters, "abc"},
			{(*LookaheadReader).MatchDigits, "123"},
		}},
		{"123 \t\r", []match{
			{(*LookaheadReader).MatchDigits, "123"},
			{(*LookaheadReader).MatchSpaces, " \t\r"},
		}},
		{" \t\rabc", []match{
			{(*LookaheadReader).MatchSpaces, " \t\r"},
			{(*LookaheadReader).MatchLetters, "abc"},
		}},
	}

	for i, c := range cases {
		lar, err := NewLookaheadReader(strings.NewReader(c.input))
		if err != nil {
			t.Errorf("couldn't create lookahead reader: %s", err)
		}

		for j, v := range c.terms {
			if !v.matchFunc(&lar) {
				t.Errorf("case %d, %d, matching method failed", i, j)
			}
			if string(lar.Result.Value) != v.result {
				t.Errorf("case %d, %d, got %s, want %s", i, j,
					string(lar.Result.Value), v.result)
			}
		}

		if result := lar.EndOfInput(); !result {
			t.Errorf("case %d, end of input not found when expected", i)
		}
	}
}

func TestNonMatching(t *testing.T) {
	cases := []struct {
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

	for i, c := range cases {
		lar, err := NewLookaheadReader(strings.NewReader(c.input))
		if err != nil {
			t.Errorf("couldn't create lookahead reader: %s", err)
		}

		for j, v := range c.terms {
			if v(&lar) {
				t.Errorf("case %d, %d, matching method succeeded", i, j)
			}
		}
	}
}

func TestMatchOneOfPresent(t *testing.T) {
	lar, err := NewLookaheadReader(strings.NewReader(":"))
	if err != nil {
		t.Errorf("couldn't create lookahead reader: %s", err)
	}

	expectedStatus := true
	if result := lar.MatchOneOf('.', ';', ':'); result != expectedStatus {
		t.Errorf("got %v, want %v", result, expectedStatus)
	}

	expectedResult := ":"
	if result := lar.Result.Value; string(result) != expectedResult {
		t.Errorf("got %v, want %v", string(result), expectedResult)
	}

	if !lar.EndOfInput() {
		t.Errorf("end of input not found when expected")
	}
}

func TestMatchOneOfNotPresent(t *testing.T) {
	lar, err := NewLookaheadReader(strings.NewReader("?"))
	if err != nil {
		t.Errorf("couldn't create lookahead reader: %s", err)
	}

	expectedStatus := false
	if result := lar.MatchOneOf('.', ';', ':'); result != expectedStatus {
		t.Errorf("got %v, want %v", result, expectedStatus)
	}

	if lar.EndOfInput() {
		t.Errorf("end of input found when not expected")
	}
}

func TestNext(t *testing.T) {
	input := "a1\t?"
	lar, err := NewLookaheadReader(strings.NewReader(input))
	if err != nil {
		t.Errorf("couldn't create lookahead reader: %s", err)
	}

	for n, c := range input {
		result, err := lar.Next()
		if err != nil {
			t.Errorf("index %d, got %v, want %v", n, err, nil)
		}
		if string(result) != string(c) {
			t.Errorf("got %s, want %s", string(result), string(c))
		}
	}

	if !lar.EndOfInput() {
		t.Errorf("end of input not found when expected")
	}
}

func TestPosNext(t *testing.T) {
	input := "ab\n123\n?!*:\n"
	results := []struct {
		value byte
		pos   int
		line  int
	}{
		{'a', 0, 1},
		{'b', 1, 1},
		{'\n', 2, 1},
		{'1', 0, 2},
		{'2', 1, 2},
		{'3', 2, 2},
		{'\n', 3, 2},
		{'?', 0, 3},
		{'!', 1, 3},
		{'*', 2, 3},
		{':', 3, 3},
		{'\n', 4, 3},
	}

	lar, err := NewLookaheadReader(strings.NewReader(input))
	if err != nil {
		t.Errorf("couldn't create lookahead reader: %s", err)
	}

	for n, r := range results {
		b, err := lar.Next()
		if err != nil {
			t.Errorf("case %d, couldn't get next byte: %s", n, err)
		}

		if b != r.value {
			t.Errorf("case %d, unexpected character, got %q, want %q",
				n, b, r.value)
		}

		pos := lar.pos
		if pos.line != r.line {
			t.Errorf("case %d, unexpected line, got %d, want %d",
				n, pos.line, r.line)
		}
		if pos.ch != r.pos {
			t.Errorf("case %d, unexpected position, got %d, want %d",
				n, pos.ch, r.pos)
		}
	}
}

type lineMatch struct {
	m         match
	pos, line int
}

func TestPosMatch(t *testing.T) {
	input := "ab12\n456 \t\r\n def\n?!"
	results := []lineMatch{
		{match{(*LookaheadReader).MatchLetters, "ab"}, 0, 1},
		{match{(*LookaheadReader).MatchDigits, "12"}, 2, 1},
		{match{(*LookaheadReader).MatchNewline, "\n"}, 4, 1},
		{match{(*LookaheadReader).MatchDigits, "456"}, 0, 2},
		{match{(*LookaheadReader).MatchSpaces, " \t\r"}, 3, 2},
		{match{(*LookaheadReader).MatchNewline, "\n"}, 6, 2},
		{match{(*LookaheadReader).MatchSpaces, " "}, 0, 3},
		{match{(*LookaheadReader).MatchLetters, "def"}, 1, 3},
		{match{(*LookaheadReader).MatchNewline, "\n"}, 4, 3},
	}

	lar, err := NewLookaheadReader(strings.NewReader(input))
	if err != nil {
		t.Errorf("couldn't create lookahead reader: %s", err)
	}

	for n, r := range results {
		if !r.m.matchFunc(&lar) {
			t.Errorf("case %d, matching method failed", n)
		}

		pos := lar.Result.Pos
		if pos.line != r.line {
			t.Errorf("case %d, unexpected line, got %d, want %d",
				n, pos.line, r.line)
		}
		if pos.ch != r.pos {
			t.Errorf("case %d, unexpected position, got %d, want %d",
				n, pos.ch, r.pos)
		}
	}

	if !lar.MatchOneOf('?') {
		t.Errorf("matching method failed")
	}
	pos := lar.Result.Pos
	if pos.line != 4 {
		t.Errorf("unexpected line, got %d, want %d", pos.line, 4)
	}
	if pos.ch != 0 {
		t.Errorf("unexpected position, got %d, want %d", pos.ch, 0)
	}

	if !lar.MatchOneOf('!') {
		t.Errorf("matching method failed")
	}
	pos = lar.Result.Pos
	if pos.line != 4 {
		t.Errorf("unexpected line, got %d, want %d", pos.line, 4)
	}
	if pos.ch != 1 {
		t.Errorf("unexpected position, got %d, want %d", pos.ch, 1)
	}
}
