package pkg

import "net/url"

// IsCorrectURL checks if the provided URL is valid.
func IsCorrectURL(urlStr string) bool {
	_, err := url.ParseRequestURI(urlStr)
	return err == nil
}
