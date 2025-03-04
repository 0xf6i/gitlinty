package config

type FailureAllowances struct {
	Gitignore bool `json:"gitignore"`
	License   bool `json:"license"`
	Readme    bool `json:"readme"`
	Workflow  bool `json:"workflow"`
}
