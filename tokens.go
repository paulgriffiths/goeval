package eval

type Token struct {
    tokenType int
    value string
}

const (
    operatorToken int = iota
    numberToken
    functionToken
    leftParenToken
    rightParenToken
    illegalToken
    nullToken
)

func NullToken() Token {
    return Token{nullToken, ""}
}

func LeftParenToken() Token {
    return Token{leftParenToken, "("}
}

func RightParenToken() Token {
    return Token{rightParenToken, ")"}
}

func NewToken(tokenType int, value string) Token {
    return Token{tokenType, value}
}

func NewOperatorToken(value string) Token {
    return Token{operatorToken, value}
}

func NewNumberToken(value string) Token {
    return Token{numberToken, value}
}

func NewFunctionToken(value string) Token {
    return Token{functionToken, value}
}

func NewIllegalToken(value string) Token {
    return Token{illegalToken, value}
}

func (t Token) Value() string {
    return t.value
}

func (t Token) IsOperatorWithValue(value string) bool {
    return t.tokenType == operatorToken && t.value == value
}

func (t Token) IsNumber() bool {
    return t.tokenType == numberToken
}

func (t Token) IsLeftParen() bool {
    return t.tokenType == leftParenToken
}

func (t Token) IsRightParen() bool {
    return t.tokenType == rightParenToken
}

func (t Token) IsIllegal() bool {
    return t.tokenType == illegalToken
}

