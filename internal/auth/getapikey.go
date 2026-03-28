package auth

import (
	"errors"
	"net/http"
	"strings"
)

func GetAPIKey(headers http.Header) (string, error) {
	// look for the Authorization header
	authHeader := headers.Get("Authorization")

	// check if header exists
	if authHeader == "" {
		return "", errors.New("authorization header not found")
	}

	// strip prefix and whitespace
	apiKey := strings.TrimPrefix(authHeader, "ApiKey ")

	// check header again in case it's not an ApiKey
	if apiKey == authHeader {
		return "", errors.New("authorization header is not an ApiKey")
	}

	return apiKey, nil
}
