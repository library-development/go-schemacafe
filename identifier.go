package schemacafe

import (
	"github.com/library-development/go-nameconv"
)

type Identifier struct {
	Path Path          `json:"path"`
	Name nameconv.Name `json:"name"`
}

func (i *Identifier) Golang(currentPackage Path) string {
	if i.Path.Length() == 0 {
		return i.Name.PascalCase()
	}
	id := i.Path.Last().AllLowerNoSpaces()
	if i.Path.String() == currentPackage.String() {
		return i.Name.PascalCase()
	} else {
		return id + "." + i.Name.PascalCase()
	}
}
