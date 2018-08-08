// Package tokens provides a lexeme token type and associated functions.
package tokens

// Token contains a lexeme token type and its value.
type Token struct {
    tokenType int
    value string
}

// Enumerated values for token types.
const (
    operatorToken int = iota
    numberToken
    wordToken
    leftParenToken
    rightParenToken
    illegalToken
    nullToken
)

// NullToken returns a token representing a null token.
func NullToken() Token {
    return Token{nullToken, ""}
}

// LeftParenToken returns a left parenthesis token.
func LeftParenToken() Token {
    return Token{leftParenToken, "("}
}

// RightParenToken returns a right parenthesis token.
func RightParenToken() Token {
    return Token{rightParenToken, ")"}
}

// ZeroNumberToken returns a token representing the number zero.
func ZeroNumberToken() Token {
    return Token{numberToken, "0"}
}

// ZeroNumberToken returns a token representing the number zero.
func EmptyWordToken() Token {
    return Token{wordToken, ""}
}

// OperatorToken returns an operator token with the specified value.
func OperatorToken(value string) Token {
    return Token{operatorToken, value}
}

// NumberToken returns a number token with the specified value.
func NumberToken(value string) Token {
    return Token{numberToken, value}
}

// WordToken returns a word token with the specified value.
func WordToken(value string) Token {
    return Token{wordToken, value}
}

// IllegalToken returns an illegal token with the specified value.
func IllegalToken(value string) Token {
    return Token{illegalToken, value}
}

// Value returns a token's value.
func (t Token) Value() string {
    return t.value
}

// IsOperatorWithValue returns true if the token is an operator token
// and its value is equal to the value specified.
func (t Token) IsOperatorWithValue(value string) bool {
    return t.tokenType == operatorToken && t.value == value
}

// IsOperator returns true if the token is an operator token.
func (t Token) IsOperator() bool {
    return t.tokenType == operatorToken
}

// IsNumber returns true if the token is a number token.
func (t Token) IsNumber() bool {
    return t.tokenType == numberToken
}

// IsWord returns true if the token is a word token.
func (t Token) IsWord() bool {
    return t.tokenType == wordToken
}

// IsLeftParen returns true if the token is a left parenthesis token.
func (t Token) IsLeftParen() bool {
    return t.tokenType == leftParenToken
}

// IsRightParen returns true if the token is a right parenthesis token.
func (t Token) IsRightParen() bool {
    return t.tokenType == rightParenToken
}

// IsIllegal returns true if the token is an illegal token.
func (t Token) IsIllegal() bool {
    return t.tokenType == illegalToken
}
