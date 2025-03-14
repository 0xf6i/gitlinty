package config // Update this to match your actual package name

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	// Test case 1: Valid config file
	t.Run("ValidConfigFile", func(t *testing.T) {
		// Create a temporary test config file
		tempDir, err := ioutil.TempDir("", "config-test")
		if err != nil {
			t.Fatalf("Failed to create temp directory: %v", err)
		}
		defer os.RemoveAll(tempDir)

		configPath := filepath.Join(tempDir, "test-config.json")
		configContent := `{
			"directory_path": "/Users/0x/tmp/gitlinty",
			"firstRun": false,
			"filesToSkip": [],
			"failureAllowances": {
				"gitignore": false,
				"license": false,
				"readme": false,
				"workflow": false,
				"tests": false
			}
		}`

		err = os.WriteFile(configPath, []byte(configContent), 0644)
		if err != nil {
			t.Fatalf("Failed to write test config file: %v", err)
		}

		// Test the function
		config, err := LoadConfig(configPath)

		// Check results
		if err != nil {
			t.Errorf("LoadConfig returned unexpected error: %v", err)
		}
		if config == nil {
			t.Fatal("LoadConfig returned nil config")
		}

		// Verify config values
		if config.DirectoryPath != "/Users/0x/tmp/gitlinty" {
			t.Errorf("Expected DirectoryPath to be '/Users/0x/tmp/gitlinty', got '%s'", config.DirectoryPath)
		}
		if config.FirstRun != false {
			t.Errorf("Expected FirstRun to be false, got %v", config.FirstRun)
		}
		if len(config.FilesToSkip) != 0 {
			t.Errorf("Expected FilesToSkip to be empty, got %v", config.FilesToSkip)
		}
		if config.FailureAllowances.Gitignore != false {
			t.Errorf("Expected FailureAllowances.GitIgnore to be false, got %v", config.FailureAllowances.Gitignore)
		}
	})

	// Test case 2: Non-existent config file
	t.Run("NonExistentConfigFile", func(t *testing.T) {
		config, err := LoadConfig("non-existent-config.json")

		if err == nil {
			t.Error("LoadConfig should return an error for non-existent file")
		}
		if config != nil {
			t.Error("LoadConfig should return nil config for non-existent file")
		}
	})
}
