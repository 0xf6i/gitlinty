package files

import (
	"bufio"
	"fmt"
	"io"
	"linty/src/summary"
	"log"
	"os"
	"path/filepath"
	"strings"
	"unicode"
)

// Track ignored files to avoid duplicate logs
var ignoredFilesMap = make(map[string]bool)

// isFileEmpty checks if a file is effectively empty (only whitespace or no content)
func isFileEmpty(filePath string) (bool, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return false, err
	}
	defer file.Close()

	// First check: Quick size check
	fileInfo, err := file.Stat()
	if err != nil {
		return false, err
	}
	if fileInfo.Size() == 0 {
		return true, nil
	}

	// Second check: Look for non-whitespace content
	reader := bufio.NewReader(file)
	hasContent := false

	// Read the file character by character
	for {
		r, _, err := reader.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			return false, err
		}

		// If we find any non-whitespace character, the file isn't empty
		if !unicode.IsSpace(r) {
			hasContent = true
			break
		}
	}

	return !hasContent, nil
}

// CheckFileContent scans a directory while respecting .gitignore rules.
func CheckFileContent(dirPath string, typeOfFile string, gitignorePaths []string, gitignorePatterns []string) ([]summary.File, error) {
	var matchedFiles []summary.File
	var ignoredPaths []string
	lowerFileType := strings.ToLower(typeOfFile)

	err := filepath.WalkDir(dirPath, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return fmt.Errorf("error occurred: %s", err)
		}

		// Skip ignored directories
		if d.IsDir() {
			absPath, err := filepath.Abs(path)
			if err != nil {
				return fmt.Errorf("error occurred: %s", err)
			}
			if isIgnored(absPath, gitignorePaths, gitignorePatterns) {
				return filepath.SkipDir
			}
			return nil
		}

		absPath, err := filepath.Abs(path)
		if err != nil {
			return fmt.Errorf("error occurred: %s", err)
		}

		lowerFileName := strings.ToLower(d.Name())
		lowerPath := strings.ToLower(path)

		// Check if the file should be ignored
		if isIgnored(absPath, gitignorePaths, gitignorePatterns) {
			ignoredPaths = append(ignoredPaths, absPath)
			return nil
		}

		switch lowerFileType {
		case "gitignore":
			if lowerFileName == ".gitignore" {
				isEmpty, err := isFileEmpty(absPath)
				if err != nil {
					log.Printf("Warning: couldn't check if file is empty: %s - %v", absPath, err)
					isEmpty = false
				}
				note := ""
				if isEmpty {
					note = "Empty file"
				}
				matchedFiles = append(matchedFiles, summary.File{Type: "gitignore", Path: absPath, Note: note})
			}

		case "readme":
			if lowerFileName == "readme.md" {
				isEmpty, err := isFileEmpty(absPath)
				if err != nil {
					log.Printf("Warning: couldn't check if file is empty: %s - %v", absPath, err)
					isEmpty = false
				}
				note := ""
				if isEmpty {
					note = "Empty file"
				}
				matchedFiles = append(matchedFiles, summary.File{Type: "readme", Path: absPath, Note: note})
			}

		case "license":
			if lowerFileName == "license" {
				isEmpty, err := isFileEmpty(absPath)
				if err != nil {
					log.Printf("Warning: couldn't check if file is empty: %s - %v", absPath, err)
					isEmpty = false
				}
				note := ""
				if isEmpty {
					note = "Empty file"
				}
				matchedFiles = append(matchedFiles, summary.File{Type: "license", Path: absPath, Note: note})
			}

		case "workflow":
			if strings.Contains(lowerPath, ".github/workflows/") &&
				(strings.HasSuffix(lowerFileName, ".yml") || strings.HasSuffix(lowerFileName, ".yaml")) {
				isEmpty, err := isFileEmpty(absPath)
				if err != nil {
					log.Printf("Warning: couldn't check if file is empty: %s - %v", absPath, err)
					isEmpty = false
				}
				note := ""
				if isEmpty {
					note = "Empty file"
				}
				matchedFiles = append(matchedFiles, summary.File{Type: "workflow", Path: absPath, Note: note})
			}

		case "test":
			// Improved test file detection
			isTestFile := false

			// Check for test files with explicit patterns
			if strings.Contains(lowerFileName, "_test.go") ||
				strings.Contains(lowerFileName, "/test") || (strings.Contains(lowerPath, "/test/") && !strings.Contains(lowerPath, "/node_modules/")) { // In a test directory
				isTestFile = true
			}

			// If no test pattern matches, check for special test frameworks
			if !isTestFile {
				if strings.Contains(lowerPath, "/tests/") ||
					strings.Contains(lowerPath, "/unittest/") ||
					strings.Contains(lowerPath, "/__tests__/") {
					isTestFile = true
				}
			}

			if isTestFile {
				isEmpty, err := isFileEmpty(absPath)
				if err != nil {
					log.Printf("Warning: couldn't check if file is empty: %s - %v", absPath, err)
					isEmpty = false
				}
				note := ""
				if isEmpty {
					note = "Empty file"
				}
				matchedFiles = append(matchedFiles, summary.File{Type: "test", Path: absPath, Note: note})
			}
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return matchedFiles, nil
}
