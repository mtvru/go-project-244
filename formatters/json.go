package formatters

import (
	"encoding/json"

	"code/differ"
)

func toJSON(nodes []differ.Node) (string, error) {
	result := nodesToSlice(nodes)

	bytes, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func nodesToSlice(nodes []differ.Node) []map[string]interface{} {
	result := make([]map[string]interface{}, 0, len(nodes))
	for _, node := range nodes {
		result = append(result, nodeToMap(node))
	}
	return result
}

func nodeToMap(node differ.Node) map[string]interface{} {
	item := map[string]interface{}{
		"key":  node.Key,
		"type": string(node.Type),
	}

	switch node.Type {
	case differ.Nested:
		item["children"] = nodesToSlice(node.Children)
	case differ.Added:
		item["value"] = node.NewValue
	case differ.Removed:
		item["value"] = node.OldValue
	case differ.Unchanged:
		item["value"] = node.OldValue
	case differ.Updated:
		item["oldValue"] = node.OldValue
		item["newValue"] = node.NewValue
	}

	return item
}
