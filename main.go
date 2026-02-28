package main

import (
	"fmt"
	"os"
)

func main() {
	input := `<div><h2>Genetec</h2><button id="myBtn">Click me</button></div>`

	switch os.Getenv("STAGE") {

	case "2":
		// ── Verify Stage 2: Lexer ──────────────────────────────────────
		// Run with: STAGE=2 go run .
		lexer := NewLexer(input)
		fmt.Println("=== Token Stream ===")
		for {
			tok := lexer.NextToken()
			fmt.Printf("%-15s %q\n", tok.Type, tok.Value)
			if tok.Type == EOF {
				break
			}
		}

	case "3":
		// ── Verify Stage 3: Parser ─────────────────────────────────────
		// Run with: STAGE=3 go run .
		lexer := NewLexer(input)
		parser := NewParser(lexer)
		element := parser.parseElement()
		fmt.Printf("=== AST ===\n%+v\n", element)

	default:
		// ── Stage 4: Full pipeline → output.html ───────────────────────
		// Run with: go run .
		lexer := NewLexer(input)
		parser := NewParser(lexer)
		element := parser.parseElement()
		output := generate(element)

		html := `<!DOCTYPE html>
<html>
<head>
    <title>JSX Compiler Output</title>
</head>
<body>
    <div id="root"></div>
    <script src="https://unpkg.com/react@18/umd/react.development.js"></script>
    <script src="https://unpkg.com/react-dom@18/umd/react-dom.development.js"></script>
    <script>
        const element = ` + output + `;
        const root = ReactDOM.createRoot(document.getElementById("root"));
        root.render(element);
    </script>
</body>
</html>`

		os.WriteFile("output.html", []byte(html), 0644)
		fmt.Println("✓ output.html written — open it in your browser!")
	}
}
