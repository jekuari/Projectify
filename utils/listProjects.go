package utils

import (
	"fmt"
	"os"
	"strings"

	"github.com/jekuari/Projectify/constants"
)

func ListProjects() {
	configDir, err := ProjectifyConfigDir()
	if err != nil {
		panic(err)
	}
	projectsFilePath := fmt.Sprintf("%s/%s", configDir, constants.PROJECTS_FILE_NAME)

	projectsFile, err := os.OpenFile(projectsFilePath, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		fmt.Println("Error opening file", err)
	}

	defer projectsFile.Close()

	projectsFileRaw, err := os.ReadFile(projectsFilePath)
	if err != nil {
		fmt.Println("Error reading file", err)
	}
	projectsFileContent := string(projectsFileRaw)
	keyValuePairs := SplitProjectsFile(projectsFileContent)

	for key, value := range keyValuePairs {
		fmt.Println(key, value)
	}
}

func SplitProjectsFile(fileContent string) map[string]string {
	projects := make(map[string]string)
	for _, line := range strings.Split(fileContent, "\n") {
		if line != "" {
			keyValue := strings.Split(line, ",")
			projects[keyValue[0]] = keyValue[1]
		}
	}
	return projects
}
