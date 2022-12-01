package schemacafe

import (
	"github.com/library-development/go-nameconv"
)

type Identifier struct {
	Path Path          `json:"path"`
	Name nameconv.Name `json:"name"`
}

func (i *Identifier) Golang() string {
	if i.Path.Length() == 0 {
		return i.Name.PascalCase()
	}
	return i.Path.Last().AllLowerNoSpaces() + "." + i.Name.PascalCase()
}
