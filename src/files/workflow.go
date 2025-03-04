package files

import (
	"os"
)

func DoesWorkflowExist(repositoryPath string) {
	//workflowsPath := filepath.Join(repositoryPath, ".github", "workflows")
	info, err := os.Stat(repositoryPath)
	if err != nil {
		if os.IsNotExist(err) {
			//return false, "", errors.New("No workflow files exist.")
		}

	}

	if info.IsDir() {
		//return true, "green", nil
	}

}
