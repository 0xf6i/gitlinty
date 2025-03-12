package config

import (
	"encoding/json"
	"errors"
	"os"
)

type FailureAllowances struct {
	Gitignore bool `json:"gitignore"`
	License   bool `json:"license"`
	Readme    bool `json:"readme"`
	Workflow  bool `json:"workflow"`
	Tests     bool `json:"tests"`
}

type Config struct {
	DirectoryPath     string            `json:"directory_path"`
	FirstRun          bool              `json:"firstRun"`
	FilesToSkip       []interface{}     `json:"filesToSkip"`
	FailureAllowances FailureAllowances `json:"failureAllowances"`
}

// https://yetanotherprogrammingblog.medium.com/using-json-config-in-go-67e824ca46cc
func LoadConfig(configPath string) (*Config, error) {

	config := Config{}

	configFile, err := os.Open(configPath)
	if err != nil {
		return nil, errors.New("failed to open config file")
	}
	defer configFile.Close()

	jsonParse := json.NewDecoder(configFile)
	jsonParse.Decode(&config)

	return &config, nil
}
func WriteConfig(config *Config, configPath string) (*Config, error) {
	newConfig := Config{
		DirectoryPath:     config.DirectoryPath,
		FirstRun:          config.FirstRun,
		FilesToSkip:       config.FilesToSkip,
		FailureAllowances: config.FailureAllowances,
	}

	jsonWriter, err := json.MarshalIndent(newConfig, "", "  ")
	if err != nil {
		return nil, errors.New("could not convert struct to json object")
	}

	err = os.WriteFile(configPath, jsonWriter, os.ModePerm)
	if err != nil {
		return nil, errors.New("failed to write the new config ")
	}

	return &newConfig, err
}
