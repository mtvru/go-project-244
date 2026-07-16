package parsers_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"code/parsers"
)

func TestParseJSON(t *testing.T) {
	content := []byte(`{"host": "hexlet.io", "follow": false}`)

	data, err := parsers.Parse(content, "json")

	require.NoError(t, err)
	assert.Equal(t, "hexlet.io", data["host"])
	assert.Equal(t, false, data["follow"])
}

func TestParseYAML(t *testing.T) {
	content := []byte("host: hexlet.io\nfollow: false\n")

	for _, format := range []string{"yaml", "yml"} {
		data, err := parsers.Parse(content, format)

		require.NoError(t, err)
		assert.Equal(t, "hexlet.io", data["host"])
		assert.Equal(t, false, data["follow"])
	}
}

func TestParseUnsupportedFormat(t *testing.T) {
	_, err := parsers.Parse([]byte("{}"), "xml")
	assert.Error(t, err)
}

func TestParseInvalidContent(t *testing.T) {
	_, err := parsers.Parse([]byte("{not valid json"), "json")
	assert.Error(t, err)
}
