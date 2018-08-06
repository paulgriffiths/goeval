package lar

import (
    "testing"
    "strings"
)

func TestEmptyStringEndOfInput(t *testing.T) {
    lar, err := NewLookaheadReader(strings.NewReader(""))
    if err != nil {
        t.Errorf("couldn't create lookahead reader: %v", err)
    }

    expected := true
    if result := lar.EndOfInput(); result != expected {
        t.Errorf("got %v, wanted %v", result, expected)
    }
}

func TestEmptyStringDoesntMatch(t *testing.T) {
    lar, err := NewLookaheadReader(strings.NewReader(""))
    if err != nil {
        t.Errorf("couldn't create lookahead reader: %v", err)
    }

    expected := false
    if result := lar.MatchLetters(); result != expected {
        t.Errorf("got %v, wanted %v", result, expected)
    }

    if result := lar.MatchDigits(); result != expected {
        t.Errorf("got %v, wanted %v", result, expected)
    }

    if result := lar.MatchSpaces(); result != expected {
        t.Errorf("got %v, wanted %v", result, expected)
    }

    if result := lar.MatchOneOf('+', '-', '*', '/'); result != expected {
        t.Errorf("got %v, wanted %v", result, expected)
    }
}

func TestMatchLetterOneLetter(t *testing.T) {
    lar, err := NewLookaheadReader(strings.NewReader("a"))
    if err != nil {
        t.Errorf("couldn't create lookahead reader: %v", err)
    }

    expectedStatus := true
    if result := lar.MatchLetter(); result != expectedStatus {
        t.Errorf("got %v, wanted %v", result, expectedStatus)
    }

    expectedResult := "a"
    if result := lar.Result(); string(result) != expectedResult {
        t.Errorf("got %v, wanted %v", string(result), expectedResult)
    }

    expectedStatus = true
    if result := lar.EndOfInput(); result != expectedStatus {
        t.Errorf("got %v, wanted %v", result, expectedStatus)
    }
}

func TestMatchLetterTwoLetters(t *testing.T) {
    lar, err := NewLookaheadReader(strings.NewReader("ab"))
    if err != nil {
        t.Errorf("couldn't create lookahead reader: %v", err)
    }

    expectedStatus := true
    if result := lar.MatchLetter(); result != expectedStatus {
        t.Errorf("got %v, wanted %v", result, expectedStatus)
    }

    expectedResult := "a"
    if result := lar.Result(); string(result) != expectedResult {
        t.Errorf("got %v, wanted %v", string(result), expectedResult)
    }

    expectedStatus = true
    if result := lar.MatchLetter(); result != expectedStatus {
        t.Errorf("got %v, wanted %v", result, expectedStatus)
    }

    expectedResult = "b"
    if result := lar.Result(); string(result) != expectedResult {
        t.Errorf("got %v, wanted %v", string(result), expectedResult)
    }

    expectedStatus = true
    if result := lar.EndOfInput(); result != expectedStatus {
        t.Errorf("got %v, wanted %v", result, expectedStatus)
    }
}

func TestMatchLettersOneLetter(t *testing.T) {
    lar, err := NewLookaheadReader(strings.NewReader("a"))
    if err != nil {
        t.Errorf("couldn't create lookahead reader: %v", err)
    }

    expectedStatus := true
    if result := lar.MatchLetters(); result != expectedStatus {
        t.Errorf("got %v, wanted %v", result, expectedStatus)
    }

    expectedResult := "a"
    if result := lar.Result(); string(result) != expectedResult {
        t.Errorf("got %v, wanted %v", string(result), expectedResult)
    }

    expectedStatus = true
    if result := lar.EndOfInput(); result != expectedStatus {
        t.Errorf("got %v, wanted %v", result, expectedStatus)
    }
}

func TestMatchLettersTwoLetters(t *testing.T) {
    lar, err := NewLookaheadReader(strings.NewReader("ab"))
    if err != nil {
        t.Errorf("couldn't create lookahead reader: %v", err)
    }

    expectedStatus := true
    if result := lar.MatchLetters(); result != expectedStatus {
        t.Errorf("got %v, wanted %v", result, expectedStatus)
    }

    expectedResult := "ab"
    if result := lar.Result(); string(result) != expectedResult {
        t.Errorf("got %v, wanted %v", string(result), expectedResult)
    }

    expectedStatus = true
    if result := lar.EndOfInput(); result != expectedStatus {
        t.Errorf("got %v, wanted %v", result, expectedStatus)
    }
}

func TestMatchDigitOneDigit(t *testing.T) {
    lar, err := NewLookaheadReader(strings.NewReader("1"))
    if err != nil {
        t.Errorf("couldn't create lookahead reader: %v", err)
    }

    expectedStatus := true
    if result := lar.MatchDigit(); result != expectedStatus {
        t.Errorf("got %v, wanted %v", result, expectedStatus)
    }

    expectedResult := "1"
    if result := lar.Result(); string(result) != expectedResult {
        t.Errorf("got %v, wanted %v", string(result), expectedResult)
    }

    expectedStatus = true
    if result := lar.EndOfInput(); result != expectedStatus {
        t.Errorf("got %v, wanted %v", result, expectedStatus)
    }
}

func TestMatchDigitTwoDigits(t *testing.T) {
    lar, err := NewLookaheadReader(strings.NewReader("12"))
    if err != nil {
        t.Errorf("couldn't create lookahead reader: %v", err)
    }

    expectedStatus := true
    if result := lar.MatchDigit(); result != expectedStatus {
        t.Errorf("got %v, wanted %v", result, expectedStatus)
    }

    expectedResult := "1"
    if result := lar.Result(); string(result) != expectedResult {
        t.Errorf("got %v, wanted %v", string(result), expectedResult)
    }

    expectedStatus = true
    if result := lar.MatchDigit(); result != expectedStatus {
        t.Errorf("got %v, wanted %v", result, expectedStatus)
    }

    expectedResult = "2"
    if result := lar.Result(); string(result) != expectedResult {
        t.Errorf("got %v, wanted %v", string(result), expectedResult)
    }

    expectedStatus = true
    if result := lar.EndOfInput(); result != expectedStatus {
        t.Errorf("got %v, wanted %v", result, expectedStatus)
    }
}

func TestMatchDigitsOneDigit(t *testing.T) {
    lar, err := NewLookaheadReader(strings.NewReader("1"))
    if err != nil {
        t.Errorf("couldn't create lookahead reader: %v", err)
    }

    expectedStatus := true
    if result := lar.MatchDigits(); result != expectedStatus {
        t.Errorf("got %v, wanted %v", result, expectedStatus)
    }

    expectedResult := "1"
    if result := lar.Result(); string(result) != expectedResult {
        t.Errorf("got %v, wanted %v", string(result), expectedResult)
    }

    expectedStatus = true
    if result := lar.EndOfInput(); result != expectedStatus {
        t.Errorf("got %v, wanted %v", result, expectedStatus)
    }
}

func TestMatchDigitsTwoDigits(t *testing.T) {
    lar, err := NewLookaheadReader(strings.NewReader("12"))
    if err != nil {
        t.Errorf("couldn't create lookahead reader: %v", err)
    }

    expectedStatus := true
    if result := lar.MatchDigits(); result != expectedStatus {
        t.Errorf("got %v, wanted %v", result, expectedStatus)
    }

    expectedResult := "12"
    if result := lar.Result(); string(result) != expectedResult {
        t.Errorf("got %v, wanted %v", string(result), expectedResult)
    }

    expectedStatus = true
    if result := lar.EndOfInput(); result != expectedStatus {
        t.Errorf("got %v, wanted %v", result, expectedStatus)
    }
}

func TestMatchSpaceOneSpace(t *testing.T) {
    lar, err := NewLookaheadReader(strings.NewReader(" "))
    if err != nil {
        t.Errorf("couldn't create lookahead reader: %v", err)
    }

    expectedStatus := true
    if result := lar.MatchSpace(); result != expectedStatus {
        t.Errorf("got %v, wanted %v", result, expectedStatus)
    }

    expectedResult := " "
    if result := lar.Result(); string(result) != expectedResult {
        t.Errorf("got %v, wanted %v", string(result), expectedResult)
    }

    expectedStatus = true
    if result := lar.EndOfInput(); result != expectedStatus {
        t.Errorf("got %v, wanted %v", result, expectedStatus)
    }
}

func TestMatchSpaceTwoSpaces(t *testing.T) {
    lar, err := NewLookaheadReader(strings.NewReader(" \t"))
    if err != nil {
        t.Errorf("couldn't create lookahead reader: %v", err)
    }

    expectedStatus := true
    if result := lar.MatchSpace(); result != expectedStatus {
        t.Errorf("got %v, wanted %v", result, expectedStatus)
    }

    expectedResult := " "
    if result := lar.Result(); string(result) != expectedResult {
        t.Errorf("got %v, wanted %v", string(result), expectedResult)
    }

    expectedStatus = true
    if result := lar.MatchSpace(); result != expectedStatus {
        t.Errorf("got %v, wanted %v", result, expectedStatus)
    }

    expectedResult = "\t"
    if result := lar.Result(); string(result) != expectedResult {
        t.Errorf("got %v, wanted %v", string(result), expectedResult)
    }

    expectedStatus = true
    if result := lar.EndOfInput(); result != expectedStatus {
        t.Errorf("got %v, wanted %v", result, expectedStatus)
    }
}

func TestMatchSpacesOneSpace(t *testing.T) {
    lar, err := NewLookaheadReader(strings.NewReader(" "))
    if err != nil {
        t.Errorf("couldn't create lookahead reader: %v", err)
    }

    expectedStatus := true
    if result := lar.MatchSpaces(); result != expectedStatus {
        t.Errorf("got %v, wanted %v", result, expectedStatus)
    }

    expectedResult := " "
    if result := lar.Result(); string(result) != expectedResult {
        t.Errorf("got %v, wanted %v", string(result), expectedResult)
    }

    expectedStatus = true
    if result := lar.EndOfInput(); result != expectedStatus {
        t.Errorf("got %v, wanted %v", result, expectedStatus)
    }
}

func TestMatchSpacesTwoSpaces(t *testing.T) {
    lar, err := NewLookaheadReader(strings.NewReader(" \t"))
    if err != nil {
        t.Errorf("couldn't create lookahead reader: %v", err)
    }

    expectedStatus := true
    if result := lar.MatchSpaces(); result != expectedStatus {
        t.Errorf("got %v, wanted %v", result, expectedStatus)
    }

    expectedResult := " \t"
    if result := lar.Result(); string(result) != expectedResult {
        t.Errorf("got %v, wanted %v", string(result), expectedResult)
    }

    expectedStatus = true
    if result := lar.EndOfInput(); result != expectedStatus {
        t.Errorf("got %v, wanted %v", result, expectedStatus)
    }
}

func TestMatchSingleLetterDigit(t *testing.T) {
    lar, err := NewLookaheadReader(strings.NewReader("a1"))
    if err != nil {
        t.Errorf("couldn't create lookahead reader: %v", err)
    }

    expectedStatus := true
    if result := lar.MatchLetter(); result != expectedStatus {
        t.Errorf("got %v, wanted %v", result, expectedStatus)
    }

    expectedResult := "a"
    if result := lar.Result(); string(result) != expectedResult {
        t.Errorf("got %v, wanted %v", string(result), expectedResult)
    }

    expectedStatus = true
    if result := lar.MatchDigit(); result != expectedStatus {
        t.Errorf("got %v, wanted %v", result, expectedStatus)
    }

    expectedResult = "1"
    if result := lar.Result(); string(result) != expectedResult {
        t.Errorf("got %v, wanted %v", string(result), expectedResult)
    }

    expectedStatus = true
    if result := lar.EndOfInput(); result != expectedStatus {
        t.Errorf("got %v, wanted %v", result, expectedStatus)
    }
}

func TestMatchSingleDigitSpace(t *testing.T) {
    lar, err := NewLookaheadReader(strings.NewReader("1 "))
    if err != nil {
        t.Errorf("couldn't create lookahead reader: %v", err)
    }

    expectedStatus := true
    if result := lar.MatchDigit(); result != expectedStatus {
        t.Errorf("got %v, wanted %v", result, expectedStatus)
    }

    expectedResult := "1"
    if result := lar.Result(); string(result) != expectedResult {
        t.Errorf("got %v, wanted %v", string(result), expectedResult)
    }

    expectedStatus = true
    if result := lar.MatchSpace(); result != expectedStatus {
        t.Errorf("got %v, wanted %v", result, expectedStatus)
    }

    expectedResult = " "
    if result := lar.Result(); string(result) != expectedResult {
        t.Errorf("got %v, wanted %v", string(result), expectedResult)
    }

    expectedStatus = true
    if result := lar.EndOfInput(); result != expectedStatus {
        t.Errorf("got %v, wanted %v", result, expectedStatus)
    }
}

func TestMatchSingleSpaceLetter(t *testing.T) {
    lar, err := NewLookaheadReader(strings.NewReader("\tz"))
    if err != nil {
        t.Errorf("couldn't create lookahead reader: %v", err)
    }

    expectedStatus := true
    if result := lar.MatchSpace(); result != expectedStatus {
        t.Errorf("got %v, wanted %v", result, expectedStatus)
    }

    expectedResult := "\t"
    if result := lar.Result(); string(result) != expectedResult {
        t.Errorf("got %v, wanted %v", string(result), expectedResult)
    }

    expectedStatus = true
    if result := lar.MatchLetter(); result != expectedStatus {
        t.Errorf("got %v, wanted %v", result, expectedStatus)
    }

    expectedResult = "z"
    if result := lar.Result(); string(result) != expectedResult {
        t.Errorf("got %v, wanted %v", string(result), expectedResult)
    }

    expectedStatus = true
    if result := lar.EndOfInput(); result != expectedStatus {
        t.Errorf("got %v, wanted %v", result, expectedStatus)
    }
}

func TestMatchMultipleLettersDigits(t *testing.T) {
    lar, err := NewLookaheadReader(strings.NewReader("abc123"))
    if err != nil {
        t.Errorf("couldn't create lookahead reader: %v", err)
    }

    expectedStatus := true
    if result := lar.MatchLetters(); result != expectedStatus {
        t.Errorf("got %v, wanted %v", result, expectedStatus)
    }

    expectedResult := "abc"
    if result := lar.Result(); string(result) != expectedResult {
        t.Errorf("got %v, wanted %v", string(result), expectedResult)
    }

    expectedStatus = true
    if result := lar.MatchDigits(); result != expectedStatus {
        t.Errorf("got %v, wanted %v", result, expectedStatus)
    }

    expectedResult = "123"
    if result := lar.Result(); string(result) != expectedResult {
        t.Errorf("got %v, wanted %v", string(result), expectedResult)
    }

    expectedStatus = true
    if result := lar.EndOfInput(); result != expectedStatus {
        t.Errorf("got %v, wanted %v", result, expectedStatus)
    }
}

func TestMatchMultipleDigitsSpaces(t *testing.T) {
    lar, err := NewLookaheadReader(strings.NewReader("123 \t\r"))
    if err != nil {
        t.Errorf("couldn't create lookahead reader: %v", err)
    }

    expectedStatus := true
    if result := lar.MatchDigits(); result != expectedStatus {
        t.Errorf("got %v, wanted %v", result, expectedStatus)
    }

    expectedResult := "123"
    if result := lar.Result(); string(result) != expectedResult {
        t.Errorf("got %v, wanted %v", string(result), expectedResult)
    }

    expectedStatus = true
    if result := lar.MatchSpaces(); result != expectedStatus {
        t.Errorf("got %v, wanted %v", result, expectedStatus)
    }

    expectedResult = " \t\r"
    if result := lar.Result(); string(result) != expectedResult {
        t.Errorf("got %v, wanted %v", string(result), expectedResult)
    }

    expectedStatus = true
    if result := lar.EndOfInput(); result != expectedStatus {
        t.Errorf("got %v, wanted %v", result, expectedStatus)
    }
}

func TestMatchMultipleSpacesLetters(t *testing.T) {
    lar, err := NewLookaheadReader(strings.NewReader(" \t\rabc"))
    if err != nil {
        t.Errorf("couldn't create lookahead reader: %v", err)
    }

    expectedStatus := true
    if result := lar.MatchSpaces(); result != expectedStatus {
        t.Errorf("got %v, wanted %v", result, expectedStatus)
    }

    expectedResult := " \t\r"
    if result := lar.Result(); string(result) != expectedResult {
        t.Errorf("got %v, wanted %v", string(result), expectedResult)
    }

    expectedStatus = true
    if result := lar.MatchLetters(); result != expectedStatus {
        t.Errorf("got %v, wanted %v", result, expectedStatus)
    }

    expectedResult = "abc"
    if result := lar.Result(); string(result) != expectedResult {
        t.Errorf("got %v, wanted %v", string(result), expectedResult)
    }

    expectedStatus = true
    if result := lar.EndOfInput(); result != expectedStatus {
        t.Errorf("got %v, wanted %v", result, expectedStatus)
    }
}

func TestLetterNonMatches(t *testing.T) {
    lar, err := NewLookaheadReader(strings.NewReader("a"))
    if err != nil {
        t.Errorf("couldn't create lookahead reader: %v", err)
    }

    expectedStatus := false
    if result := lar.MatchDigit(); result != expectedStatus {
        t.Errorf("got %v, wanted %v", result, expectedStatus)
    }

    expectedStatus = false
    if result := lar.MatchSpace(); result != expectedStatus {
        t.Errorf("got %v, wanted %v", result, expectedStatus)
    }

    expectedStatus = false
    if result := lar.EndOfInput(); result != expectedStatus {
        t.Errorf("got %v, wanted %v", result, expectedStatus)
    }
}

func TestDigitNonMatches(t *testing.T) {
    lar, err := NewLookaheadReader(strings.NewReader("1"))
    if err != nil {
        t.Errorf("couldn't create lookahead reader: %v", err)
    }

    expectedStatus := false
    if result := lar.MatchLetter(); result != expectedStatus {
        t.Errorf("got %v, wanted %v", result, expectedStatus)
    }

    expectedStatus = false
    if result := lar.MatchSpace(); result != expectedStatus {
        t.Errorf("got %v, wanted %v", result, expectedStatus)
    }

    expectedStatus = false
    if result := lar.EndOfInput(); result != expectedStatus {
        t.Errorf("got %v, wanted %v", result, expectedStatus)
    }
}

func TestSpaceNonMatches(t *testing.T) {
    lar, err := NewLookaheadReader(strings.NewReader(" "))
    if err != nil {
        t.Errorf("couldn't create lookahead reader: %v", err)
    }

    expectedStatus := false
    if result := lar.MatchLetter(); result != expectedStatus {
        t.Errorf("got %v, wanted %v", result, expectedStatus)
    }

    expectedStatus = false
    if result := lar.MatchDigit(); result != expectedStatus {
        t.Errorf("got %v, wanted %v", result, expectedStatus)
    }

    expectedStatus = false
    if result := lar.EndOfInput(); result != expectedStatus {
        t.Errorf("got %v, wanted %v", result, expectedStatus)
    }
}

func TestMatchOneOfPresent(t *testing.T) {
    lar, err := NewLookaheadReader(strings.NewReader(":"))
    if err != nil {
        t.Errorf("couldn't create lookahead reader: %v", err)
    }

    expectedStatus := true
    if result := lar.MatchOneOf('.', ';', ':'); result != expectedStatus {
        t.Errorf("got %v, wanted %v", result, expectedStatus)
    }

    expectedResult := ":"
    if result := lar.Result(); string(result) != expectedResult {
        t.Errorf("got %v, wanted %v", string(result), expectedResult)
    }

    expectedStatus = true
    if result := lar.EndOfInput(); result != expectedStatus {
        t.Errorf("got %v, wanted %v", result, expectedStatus)
    }
}

func TestMatchOneOfNotPresent(t *testing.T) {
    lar, err := NewLookaheadReader(strings.NewReader("?"))
    if err != nil {
        t.Errorf("couldn't create lookahead reader: %v", err)
    }

    expectedStatus := false
    if result := lar.MatchOneOf('.', ';', ':'); result != expectedStatus {
        t.Errorf("got %v, wanted %v", result, expectedStatus)
    }

    expectedStatus = false
    if result := lar.EndOfInput(); result != expectedStatus {
        t.Errorf("got %v, wanted %v", result, expectedStatus)
    }
}

func TestNext(t *testing.T) {
    lar, err := NewLookaheadReader(strings.NewReader("a1\t?"))
    if err != nil {
        t.Errorf("couldn't create lookahead reader: %v", err)
    }

    var expectedError error = nil
    expectedResult := "a"
    result, err := lar.Next();
    if err != nil {
        t.Errorf("got %v, wanted %v", err, expectedError)
    }
    if string(result) != expectedResult {
        t.Errorf("got %v, wanted %v", string(result), expectedResult)
    }

    expectedResult = "1"
    result, err = lar.Next();
    if err != nil {
        t.Errorf("got %v, wanted %v", err, expectedError)
    }
    if string(result) != expectedResult {
        t.Errorf("got %v, wanted %v", string(result), expectedResult)
    }

    expectedResult = "\t"
    result, err = lar.Next();
    if err != nil {
        t.Errorf("got %v, wanted %v", err, expectedError)
    }
    if string(result) != expectedResult {
        t.Errorf("got %v, wanted %v", string(result), expectedResult)
    }

    expectedResult = "?"
    result, err = lar.Next();
    if err != nil {
        t.Errorf("got %v, wanted %v", err, expectedError)
    }
    if string(result) != expectedResult {
        t.Errorf("got %v, wanted %v", string(result), expectedResult)
    }

    expectedStatus := true
    if result := lar.EndOfInput(); result != expectedStatus {
        t.Errorf("got %v, wanted %v", result, expectedStatus)
    }
}

