package cfl

import (
	"github.com/paulgriffiths/goeval/lar"
	"io"
)

// lex extracts a list of tokens from a context-free grammar.
func lex(input io.Reader) (tokenList, lexErr) {
	reader, err := lar.NewLookaheadReader(input)
	if err != nil {
		return nil, lexError{lexErrBadInput, "", lar.FilePos{0, 0}}
	}

	tokens := []token{}
	startOfLine := true // True if no tokens yet on current line

	for {

		// Ignore leading whitespace

		reader.MatchSpaces()
		if reader.EndOfInput() {
			break
		}

		// Only return an end-of-line token for lines which
		// have some other token on them (i.e. don't return
		// end-of-line tokens for blank lines and comment-only
		// lines).

		if startOfLine && reader.MatchOneOf('#') {
			for reader.MatchAnyExcept('\n') {
			}
		}

		if startOfLine && reader.MatchNewline() {
			continue
		} else {
			startOfLine = false
		}

		// Extract the next token

		switch {
		case reader.MatchOneOf('#'):
			for reader.MatchAnyExcept('\n') {
			}
		case reader.MatchNewline():
			tokens = append(tokens,
				token{tokenEndOfLine, "", reader.Result.Pos})
			startOfLine = true
		case reader.MatchOneOf(':'):
			tokens = append(tokens,
				token{tokenArrow, ":", reader.Result.Pos})
		case reader.MatchOneOf('|'):
			tokens = append(tokens,
				token{tokenAlt, "|", reader.Result.Pos})
		case reader.MatchLetter():
			t := string(reader.Result.Value[0])
			pos := reader.Result.Pos
			for {
				if reader.MatchLetters() {
					t = t + string(reader.Result.Value)
				} else if reader.MatchDigits() {
					t = t + string(reader.Result.Value)
				} else if reader.MatchOneOf('\'') {
					t = t + string(reader.Result.Value)
				} else {
					break
				}
			}
			if t == "e" {
				tokens = append(tokens, token{tokenEmpty, t, pos})
			} else {
				tokens = append(tokens, token{tokenNonTerminal, t, pos})
			}
		case reader.MatchOneOf('`'):
			t := ""
			pos := lar.FilePos{}
			startPos := reader.Result.Pos
			for reader.MatchAnyExcept('`', '\n') {
				t += string(reader.Result.Value)
				pos = reader.Result.Pos
			}
			if !reader.MatchOneOf('`') {
				return nil, lexError{lexErrUnterminatedTerminal, "", pos}
			}
			tokens = append(tokens, token{tokenTerminal, t, startPos})
		default:
			reader.MatchAnyExcept()
			return nil, lexError{lexErrIllegalCharacter,
				string(reader.Result.Value[0]), reader.Result.Pos}
		}
	}
	return tokens, nil
}
