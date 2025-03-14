package utils

import (
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestGenerateBase64(t *testing.T) {
	// Get two consecutive base64 strings
	first := GenerateBase64()
	time.Sleep(time.Second) // Ensure different timestamps
	second := GenerateBase64()

	// Test that we get different values
	if first == second {
		t.Errorf("GenerateBase64() generated identical strings: %v", first)
	}

	// Test that the output is valid base64
	decoded := DecodeBase64(first)
	if len(decoded) == 0 {
		t.Error("GenerateBase64() generated invalid base64 string")
	}
}

func TestDecodeBase64(t *testing.T) {
	// Test cases
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Valid timestamp base64",
			input:    "MjAyNC0wMy0yNy0xNTowNDowNQ==", // "2024-03-27-15:04:05" encoded
			expected: "2024-03-27-15:04:05",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := DecodeBase64(tt.input)
			if result != tt.expected {
				t.Errorf("DecodeBase64() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestFolderExists(t *testing.T) {
	// Create a temporary directory for testing
	tmpDir, err := os.MkdirTemp("", "test-folder")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	tests := []struct {
		name     string
		path     string
		expected bool
	}{
		{
			name:     "Existing folder",
			path:     tmpDir,
			expected: true,
		},
		{
			name:     "Non-existing folder",
			path:     filepath.Join(tmpDir, "nonexistent"),
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FolderExists(tt.path)
			if result != tt.expected {
				t.Errorf("FolderExists() = %v, want %v", result, tt.expected)
			}
		})
	}
}
