package main

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
