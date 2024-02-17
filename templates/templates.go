package templates

var BaseTmuxpYaml = `session_name: %s
windows:
- focus: 'true'
  panes:
  - focus: 'true'
    shell_command: 
  start_directory: %s
  window_name: zsh`
