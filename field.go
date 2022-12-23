package schemacafe

import "lib.dev/nameconv"

type Field struct {
	Name nameconv.Name `json:"name"`
	Type Type          `json:"type"`
}
