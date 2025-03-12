package summary

import (
	"os/exec"
	"strings"
)

func ScanWithGitleaks(repoPath string) (bool, error) {
	cmd := exec.Command("gitleaks", "detect", "-s", repoPath)
	output, err := cmd.CombinedOutput()

	if err != nil {
		// Exit code 1 means leaks were found
		if strings.Contains(string(output), "leaks found") {
			return true, nil
		}
		return false, err
	}

	return false, nil
}
