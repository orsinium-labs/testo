package parser

// Node represents a node in the AST.
type Node interface {
	TokenLiteral() string
}

// Object represents an object value with key-value pairs.
type Object struct {
	Pairs map[string]Node
}

func (o *Object) TokenLiteral() string {
	return "{"
}

// Array represents an array value with elements.
type Array struct {
	Elements []Node
}

func (a *Array) TokenLiteral() string {
	return "["
}

// String represents a string value.
type String struct {
	Value string
}

func (s *String) TokenLiteral() string {
	return s.Value
}

// Number represents a number value.
type Number struct {
	Value float64
}

func (n *Number) TokenLiteral() string {
	return "number"
}

// Boolean represents a boolean value.
type Boolean struct {
	Value bool
}

func (b *Boolean) TokenLiteral() string {
	return "boolean"
}

// Null represents a null value.
type Null struct{}

func (n *Null) TokenLiteral() string {
	return "null"
}
