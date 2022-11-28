package schemacafe

type Event struct {
	UserID    string            `json:"userId"`
	Timestamp int64             `json:"timestamp"`
	Command   string            `json:"command"`
	Input     map[string]string `json:"input"`
}
