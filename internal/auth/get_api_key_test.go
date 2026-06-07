package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	var httpHeaders = http.Header{}
	_, err := GetAPIKey(httpHeaders)
	if err != ErrNoAuthHeaderIncluded {
		t.Errorf("Expected error: %v, got: %v", ErrNoAuthHeaderIncluded, err)
	}

	httpHeaders.Set("Authorization", "")
	_, err = GetAPIKey(httpHeaders)
	if err != ErrNoAuthHeaderIncluded {
		t.Errorf("Expected error: %v, got: %v", ErrNoAuthHeaderIncluded, err)
	}

	httpHeaders.Set("Authorization", "ApiKeysome_token")
	_, err = GetAPIKey(httpHeaders)
	if err == nil {
		t.Errorf("Expected malformed authorization header error")
	}

	httpHeaders.Set("Authorization", "Brearer some_token")
	_, err = GetAPIKey(httpHeaders)
	if err == nil {
		t.Errorf("Expected malformed authorization header error")
	}

	httpHeaders.Set("Authorization", "ApiKey some_token")
	_, err = GetAPIKey(httpHeaders)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}
