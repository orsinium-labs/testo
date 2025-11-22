package parser

import (
	"fmt"
	"strconv"

	"github.com/orsinium-labs/testo/internal/lexer"
	"github.com/orsinium-labs/valdo/valdo"
)

func Validate(given any, expected string) error {
	validator, err := Parse(expected)
	if err != nil {
		return err
	}
	return validator.Validate(given)
}

func Parse(input string) (valdo.Validator, error) {
	return New(lexer.New(input)).Parse()
}

// Parser is responsible for parsing tokens into a structured format.
type Parser struct {
	l         *lexer.Lexer
	curToken  lexer.Token
	peekToken lexer.Token
}

// New creates a new Parser instance.
func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}
	// Initialize curToken and peekToken
	p.nextToken()
	p.nextToken()
	return p
}

// nextToken advances the parser to the next token.
func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

// Parse parses the input starting from the root and returns the root valdo.Validator.
func (p *Parser) Parse() (valdo.Validator, error) {
	validator, err := p.parseValue()
	if err != nil {
		return nil, err
	}
	p.nextToken()
	if p.curToken.Type != lexer.EOF {
		return nil, fmt.Errorf("expected EOF, found, %s", p.curToken.Type)
	}
	return validator, nil
}

// parseObject parses an object and returns an ObjectValue node.
func (p *Parser) parseObject() (valdo.Validator, error) {
	props := make([]valdo.PropertyType, 0)

	p.nextToken()

	// Handle an empty object
	if p.curToken.Type == lexer.RBRACE {
		p.nextToken()
		return valdo.O(), nil
	}

	// Parse object contents.
	for p.curToken.Type != lexer.EOF {
		key, err := p.parseKey()
		if err != nil {
			return nil, err
		}

		if p.curToken.Type != lexer.COLON {
			return nil, fmt.Errorf("expected ':', got %s at line %d, column %d", p.curToken.Type, p.curToken.Line, p.curToken.Column)
		}
		p.nextToken()

		value, err := p.parseValue()
		if err != nil {
			return nil, err
		}

		props = append(props, valdo.P(key, value))
		if p.curToken.Type == lexer.RBRACE {
			p.nextToken()
			return valdo.O(props...), nil
		}

		if p.curToken.Type != lexer.COMMA {
			return nil, fmt.Errorf("expected ',' or '}', got %s at line %d, column %d", p.curToken.Type, p.curToken.Line, p.curToken.Column)
		}
		p.nextToken()
	}

	return nil, fmt.Errorf("unexpected end of input")
}

// parseKey parses a key in an object.
func (p *Parser) parseKey() (string, error) {
	if p.curToken.Type != lexer.STRING {
		return "", fmt.Errorf("expected string key, got %s at line %d, column %d", p.curToken.Type, p.curToken.Line, p.curToken.Column)
	}
	key := p.curToken.Literal
	p.nextToken()
	return key, nil
}

// parseValue parses a value in an object or array and returns a Value node.
func (p *Parser) parseValue() (valdo.Validator, error) {
	switch p.curToken.Type {
	case lexer.STRING:
		value := valdo.Const(p.curToken.Literal)
		p.nextToken()
		return value, nil
	case lexer.NUMBER:
		// Convert the string literal to a float64
		numValue, err := strconv.ParseFloat(p.curToken.Literal, 64)
		if err != nil {
			return nil, fmt.Errorf("could not parse number: %v", err)
		}
		// TODO: support float.
		value := valdo.IntConst(int(numValue))
		p.nextToken()
		return value, nil
	case lexer.TRUE:
		value := valdo.BoolConst(true)
		p.nextToken()
		return value, nil
	case lexer.FALSE:
		value := valdo.BoolConst(false)
		p.nextToken()
		return value, nil
	case lexer.NULL:
		value := valdo.Null()
		p.nextToken()
		return value, nil
	case lexer.LBRACE:
		return p.parseObject()
	case lexer.LBRACKET:
		return p.parseArray()
	case lexer.TYPE_ANY:
		value := valdo.Any()
		p.nextToken()
		return value, nil
	case lexer.TYPE_STRING:
		value := valdo.String()
		p.nextToken()
		return value, nil
	case lexer.TYPE_INT:
		value := valdo.Int()
		p.nextToken()
		return value, nil
	case lexer.TYPE_UINT:
		value := valdo.Int(valdo.Min(0))
		p.nextToken()
		return value, nil
	case lexer.TYPE_FLOAT:
		value := valdo.Float64()
		p.nextToken()
		return value, nil
	case lexer.TYPE_BOOL:
		value := valdo.Bool()
		p.nextToken()
		return value, nil
	case lexer.TYPE_OBJECT:
		value := valdo.Map(valdo.Any())
		p.nextToken()
		return value, nil
	case lexer.TYPE_ARRAY:
		value := valdo.Array(valdo.Any())
		p.nextToken()
		return value, nil
	case lexer.TYPE_STRINGS:
		value := valdo.Array(valdo.String())
		p.nextToken()
		return value, nil
	case lexer.TYPE_INTS:
		value := valdo.Array(valdo.Int())
		p.nextToken()
		return value, nil
	case lexer.TYPE_UINTS:
		value := valdo.Array(valdo.Int(valdo.Min(0)))
		p.nextToken()
		return value, nil
	case lexer.TYPE_FLOATS:
		value := valdo.Array(valdo.Float64())
		p.nextToken()
		return value, nil
	case lexer.TYPE_BOOLS:
		value := valdo.Array(valdo.Bool())
		p.nextToken()
		return value, nil
	case lexer.TYPE_OBJECTS:
		value := valdo.Array(valdo.Map(valdo.Any()))
		p.nextToken()
		return value, nil
	default:
		return nil, fmt.Errorf("unexpected token %s at line %d, column %d", p.curToken.Type, p.curToken.Line, p.curToken.Column)
	}
}

func (p *Parser) parseArray() (valdo.Validator, error) {
	items := make([]valdo.Validator, 0)

	p.nextToken()

	// Handle an empty array.
	if p.curToken.Type == lexer.RBRACKET {
		p.nextToken()
		return valdo.T(), nil
	}

	for {
		value, err := p.parseValue()
		if err != nil {
			return nil, err
		}
		items = append(items, value)

		if p.curToken.Type == lexer.RBRACKET {
			p.nextToken()
			return valdo.T(items...), nil
		}

		if p.curToken.Type != lexer.COMMA {
			return nil, fmt.Errorf("expected ',' or ']', got %s at line %d, column %d", p.curToken.Type, p.curToken.Line, p.curToken.Column)
		}
		p.nextToken()
	}
}
