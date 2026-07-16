package formatter

import (
	"fmt"
	"strings"

	"code/internal/differ"
)

func plain(nodes []*differ.DiffNode) string {
	return strings.Join(plainLines(nodes, ""), "\n")
}

func plainLines(nodes []*differ.DiffNode, path string) []string {
	var lines []string

	for _, node := range nodes {
		property := node.Key
		if path != "" {
			property = path + "." + node.Key
		}

		switch node.Type {
		case differ.TypeAdded:
			lines = append(lines, fmt.Sprintf(
				"Property '%s' was added with value: %s",
				property, plainValue(node.Value2)))
		case differ.TypeDeleted:
			lines = append(lines, fmt.Sprintf(
				"Property '%s' was removed", property))
		case differ.TypeChanged:
			lines = append(lines, fmt.Sprintf(
				"Property '%s' was updated. From %s to %s",
				property, plainValue(node.Value1), plainValue(node.Value2)))
		case differ.TypeNested:
			lines = append(lines, plainLines(node.Children, property)...)
		case differ.TypeUnchanged:
		}
	}

	return lines
}

func plainValue(value interface{}) string {
	switch typed := value.(type) {
	case map[string]interface{}:
		return "[complex value]"
	case string:
		return "'" + typed + "'"
	case nil:
		return "null"
	default:
		return fmt.Sprintf("%v", typed)
	}
}
