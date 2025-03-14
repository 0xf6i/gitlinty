package repository

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCheckRepository(t *testing.T) {
	// Create a temporary directory structure
	tmpDir, err := os.MkdirTemp("", "test-repo")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	// Create a fake .git directory
	gitDir := filepath.Join(tmpDir, ".git")
	if err := os.Mkdir(gitDir, 0755); err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name    string
		path    string
		wantErr bool
	}{
		{
			name:    "Valid git repository",
			path:    tmpDir,
			wantErr: false,
		},
		{
			name:    "Non-existent directory",
			path:    filepath.Join(tmpDir, "nonexistent"),
			wantErr: true,
		},
		{
			name:    "Directory without .git",
			path:    os.TempDir(),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPath, err := CheckRepository(tt.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckRepository() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && gotPath != filepath.Join(tt.path, ".git") {
				t.Errorf("CheckRepository() = %v, want %v", gotPath, filepath.Join(tt.path, ".git"))
			}
		})
	}
}

func TestGetRepoInfo(t *testing.T) {
	// Create a temporary directory structure
	tmpDir, err := os.MkdirTemp("", "test-repo")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	tests := []struct {
		name         string
		path         string
		wantAuthor   string
		wantRepoName string
		wantErr      bool
	}{
		{
			name:         "Local repository",
			path:         tmpDir,
			wantAuthor:   "unknown-author",
			wantRepoName: filepath.Base(tmpDir),
			wantErr:      true, // Will error because it's not a real git repo
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotAuthor, gotRepo, err := GetRepoInfo(tt.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetRepoInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if gotAuthor != tt.wantAuthor {
					t.Errorf("GetRepoInfo() gotAuthor = %v, want %v", gotAuthor, tt.wantAuthor)
				}
				if gotRepo != tt.wantRepoName {
					t.Errorf("GetRepoInfo() gotRepo = %v, want %v", gotRepo, tt.wantRepoName)
				}
			}
		})
	}
}
