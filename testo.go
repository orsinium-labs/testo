package testo

import (
	"encoding/json"

	"github.com/orsinium-labs/testo/internal/parser"
	"github.com/orsinium-labs/valdo/valdo"
)

// Validate that the given JSON message matches the expected pattern.
func ValidateJSON[T []byte | string](given T, expected string) error {
	var parsed any
	err := json.Unmarshal([]byte(given), &parsed)
	if err != nil {
		return err
	}
	return parser.Validate(parsed, expected)
}

// Validate that the given Go value matches the expected pattern.
func Validate(given any, expected string) error {
	return parser.Validate(given, expected)
}

// Convert the pattern to a [valdo.Validator].
func Parse(input string) (valdo.Validator, error) {
	return parser.Parse(input)
}
