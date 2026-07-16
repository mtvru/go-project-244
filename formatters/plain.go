package formatters

import (
	"fmt"
	"strings"

	"code/differ"
)

func plain(nodes []differ.Node) string {
	return strings.Join(plainLines(nodes, ""), "\n")
}

func plainLines(nodes []differ.Node, path string) []string {
	var lines []string

	for _, node := range nodes {
		property := node.Key
		if path != "" {
			property = path + "." + node.Key
		}

		switch node.Type {
		case differ.Added:
			lines = append(lines, fmt.Sprintf(
				"Property '%s' was added with value: %s",
				property, plainValue(node.NewValue)))
		case differ.Removed:
			lines = append(lines, fmt.Sprintf(
				"Property '%s' was removed", property))
		case differ.Updated:
			lines = append(lines, fmt.Sprintf(
				"Property '%s' was updated. From %s to %s",
				property, plainValue(node.OldValue), plainValue(node.NewValue)))
		case differ.Nested:
			lines = append(lines, plainLines(node.Children, property)...)
		case differ.Unchanged:
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
