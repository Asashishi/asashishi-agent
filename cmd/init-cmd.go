package cmd

import (
	"os"
	"os/exec"
)

var CmdTools map[string]func() = map[string]func(){
	"exit": func() {
		os.Exit(0)
	},
	"cls": func() {
		exec.Command("powershell", "-Command", Clear).Run()
	},
}
