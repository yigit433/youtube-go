package utils

import (
    "net/url"
    "fmt"
)

func BuildURLWithQueryParams(baseURL string, params map[string]string) (string, error) {
    parsedURL, err := url.Parse(baseURL)
    if err != nil {
        return "", fmt.Errorf("invalid URL: %v", err)
    }

    query := parsedURL.Query()

    for key, value := range params {
        query.Set(key, value)
    }

    parsedURL.RawQuery = query.Encode()

    return parsedURL.String(), nil
}