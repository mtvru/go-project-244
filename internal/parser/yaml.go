package parser

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

func parseYAML(content []byte) (map[string]interface{}, error) {
	data := make(map[string]interface{})
	if err := yaml.Unmarshal(content, &data); err != nil {
		return nil, fmt.Errorf("failed to parse YAML: %w", err)
	}

	return data, nil
}
