package schemacafe

import (
	"encoding/json"
)

func ParseType(s string) Type {
	t := Type{
		BaseType: Identifier{
			Name: "string",
		},
	}
	json.Unmarshal([]byte(s), &t)
	return t
}
