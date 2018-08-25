package tokens

// Token contains a lexeme token type and its value.
type Token struct {
	Type  TokenType
	Value string
}

// TokenType represents the type of a token.
type TokenType int

// Enumerated values for token types.
const (
	Number TokenType = iota
	String
	Word
	Keyword
	Identifier
	LeftParen
	RightParen
	LeftBrace
	RightBrace
	LeftBracket
	RightBracket
	AddOperator
	SubOperator
	MulOperator
	DivOperator
	PowOperator
	NegOperator
	AndOperator
	OrOperator
	XorOperator
	NandOperator
	NorOperator
	NotOperator
	AssignmentOperator
	EqualityOperator
	NonEqualityOperator
	LessOperator
	LessEqualOperator
	GreaterOperator
	GreaterEqualOperator
	Newline
	StatementSeparator
	Illegal
	nullTokenID
)

// NullToken returns a token representing a null token.
// This function is not exported, because the lookahead token
// channel uses this as it's "no token" value, and behaves strangely
// if the null token is actually passed to it as input, so we
// don't export the symbol to avoid that.
func nullToken() Token {
	return Token{nullTokenID, ""}
}

// New creates a new token with the specified type and value.
func New(tokenType TokenType, value string) Token {
	return Token{tokenType, value}
}
