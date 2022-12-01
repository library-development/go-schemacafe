package schemacafe

import "bytes"

type Type struct {
	IsArray   bool       `json:"isArray"`
	IsMap     bool       `json:"isMap"`
	IsPointer bool       `json:"isPointer"`
	BaseType  Identifier `json:"baseType"`
}

func (t *Type) Golang(currentPackage Path) string {
	var buf bytes.Buffer

	if t.IsPointer {
		buf.WriteString("*")
	}

	if t.IsArray {
		buf.WriteString("[]")
	}

	if t.IsMap {
		buf.WriteString("map[string]")
	}

	buf.WriteString(t.BaseType.Golang(currentPackage))

	return buf.String()
}
