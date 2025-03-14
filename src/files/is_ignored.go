package files

import (
	"log"
	"path/filepath"
	"strings"
)

var ignoredFilesMap = make(map[string]bool)

func IsIgnored(filePath string, gitignorePaths []string, gitignorePatterns []string) bool {
	baseFileName := filepath.Base(filePath)

	if baseFileName == ".gitignore" {
		absFilePath, err := filepath.Abs(filePath)
		if err != nil {
			log.Fatal(err)
		}

		for _, ignoredPath := range gitignorePaths {
			ignoreAbs, err := filepath.Abs(ignoredPath)
			if err != nil {
				log.Fatal(err)
			}

			if filepath.Base(filePath) == ".gitignore" {
				return false
			}

			absFilePath = filepath.Clean(absFilePath)
			ignoreAbs = filepath.Clean(ignoreAbs)

			relPath, err := filepath.Rel(ignoreAbs, absFilePath)
			if err == nil && !strings.HasPrefix(relPath, "..") {
				if !ignoredFilesMap[filePath] {
					ignoredFilesMap[filePath] = true
				}
				return true
			}
		}
		return false
	}

	for _, ignoredPath := range gitignorePaths {

		if filePath == ignoredPath {
			if !ignoredFilesMap[filePath] {
				ignoredFilesMap[filePath] = true

			}
			return true
		}
	}

	for _, dir := range []string{"env", "/node*modules/", "/.git/"} {

		if strings.Contains(filePath, dir) {
			if !ignoredFilesMap[filePath] {
				ignoredFilesMap[filePath] = true
			}
			return true
		}
	}

	for _, pattern := range gitignorePatterns {

		if strings.HasSuffix(pattern, "/*") {
			dirPattern := pattern[:len(pattern)-2]
			if strings.Contains(filepath.Dir(filePath), dirPattern) {
				if !ignoredFilesMap[filePath] {
					ignoredFilesMap[filePath] = true
				}
				return true
			}
		}

		if strings.HasSuffix(pattern, "/") {
			dirPattern := pattern[:len(pattern)-1]
			if strings.Contains(filePath, "/"+dirPattern+"/") {
				if !ignoredFilesMap[filePath] {
					ignoredFilesMap[filePath] = true
				}
				return true
			}
		}

		if strings.HasPrefix(pattern, "*.") {
			ext := pattern[1:] // Get the extension part
			if strings.HasSuffix(baseFileName, ext) {
				if !ignoredFilesMap[filePath] {
					ignoredFilesMap[filePath] = true
				}
				return true
			}
		}

		matched, err := filepath.Match(pattern, baseFileName)
		if err == nil && matched {
			if !ignoredFilesMap[filePath] {
				ignoredFilesMap[filePath] = true
			}
			return true
		}
	}
	return false
}
