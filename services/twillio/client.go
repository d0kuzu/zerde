package twilio

import (
	"encoding/base64"
	"net/http"
)

type Client struct {
	AccountSID string
	AuthToken  string
	Client     *http.Client
}

func NewClient(accountSID, authToken string) *Client {
	return &Client{
		AccountSID: accountSID,
		AuthToken:  authToken,
		Client:     &http.Client{},
	}
}

func (c *Client) authHeader() string {
	auth := c.AccountSID + ":" + c.AuthToken
	return "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))
}
