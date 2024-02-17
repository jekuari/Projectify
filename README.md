# Projectify
My personal workflow relies mostly on using the terminal and working on multiple projects at once. 
This project aims to build a CLI tool that quickly launches tmuxp saved sessions.

## Pre-requisites
Sessions are saved in Tmux, an extremely powerful terminal multiplexer.
- [tmux](https://github.com/tmux/tmux/wiki)

To save the session it uses Tmuxp a very popular session manager.
- [tmuxp](https://github.com/tmux-python/tmuxp)

## Installation

TODO, for now, build the go files and put it in your binaries directory.

## Usage

### Available Commands:
- create
It creates a new project from a template, and it'll add it to .gitignore
```bash
pfy c
```
- delete
Deletes a project from the list, NOTE: it won't delete the .tmuxp.yaml file
```bash
pfy delete Your_Awesome_Project
```
- list
Lists all available projects.
```bash
pfy l
```
- open
Opens the specified project in a new terminal
```bash
pfy o Your_Awesome_Project
```
- help
Shows information about these commands
```bash
pfy h
```
