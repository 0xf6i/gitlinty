package utils

import (
	"fmt"
	"linty/input"
	"log"
	"os"
)

// CheckIfGitlintyExists checks if the directory exists and performs initial setup if needed
func CloneFolderExists(path string) (bool, string) {
	configFileName := "config.json"

	// Load existing config (if any)
	config, err := LoadConfig(configFileName)
	if err != nil {
		log.Fatal("Error loading config:", err)
	}

	// If no config exists, create a default one
	if config == nil {
		config = &Config{
			DirectoryPath: path,
			FirstRun:      true,
		}
	}

	// If firstTime is 1, ask setup questions
	if config.FirstRun == 1 {
		fmt.Println("No gitlinty folder available.")

		u, err := input.UserChoice("Would you like to use the default route?")
		if err != nil {
			log.Fatal(err)
		}

		var userClonePath string
		if u {
			userClonePath = "/tmp/gitlinty"
		} else {
			userClonePath, err = input.UserInput("Please specify the path where you want the gitlinty folder to be stored: ")
			if err != nil {
				log.Fatal(err)
			}
		}

		// Create the directory
		err = os.MkdirAll(userClonePath, 0755)
		if err != nil {
			log.Fatal(err)
		}

		// Update config
		config.DirectoryPath = userClonePath
		config.FirstTime = 0 // Mark setup as complete

		// Save config
		err = SaveConfig(configFileName, config)
		if err != nil {
			log.Fatal("Error saving config:", err)
		}

		fmt.Printf("Setup complete. Directory: %s, Config saved to: %s\n", userClonePath, configFileName)
		return true, userClonePath
	}

	// Ensure the directory exists
	if _, err := os.Stat(config.DirectoryPath); os.IsNotExist(err) {
		log.Fatalf("Config path %s does not exist. Please check your config.", config.DirectoryPath)
	}

	return true, config.DirectoryPath
}
