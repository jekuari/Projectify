package utils

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/jekuari/Projectify/constants"
	"github.com/jekuari/Projectify/templates"
)

func CreateProject() {
	workingDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	targetDir := fmt.Sprintf("%s/%s", workingDir, constants.TMUX_FILENAME)

	if _, err := os.Stat(targetDir); err == nil {
		fmt.Print("File already exists, if you continue, it will be overwritten, continue? (y/N) ")
		var response string
		fmt.Scan(&response)
		if response != "y" {
			fmt.Println("Aborted")
			return
		}
	} else {
		// create file
		_, err := os.Create(targetDir)
		if err != nil {
			panic(err)
		}
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

	// check if the project already exists
	fileContentRaw, err := os.ReadFile(projectsFilePath)
	if err != nil {
		fmt.Println("Error reading file", err)
	}
	fileContent := string(fileContentRaw)

	// check if there is a project with the same name
	if strings.Contains(fileContent, folderName) {
		fmt.Printf("A project extists with the same name: %s\n", folderName)
		fmt.Println("Enter a different name or (C) to cancel")
		var response string
		fmt.Scan(&response)
		if response == "C" {
			fmt.Println("Aborted")
			return
		}
		if strings.Contains(response, ",") {
			fmt.Println("Name cannot contain a comma")
			return
		}

		folderName = response
	}

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

	baseOrCurrentLayout := ""
	fmt.Print("Do you want to use the base layout? (y/N) ")
	var layoutResponse string
	fmt.Scan(&layoutResponse)
	if layoutResponse != "y" {
		baseOrCurrentLayout = "base"
	} else {
		baseOrCurrentLayout = "current"
	}

	if baseOrCurrentLayout == "base" {
		baseTmuxpFileStr := fmt.Sprintf(templates.BaseTmuxpYaml, folderName, workingDir)

		//  write to tmuxp file
		if err = os.WriteFile(fmt.Sprintf("%s/%s", workingDir, constants.TMUX_FILENAME), []byte(baseTmuxpFileStr), 0755); err != nil {
			fmt.Println("Error writing to file", err)
		}
	} else {
		err := exec.Command("tmuxp", "freeze", "--yes", fmt.Sprintf("%s/%s", workingDir, constants.TMUX_FILENAME)).Run()
		if err != nil {
			fmt.Println("Error saving current layout:\n", err)
		}
	}

	// check if .gitignore exists
	gitignorePath := fmt.Sprintf("%s/%s", workingDir, constants.GITIGNORE_FILENAME)
	if _, err := os.Stat(gitignorePath); err != nil {
		// finish
		return
	}

	// add ignore .tmuxp.yaml to .gitignore
	gitignoreFile, err := os.OpenFile(gitignorePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file", err)
	}

	defer gitignoreFile.Close()

	// check if gitignore file already contains the tmux file
	gitIgnoreContent, err := os.ReadFile(gitignorePath)
	if err != nil {
		fmt.Println("Error reading file", err)
	}

	if strings.Contains(string(gitIgnoreContent), constants.TMUX_FILENAME) {
		return
	}

	// add .tmuxp.yaml to .gitignore
	if _, err = gitignoreFile.WriteString(fmt.Sprintf("\n%s\n", constants.TMUX_FILENAME)); err != nil {
		fmt.Println("Error writing to file", err)
	}

	fmt.Println("Project added successfully")
}
