package schemacafe

type Identifier struct {
	Path Path `json:"path"`
}

func (i *Identifier) Golang(currentPackage Path) string {
	if i.Path.Length() == 0 {
		panic("identifier has no path")
	}
	id := i.Path.Last().AllLowerNoSpaces()
	if i.Path.Pop().String() == currentPackage.String() {
		return i.Path.Last().PascalCase()
	} else {
		return id + "." + i.Path.Last().PascalCase()
	}
}

func (i *Identifier) Typescript() string {
	return "I" + i.Path.Last().PascalCase()
}
