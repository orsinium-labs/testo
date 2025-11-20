package lexer

// TokenType represents the type of lexical token.
type TokenType string

// TokenType constants define all possible token types in the lexer.
const (
	ILLEGAL TokenType = "ILLEGAL" // Represents an invalid or unrecognized token
	EOF     TokenType = "EOF"     // End of input

	LBRACE   TokenType = "{"
	RBRACE   TokenType = "}"
	LBRACKET TokenType = "["
	RBRACKET TokenType = "]"
	COLON    TokenType = ":"
	COMMA    TokenType = ","

	STRING TokenType = "STRING"
	NUMBER TokenType = "NUMBER"

	TRUE  TokenType = "TRUE"
	FALSE TokenType = "FALSE"
	NULL  TokenType = "NULL"

	TYPE_ANY    TokenType = "MATCH_ANY"
	TYPE_STRING TokenType = "MATCH_STRING"
	TYPE_BOOL   TokenType = "MATCH_BOOL"
	TYPE_INT    TokenType = "MATCH_INT"
)

// Token represents a single lexical token with its type, literal value, and position in the input.
type Token struct {
	Type    TokenType // The type of the token
	Literal string    // The literal value of the token
	Line    int       // Line number where the token appears
	Column  int       // Column number where the token appears
}
