package main

func generate(element JSXElement) string {
	result := "React.createElement("
	result += element.Tag

	if len(element.Props) > 0 {
		result += ", { "
		for i, prop := range element.Props {
			result += prop.Name + `: "` + prop.Value + `"`
			if i < len(element.Props)-1 {
				result += ", "
			}
		}
		result += " }"
	} else {
		result += ", null"
	}

	for _, child := range element.Children {
		switch c := child.(type) {
		case TextNode:
			result += `, "` + c.Value + `"`
		case JSXElement:
			result += ", " + generate(c)
		}
	}

	result += ")"
	return result
}
