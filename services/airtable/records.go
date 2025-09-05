package airtable

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListRecords(table string) ([]Record, error) {
	url := fmt.Sprintf("https://api.airtable.com/v0/%s/%s", c.BaseID, table)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+c.APIKey)

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var data ListResponse
	if err = json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	return data.Records, nil
}
