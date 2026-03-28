package auth

import (
	"errors"
	"net/http"
	"strings"
)

func GetBearerToken(headers http.Header) (string, error) {
	// look for the Authorization header
	authHeader := headers.Get("Authorization")

	// check if header exists
	if authHeader == "" {
		return "", errors.New("authorization header not found")
	}

	// strip prefix and whitespace
	token := strings.TrimPrefix(authHeader, "Bearer ")

	// check header again in case it's not a Bearer token
	if token == authHeader {
		return "", errors.New("authorization header is not a bearer token")
	}

	return token, nil
}
