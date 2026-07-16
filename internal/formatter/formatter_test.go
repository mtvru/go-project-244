package formatter_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"code/internal/differ"
	"code/internal/formatter"
)

func sampleNodes() []*differ.DiffNode {
	return []*differ.DiffNode{
		differ.NewAddedNode("added", false),
		differ.NewDeletedNode("removed", "gone"),
		differ.NewUnchangedNode("same", "keep"),
		differ.NewChangedNode("updated", 1, 2),
	}
}

func TestFormatStylish(t *testing.T) {
	result, err := formatter.Format(sampleNodes(), "stylish")

	require.NoError(t, err)
	expected := "{\n" +
		"  + added: false\n" +
		"  - removed: gone\n" +
		"    same: keep\n" +
		"  - updated: 1\n" +
		"  + updated: 2\n" +
		"}"
	assert.Equal(t, expected, result)
}

func TestFormatPlain(t *testing.T) {
	result, err := formatter.Format(sampleNodes(), "plain")

	require.NoError(t, err)
	expected := "Property 'added' was added with value: false\n" +
		"Property 'removed' was removed\n" +
		"Property 'updated' was updated. From 1 to 2"
	assert.Equal(t, expected, result)
}

func TestFormatJSON(t *testing.T) {
	result, err := formatter.Format(sampleNodes(), "json")

	require.NoError(t, err)
	assert.Contains(t, result, `"type": "root"`)
	assert.Contains(t, result, `"type": "added"`)
	assert.Contains(t, result, `"key": "updated"`)
	assert.Contains(t, result, `"value1": 1`)
	assert.Contains(t, result, `"value2": 2`)
}

func nestedNodes() []*differ.DiffNode {
	return []*differ.DiffNode{
		differ.NewAddedNode("complex", map[string]interface{}{
			"nested": map[string]interface{}{"deep": "x"},
		}),
		differ.NewNestedNode("group", []*differ.DiffNode{
			differ.NewUnchangedNode("inner", "value"),
		}),
	}
}

func TestFormatStylishNested(t *testing.T) {
	result, err := formatter.Format(nestedNodes(), "stylish")

	require.NoError(t, err)
	expected := "{\n" +
		"  + complex: {\n" +
		"        nested: {\n" +
		"            deep: x\n" +
		"        }\n" +
		"    }\n" +
		"    group: {\n" +
		"        inner: value\n" +
		"    }\n" +
		"}"
	assert.Equal(t, expected, result)
}

func TestFormatPlainNested(t *testing.T) {
	result, err := formatter.Format(nestedNodes(), "plain")

	require.NoError(t, err)
	expected := "Property 'complex' was added with value: [complex value]"
	assert.Equal(t, expected, result)
}

func TestFormatUnknown(t *testing.T) {
	_, err := formatter.Format(sampleNodes(), "toml")
	assert.Error(t, err)
}
