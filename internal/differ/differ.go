package differ

import (
	"reflect"
	"sort"
)

type NodeType string

const (
	TypeAdded     NodeType = "added"
	TypeDeleted   NodeType = "deleted"
	TypeChanged   NodeType = "changed"
	TypeUnchanged NodeType = "unchanged"
	TypeNested    NodeType = "nested"
	TypeRoot      NodeType = "root"
)

type DiffNode struct {
	Key      string      `json:"key"`
	Type     NodeType    `json:"type"`
	Value1   interface{} `json:"value1,omitempty"`
	Value2   interface{} `json:"value2,omitempty"`
	Children []*DiffNode `json:"children,omitempty"`
}

func NewAddedNode(key string, value interface{}) *DiffNode {
	return &DiffNode{Key: key, Type: TypeAdded, Value2: value}
}

func NewDeletedNode(key string, value interface{}) *DiffNode {
	return &DiffNode{Key: key, Type: TypeDeleted, Value1: value}
}

func NewChangedNode(key string, oldValue, newValue interface{}) *DiffNode {
	return &DiffNode{Key: key, Type: TypeChanged, Value1: oldValue, Value2: newValue}
}

func NewUnchangedNode(key string, value interface{}) *DiffNode {
	return &DiffNode{Key: key, Type: TypeUnchanged, Value1: value}
}

func NewNestedNode(key string, children []*DiffNode) *DiffNode {
	return &DiffNode{Key: key, Type: TypeNested, Children: children}
}

func NewRootNode(children []*DiffNode) *DiffNode {
	return &DiffNode{Type: TypeRoot, Children: children}
}

func Diff(data1, data2 map[string]interface{}) []*DiffNode {
	keys := collectKeys(data1, data2)

	nodes := make([]*DiffNode, 0, len(keys))
	for _, key := range keys {
		value1, ok1 := data1[key]
		value2, ok2 := data2[key]

		nodes = append(nodes, buildNode(key, value1, ok1, value2, ok2))
	}

	return nodes
}

func buildNode(key string, value1 interface{}, ok1 bool, value2 interface{}, ok2 bool) *DiffNode {
	switch {
	case !ok2:
		return NewDeletedNode(key, value1)
	case !ok1:
		return NewAddedNode(key, value2)
	case isObject(value1) && isObject(value2):
		return NewNestedNode(key, Diff(
			value1.(map[string]interface{}),
			value2.(map[string]interface{}),
		))
	case reflect.DeepEqual(value1, value2):
		return NewUnchangedNode(key, value1)
	default:
		return NewChangedNode(key, value1, value2)
	}
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
