package parser

import (
	"testing"

	"github.com/orsinium-labs/testo/internal/lexer"
)

func TestParser(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		hasError bool
		expected Node
	}{
		{
			name:     "Empty Object",
			input:    "{}",
			hasError: false,
			expected: &Object{Pairs: map[string]Node{}},
		},
		{
			name:     "Simple Key-Value Pair",
			input:    `{"key": "value"}`,
			hasError: false,
			expected: &Object{Pairs: map[string]Node{
				"key": &String{Value: "value"},
			}},
		},
		{
			name: "Nested Structure",
			input: `{
                "object": {},
                "array": [],
                "nested": {"key": ["value", 123, true, null]}
            }`,
			hasError: false,
			expected: &Object{Pairs: map[string]Node{
				"object": &Object{Pairs: map[string]Node{}},
				"array":  &Array{Elements: []Node{}},
				"nested": &Object{Pairs: map[string]Node{
					"key": &Array{Elements: []Node{
						&String{Value: "value"},
						&Number{Value: 123},
						&Boolean{Value: true},
						&Null{},
					}},
				}},
			}},
		},
		{
			name:     "Incomplete Object",
			input:    "{",
			hasError: true,
		},
		{
			name:     "Missing Value",
			input:    `{"key"}`,
			hasError: true,
		},
		{
			name:     "Trailing Comma",
			input:    `{"key": "value",}`,
			hasError: true,
		},
		{
			name:     "Invalid Value",
			input:    `{"key": undefined}`,
			hasError: true,
		},
		{
			name:     "Valid Key-Value Pair",
			input:    `{"key": "value"}`,
			hasError: false,
			expected: &Object{Pairs: map[string]Node{
				"key": &String{Value: "value"},
			}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := lexer.New(tt.input)
			p := New(l)
			node, err := p.Parse()

			if tt.hasError && err == nil {
				t.Errorf("expected error for input %q, got none", tt.input)
			}
			if !tt.hasError && err != nil {
				t.Errorf("unexpected error for input %q: %v", tt.input, err)
			}
			if !tt.hasError && err == nil {
				if !compareNodes(node, tt.expected) {
					t.Errorf("expected node %v, got %v", tt.expected, node)
				}
			}
		})
	}
}

// compareNodes is a helper function to compare two AST nodes for equality.
func compareNodes(node1, node2 Node) bool {
	switch n1 := node1.(type) {
	case *Object:
		n2, ok := node2.(*Object)
		if !ok {
			return false
		}
		if len(n1.Pairs) != len(n2.Pairs) {
			return false
		}
		for key, value1 := range n1.Pairs {
			value2, exists := n2.Pairs[key]
			if !exists || !compareNodes(value1, value2) {
				return false
			}
		}
		return true
	case *Array:
		n2, ok := node2.(*Array)
		if !ok {
			return false
		}
		if len(n1.Elements) != len(n2.Elements) {
			return false
		}
		for i, elem1 := range n1.Elements {
			if !compareNodes(elem1, n2.Elements[i]) {
				return false
			}
		}
		return true
	case *String:
		n2, ok := node2.(*String)
		return ok && n1.Value == n2.Value
	case *Number:
		n2, ok := node2.(*Number)
		return ok && n1.Value == n2.Value
	case *Boolean:
		n2, ok := node2.(*Boolean)
		return ok && n1.Value == n2.Value
	case *Null:
		_, ok := node2.(*Null)
		return ok
	default:
		return false
	}
}
