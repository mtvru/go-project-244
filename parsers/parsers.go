package parsers

import (
	"encoding/json"
	"fmt"

	"gopkg.in/yaml.v3"
)

func Parse(content []byte, format string) (map[string]interface{}, error) {
	data := make(map[string]interface{})

	switch format {
	case "json":
		if err := json.Unmarshal(content, &data); err != nil {
			return nil, fmt.Errorf("failed to parse JSON: %w", err)
		}
	case "yaml", "yml":
		if err := yaml.Unmarshal(content, &data); err != nil {
			return nil, fmt.Errorf("failed to parse YAML: %w", err)
		}
	default:
		return nil, fmt.Errorf("unsupported format: %q", format)
	}

	return data, nil
}
