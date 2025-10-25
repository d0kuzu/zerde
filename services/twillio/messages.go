package twilio

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"
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

func (c *Client) fetchChat(from, to string, limit int) ([]Message, error) {
	from = strings.ReplaceAll(from, "-", "")

	if len(from) > 12 {
		from = "+" + "1" + from[3:]
	}

	incoming, err := c.fetchMessages(from, to, limit)
	if err != nil {
		return nil, err
	}

	outgoing, err := c.fetchMessages(to, from, limit)
	if err != nil {
		return nil, err
	}

	all := append(incoming, outgoing...)

	return all, nil
}

func (c *Client) GetConversation(clientNumber, botNumber string, limit int) ([]Message, error) {
	all, err := c.fetchChat(clientNumber, botNumber, limit)
	if err != nil {
		return nil, err
	}

	sort.Slice(all, func(i, j int) bool {
		layoutZ := time.RFC1123Z // с часовым смещением +0000
		layoutG := time.RFC1123  // с GMT

		ti, err := time.Parse(layoutZ, all[i].DateCreated)
		if err != nil {
			ti, _ = time.Parse(layoutG, all[i].DateCreated)
		}

		tj, err := time.Parse(layoutZ, all[j].DateCreated)
		if err != nil {
			tj, _ = time.Parse(layoutG, all[j].DateCreated)
		}

		return ti.Before(tj)
	})

	if len(all) > limit {
		all = all[len(all)-limit:]
	}

	return all, nil
}

func (c *Client) GetMessagesCounter(clientNumber, botNumber string, limit int) (int, error) {
	all, err := c.fetchChat(clientNumber, botNumber, limit)
	if err != nil {
		return 0, err
	}

	if len(all) > limit {
		all = all[len(all)-limit:]
	}

	return len(all), nil
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
