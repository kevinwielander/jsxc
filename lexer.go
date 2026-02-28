package main

// Lexer tokenizes a JSX input string into a stream of Tokens.
type Lexer struct {
	input string
	pos   int
	inTag bool // true when we're inside a tag (between < and >)
}

// NewLexer creates a new lexer for the given input string.
func NewLexer(input string) *Lexer {
	return &Lexer{input: input, pos: 0}
}

// NextToken returns the next token from the input.
func (l *Lexer) NextToken() Token {
	// TODO (Stage 2): Implement the tokenizer state machine.
	//
	// Key insight: the same character means different things depending on context.
	//   'G' outside a tag → Text content ("Genetec")
	//   'G' inside a tag  → would be an Identifier (tag name or attr name)
	//   Use l.inTag to track which context we're in.
	//
	// State machine:
	//   When outside a tag (!l.inTag):
	//     '<'  → set l.inTag = true, return OpenAngle
	//     else → return l.readText()
	//
	//   When inside a tag (l.inTag):
	//     '>'  → set l.inTag = false, return CloseAngle
	//     '/'  → return Slash
	//     '='  → return Equals
	//     '"'  → return String (use l.readString())
	//     ' ', '\n', '\t' → skip whitespace (l.pos++, call NextToken again)
	//     else → return Identifier (use l.readIdentifier())
	//
	//   End of input → return Token{EOF, ""}
	panic("not implemented")
}

// readString reads a quoted string value (the part between the quotes).
// Pre-provided — you don't need to implement this.
func (l *Lexer) readString() string {
	l.pos++ // skip opening quote
	start := l.pos
	for l.pos < len(l.input) && l.input[l.pos] != '"' {
		l.pos++
	}
	result := l.input[start:l.pos]
	l.pos++ // skip closing quote
	return result
}

// readIdentifier reads a sequence of identifier characters (letters and digits).
// Pre-provided — you don't need to implement this.
func (l *Lexer) readIdentifier() string {
	start := l.pos
	for l.pos < len(l.input) && isIdentifierChar(l.input[l.pos]) {
		l.pos++
	}
	return l.input[start:l.pos]
}

// readText reads plain text content up to the next '<'.
// Pre-provided — you don't need to implement this.
func (l *Lexer) readText() string {
	start := l.pos
	for l.pos < len(l.input) && l.input[l.pos] != '<' {
		l.pos++
	}
	return l.input[start:l.pos]
}

// isIdentifierChar returns true for letters and digits.
func isIdentifierChar(ch byte) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || (ch >= '0' && ch <= '9')
}
