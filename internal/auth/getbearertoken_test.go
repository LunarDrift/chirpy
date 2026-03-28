package auth

import (
	"net/http"
	"testing"
)

func TestGetBearerToken(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "Bearer mytoken123")

	token, err := GetBearerToken(headers)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if token != "mytoken123" {
		t.Errorf("expected mytoken123, got %v", token)
	}
}

func TestGetBearerTokenMissingHeader(t *testing.T) {
	headers := http.Header{}

	_, err := GetBearerToken(headers)
	if err == nil {
		t.Error("expected error, got nil")
	}
}

func TestGetBearerTokenNonBearerToken(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "Access mytoken123")

	_, err := GetBearerToken(headers)
	if err == nil {
		t.Error("expected error, got nil")
	}
}
