package schemacafe

import (
	"encoding/json"
	"net/http"
	"path/filepath"
	"time"

	"github.com/library-development/go-auth"
)

type Service struct {
	DataDir    string      `json:"dataDir"`
	AuthClient auth.Client `json:"authClient"`
}

func (s *Service) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		http.ServeFile(w, r, filepath.Join(s.DataDir, "public", r.URL.Path))
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

		event := Event{
			Timestamp: time.Now().Unix(),
			UserID:    req.Auth.Email,
			Command:   req.Command,
			Input:     req.Input,
		}
		eventsDir := filepath.Join(s.DataDir, "public/events")
		err = WriteEvent(eventsDir, &event)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
