package code_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"code"
)

func readFixture(t *testing.T, name string) string {
	t.Helper()

	content, err := os.ReadFile(filepath.Join("testdata", "fixtures", name))
	require.NoError(t, err)

	return strings.TrimSuffix(string(content), "\n")
}

func fixture(name string) string {
	return filepath.Join("testdata", "fixtures", name)
}

func TestGenDiffStylish(t *testing.T) {
	expected := readFixture(t, "expected_stylish.txt")

	testCases := []struct {
		name  string
		file1 string
		file2 string
	}{
		{"json", "nested1.json", "nested2.json"},
		{"yaml", "nested1.yml", "nested2.yml"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := code.GenDiff(fixture(tc.file1), fixture(tc.file2), "stylish")
			require.NoError(t, err)
			assert.Equal(t, expected, actual)
		})
	}
}

func TestGenDiffStylishIsDefault(t *testing.T) {
	expected := readFixture(t, "expected_stylish.txt")

	actual, err := code.GenDiff(fixture("nested1.json"), fixture("nested2.json"), "stylish")
	require.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func TestGenDiffFlat(t *testing.T) {
	expected := readFixture(t, "expected_flat.txt")

	testCases := []struct {
		name  string
		file1 string
		file2 string
	}{
		{"json", "file1.json", "file2.json"},
		{"yaml", "file1.yml", "file2.yml"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := code.GenDiff(fixture(tc.file1), fixture(tc.file2), "stylish")
			require.NoError(t, err)
			assert.Equal(t, expected, actual)
		})
	}
}

func TestGenDiffPlain(t *testing.T) {
	expected := readFixture(t, "expected_plain.txt")

	testCases := []struct {
		name  string
		file1 string
		file2 string
	}{
		{"json", "nested1.json", "nested2.json"},
		{"yaml", "nested1.yml", "nested2.yml"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := code.GenDiff(fixture(tc.file1), fixture(tc.file2), "plain")
			require.NoError(t, err)
			assert.Equal(t, expected, actual)
		})
	}
}

func TestGenDiffJSON(t *testing.T) {
	expected := readFixture(t, "expected_json.txt")

	actual, err := code.GenDiff(fixture("nested1.json"), fixture("nested2.json"), "json")
	require.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func TestGenDiffErrors(t *testing.T) {
	t.Run("missing file", func(t *testing.T) {
		_, err := code.GenDiff(fixture("does_not_exist.json"), fixture("file2.json"), "stylish")
		assert.Error(t, err)
	})

	t.Run("unknown format", func(t *testing.T) {
		_, err := code.GenDiff(fixture("file1.json"), fixture("file2.json"), "unknown")
		assert.Error(t, err)
	})
}
