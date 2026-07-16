package app

import (
	"fmt"
	"os"
	"path/filepath"

	"code/internal/differ"
	"code/internal/formatter"
	"code/internal/parser"
)

func Run(filePath1, filePath2, format string) (string, error) {
	format1 := getFileFormat(filePath1)
	format2 := getFileFormat(filePath2)

	if format1 != format2 {
		return "", fmt.Errorf("cannot compare files of different formats: %s vs %s", format1, format2)
	}

	data1, err := readAndParse(filePath1, format1)
	if err != nil {
		return "", fmt.Errorf("file1: %w", err)
	}

	data2, err := readAndParse(filePath2, format2)
	if err != nil {
		return "", fmt.Errorf("file2: %w", err)
	}

	diff := differ.Diff(data1, data2)

	return formatter.Format(diff, format)
}

func readAndParse(filePath, format string) (map[string]interface{}, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %q: %w", filePath, err)
	}

	return parser.Parse(content, format)
}

func getFileFormat(filePath string) string {
	switch filepath.Ext(filePath) {
	case ".json":
		return parser.FormatJSON
	case ".yaml", ".yml":
		return parser.FormatYAML
	default:
		return ""
	}
}
