package schemacafe

import "github.com/library-development/go-nameconv"

type Field struct {
	Name nameconv.Name `json:"name"`
	Type Type          `json:"type"`
}
