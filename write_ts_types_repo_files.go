package schemacafe

import (
	"os"
	"path/filepath"
)

func writeTSTypesRepoFiles(dir string, version string) error {
	files := map[string]string{
		"package.json": `{
  "name": "@schema.cafe/types",
  "version": "` + version + `",
  "description": "Autogenerated types from schema.cafe",
  "scripts": {
    "test": "echo \"Error: no test specified\" && exit 1"
  },
  "repository": {
    "type": "git",
    "url": "git+https://github.com/schema-cafe/ts-types.git"
  },
  "author": "",
  "license": "ISC",
  "bugs": {
    "url": "https://github.com/schema-cafe/ts-types/issues"
  },
  "homepage": "https://github.com/schema-cafe/ts-types#readme",
  "devDependencies": {
    "typescript": "^4.9.3"
  }
}`,
	}

	for path, contents := range files {
		err := os.MkdirAll(filepath.Dir(path), 0755)
		if err != nil {
			return err
		}
		err = os.WriteFile(path, []byte(contents), os.ModePerm)
		if err != nil {
			return err
		}
	}

	return nil
}
