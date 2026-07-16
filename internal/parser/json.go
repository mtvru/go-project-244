package parser

import (
	"encoding/json"
	"fmt"
)

func parseJSON(content []byte) (map[string]interface{}, error) {
	data := make(map[string]interface{})
	if err := json.Unmarshal(content, &data); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %w", err)
	}

	return data, nil
}
