package twilio

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sort"
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
