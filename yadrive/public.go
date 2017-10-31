package yadrive

import (
	"fmt"
	"net/http"

	"github.com/google/go-querystring/query"
)

// DownloadPublicLink get download link of a public resource by its public key or link.
// Please, see https://tech.yandex.com/disk/api/reference/public-docpage/#download
func (c *Client) DownloadPublicLink(key string, path *string) (string, error) {
	qs := struct {
		Key  string  `url:"public_key"`
		Path *string `url:"path"`
	}{}

	qs.Key = key
	if path != nil {
		qs.Path = path
	}

	values, err := query.Values(qs)
	if err != nil {
		return "", err
	}

	u := fmt.Sprintf("/v1/disk/public/resources/download?%s", values.Encode())
	req, err := c.request("GET", u, nil)
	if err != nil {
		return "", err
	}

	data := struct {
		Href      string `json:"href"`
		Method    string `json:"method"`
		Templated bool   `json:"templated"`

		Message     string `json:"message"`
		Description string `json:"description"`
		Error       string `json:"error"`
	}{}

	resp, err := c.do(req, &data)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("%s: %s", resp.Status, data.Message)
	}

	return data.Href, nil
}
