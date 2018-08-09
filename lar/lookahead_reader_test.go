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
