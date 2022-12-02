package schemacafe

import (
	"encoding/json"
	"os"
)

func tsTypesVersion(tsDir string) string {
	f, err := os.Open(tsDir)
	if err != nil {
		return "0.0.1"
	}
	defer f.Close()
	var pkg struct {
		Version string `json:"version"`
	}
	err = json.NewDecoder(f).Decode(&pkg)
	if err != nil {
		return "0.0.1"
	}
	return pkg.Version
}
