package schemacafe

import "github.com/library-development/go-auth"

type Request struct {
	Command string            `json:"command"`
	Auth    *auth.Credentials `json:"auth"`
	Input   map[string]string `json:"args"`
}
