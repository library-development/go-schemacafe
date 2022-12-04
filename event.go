package schemacafe

import (
	"strings"
	"time"
)

type Event struct {
	UserID    string            `json:"userId"`
	Timestamp int64             `json:"timestamp"`
	Command   string            `json:"command"`
	Context   string            `json:"context"`
	Input     map[string]string `json:"input"`
}

func (e *Event) CommitMessage() string {
	var msg strings.Builder
	msg.WriteString(e.Command)
	for k, v := range e.Input {
		msg.WriteString(" ")
		msg.WriteString(k)
		msg.WriteString("=")
		msg.WriteString(v)
	}
	msg.WriteString("\n\nCreated by ")
	msg.WriteString(e.UserID)
	msg.WriteString(" at ")
	msg.WriteString(time.Unix(0, e.Timestamp).Format(time.RFC3339))
	return msg.String()
}
