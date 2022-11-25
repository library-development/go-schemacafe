package schemacafe

import (
	"encoding/json"
	"net/http"
	"path/filepath"

	"github.com/library-development/go-auth"
)

type Service struct {
	DataDir    string      `json:"dataDir"`
	AuthClient auth.Client `json:"authClient"`
}

func (s *Service) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		http.ServeFile(w, r, filepath.Join(s.DataDir, r.URL.Path))
	case "POST":
		req := Request{}
		json.NewDecoder(r.Body).Decode(&req)
		ok, err := s.AuthClient.Validate(req.Auth)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if !ok {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		switch req.Command {
		// TODO
		}
	}
}
