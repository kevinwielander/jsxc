package main

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
