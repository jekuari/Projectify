package utils

import (
	"fmt"
	"os"
	"strings"

	"github.com/jekuari/Projectify/constants"
)

func DeleteProject(which string) {
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
	dir, ok := keyValuePairs[which]
	if !ok {
		fmt.Println("Project not found")
		return
	}

	deleteString := fmt.Sprintf("%s,%s\n", which, dir)
	projectsFileContent = strings.Replace(projectsFileContent, deleteString, "", -1)
	err = os.WriteFile(projectsFilePath, []byte(projectsFileContent), 0755)
	if err != nil {
		fmt.Println("Error deleting project", err)
	}

	fmt.Println("Project deleted")
}
