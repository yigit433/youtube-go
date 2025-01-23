package httpclient

import (
    "net/http"
    "time"
)

type Client struct {
    httpClient *http.Client
}

func New(timeout time.Duration) *Client {
    return &Client{
        httpClient: &http.Client{
            Timeout: timeout,
        },
    }
}

func (c *Client) Get(url string) (*http.Response, error) {
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return nil, err
    }

    req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; rv:78.0) Gecko/20100101 Firefox/78.0")
    return c.httpClient.Do(req)
}