package schemacafe

import "lib.dev/nameconv"

type FolderEntry struct {
	Name nameconv.Name `json:"name"`
	Type string        `json:"type"`
}
