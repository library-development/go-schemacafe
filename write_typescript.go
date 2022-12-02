package schemacafe

import (
	"bytes"
	"fmt"
	"os"
)

func WriteTypescript(from Path, to string) error {
	client := NewClient()
	fmt.Printf("Pulling from %s%s...", client.APIURL, from.String())
	r := client.Get(from)

	if r.IsFolder {
		fmt.Println(" folder found")
		fmt.Printf("Making sure directory %s exists...", to)
		err := os.MkdirAll(to, 0755)
		if err != nil {
			return err
		}
		fmt.Println(" done")
		for _, entry := range r.Folder.Contents {
			err := WriteTypescript(from.Append(entry.Name), to+"/"+entry.Name.SnakeCase())
			if err != nil {
				return err
			}
		}
	}

	if r.IsSchema {
		fmt.Println(" schema found")
		if len(from) == 0 {
			return fmt.Errorf("no schema at root")
		}

		name := from.Last()

		var typeDef bytes.Buffer
		typeDef.WriteString("export default interface I")
		typeDef.WriteString(name.PascalCase())
		typeDef.WriteString(" {\n")

		imports := map[string]string{}
		for _, field := range r.Schema.Fields {
			if field.Type.BaseType.Path.Length() > 0 {
				importPath := "@schema.cafe/types" + field.Type.BaseType.Path.String() + "/" + field.Type.BaseType.Name.SnakeCase()
				name := field.Type.BaseType.Typescript()
				if n, ok := imports[importPath]; ok && n != name {
					// TODO: handle this
					return fmt.Errorf("import name conflict: %s and %s both import %s", name, n, importPath)
				}
				imports[name] = importPath
			}

			typeDef.WriteString("  ")
			typeDef.WriteString(field.Name.CamelCase())
			typeDef.WriteString(": ")
			typeDef.WriteString(field.Type.Typescript())
			typeDef.WriteString(";\n")
		}

		typeDef.WriteString("}\n")

		var buf bytes.Buffer
		if len(imports) > 0 {
			for name, path := range imports {
				buf.WriteString("import ")
				buf.WriteString(name)
				buf.WriteString(" from \"")
				buf.WriteString(path)
				buf.WriteString("\";\n")
			}
			buf.WriteString("\n")
		}

		buf.WriteString(typeDef.String())

		path := to + ".ts"
		fmt.Printf("Writing TypeScript file to %s...", path)
		err := os.WriteFile(path, buf.Bytes(), os.ModePerm)
		if err != nil {
			return err
		}
		fmt.Println(" done")
	}

	if !r.IsFolder && !r.IsSchema {
		fmt.Println(" not found")
	}

	return nil
}
