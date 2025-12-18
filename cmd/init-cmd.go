package cmd

import (
	"asashishi-agent/conf"
	"asashishi-agent/global"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var CmdTools = map[string]func(...string) error{
	"-exit": func(cmd ...string) error {
		os.Exit(0)
		return nil
	},
	"-cls": func(cmd ...string) error {
		if conf.Env.System == conf.Windows {
			exec.Command("powershell", "-Command", Clear).Run()
		} else {
			var shell *exec.Cmd = exec.Command("clear")
			shell.Stdout = os.Stdout
			shell.Run()
		}
		return nil
	},
	"-rfile": func(cmd ...string) error {
		var (
			err    error
			dParam []string
		)
		if cmd[0] == "" {
			fmt.Println(global.GetStyledWarn(ExceptionAtReadFile))
			return nil
		} else if dParam = strings.Split(cmd[0], global.SpaceString); len(dParam) < 3 {
			fmt.Println(global.GetStyledWarn(ExceptionAtReadFile))
			return nil
		}
		if err = RenderFileToTerminal(dParam[2]); err != nil {
			fmt.Println(global.GetStyledWarn(ExceptionAtReadFile))
		}
		return nil
	},
}
