package lexer_test

import (
	"testing"

	"github.com/orsinium-labs/testo/internal/lexer"
)

func TestNextToken_EmptyInput(t *testing.T) {
	// Test empty input string
	input := ""

	// Expected token is EOF
	tests := []struct {
		expectedType    lexer.TokenType
		expectedLiteral string
	}{
		{lexer.EOF, ""}, // End of input
	}

	// Initialize the lexer with the empty input
	lexer := lexer.New(input)

	// Check for EOF token
	tok := lexer.NextToken()

	// Log the token for debugging purposes
	t.Logf("Token: Type=%q, Literal=%q, Line=%d, Column=%d", tok.Type, tok.Literal, tok.Line, tok.Column)

	// Validate the token type
	if tok.Type != tests[0].expectedType {
		t.Fatalf("expected=%q, got=%q (literal=%q)", tests[0].expectedType, tok.Type, tok.Literal)
	}

	// Validate the token literal
	if tok.Literal != tests[0].expectedLiteral {
		t.Fatalf("expected=%q, got=%q (type=%q)", tests[0].expectedLiteral, tok.Literal, tok.Type)
	}
}

func TestNextToken_ValidCharacters(t *testing.T) {
	// Define the input string with various edge cases and scenarios
	input := `{"name":"Alice","age":25,"isStudent":false,"courses":["Math","Science","History"],"address":{"street":"123 Main St","city":"Metropolis","zipcode":"12345"},"graduationYear":null,"isGraduated":true,"height":5.7,"siblings":[{"name":"Bob","age":22},{"name":"Charlie","age":28}]}`

	// Define the expected sequence of tokens for these scenarios
	tests := []struct {
		expectedType    lexer.TokenType
		expectedLiteral string
	}{
		{lexer.LBRACE, "{"},              // Start of object
		{lexer.STRING, "name"},           // String key
		{lexer.COLON, ":"},               // Colon separator
		{lexer.STRING, "Alice"},          // String value
		{lexer.COMMA, ","},               // Comma separator
		{lexer.STRING, "age"},            // String key
		{lexer.COLON, ":"},               // Colon separator
		{lexer.NUMBER, "25"},             // Number
		{lexer.COMMA, ","},               // Comma separator
		{lexer.STRING, "isStudent"},      // String key
		{lexer.COLON, ":"},               // Colon separator
		{lexer.FALSE, "false"},           // Boolean value (false)
		{lexer.COMMA, ","},               // Comma separator
		{lexer.STRING, "courses"},        // String key
		{lexer.COLON, ":"},               // Colon separator
		{lexer.LBRACKET, "["},            // Start of array
		{lexer.STRING, "Math"},           // String in array
		{lexer.COMMA, ","},               // Comma separator
		{lexer.STRING, "Science"},        // String in array
		{lexer.COMMA, ","},               // Comma separator
		{lexer.STRING, "History"},        // String in array
		{lexer.RBRACKET, "]"},            // End of array
		{lexer.COMMA, ","},               // Comma separator
		{lexer.STRING, "address"},        // String key
		{lexer.COLON, ":"},               // Colon separator
		{lexer.LBRACE, "{"},              // Start of nested object
		{lexer.STRING, "street"},         // String key
		{lexer.COLON, ":"},               // Colon separator
		{lexer.STRING, "123 Main St"},    // String value
		{lexer.COMMA, ","},               // Comma separator
		{lexer.STRING, "city"},           // String key
		{lexer.COLON, ":"},               // Colon separator
		{lexer.STRING, "Metropolis"},     // String value
		{lexer.COMMA, ","},               // Comma separator
		{lexer.STRING, "zipcode"},        // String key
		{lexer.COLON, ":"},               // Colon separator
		{lexer.STRING, "12345"},          // String value
		{lexer.RBRACE, "}"},              // End of nested object
		{lexer.COMMA, ","},               // Comma separator
		{lexer.STRING, "graduationYear"}, // String key
		{lexer.COLON, ":"},               // Colon separator
		{lexer.NULL, "null"},             // Null value
		{lexer.COMMA, ","},               // Comma separator
		{lexer.STRING, "isGraduated"},    // String key
		{lexer.COLON, ":"},               // Colon separator
		{lexer.TRUE, "true"},             // Boolean value (true)
		{lexer.COMMA, ","},               // Comma separator
		{lexer.STRING, "height"},         // String key
		{lexer.COLON, ":"},               // Colon separator
		{lexer.NUMBER, "5.7"},            // Number value
		{lexer.COMMA, ","},               // Comma separator
		{lexer.STRING, "siblings"},       // String key
		{lexer.COLON, ":"},               // Colon separator
		{lexer.LBRACKET, "["},            // Start of nested array
		{lexer.LBRACE, "{"},              // Start of inner object
		{lexer.STRING, "name"},           // String key
		{lexer.COLON, ":"},               // Colon separator
		{lexer.STRING, "Bob"},            // String value
		{lexer.COMMA, ","},               // Comma separator
		{lexer.STRING, "age"},            // String key
		{lexer.COLON, ":"},               // Colon separator
		{lexer.NUMBER, "22"},             // Number value
		{lexer.RBRACE, "}"},              // End of inner object
		{lexer.COMMA, ","},               // Comma separator
		{lexer.LBRACE, "{"},              // Start of another inner object
		{lexer.STRING, "name"},           // String key
		{lexer.COLON, ":"},               // Colon separator
		{lexer.STRING, "Charlie"},        // String value
		{lexer.COMMA, ","},               // Comma separator
		{lexer.STRING, "age"},            // String key
		{lexer.COLON, ":"},               // Colon separator
		{lexer.NUMBER, "28"},             // Number value
		{lexer.RBRACE, "}"},              // End of inner object
		{lexer.RBRACKET, "]"},            // End of nested array
		{lexer.RBRACE, "}"},              // End of main object
		{lexer.EOF, ""},                  // End of input
	}

	// Initialize the lexer with the input
	lexer := lexer.New(input)

	// Iterate over the expected tokens and compare with lexer output
	for i, tt := range tests {
		tok := lexer.NextToken()

		// Log the token for debugging purposes
		t.Logf("Token[%d]: Type=%q, Literal=%q, Line=%d, Column=%d", i, tok.Type, tok.Literal, tok.Line, tok.Column)

		// Validate the token type
		if tok.Type != tt.expectedType {
			t.Fatalf(
				"tests[%d] - wrong token type. expected=%q, got=%q (literal=%q)",
				i, tt.expectedType, tok.Type, tok.Literal,
			)
		}

		// Validate the token literal
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf(
				"tests[%d] - wrong literal. expected=%q, got=%q (type=%q)",
				i, tt.expectedLiteral, tok.Literal, tok.Type,
			)
		}
	}
}
