package main

import (
	"net/url"
	"strings"
)

func normalizeURL(inputURL string) (string, error) {
	parsedURL, err := url.Parse(inputURL)
	if err != nil {
		return "", err
	}

	newURL := parsedURL.Host + parsedURL.Path
	newURL = strings.ToLower(newURL)
	newURL = strings.TrimSuffix(newURL, "/")
	return newURL, nil
}
