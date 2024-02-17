package models

type Config struct {
	ProjectsFileDir   string `json:"projectsFile"`
	TmuxFilename      string `json:"tmuxpFilename"`
	BaseTmuxpFilename string `json:"baseTmuxpFilename"`
}
