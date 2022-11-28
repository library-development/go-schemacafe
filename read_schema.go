package schemacafe

import (
	"encoding/json"
	"os"
)

func ReadSchema(path string) (*Schema, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	schema := Schema{}
	err = json.NewDecoder(f).Decode(&schema)
	if err != nil {
		return nil, err
	}

	return &schema, nil
}
