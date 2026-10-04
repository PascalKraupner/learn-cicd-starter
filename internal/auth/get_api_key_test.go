package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "ApiKey test-key")

	apiKey, err := GetAPIKey(headers)
	if err != nil {
		t.Fatalf("GetAPIKey() unexpected error: %v", err)
	}
	if apiKey != "test-key" {
		t.Errorf("GetAPIKey() = %q, want %q", apiKey, "test-key")
	}
}

func TestGetAPIKeyWithoutAuthorizationHeader(t *testing.T) {
	apiKey, err := GetAPIKey(http.Header{})

	if !errors.Is(err, ErrNoAuthHeaderIncluded) {
		t.Fatalf("GetAPIKey() error = %v, want %v", err, ErrNoAuthHeaderIncluded)
	}
	if apiKey != "" {
		t.Errorf("GetAPIKey() = %q, want an empty API key", apiKey)
	}
}
