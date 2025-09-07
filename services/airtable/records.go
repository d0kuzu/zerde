package airtable

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func (c *Client) ListRecords(table string) ([]byte, error) {
	baseURL := fmt.Sprintf("https://api.airtable.com/v0/%s/%s", c.BaseID, table)

	// Добавляем query-параметры для выбора только нужных колонок
	u, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}

	q := u.Query()
	q.Add("fields[]", "Mobile Number")
	q.Add("fields[]", "Status")
	q.Add("maxRecords", "5")
	u.RawQuery = q.Encode()

	// Формируем запрос
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+c.APIKey)

	// Выполняем
	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
