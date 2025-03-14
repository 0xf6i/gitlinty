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

	scanner := bufio.NewScanner(file)
	hasGitignorePattern := false

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if strings.HasPrefix(line, "#") || len(line) == 0 {
			continue
		}

		if strings.HasPrefix(line, "!") {
			continue
		}

		if line == ".gitignore" || line == "*.gitignore" {
			hasGitignorePattern = true
		}

		ignorePatterns = append(ignorePatterns, line)

		if !strings.Contains(line, "*") && !strings.Contains(line, "?") {
			absPath := filepath.Join(clonePath, line)
			ignoreFiles = append(ignoreFiles, absPath)
		}
	}

	if !hasGitignorePattern {
		ignorePatterns = append(ignorePatterns, ".gitignore")
	}

	if err := scanner.Err(); err != nil {

		return nil, nil, err
	}
	return ignoreFiles, ignorePatterns, nil
}
