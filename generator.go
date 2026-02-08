package main

import "strings"

func generate(element JSXElement) string {
	var result strings.Builder
	result.WriteString("React.createElement(")
	result.WriteString(element.Tag)

	if len(element.Props) > 0 {
		result.WriteString(", { ")
		for i, prop := range element.Props {
			result.WriteString(prop.Name + `: "` + prop.Value + `"`)
			if i < len(element.Props)-1 {
				result.WriteString(", ")
			}
		}
		result.WriteString(" }")
	} else {
		result.WriteString(", null")
	}

	for _, child := range element.Children {
		switch c := child.(type) {
		case TextNode:
			result.WriteString(`, "` + c.Value + `"`)
		case JSXElement:
			result.WriteString(", " + generate(c))
		}
	}

	result.WriteString(")")
	return result.String()
}
