package schemacafe

import (
	"encoding/json"
	"net/http"
	"path/filepath"
	"sync"
	"time"

	"github.com/library-development/go-auth"
)

type Service struct {
	DataDir    string      `json:"dataDir"`
	AuthClient auth.Client `json:"authClient"`
	writeLock  sync.Mutex
}

func (s *Service) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		http.ServeFile(w, r, filepath.Join(s.DataDir, "public", r.URL.Path))
	case "POST":
		req := Request{}
		json.NewDecoder(r.Body).Decode(&req)
		ok, err := s.AuthClient.VerifyToken(req.Auth)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if !ok {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		s.writeLock.Lock()
		defer s.writeLock.Unlock()

		event := Event{
			Timestamp: time.Now().UnixNano(),
			UserID:    req.Auth.Email,
			Command:   req.Command,
			Input:     req.Input,
		}

		schemasDir := filepath.Join(s.DataDir, "public/schemas")
		err = ApplyEvent(schemasDir, &event)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		eventsDir := filepath.Join(s.DataDir, "public/events")
		err = WriteEvent(eventsDir, &event)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tsDir := filepath.Join(s.DataDir, "src/main/ts-types")
		version := tsTypesVersion(tsDir)
		err = emptyGitDir(tsDir)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = WriteTypescript(Path{}, tsDir)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		newVersion, err := incrementVersion(version)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = writeTSTypesRepoFiles(tsDir, newVersion)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = gitCommitAndPush(tsDir, &event)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// TODO: publish to npm
	}
}
