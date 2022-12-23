package schemacafe

import (
	"encoding/json"

	"lib.dev/nameconv"
)

func ParseType(s string) Type {
	t := Type{
		BaseType: Identifier{
			Path: []nameconv.Name{{"string"}},
		},
	}
	json.Unmarshal([]byte(s), &t)
	return t
}
