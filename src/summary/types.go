package summary

type Contributor struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Commits int    `json:"commits"`
}

type Contributions struct {
	TotalCommits int           `json:"totalCommits"`
	Contributors []Contributor `json:"contributors"`
}

type Repository struct {
	Author string `json:"author"`
	Name   string `json:"name"`
}

type File struct {
	Path   string
	Type   string
	Note   string
	Status string
}

type Summary struct {
	Repository    Repository    `json:"repository"`
	Passed        []File        `json:"passed"`
	Warning       []File        `json:"warning"`
	Failed        []File        `json:"failed"`
	License       bool          `json:"license"`
	Gitignore     bool          `json:"gitignore"`
	Readme        bool          `json:"readme"`
	Workflow      bool          `json:"workflow"`
	Tests         bool          `json:"tests"`
	Contributions Contributions `json:"contributions"`
}

type SummaryResult struct {
	Summary         *Summary
	CategoryResults map[string][]File
	Status          map[string]string
	Reason          map[string]string
	RedLight        bool
}
