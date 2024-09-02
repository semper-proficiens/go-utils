package urlcleaner

import (
	"fmt"
	"net/url"
)

// UrlParser takes a query, an url endpoint, and an allowed query size, and cleans the url
// to make sure this can be used to construct a valid request.
func UrlParser(query, endpoint string, allowedQueryLength int) (*url.URL, error) {
	encodedQuery := url.QueryEscape(query)
	if len(encodedQuery) > allowedQueryLength {
		return nil, fmt.Errorf("encoded query exceeds the maximum length of 500 characters")
	}
	baseURL, err := url.Parse(endpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to parse base URL: %v", err)
	}
	return baseURL, nil
}
