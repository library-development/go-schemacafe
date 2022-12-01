package schemacafe

import "github.com/library-development/go-nameconv"

type FolderEntry struct {
	Name nameconv.Name `json:"name"`
	Type string        `json:"type"`
}
