package auth

import (
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetAPIKey(t *testing.T) {
	tests := map[string]struct {
		input         http.Header
		expectedValue string
	}{
		"simple": {input: http.Header{"Authorization": []string{"ApiKey 1234"}}, expectedValue: "ApiKey 1234"},
	}

	for test, tc := range tests {
		receivedValue, _ := GetAPIKey(tc.input)
		diff := cmp.Diff(tc.input, receivedValue)
	}
}
