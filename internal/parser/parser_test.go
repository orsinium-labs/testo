package parser_test

import (
	"encoding/json"
	"testing"

	"github.com/orsinium-labs/testo/internal/parser"
)

func validate(given, expected string) error {
	var parsed any
	err := json.Unmarshal([]byte(given), &parsed)
	if err != nil {
		return err
	}
	return parser.Validate(parsed, expected)
}

func TestIdentity(t *testing.T) {
	inputs := []string{
		`true`,
		`false`,
		`null`,
		`13`,
		// `13.3`,
		`"hi"`,
		`[]`,
		`[1]`,
		`[1, 2, 3]`,
		`[1, null, false, "hello"]`,
		`{}`,
		`{"name": "aragorn"}`,
		`{"name": "aragorn", "age": 82}`,
	}
	for _, input := range inputs {
		err := validate(input, input)
		if err != nil {
			t.Fatalf("unexpected error in `%s`: %v", input, err)
		}
	}
}

func TestBadJson(t *testing.T) {
	inputs := []string{
		`True`,
		`nil`,
		// `1.2.3`,
		`!`,
		`{`,
		`[`,
		`]`,
		`}`,
		`}{`,
		`][`,
		`"`,
		`"hello`,
		`{,,}`,
		`{,"hello":""}`,
		`{"hello"}`,
		`["hello": "world"]`,
	}
	for _, input := range inputs {
		_, err := parser.Parse(input)
		if err == nil {
			t.Fatalf("expected error in `%s`", input)
		}
	}
}
