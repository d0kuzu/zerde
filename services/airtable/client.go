package airtable

import "net/http"

type Client struct {
	APIKey string
	BaseID string
	Client *http.Client
}

func NewClient(apiKey, baseID string) *Client {
	return &Client{
		APIKey: apiKey,
		BaseID: baseID,
		Client: &http.Client{},
	}
}
