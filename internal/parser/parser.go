package parser

import "fmt"

const (
	FormatJSON = "json"
	FormatYAML = "yaml"
)

func Parse(content []byte, format string) (map[string]interface{}, error) {
	switch format {
	case FormatJSON:
		return parseJSON(content)
	case FormatYAML:
		return parseYAML(content)
	default:
		return nil, fmt.Errorf("unsupported format: %q", format)
	}
}
