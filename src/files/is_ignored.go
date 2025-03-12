package files

import (
	"path/filepath"
	"strings"
)

func isIgnored(filePath string, gitignorePaths []string, gitignorePatterns []string) bool {
	baseFileName := filepath.Base(filePath)
	// Always allow .gitignore to be processed
	if baseFileName == ".gitignore" {
		return false
	}
	// Check if the absolute path is in ignored paths
	for _, ignoredPath := range gitignorePaths {
		if filePath == ignoredPath {
			if !ignoredFilesMap[filePath] {
				ignoredFilesMap[filePath] = true
				// fmt.Println("[IGNORED] Absolute Path:", filePath)
			}
			return true
		}
	}
	// Check for directories that should be completely ignored
	for _, dir := range []string{"/myenv/", "/node*modules/", "/.git/"} {
		if strings.Contains(filePath, dir) {
			if !ignoredFilesMap[filePath] {
				ignoredFilesMap[filePath] = true
				// fmt.Println("[IGNORED] Directory Match:", filePath)
			}
			return true
		}
	}
	// Check if path matches any pattern
	for _, pattern := range gitignorePatterns {
		// Handle patterns ending with /* (directory contents)
		if strings.HasSuffix(pattern, "/*") {
			dirPattern := pattern[:len(pattern)-2]
			if strings.Contains(filepath.Dir(filePath), dirPattern) {
				if !ignoredFilesMap[filePath] {
					ignoredFilesMap[filePath] = true
					// fmt.Println("[IGNORED] Directory Content Match:", filePath, "-> Pattern:", pattern)
				}
				return true
			}
		}
		// Handle patterns ending with / (directories)
		if strings.HasSuffix(pattern, "/") {
			dirPattern := pattern[:len(pattern)-1]
			if strings.Contains(filePath, "/"+dirPattern+"/") {
				if !ignoredFilesMap[filePath] {
					ignoredFilesMap[filePath] = true
					// fmt.Println("[IGNORED] Directory Match:", filePath, "-> Pattern:", pattern)
				}
				return true
			}
		}
		// Handle *.extension patterns
		if strings.HasPrefix(pattern, "*.") {
			ext := pattern[1:] // Get the extension part
			if strings.HasSuffix(baseFileName, ext) {
				if !ignoredFilesMap[filePath] {
					ignoredFilesMap[filePath] = true
					// fmt.Println("[IGNORED] Extension Match:", filePath, "-> Pattern:", pattern)
				}
				return true
			}
		}
		// Try standard pattern matching on the basename
		matched, err := filepath.Match(pattern, baseFileName)
		if err == nil && matched {
			if !ignoredFilesMap[filePath] {
				ignoredFilesMap[filePath] = true
				// fmt.Println("[IGNORED] Pattern Match:", filePath, "-> Pattern:", pattern)
			}
			return true
		}
	}
	return false
}
