package main

import "os"

func main() {
	input := `<div><h2>Genetec</h2><button id="myBtn">Click me</button></div>`
	lexer := NewLexer(input)
	parser := NewParser(lexer)
	element := parser.parseElement()
	output := generate(element)

	html := `<!DOCTYPE html>
<html>
<head>
    <title>JSX Compiler Test</title>
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
}
