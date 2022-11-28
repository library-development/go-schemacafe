package schemacafe

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strconv"
)

func WriteEvent(eventsDir string, event *Event) error {
	path := filepath.Join(eventsDir, strconv.FormatInt(event.Timestamp, 10))
	d, err := json.Marshal(event)
	if err != nil {
		return err
	}
	return os.WriteFile(path, d, os.ModePerm)
}
