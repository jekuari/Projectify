package constants

const (
	PROJECTS_FILE_NAME = ".projects"
	TMUX_FILENAME      = ".tmuxp.yaml"
	GITIGNORE_FILENAME = ".gitignore"
)

const HELP_STRING = `
This is tool help you manage your projects and quickly open them in tmux.
Usage:
  projectify [a,c,d,l,o,h] [project?]
Available Commands:
  create(c)              If you already have a .tmuxp.yaml file in the current directory, add it to the list of projects
  delete(d) [project]    Delete a project from the list
  list(l)                List all projects
  open(o) [project]      Open a project
  help(h)                Help about any command
`
