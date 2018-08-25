package tokens

import "testing"

var cases = []Token{
	{Word, "sausage"},
	{Number, "2.42"},
	{SubOperator, "-"},
	{LeftParen, "("},
	{RightParen, "}"},
	{Identifier, "chips"},
	{Illegal, "?"},
}

func makeTestChannel() chan Token {
	ch := make(chan Token)

	go func() {
		for _, token := range cases {
			ch <- token
		}
		close(ch)
	}()

	return ch
}

func TestLTChanNext(t *testing.T) {
	ltc := NewLTChan(makeTestChannel())

	for _, token := range cases {
		result, err := ltc.Next()
		if err != nil {
			t.Errorf("couldn't retrieve token: %v", err)
		} else if result != token {
			t.Errorf("got %v, want %v", result, token)
		}
	}

	if !ltc.IsEmpty() {
		t.Errorf("ltchan not empty when expected")
	}
}

func TestLTMatch(t *testing.T) {
	ltc := NewLTChan(makeTestChannel())

	for _, token := range cases {
		result := ltc.MatchToken(token)
		if !result {
			t.Errorf("couldn't match %v when expected", token)
		}
	}

	if !ltc.IsEmpty() {
		t.Errorf("ltchan not empty when expected")
	}
}

func TestLTDontMatch(t *testing.T) {
	ltc := NewLTChan(makeTestChannel())

	for _, token := range cases {
		result := ltc.MatchToken(nullToken())
		if result {
			t.Errorf("matched %v when not expected", token)
		}
	}

	if ltc.IsEmpty() {
		t.Errorf("ltchan empty when not expected")
	}
}

func TestLTMatchType(t *testing.T) {
	var matches = []Token{
		{Word, "sausage"},
		{Number, "2.42"},
		{SubOperator, "-"},
		{LeftParen, "("},
		{RightParen, "}"},
		{Identifier, "chips"},
		{Illegal, "?"},
	}

	ltc := NewLTChan(makeTestChannel())

	for _, token := range matches {
		result := ltc.Match(token.Type)
		if !result {
			t.Errorf("couldn't match %v when expected", token)
		}
	}

	if !ltc.IsEmpty() {
		t.Errorf("ltchan not empty when expected")
	}
}

func TestLTDontMatchType(t *testing.T) {
	ltc := NewLTChan(makeTestChannel())

	for _, token := range cases {
		result := ltc.MatchToken(nullToken())
		if result {
			t.Errorf("matched %v when not expected", token)
		}
	}

	if ltc.IsEmpty() {
		t.Errorf("ltchan empty when not expected")
	}
}

func TestFlush(t *testing.T) {
	ltc := NewLTChan(makeTestChannel())

	if ltc.IsEmpty() {
		t.Errorf("ltchan empty when not expected")
	}

	ltc.Flush()

	if !ltc.IsEmpty() {
		t.Errorf("ltchan not empty when expected")
	}
}
