package formatter

import (
	"fmt"

	"code/internal/differ"
)

const (
	OutputFormatStylish = "stylish"
	OutputFormatPlain   = "plain"
	OutputFormatJSON    = "json"
)

func Format(nodes []*differ.DiffNode, format string) (string, error) {
	switch format {
	case OutputFormatStylish:
		return stylish(nodes), nil
	case OutputFormatPlain:
		return plain(nodes), nil
	case OutputFormatJSON:
		return toJSON(nodes)
	default:
		return "", fmt.Errorf("unknown format: %q", format)
	}
}
