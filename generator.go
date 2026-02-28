package main

import "strings"

func generate(element JSXElement) string {
	// TODO (Stage 4): Walk the AST and emit React.createElement(...) calls.
	//
	// Structure:
	//   React.createElement(tag, props, ...children)
	//
	// Rules:
	//   - Lowercase tag  → pass as string:     "div"
	//   - Uppercase tag  → pass as identifier:  MyComponent
	//   - No props       → pass null
	//   - Props          → pass as JS object:   { name: "value" }
	//   - Children       → additional args, use a type switch on Child
	//
	// Hint: use strings.Builder for efficient string building.
	// Hint: call generate(child) recursively for JSXElement children.
	_ = strings.Builder{} // remove this line when you start implementing
	panic("not implemented")
}

// isHTMLTag returns true if the tag starts with a lowercase letter.
// Lowercase = HTML element (pass as string), Uppercase = React component (pass as identifier).
func isHTMLTag(tag string) bool {
	return tag[0] >= 'a' && tag[0] <= 'z'
}
