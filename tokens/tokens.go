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
    nullTokenId
)

// NullToken returns a token representing a null token.
// This function is not exported, because the lookahead token
// channel uses this as it's "no token" value, and behaves strangely
// if the null token is actually passed to it as input, so we
// don't export the symbol to avoid that.
func nullToken() Token {
    return Token{nullTokenId, ""}
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
// It can be passed to a function which compares the types of tokens,
// but which doesn't care about the values.
func ZeroNumberToken() Token {
    return Token{numberToken, "0"}
}

// EmptyWordToken returns a word token containing the empty string.
// It can be passed to a function which compares the types of tokens,
// but which doesn't care about the values.
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
