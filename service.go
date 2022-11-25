package schemacafe

import (
	"net/http"
	"path/filepath"
)

type Service struct {
	DataDir      string `json:"dataDir"`
	AuthEndpoint string `json:"authEndpoint"`
}

func (s *Service) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		http.ServeFile(w, r, filepath.Join(s.DataDir, r.URL.Path))
	}
}
