package schemacafe

import (
	"fmt"
	"strconv"
	"strings"
)

func incrementVersion(v string) (string, error) {
	versionNumber, err := strconv.Atoi(v[strings.LastIndex(v, "."):])
	if err != nil {
		return "", fmt.Errorf("failed to parse version number: %w", err)
	}
	return fmt.Sprintf("0.0.%d", versionNumber+1), nil
}
