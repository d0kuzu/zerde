package airtable

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

func (c *Client) ListPageRecords(table string, page, pageSize int64) ([]Record, error) {
	baseURL := fmt.Sprintf("https://api.airtable.com/v0/%s/%s", c.BaseID, table)

	var records []Record
	var offset string
	var currentPage int64 = 1

	for {
		u, err := url.Parse(baseURL)
		if err != nil {
			return nil, err
		}

		q := u.Query()
		q.Add("fields[]", "Mobile Number")
		q.Add("fields[]", "Status")
		q.Add("sort[0][field]", "Created")
		q.Add("sort[0][direction]", "desc")
		q.Add("pageSize", strconv.FormatInt(pageSize, 10))

		if offset != "" {
			q.Add("offset", offset)
		}
		u.RawQuery = q.Encode()

		req, err := http.NewRequest("GET", u.String(), nil)
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
		if err := json.Unmarshal(body, &data); err != nil {
			return nil, err
		}

		if currentPage == page {
			if len(data.Records) > int(pageSize) {
				return data.Records[:pageSize], nil
			}
			return data.Records, nil
		}

		if data.Offset == "" {
			break
		}
		offset = data.Offset
		currentPage++
	}

	return records, nil
}

func (c *Client) GetTotalPages(table string, pageSize int64) (int64, error) {
	baseURL := fmt.Sprintf("https://api.airtable.com/v0/%s/%s", c.BaseID, table)

	var offset string
	var total int64

	for {
		u, err := url.Parse(baseURL)
		if err != nil {
			return 0, err
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
			return 0, err
		}
		req.Header.Set("Authorization", "Bearer "+c.APIKey)

		resp, err := c.Client.Do(req)
		if err != nil {
			return 0, err
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return 0, err
		}

		var data ListResponse
		if err := json.Unmarshal(body, &data); err != nil {
			return 0, err
		}

		total += int64(len(data.Records))

		if data.Offset == "" {
			break
		}
		offset = data.Offset
	}

	totalPages := (total + pageSize - 1) / pageSize
	return totalPages, nil
}
