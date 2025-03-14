package files

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCheckFileContent(t *testing.T) {
	// Create a temporary directory structure
	tmpDir, err := os.MkdirTemp("", "test-files")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	// Create test files
	files := map[string]string{
		".gitignore": "*.log\nnode_modules/",
		"README.md":  "# Test Project",
		"LICENSE":    "MIT License",
	}

	for name, content := range files {
		err := os.WriteFile(filepath.Join(tmpDir, name), []byte(content), 0644)
		if err != nil {
			t.Fatal(err)
		}
	}

	tests := []struct {
		name            string
		fileType        string
		wantNumFiles    int
		ignoredPaths    []string
		ignoredPatterns []string
	}{
		{
			name:         "Find gitignore",
			fileType:     "gitignore",
			wantNumFiles: 1,
		},
		{
			name:         "Find readme",
			fileType:     "readme",
			wantNumFiles: 1,
		},
		{
			name:         "Find license",
			fileType:     "license",
			wantNumFiles: 1,
		},
		{
			name:         "Find workflow (none)",
			fileType:     "workflow",
			wantNumFiles: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CheckFileContent(tmpDir, tt.fileType, tt.ignoredPaths, tt.ignoredPatterns)
			if err != nil {
				t.Errorf("CheckFileContent() error = %v", err)
				return
			}
			if len(got) != tt.wantNumFiles {
				t.Errorf("CheckFileContent() got %v files, want %v", len(got), tt.wantNumFiles)
			}
		})
	}
}

func TestIsIgnored(t *testing.T) {
	tests := []struct {
		name            string
		filePath        string
		ignoredPaths    []string
		ignoredPatterns []string
		want            bool
	}{
		{
			name:            "Gitignore file should never be ignored",
			filePath:        "/path/to/.gitignore",
			ignoredPaths:    []string{"/path/to/.gitignore"},
			ignoredPatterns: []string{},
			want:            false,
		},
		{
			name:            "Ignored path",
			filePath:        "/path/to/node_modules/file.js",
			ignoredPaths:    []string{},
			ignoredPatterns: []string{"node_modules/"},
			want:            true,
		},
		{
			name:            "Not ignored path",
			filePath:        "/path/to/src/file.go",
			ignoredPaths:    []string{},
			ignoredPatterns: []string{"node_modules/"},
			want:            false,
		},
		{
			name:            "Ignored by extension",
			filePath:        "file.log",
			ignoredPaths:    []string{},
			ignoredPatterns: []string{"*.log"},
			want:            true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsIgnored(tt.filePath, tt.ignoredPaths, tt.ignoredPatterns)
			if got != tt.want {
				t.Errorf("isIgnored() = %v, want %v", got, tt.want)
			}
		})
	}
}
