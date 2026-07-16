package formatters

import (
	"fmt"
	"sort"
	"strings"

	"code/differ"
)

const indentSize = 4

func stylish(nodes []differ.Node) string {
	return "{\n" + formatNodes(nodes, 1) + "}"
}

func formatNodes(nodes []differ.Node, depth int) string {
	var builder strings.Builder
	indent := strings.Repeat(" ", depth*indentSize-2)

	for _, node := range nodes {
		switch node.Type {
		case differ.Nested:
			fmt.Fprintf(&builder, "%s  %s: %s\n",
				indent, node.Key, formatChildren(node.Children, depth))
		case differ.Unchanged:
			fmt.Fprintf(&builder, "%s  %s: %s\n",
				indent, node.Key, formatValue(node.OldValue, depth))
		case differ.Added:
			fmt.Fprintf(&builder, "%s+ %s: %s\n",
				indent, node.Key, formatValue(node.NewValue, depth))
		case differ.Removed:
			fmt.Fprintf(&builder, "%s- %s: %s\n",
				indent, node.Key, formatValue(node.OldValue, depth))
		case differ.Updated:
			fmt.Fprintf(&builder, "%s- %s: %s\n",
				indent, node.Key, formatValue(node.OldValue, depth))
			fmt.Fprintf(&builder, "%s+ %s: %s\n",
				indent, node.Key, formatValue(node.NewValue, depth))
		}
	}

	return builder.String()
}

func formatChildren(nodes []differ.Node, depth int) string {
	closingIndent := strings.Repeat(" ", depth*indentSize)
	return "{\n" + formatNodes(nodes, depth+1) + closingIndent + "}"
}

func formatValue(value interface{}, depth int) string {
	object, ok := value.(map[string]interface{})
	if !ok {
		return toString(value)
	}

	closingIndent := strings.Repeat(" ", depth*indentSize)
	childIndent := strings.Repeat(" ", (depth+1)*indentSize)

	keys := make([]string, 0, len(object))
	for key := range object {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	var builder strings.Builder
	builder.WriteString("{\n")
	for _, key := range keys {
		fmt.Fprintf(&builder, "%s%s: %s\n",
			childIndent, key, formatValue(object[key], depth+1))
	}
	builder.WriteString(closingIndent + "}")

	return builder.String()
}

func toString(value interface{}) string {
	if value == nil {
		return "null"
	}
	return fmt.Sprintf("%v", value)
}
