package auth

import (
	"errors"
	"net/http"
	"strings"
)

// Extracts the API key from the HTTP request headers and returns it.
// Eg. Auth: ApiKey {insert API Key Here}
func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Auth")
	if val == "" {
		return "", errors.New("authentication info missing")
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2 || vals[0] != "ApiKey" {
		return "", errors.New("invalid authentication info format")
	}
	return vals[1], nil
}
