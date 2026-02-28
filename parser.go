package main

// Child is the interface for anything that can appear inside a JSX element.
// Both JSXElement and TextNode satisfy this interface.
// (In TypeScript terms: type Child = JSXElement | TextNode)
type Child interface {
	childNode()
}

// JSXElement represents a JSX tag: <tag props>children</tag>
type JSXElement struct {
	Tag      string
	Props    []Prop
	Children []Child
}

// Prop represents a single attribute: name="value"
type Prop struct {
	Name  string
	Value string
}

// TextNode represents plain text content between tags.
type TextNode struct {
	Value string
}

// Marker methods — these make JSXElement and TextNode satisfy the Child interface.
func (t TextNode) childNode()   {}
func (e JSXElement) childNode() {}

// Parser holds the lexer and one token of lookahead.
type Parser struct {
	lexer   *Lexer
	current Token
}

// NewParser creates a parser and reads the first token.
func NewParser(lexer *Lexer) *Parser {
	p := &Parser{lexer: lexer}
	p.current = lexer.NextToken()
	return p
}

// eat consumes the current token (asserting it matches the expected type) and
// advances to the next token. This is the canonical function in recursive descent parsers.
func (p *Parser) eat(expected TokenType) Token {
	token := p.current
	if token.Type != expected {
		panic("expected " + string(expected) + " but got " + string(token.Type))
	}
	p.current = p.lexer.NextToken()
	return token
}

// parseElement parses a complete JSX element: <tag props>children</tag>
func (p *Parser) parseElement() JSXElement {
	// TODO (Stage 3): Parse an element.
	//
	// Sequence of tokens for <div id="x">Hello</div>:
	//   OpenAngle  Identifier("div")  Identifier("id") Equals String("x")  CloseAngle
	//   Text("Hello")
	//   OpenAngle  Slash  Identifier("div")  CloseAngle
	//
	// Steps:
	//   1. eat(OpenAngle)
	//   2. eat(Identifier) → tag name
	//   3. parseProps()
	//   4. eat(CloseAngle)
	//   5. parseChildren()
	//   6. eat(OpenAngle), eat(Slash), eat(Identifier), eat(CloseAngle)  ← closing tag
	//   7. return JSXElement{...}
	panic("not implemented")
}

// parseProps parses zero or more name="value" attribute pairs.
func (p *Parser) parseProps() []Prop {
	// TODO (Stage 3): Parse attributes until you hit CloseAngle (or Slash for self-closing).
	//
	// Each prop is: Identifier Equals String
	// Stop when: p.current.Type == CloseAngle
	panic("not implemented")
}

// parseChildren parses the children of an element until the closing tag is detected.
func (p *Parser) parseChildren() []Child {
	// TODO (Stage 3): Parse children until you see "</" (OpenAngle followed by Slash in input).
	//
	// A child is either:
	//   - A TextNode  (p.current.Type == Text)
	//   - A nested JSXElement  (p.current.Type == OpenAngle, and next char is NOT '/')
	//
	// Hint to detect closing tag: check p.current.Type == OpenAngle &&
	//   p.lexer.input[p.lexer.pos] == '/'
	//
	// Hint: this function calls parseElement() for nested elements. That recursion
	//   is what makes this a recursive descent parser.
	panic("not implemented")
}
