package code

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"code/differ"
	"code/formatters"
	"code/parsers"
)

func GenDiff(filepath1, filepath2, format string) (string, error) {
	data1, err := readFile(filepath1)
	if err != nil {
		return "", err
	}

	data2, err := readFile(filepath2)
	if err != nil {
		return "", err
	}

	diff := differ.BuildDiff(data1, data2)

	return formatters.Format(diff, format)
}

func readFile(path string) (map[string]interface{}, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %q: %w", path, err)
	}

	format := strings.TrimPrefix(filepath.Ext(path), ".")

	return parsers.Parse(content, format)
}
