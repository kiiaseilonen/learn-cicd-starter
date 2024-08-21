package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "ApiKey my_test_key")
	got, err := GetAPIKey(headers)
	if got != "my_test_key" || err != nil {
		t.Errorf("GetAPIKey() = %s, %v; want %s, nil", got, err, "my_test_key")
	}
}
