package schemacafe

func NewClient() *Client {
	return &Client{
		APIURL: "https://api.schema.cafe",
	}
}
