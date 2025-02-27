package utils

import (
	"encoding/json"
	"os"
)

type Config struct {
	DirectoryPath string `json:"directory_path"`
	FirstRun      bool   `json:"firstRun"`
}

// load config file into memory
func LoadConfig(configDirectory string) (*Config, error) {
	config := &Config{}
	file, err := os.ReadFile(configDirectory + "/config.json")
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(file, config)
	if err != nil {
		return nil, err
	}

	return config, nil

}

func SaveConfig(filename string, config *Config) error {
	configData, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(filename, configData, 0777)
}
