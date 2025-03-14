package input

import (
	"testing"
)

func TestCheckUrl(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		wantBool bool
		wantErr  bool
	}{
		{
			name:     "Valid GitHub URL",
			path:     "https://github.com/username/repo",
			wantBool: true,
			wantErr:  false,
		},
		{
			name:     "Invalid GitHub URL",
			path:     "https://github.com/invalid/notexist",
			wantBool: false,
			wantErr:  true,
		},
		{
			name:     "Non-GitHub URL",
			path:     "https://example.com",
			wantBool: false,
			wantErr:  true,
		},
		{
			name:     "Invalid path",
			path:     "/nonexistent/path",
			wantBool: false,
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotBool, err := CheckUrl(tt.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckUrl() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotBool != tt.wantBool {
				t.Errorf("CheckUrl() = %v, want %v", gotBool, tt.wantBool)
			}
		})
	}
}
