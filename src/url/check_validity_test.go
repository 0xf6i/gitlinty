package url

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCheckValidity(t *testing.T) {
	// Create a mock GitHub server
	mockGitHub := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Server", "GitHub.com")
		w.WriteHeader(http.StatusOK)
	}))
	defer mockGitHub.Close()

	// Create a mock non-GitHub server
	mockNonGitHub := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer mockNonGitHub.Close()

	tests := []struct {
		name     string
		url      string
		wantBool bool
		wantErr  bool
	}{
		{
			name:     "Valid GitHub URL",
			url:      mockGitHub.URL,
			wantBool: true,
			wantErr:  false,
		},
		{
			name:     "Non-GitHub URL",
			url:      mockNonGitHub.URL,
			wantBool: false,
			wantErr:  true,
		},
		{
			name:     "Invalid URL",
			url:      "http://invalid.url",
			wantBool: false,
			wantErr:  true,
		},
		{
			name:     "Malformed URL",
			url:      "not-a-url",
			wantBool: false,
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotBool, err := CheckValidity(tt.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckValidity() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotBool != tt.wantBool {
				t.Errorf("CheckValidity() = %v, want %v", gotBool, tt.wantBool)
			}
		})
	}
}
