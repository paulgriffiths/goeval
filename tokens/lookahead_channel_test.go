package tokens

import "testing"

func TestLTChanNext(t *testing.T) {
    var cases = []Token {
        WordToken("foobar"),
        NumberToken("67.89"),
        OperatorToken("+"),
        LeftParenToken(),
        RightParenToken(),
        IllegalToken("dracula"),
    }

    ch := make(chan Token)

    go func() {
        for _, token := range cases {
            ch <- token
        }
        close(ch)
    } ()

    ltc := NewLTChan(ch)

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
    var cases = []Token {
        WordToken("foobar"),
        NumberToken("67.89"),
        OperatorToken("+"),
        LeftParenToken(),
        RightParenToken(),
        IllegalToken("dracula"),
    }

    ch := make(chan Token)

    go func() {
        for _, token := range cases {
            ch <- token
        }
        close(ch)
    } ()

    ltc := NewLTChan(ch)

    for _, token := range cases {
        result := ltc.Match(token)
        if !result {
            t.Errorf("couldn't match %v when expected", token)
        }
    }

    if !ltc.IsEmpty() {
        t.Errorf("ltchan not empty when expected")
    }
}

func TestLTDontMatch(t *testing.T) {
    var cases = []Token {
        WordToken("foobar"),
        NumberToken("67.89"),
        OperatorToken("+"),
        LeftParenToken(),
        RightParenToken(),
        IllegalToken("dracula"),
    }

    ch := make(chan Token)

    go func() {
        for _, token := range cases {
            ch <- token
        }
        close(ch)
    } ()

    ltc := NewLTChan(ch)

    for _, token := range cases {
        result := ltc.Match(nullToken())
        if result {
            t.Errorf("matched %v when not expected", token)
        }
    }

    if ltc.IsEmpty() {
        t.Errorf("ltchan empty when not expected")
    }

    ltc.Flush()

    if !ltc.IsEmpty() {
        t.Errorf("ltchan not empty when expected")
    }
}

func TestLTMatchType(t *testing.T) {
    var cases = []Token {
        WordToken("foobar"),
        NumberToken("67.89"),
        OperatorToken("+"),
        LeftParenToken(),
        RightParenToken(),
        IllegalToken("dracula"),
    }

    var matches = []Token {
        WordToken("sausage"),
        NumberToken("2.42"),
        OperatorToken("-"),
        LeftParenToken(),
        RightParenToken(),
        IllegalToken("chips"),
    }

    ch := make(chan Token)

    go func() {
        for _, token := range cases {
            ch <- token
        }
        close(ch)
    } ()

    ltc := NewLTChan(ch)

    for _, token := range matches {
        result := ltc.MatchType(token)
        if !result {
            t.Errorf("couldn't match %v when expected", token)
        }
    }

    if !ltc.IsEmpty() {
        t.Errorf("ltchan not empty when expected")
    }
}

func TestLTDontMatchType(t *testing.T) {
    var cases = []Token {
        WordToken("foobar"),
        NumberToken("67.89"),
        OperatorToken("+"),
        LeftParenToken(),
        RightParenToken(),
        IllegalToken("dracula"),
    }

    ch := make(chan Token)

    go func() {
        for _, token := range cases {
            ch <- token
        }
        close(ch)
    } ()

    ltc := NewLTChan(ch)

    for _, token := range cases {
        result := ltc.MatchType(nullToken())
        if result {
            t.Errorf("matched %v when not expected", token)
        }
    }

    if ltc.IsEmpty() {
        t.Errorf("ltchan empty when not expected")
    }

    ltc.Flush()

    if !ltc.IsEmpty() {
        t.Errorf("ltchan not empty when expected")
    }
}
