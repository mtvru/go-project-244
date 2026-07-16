package formatter

import (
	"encoding/json"

	"code/internal/differ"
)

func toJSON(nodes []*differ.DiffNode) (string, error) {
	root := differ.NewRootNode(nodes)

	bytes, err := json.MarshalIndent(root, "", "  ")
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}
