package main

// TokenType is a string label for each kind of token.
type TokenType string

const (
	// TODO (Stage 1): Define token types for each distinct thing in JSX.
	//
	// Look at this input:  <div id="myBtn">Click me</div>
	// Walk through it character by character. What groups do you see?
	//
	//   <            → ?
	//   div          → ?
	//   id           → ?
	//   =            → ?
	//   "myBtn"      → ?
	//   >            → ?
	//   Click me     → ?
	//   /            → ?
	//
	// Each group needs its own TokenType constant. Example syntax:
	//   OpenAngle TokenType = "OPEN_ANGLE"

	// EOF is pre-provided (it's used internally by the pipeline).
	EOF TokenType = "EOF"
)

// Token is a single unit of meaning produced by the lexer.
type Token struct {
	Type  TokenType
	Value string
}
