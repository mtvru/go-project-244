package differ_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"code/internal/differ"
)

func TestDiffSortsAndClassifies(t *testing.T) {
	data1 := map[string]interface{}{
		"host":    "hexlet.io",
		"timeout": 50,
		"proxy":   "123.234.53.22",
	}
	data2 := map[string]interface{}{
		"host":    "hexlet.io",
		"timeout": 20,
		"verbose": true,
	}

	nodes := differ.Diff(data1, data2)

	keys := make([]string, 0, len(nodes))
	for _, node := range nodes {
		keys = append(keys, node.Key)
	}
	assert.Equal(t, []string{"host", "proxy", "timeout", "verbose"}, keys)

	byKey := map[string]*differ.DiffNode{}
	for _, node := range nodes {
		byKey[node.Key] = node
	}

	assert.Equal(t, differ.TypeUnchanged, byKey["host"].Type)
	assert.Equal(t, differ.TypeDeleted, byKey["proxy"].Type)
	assert.Equal(t, differ.TypeChanged, byKey["timeout"].Type)
	assert.Equal(t, differ.TypeAdded, byKey["verbose"].Type)
}

func TestDiffNested(t *testing.T) {
	data1 := map[string]interface{}{
		"group": map[string]interface{}{"a": 1},
	}
	data2 := map[string]interface{}{
		"group": map[string]interface{}{"a": 2},
	}

	nodes := differ.Diff(data1, data2)

	assert.Len(t, nodes, 1)
	assert.Equal(t, differ.TypeNested, nodes[0].Type)
	assert.Len(t, nodes[0].Children, 1)
	assert.Equal(t, differ.TypeChanged, nodes[0].Children[0].Type)
}
