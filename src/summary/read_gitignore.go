package summary

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
)

func ReadGitignore(clonePath string) ([]string, []string, error) {
	var ignoreFiles []string
	var ignorePatterns []string

	file, err := os.Open(filepath.Join(clonePath, ".gitignore"))
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	// Read all lines from .gitignore
	scanner := bufio.NewScanner(file)
	hasGitignorePattern := false

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// Skip comments and empty lines
		if strings.HasPrefix(line, "#") || len(line) == 0 {
			continue
		}

		// Skip negation patterns
		if strings.HasPrefix(line, "!") {
			continue
		}

		// Check if .gitignore is mentioned
		if line == ".gitignore" || line == "*.gitignore" {
			hasGitignorePattern = true
		}

		// Add the pattern
		ignorePatterns = append(ignorePatterns, line)

		// If it's a specific path (no wildcards), add to ignoreFiles
		if !strings.Contains(line, "*") && !strings.Contains(line, "?") {
			absPath := filepath.Join(clonePath, line)
			ignoreFiles = append(ignoreFiles, absPath)
		}
	}

	// If .gitignore pattern wasn't found, add it
	if !hasGitignorePattern {
		ignorePatterns = append(ignorePatterns, ".gitignore")
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	return ignoreFiles, ignorePatterns, nil
}
