package main

import (
	"fmt"
	"io"
	"os"

	"github.com/orsinium-labs/testo/internal/lexer"
)

func main() {
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "read stdin: %v", err)
	}
	tokens := lexer.New(string(input))
	for {
		token := tokens.NextToken()
		fmt.Printf(
			"%02d:%02d   %-10s %s\n",
			token.Line, token.Column, token.Type, token.Literal,
		)
		if token.Type == lexer.EOF {
			break
		}
	}
}
