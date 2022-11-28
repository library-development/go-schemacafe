package schemacafe

import (
	"encoding/json"
	"os"
)

func WriteSchema(path string, schema *Schema) error {
	b, err := json.MarshalIndent(schema, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, b, os.ModePerm)
}
