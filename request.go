package schemacafe

import "lib.dev/auth"

type Request struct {
	Auth    *auth.Credentials `json:"auth"`
	Command string            `json:"command"`
	Context string            `json:"context"`
	Input   map[string]string `json:"input"`
}
