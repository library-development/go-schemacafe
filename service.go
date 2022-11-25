package schemacafe

import (
	"net/http"
	"path/filepath"

	"github.com/library-development/go-web"
)

type Service struct {
	DataDir string `json:"dataDir"`
}

func (s *Service) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	web.HandleCORS(w, r)
	switch r.Method {
	case "GET":
		http.ServeFile(w, r, filepath.Join(s.DataDir, r.URL.Path))
	}
}
