package differ

import (
	"reflect"
	"sort"
)

type NodeType string

const (
	Added     NodeType = "added"
	Removed   NodeType = "removed"
	Unchanged NodeType = "unchanged"
	Updated   NodeType = "updated"
	Nested    NodeType = "nested"
)

type Node struct {
	Key      string
	Type     NodeType
	OldValue interface{}
	NewValue interface{}
	Children []Node
}

func BuildDiff(data1, data2 map[string]interface{}) []Node {
	keys := collectKeys(data1, data2)

	nodes := make([]Node, 0, len(keys))
	for _, key := range keys {
		value1, ok1 := data1[key]
		value2, ok2 := data2[key]

		node := Node{Key: key}

		switch {
		case !ok2:
			node.Type = Removed
			node.OldValue = value1
		case !ok1:
			node.Type = Added
			node.NewValue = value2
		case isObject(value1) && isObject(value2):
			node.Type = Nested
			node.Children = BuildDiff(
				value1.(map[string]interface{}),
				value2.(map[string]interface{}),
			)
		case reflect.DeepEqual(value1, value2):
			node.Type = Unchanged
			node.OldValue = value1
			node.NewValue = value2
		default:
			node.Type = Updated
			node.OldValue = value1
			node.NewValue = value2
		}

		nodes = append(nodes, node)
	}

	return nodes
}

func collectKeys(data1, data2 map[string]interface{}) []string {
	unique := make(map[string]struct{}, len(data1)+len(data2))
	for key := range data1 {
		unique[key] = struct{}{}
	}
	for key := range data2 {
		unique[key] = struct{}{}
	}

	keys := make([]string, 0, len(unique))
	for key := range unique {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	return keys
}

func isObject(value interface{}) bool {
	_, ok := value.(map[string]interface{})
	return ok
}
