package utils

import (
	"fmt"
	"os"
	"strings"

	"github.com/jekuari/Projectify/constants"
)

func AddProject() {
	_, err := os.ReadFile(constants.TMUX_FILENAME)
	if err != nil {
		fmt.Printf("Tmuxp file not found: %s\n", err)
	}

	workingDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

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

	lastSlash := strings.LastIndex(workingDir, "/")
	folderName := workingDir[lastSlash+1:]

	if strings.Contains(folderName, ",") {
		fmt.Println("Name cannot contain a comma")
		return
	}

	fileContentRaw, err := os.ReadFile(projectsFilePath)
	if err != nil {
		fmt.Println("Error reading file", err)
	}
	fileContent := string(fileContentRaw)

	stringToFind := fmt.Sprintf(",%s\n", workingDir)
	if index := strings.Index(fileContent, stringToFind); index != -1 {
		name := ""
		for i := index - 1; i >= 0 && string(fileContent[i]) != "\n"; i-- {
			name += string(fileContent[i])
		}
		invertName := ""
		for i := len(name) - 1; i >= 0; i-- {
			invertName += string(name[i])
		}
		fmt.Printf("This project already exists with name: %s\n", invertName)
		return
	} else {
		// add to projects file
		fmt.Println("Adding to projects file", folderName, workingDir)
		if _, err = projectsFile.WriteString(fmt.Sprintf("%s%s,%s\n", fileContent, folderName, workingDir)); err != nil {
			fmt.Println("Error writing to file", err)
		}
	}
}
