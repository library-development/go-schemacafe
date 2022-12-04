package schemacafe

import (
	"io/fs"
	"os"
)

func emptyGitDir(dir string) error {
	return fs.WalkDir(os.DirFS(dir), ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.Name() == ".git" {
			return fs.SkipDir
		}
		return os.Remove(path)
	})
}
