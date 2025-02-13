package local

import "strings"

func CheckForGitIgnore(fileSlice []string) (bool, int, []string) {
	var gitIgnoreSlice []string
	gitIgnoreFileCount := 0

	for _, file := range fileSlice {
		if strings.Contains(file, ".gitignore") {
			gitIgnoreFileCount++
			gitIgnoreSlice = append(gitIgnoreSlice, file)
		}
	}
	return gitIgnoreFileCount > 0, gitIgnoreFileCount, gitIgnoreSlice
}
