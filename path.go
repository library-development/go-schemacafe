package schemacafe

import (
	"strings"

	"github.com/library-development/go-nameconv"
)

type Path []nameconv.Name

func (p Path) String() string {
	var buf strings.Builder
	for _, name := range p {
		buf.WriteString("/")
		buf.WriteString(name.SnakeCase())
	}
	return buf.String()
}

func (p Path) Length() int {
	return len(p)
}

func (p Path) Last() nameconv.Name {
	return p[len(p)-1]
}
