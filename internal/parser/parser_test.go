package parser_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"code/internal/parser"
)

func TestParseJSON(t *testing.T) {
	content := []byte(`{"host": "hexlet.io", "follow": false}`)

	data, err := parser.Parse(content, "json")

	require.NoError(t, err)
	assert.Equal(t, "hexlet.io", data["host"])
	assert.Equal(t, false, data["follow"])
}

func TestParseYAML(t *testing.T) {
	content := []byte("host: hexlet.io\nfollow: false\n")

	data, err := parser.Parse(content, "yaml")

	require.NoError(t, err)
	assert.Equal(t, "hexlet.io", data["host"])
	assert.Equal(t, false, data["follow"])
}

func TestParseUnsupportedFormat(t *testing.T) {
	_, err := parser.Parse([]byte("{}"), "xml")
	assert.Error(t, err)
}

func TestParseInvalidContent(t *testing.T) {
	_, err := parser.Parse([]byte("{not valid json"), "json")
	assert.Error(t, err)
}
