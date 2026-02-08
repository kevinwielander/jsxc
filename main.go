package main

import "fmt"

func main() {
	input := `<div><h1>Hello</h1><Button color="red">Click me</Button></div>`
	lexer := NewLexer(input)
	parser := NewParser(lexer)
	element := parser.parseElement()
	output := generate(element)
	fmt.Println(output)
}
