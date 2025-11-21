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
	TYPE_UINT   TokenType = "MATCH_UINT"
	TYPE_FLOAT  TokenType = "MATCH_FLOAT"
	TYPE_OBJECT TokenType = "MATCH_OBJECT"
	TYPE_ARRAY  TokenType = "MATCH_ARRAY"

	TYPE_STRINGS TokenType = "MATCH_STRINGS"
	TYPE_BOOLS   TokenType = "MATCH_BOOLS"
	TYPE_INTS    TokenType = "MATCH_INTS"
	TYPE_UINTS   TokenType = "MATCH_UINTS"
	TYPE_FLOATS  TokenType = "MATCH_FLOATS"
	TYPE_OBJECTS TokenType = "MATCH_OBJECTS"
)

// Token represents a single lexical token with its type, literal value, and position in the input.
type Token struct {
	Type    TokenType // The type of the token
	Literal string    // The literal value of the token
	Line    int       // Line number where the token appears
	Column  int       // Column number where the token appears
}
