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
	if p.Length() < 1 {
		return nil
	}
	return p[len(p)-1]
}

func (p Path) Append(name nameconv.Name) Path {
	return append(p, name)
}

func (p Path) SecondLast() nameconv.Name {
	if p.Length() < 2 {
		return nil
	}
	return p[len(p)-2]
}

func (p Path) Pop() Path {
	return p[:len(p)-1]
}
