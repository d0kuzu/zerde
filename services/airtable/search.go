package airtable

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func (c *Client) SearchChat(table, number string) ([]Record, error) {
	baseURL := fmt.Sprintf("https://api.airtable.com/v0/%s/%s", c.BaseID, table)

	var offset string
	var records []Record

	for {
		u, err := url.Parse(baseURL)
		if err != nil {
			return []Record{}, err
		}

		q := u.Query()
		q.Add("fields[]", "Mobile Number")
		q.Add("fields[]", "Status")
		q.Add("pageSize", "100")

		if offset != "" {
			q.Add("offset", offset)
		}
		u.RawQuery = q.Encode()

		req, err := http.NewRequest("GET", u.String(), nil)
		if err != nil {
			return []Record{}, err
		}
		req.Header.Set("Authorization", "Bearer "+c.APIKey)

		resp, err := c.Client.Do(req)
		if err != nil {
			return []Record{}, err
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return []Record{}, err
		}

		var data ListResponse
		if err := json.Unmarshal(body, &data); err != nil {
			return []Record{}, err
		}

		for _, record := range data.Records {
			if strings.Contains(record.Fields.MobileNumber, number) {
				records = append(records, record)
			}
		}

		if data.Offset == "" {
			break
		}
		offset = data.Offset
	}

	return records, nil
}
