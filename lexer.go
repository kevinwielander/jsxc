package main

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
