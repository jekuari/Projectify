package utils

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/jekuari/Projectify/constants"
)

func OpenProject(which string) {
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

	if val, ok := keyValuePairs[which]; ok {
		fmt.Println("Opening", val)
		pathOfTmuxpFile := fmt.Sprintf("%s/%s", val, constants.TMUX_FILENAME)

		cmd := exec.Command("tmuxp", "load", "--yes", "-d", pathOfTmuxpFile)
		_, err := cmd.CombinedOutput()

		if err != nil {
			fmt.Println("Error opening project", err)
			return
		}

		// attach tmux to the session
		cmd = exec.Command("kitty", "-d", val, "--title", which, "tmux", "attach-session", "-t", which)
		err = cmd.Start()

		if err != nil {
			fmt.Println("Error attaching to session", err)
			return
		}

		fmt.Println("Project opened")
	} else {
		fmt.Println("Project not found")
	}
}
