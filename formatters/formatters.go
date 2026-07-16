package formatters

import (
	"fmt"

	"code/differ"
)

func Format(nodes []differ.Node, format string) (string, error) {
	switch format {
	case "stylish":
		return stylish(nodes), nil
	case "plain":
		return plain(nodes), nil
	case "json":
		return toJSON(nodes)
	default:
		return "", fmt.Errorf("unknown format: %q", format)
	}
}
