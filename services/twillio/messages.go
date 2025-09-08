package twilio

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sort"
	"strings"
)

func (c *Client) fetchMessages(from, to string, limit int) ([]Message, error) {
	url := fmt.Sprintf(
		"https://api.twilio.com/2010-04-01/Accounts/%s/Messages.json?From=%s&To=%s&PageSize=%d",
		c.AccountSID, from, to, limit,
	)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", c.authHeader())

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var data MessagesResponse
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	return data.Messages, nil
}

func (c *Client) GetConversation(clientNumber, botNumber string, limit int) ([]Message, error) {
	clientNumber = strings.ReplaceAll(clientNumber, "-", "")

	if len(clientNumber) > 12 {
		clientNumber = "+" + "1" + clientNumber[3:]
	}

	incoming, err := c.fetchMessages(clientNumber, botNumber, limit)
	if err != nil {
		return nil, err
	}

	outgoing, err := c.fetchMessages(botNumber, clientNumber, limit)
	if err != nil {
		return nil, err
	}

	all := append(incoming, outgoing...)

	sort.Slice(all, func(i, j int) bool {
		return all[i].DateCreated < all[j].DateCreated
	})

	if len(all) > limit {
		all = all[len(all)-limit:]
	}

	return all, nil
}

func (c *Client) SendMessage(from, to, body string) (*Message, error) {
	from = strings.ReplaceAll(from, "-", "")

	if len(from) > 12 {
		from = "+" + "1" + from[3:]
	}

	data := url.Values{}
	data.Set("From", from)
	data.Set("To", to)
	data.Set("Body", body)

	req, err := http.NewRequest("POST", fmt.Sprintf(
		"https://api.twilio.com/2010-04-01/Accounts/%s/Messages.json",
		c.AccountSID,
	), strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", c.authHeader())
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyResp, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var msg Message
	if err := json.Unmarshal(bodyResp, &msg); err != nil {
		return nil, err
	}

	return &msg, nil
}
