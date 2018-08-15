package bot

import (
	"net/http"
	"io"
)

type client interface {
	SendRequest(method, url string, body io.Reader) (*http.Response, error)
}

type Client struct {}

func (h *Client) SendRequest(method, url string, body io.Reader) (*http.Response, error) {
	request, err := http.NewRequest(method, url, body)

	if err != nil {
		return nil, err
	}

	client := &http.Client{}

	return client.Do(request)
}

