package url

import (
	"reflect"
	"testing"
)

func TestHandler(t *testing.T) {
	tests := []struct {
		name    string
		url     string
		want    HandlerStruct
		wantErr bool
	}{
		{
			name: "Valid GitHub URL",
			url:  "https://github.com/username/repo",
			want: HandlerStruct{
				Author:     "username",
				Repository: "repo",
			},
			wantErr: false,
		},
		{
			name:    "Invalid URL format",
			url:     "https://invalid-url",
			want:    HandlerStruct{},
			wantErr: true,
		},
		{
			name:    "Missing repository",
			url:     "https://github.com/username",
			want:    HandlerStruct{},
			wantErr: true,
		},
		{
			name:    "Empty URL",
			url:     "",
			want:    HandlerStruct{},
			wantErr: true,
		},
		{
			name:    "Non-GitHub URL",
			url:     "https://gitlab.com/username/repo",
			want:    HandlerStruct{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Handler(tt.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("Handler() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Handler() = %v, want %v", got, tt.want)
			}
		})
	}
}
