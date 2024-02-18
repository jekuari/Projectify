package main

import (
	"fmt"
	"os"

	u "github.com/jekuari/Projectify/utils"
)

func main() {
	args := os.Args

	switch args[1] {
	case "c":
		u.CreateProject()
	case "l":
		u.ListProjects()
	case "d":
		which := args[2]
		u.DeleteProject(which)
	case "o":
		which := args[2]
		u.OpenProject(which)
	case "a":
		u.AddProject()
	case "h":
		u.Help()
	default:
		fmt.Println("Please specify a command, available commands are: c, l, d")
	}
}
