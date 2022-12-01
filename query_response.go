package schemacafe

type QueryResponse struct {
	IsFolder bool    `json:"isFolder"`
	Folder   *Folder `json:"folder"`
	IsSchema bool    `json:"isSchema"`
	Schema   *Schema `json:"schema"`
}
