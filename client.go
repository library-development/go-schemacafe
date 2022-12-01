package schemacafe

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

type Client struct {
	APIURL string
}

func (c *Client) Get(path string) *QueryResponse {
	r := &QueryResponse{}
	resp, err := http.Get(c.APIURL + path)
	if err != nil {
		return r
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return r
	}
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return r
	}
	if strings.HasPrefix(string(b), "<pre>") {
		r.IsFolder = true
		lines := strings.Split(string(b), "\n")
		for _, line := range lines {
			if strings.HasPrefix(line, "<a href=\"") {
				name := strings.TrimPrefix(line, "<a href=\"")
				name = name[:strings.Index(name, "\"")]
				entry := FolderEntry{}
				if strings.HasSuffix(name, "/") {
					entry.Type = "folder"
					entry.Name = name[:len(name)-1]
				} else {
					entry.Type = "schema"
					entry.Name = name
				}
				r.Folder.Contents = append(r.Folder.Contents, entry)
			}
		}
	} else {
		err = json.Unmarshal(b, r.Schema)
		if err != nil {
			return r
		}
		r.IsSchema = true
	}
	return r
}
