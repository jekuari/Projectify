package utils

import (
	"fmt"
	"os"
)

func ProjectifyConfigDir() (string, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	projectifyDir := fmt.Sprintf("%s/projectify", configDir)
	if _, err := os.Stat(projectifyDir); os.IsNotExist(err) {
		if err := os.Mkdir(projectifyDir, 0755); err != nil {
			return "", fmt.Errorf("Error creating projectify directory: %s", err)
		}
	}

	return projectifyDir, nil
}
