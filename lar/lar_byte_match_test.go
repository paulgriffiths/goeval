package lar

import (
	"strings"
	"testing"
)

func TestSuccessfulByteMatch(t *testing.T) {
	for i, testCase := range byteMatchGoodCases {
		lar, err := NewLookaheadReader(strings.NewReader(testCase.input))
		if err != nil {
			t.Errorf("couldn't create lookahead reader: %s", err)
			continue
		}

		for j, match := range testCase.matches {
			if !lar.MatchOneOf(match.args...) {
				t.Errorf("case %d, %d, matching method failed", i, j)
			}
		}
	}
}

func TestSuccessfulByteMatchResultValue(t *testing.T) {
	for i, testCase := range byteMatchGoodCases {
		lar, err := NewLookaheadReader(strings.NewReader(testCase.input))
		if err != nil {
			t.Errorf("couldn't create lookahead reader: %s", err)
			continue
		}

		for j, match := range testCase.matches {
			lar.MatchOneOf(match.args...)
			if result := string(lar.Result.Value); result != match.result {
				t.Errorf("case %d, %d, got %s, want %s", i, j,
					result, match.result)
			}
		}
	}
}

func TestSuccessfulByteMatchResultPosition(t *testing.T) {
	for i, testCase := range byteMatchGoodCases {
		lar, err := NewLookaheadReader(strings.NewReader(testCase.input))
		if err != nil {
			t.Errorf("couldn't create lookahead reader: %s", err)
			continue
		}

		for j, match := range testCase.matches {
			lar.MatchOneOf(match.args...)
			if result := lar.Result.Pos; result != match.pos {
				t.Errorf("case %d, %d, got %v, want %v", i, j,
					result, match.pos)
			}
		}
	}
}

func TestSuccessfulByteMatchEndOfInputReached(t *testing.T) {
	for i, testCase := range byteMatchGoodCases {
		lar, err := NewLookaheadReader(strings.NewReader(testCase.input))
		if err != nil {
			t.Errorf("couldn't create lookahead reader: %s", err)
			continue
		}

		for _, match := range testCase.matches {
			lar.MatchOneOf(match.args...)
		}

		if result := lar.EndOfInput(); !result {
			t.Errorf("case %d, end of input not found when expected", i)
		}
	}
}

func TestUnsuccessfulByteMatch(t *testing.T) {
	for i, testCase := range byteMatchBadCases {
		lar, err := NewLookaheadReader(strings.NewReader(testCase.input))
		if err != nil {
			t.Errorf("couldn't create lookahead reader: %s", err)
			continue
		}

		if lar.MatchOneOf(testCase.args...) {
			t.Errorf("case %d, matching method succeeded", i)
		}
	}
}
