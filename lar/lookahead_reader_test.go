package lar

import (
	"strings"
	"testing"
)

type match struct {
	matchFunc func(*lookaheadReader) bool
	result    string
}

func TestMatching(t *testing.T) {
	cases := []struct {
		input string
		terms []match
	}{
		{"", []match{}},
		{"a", []match{{(*lookaheadReader).MatchLetter, "a"}}},
		{"ab", []match{
			{(*lookaheadReader).MatchLetter, "a"},
			{(*lookaheadReader).MatchLetter, "b"},
		}},
		{"a", []match{{(*lookaheadReader).MatchLetters, "a"}}},
		{"ab", []match{{(*lookaheadReader).MatchLetters, "ab"}}},
		{"1", []match{{(*lookaheadReader).MatchDigit, "1"}}},
		{"12", []match{
			{(*lookaheadReader).MatchDigit, "1"},
			{(*lookaheadReader).MatchDigit, "2"},
		}},
		{"1", []match{{(*lookaheadReader).MatchDigits, "1"}}},
		{"12", []match{{(*lookaheadReader).MatchDigits, "12"}}},
		{" ", []match{{(*lookaheadReader).MatchSpace, " "}}},
		{" \t", []match{
			{(*lookaheadReader).MatchSpace, " "},
			{(*lookaheadReader).MatchSpace, "\t"},
		}},
		{" ", []match{{(*lookaheadReader).MatchSpaces, " "}}},
		{" \t", []match{{(*lookaheadReader).MatchSpaces, " \t"}}},
		{"a1", []match{
			{(*lookaheadReader).MatchLetter, "a"},
			{(*lookaheadReader).MatchDigit, "1"},
		}},
		{"1 ", []match{
			{(*lookaheadReader).MatchDigit, "1"},
			{(*lookaheadReader).MatchSpace, " "},
		}},
		{" a", []match{
			{(*lookaheadReader).MatchSpace, " "},
			{(*lookaheadReader).MatchLetter, "a"},
		}},
		{"abc123", []match{
			{(*lookaheadReader).MatchLetters, "abc"},
			{(*lookaheadReader).MatchDigits, "123"},
		}},
		{"123 \t\r", []match{
			{(*lookaheadReader).MatchDigits, "123"},
			{(*lookaheadReader).MatchSpaces, " \t\r"},
		}},
		{" \t\rabc", []match{
			{(*lookaheadReader).MatchSpaces, " \t\r"},
			{(*lookaheadReader).MatchLetters, "abc"},
		}},
	}

	for i, c := range cases {
		lar, err := NewLookaheadReader(strings.NewReader(c.input))
		if err != nil {
			t.Errorf("couldn't create lookahead reader: %v", err)
		}

		for j, v := range c.terms {
			if !v.matchFunc(&lar) {
				t.Errorf("case %d, %d, matching method failed", i, j)
			}
			if string(lar.result) != v.result {
				t.Errorf("case %d, %d, got %s, want %s", i, j,
					string(lar.result), v.result)
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
		terms []func(*lookaheadReader) bool
	}{
		{"a", []func(*lookaheadReader) bool{
			(*lookaheadReader).MatchDigit,
			(*lookaheadReader).MatchDigits,
			(*lookaheadReader).MatchSpace,
			(*lookaheadReader).MatchSpaces,
		}},
		{"1", []func(*lookaheadReader) bool{
			(*lookaheadReader).MatchLetter,
			(*lookaheadReader).MatchLetters,
			(*lookaheadReader).MatchSpace,
			(*lookaheadReader).MatchSpaces,
		}},
		{" ", []func(*lookaheadReader) bool{
			(*lookaheadReader).MatchLetter,
			(*lookaheadReader).MatchLetters,
			(*lookaheadReader).MatchDigit,
			(*lookaheadReader).MatchDigits,
		}},
		{"", []func(*lookaheadReader) bool{
			(*lookaheadReader).MatchLetter,
			(*lookaheadReader).MatchLetters,
			(*lookaheadReader).MatchDigit,
			(*lookaheadReader).MatchDigits,
			(*lookaheadReader).MatchSpace,
			(*lookaheadReader).MatchSpaces,
		}},
	}

	for i, c := range cases {
		lar, err := NewLookaheadReader(strings.NewReader(c.input))
		if err != nil {
			t.Errorf("couldn't create lookahead reader: %v", err)
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
		t.Errorf("couldn't create lookahead reader: %v", err)
	}

	expectedStatus := true
	if result := lar.MatchOneOf('.', ';', ':'); result != expectedStatus {
		t.Errorf("got %v, want %v", result, expectedStatus)
	}

	expectedResult := ":"
	if result := lar.Result(); string(result) != expectedResult {
		t.Errorf("got %v, want %v", string(result), expectedResult)
	}

	if !lar.EndOfInput() {
		t.Errorf("end of input not found when expected")
	}
}

func TestMatchOneOfNotPresent(t *testing.T) {
	lar, err := NewLookaheadReader(strings.NewReader("?"))
	if err != nil {
		t.Errorf("couldn't create lookahead reader: %v", err)
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
		t.Errorf("couldn't create lookahead reader: %v", err)
	}

	for n, c := range input {
		result, err := lar.Next()
		if err != nil {
			t.Errorf("index %d, got %v, want %v", n, err, nil)
		}
		if string(result) != string(c) {
			t.Errorf("got %v, want %v", string(result), c)
		}
	}

	if !lar.EndOfInput() {
		t.Errorf("end of input not found when expected")
	}
}
