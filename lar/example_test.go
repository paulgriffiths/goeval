package lar

import (
	"fmt"
	"strings"
)

func Example() {
	lar, _ := NewLookaheadReader(strings.NewReader("abc123?"))

	if lar.MatchDigits() {
		fmt.Printf("Matched digits '%s'.\n", lar.Result())
	} else {
		fmt.Printf("Didn't match any digits.\n")
	}

	if lar.MatchLetters() {
		fmt.Printf("Matched letters '%s'.\n", lar.Result())
	} else {
		fmt.Printf("Didn't match any digits.\n")
	}

	if lar.MatchDigits() {
		fmt.Printf("Matched digits '%s'.\n", lar.Result())
	} else {
		fmt.Printf("Didn't match any digits.\n")
	}

	if lar.MatchOneOf('!', '$', '%') {
		fmt.Printf("Matched character '%s'.\n", lar.Result())
	} else {
		fmt.Printf("Didn't match a character.\n")
	}

	if lar.MatchOneOf('{', '?', '#', '@') {
		fmt.Printf("Matched character '%s'.\n", lar.Result())
	} else {
		fmt.Printf("Didn't match a character.\n")
	}

	if lar.EndOfInput() {
		fmt.Printf("All input has been matched.\n")
	}

	// Output: Didn't match any digits.
	// Matched letters 'abc'.
	// Matched digits '123'.
	// Didn't match a character.
	// Matched character '?'.
	// All input has been matched.
}
