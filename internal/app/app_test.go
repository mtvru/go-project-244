package app_test

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"code/internal/app"
)

func fixture(name string) string {
	return filepath.Join("..", "..", "testdata", "fixtures", name)
}

func TestRunStylish(t *testing.T) {
	result, err := app.Run(fixture("file1.json"), fixture("file2.json"), "stylish")

	require.NoError(t, err)
	assert.Contains(t, result, "- follow: false")
	assert.Contains(t, result, "+ timeout: 20")
}

func TestRunDifferentFormats(t *testing.T) {
	_, err := app.Run(fixture("file1.json"), fixture("file2.yml"), "stylish")
	assert.Error(t, err)
}

func TestRunMissingFile(t *testing.T) {
	_, err := app.Run(fixture("does_not_exist.json"), fixture("file2.json"), "stylish")
	assert.Error(t, err)
}

func TestRunUnknownOutputFormat(t *testing.T) {
	_, err := app.Run(fixture("file1.json"), fixture("file2.json"), "unknown")
	assert.Error(t, err)
}
