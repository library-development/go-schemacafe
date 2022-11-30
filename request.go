package schemacafe

import "github.com/library-development/go-auth"

type Request struct {
	Auth    *auth.Credentials `json:"auth"`
	Command string            `json:"command"`
	Input   map[string]string `json:"input"`
}
