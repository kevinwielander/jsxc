package main

import "fmt"

type TokenType string

const (
	OpenAngle  TokenType = "OPEN_ANGLE"
	CloseAngle TokenType = "CLOSE_ANGLE"
	Slash      TokenType = "SLASH"
	Equals     TokenType = "EQUALS"
	Identifier TokenType = "IDENTIFIER"
	String     TokenType = "STRING"
	Text       TokenType = "TEXT"
	EOF        TokenType = "EOF"
)

type Token struct {
	Type  TokenType
	Value string
}

type Lexer struct {
	input string
	pos   int
	inTag bool
}

func NewLexer(input string) *Lexer {
	return &Lexer{input: input, pos: 0}
}

func (l *Lexer) NextToken() Token {
	if l.pos >= len(l.input) {
		return Token{Type: EOF, Value: ""}
	}

	// Outside a tag, read everything up to '<' as text
	if !l.inTag && l.input[l.pos] != '<' {
		text := l.readText()
		return Token{Type: Text, Value: text}
	}

	ch := l.input[l.pos]

	switch ch {
	case '<':
		l.pos++
		l.inTag = true
		return Token{Type: OpenAngle, Value: "<"}
	case '>':
		l.pos++
		l.inTag = false
		return Token{Type: CloseAngle, Value: ">"}
	case '/':
		l.pos++
		return Token{Type: Slash, Value: "/"}
	case '=':
		l.pos++
		return Token{Type: Equals, Value: "="}
	case '"':
		return Token{Type: String, Value: l.readString()}
	}
	if ch == ' ' || ch == '\n' || ch == '\t' {
		l.pos++
		return l.NextToken()
	}
	if isIdentifierChar(ch) {
		value := l.readIdentifier()
		return Token{Type: Identifier, Value: value}
	}

	panic("unexpected character: " + string(ch))
}

func (l *Lexer) readString() string {
	l.pos++
	start := l.pos
	for l.pos < len(l.input) && l.input[l.pos] != '"' {
		l.pos++
	}
	result := l.input[start:l.pos]
	l.pos++
	return result
}

func isIdentifierChar(ch byte) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || (ch >= '0' && ch <= '9')
}

func (l *Lexer) readIdentifier() string {
	start := l.pos
	for l.pos < len(l.input) && isIdentifierChar(l.input[l.pos]) {
		l.pos++
	}
	return l.input[start:l.pos]
}

func (l *Lexer) readText() string {
	start := l.pos
	for l.pos < len(l.input) && l.input[l.pos] != '<' {
		l.pos++
	}
	return l.input[start:l.pos]
}

// ----- AST nodes -----

type Child interface {
	childNode()
}

type JSXElement struct {
	Tag      string
	Props    []Prop
	Children []Child
}

type Prop struct {
	Name  string
	Value string
}

type TextNode struct {
	Value string
}

func (t TextNode) childNode()   {}
func (e JSXElement) childNode() {}

type Parser struct {
	lexer   *Lexer
	current Token
}

func NewParser(lexer *Lexer) *Parser {
	p := &Parser{lexer: lexer}
	p.current = lexer.NextToken()
	return p
}

func (p *Parser) eat(expected TokenType) Token {
	token := p.current
	if token.Type != expected {
		panic("expected " + string(expected) + " but got " + string(token.Type))
	}
	p.current = p.lexer.NextToken()
	return token
}

func (p *Parser) parseElement() JSXElement {
	p.eat(OpenAngle)
	tag := p.eat(Identifier)
	props := p.parseProps()
	p.eat(CloseAngle)

	children := p.parseChildren()

	p.eat(OpenAngle)
	p.eat(Slash)
	p.eat(Identifier)
	p.eat(CloseAngle)

	return JSXElement{
		Tag:      tag.Value,
		Props:    props,
		Children: children,
	}
}

func (p *Parser) parseProps() []Prop {
	var props []Prop
	for p.current.Type != CloseAngle {
		name := p.eat(Identifier).Value
		p.eat(Equals)
		value := p.eat(String).Value
		props = append(props, Prop{Name: name, Value: value})
	}
	return props
}

func (p *Parser) parseChildren() []Child {
	var children []Child
	for {
		if p.current.Type == EOF {
			break
		}
		if p.current.Type == OpenAngle {
			if p.lexer.input[p.lexer.pos] == '/' {
				break
			}
			children = append(children, p.parseElement())
		}
		if p.current.Type == Text {
			children = append(children, TextNode{Value: p.current.Value})
			p.current = p.lexer.NextToken()
		}
	}
	return children
}

// ----- Generation -----

func generate(element JSXElement) string {
	result := "React.createElement("
	result += element.Tag

	if len(element.Props) > 0 {
		result += ", { "
		for i, prop := range element.Props {
			result += prop.Name + `: "` + prop.Value + `"`
			if i < len(element.Props)-1 {
				result += ", "
			}
		}
		result += " }"
	} else {
		result += ", null"
	}

	for _, child := range element.Children {
		switch c := child.(type) {
		case TextNode:
			result += `, "` + c.Value + `"`
		case JSXElement:
			result += ", " + generate(c)
		}
	}

	result += ")"
	return result
}

// ----- Main -----

func main() {
	//input := `</=>`
	//input := `</Button="red">`
	//input := `<Button color="red">`
	input := `<Button color="red">Click me</Button>`
	lexer := NewLexer(input)

	for {
		token := lexer.NextToken()
		if token.Type == EOF {
			break
		}
		println("Token:", token.Type, "Value:", token.Value)
	}

	// input = `<div>`
	// lexer = NewLexer(input)
	// parser := NewParser(lexer)
	// element := parser.parseElement()
	// fmt.Println("Tag:", element.Tag)
	// fmt.Println("=========")

	// input = `<Button color="red">`
	// lexer = NewLexer(input)
	// parser = NewParser(lexer)
	// element = parser.parseElement()
	// fmt.Println("Tag:", element.Tag)
	// for _, prop := range element.Props {
	// 	fmt.Println("Prop:", prop.Name, "=", prop.Value)
	// }
	// fmt.Println("=========")
	// input = `<Button color="red">Click me</Button>`
	// lexer = NewLexer(input)
	// parser := NewParser(lexer)
	// element := parser.parseElement()
	// fmt.Println("Tag:", element.Tag)
	// for _, prop := range element.Props {
	// 	fmt.Println("Prop:", prop.Name, "=", prop.Value)
	// }
	// for _, child := range element.Children {
	// 	switch c := child.(type) {
	// 	case TextNode:
	// 		fmt.Println("Text:", c.Value)
	// 	}
	// }

	input = `<div><h1>Hello</h1><Button color="red">Click me</Button></div>`
	lexer = NewLexer(input)
	parser := NewParser(lexer)
	element := parser.parseElement()
	output := generate(element)
	fmt.Println(output)

}
