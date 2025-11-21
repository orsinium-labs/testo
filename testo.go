package testo

import (
	"encoding/json"
	"fmt"
	"io"
	"testing"

	"github.com/orsinium-labs/testo/internal/parser"
	"github.com/orsinium-labs/valdo/valdo"
)

// Fail tests if the given input doesn't match the expected pattern.
//
// The input can be one of:
//
//   - io.Reader returning JSON. For example, http.Response.Body.
//   - string containing JSON.
//   - []byte containing JSON.
//   - an arbitrary object that can be validated with [valdo].
func Assert(t *testing.T, given any, expected string) {
	t.Helper()
	parsed, err := readInput(given)
	if err != nil {
		t.Fatalf("failed to read input: %v", err)
	}
	err = parser.Validate(parsed, expected)
	if err != nil {
		givenJSONBytes, marshalErr := json.MarshalIndent(parsed, "", "  ")
		var givenJSONStr string
		if marshalErr != nil {
			givenJSONStr = fmt.Sprintf("input cannot be serialized: %v", err)
		} else {
			givenJSONStr = string(givenJSONBytes)
		}
		t.Fatalf(
			"validation error: %v\n\ninput:\n%s\n\nschema:\n%s",
			err, givenJSONStr, expected,
		)
	}
}

func readInput(raw any) (any, error) {
	switch typed := raw.(type) {
	case io.Reader:
		rawAll, err := io.ReadAll(typed)
		if err != nil {
			return nil, err
		}
		var parsed any
		err = json.Unmarshal(rawAll, &parsed)
		return parsed, err
	case string:
		var parsed any
		err := json.Unmarshal([]byte(typed), &parsed)
		return parsed, err
	case []byte:
		var parsed any
		err := json.Unmarshal(typed, &parsed)
		return parsed, err
	default:
		return raw, nil
	}
}

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
