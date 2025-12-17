package cmd

import (
	"asashishi-agent/conf"
	"asashishi-agent/global"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var CmdTools = CmdMap[any, any]{
	"-exit": func(cmd ...any) any {
		os.Exit(0)
		return nil
	},
	"-cls": func(cmd ...any) any {
		if conf.Env.System == conf.Windows {
			exec.Command("powershell", "-Command", Clear).Run()
		} else {
			var shell *exec.Cmd = exec.Command("clear")
			shell.Stdout = os.Stdout
			shell.Run()
		}
		return nil
	},
	"-rfile": func(cmd ...any) any {
		var (
			ok     bool
			err    error
			param  string
			dParam []string
		)
		if param, ok = cmd[0].(string); !ok {
			fmt.Println(global.GetStyledError(ExceptionAtReadFile))
			return nil
		} else if dParam = strings.Split(param, global.SpaceString); len(dParam) < 3 {
			fmt.Println(global.GetStyledError(ExceptionAtReadFile))
			return nil
		}
		if err = RenderFileToTerminal(dParam[2]); err != nil {
			fmt.Println(global.GetStyledError(ExceptionAtReadFile))
		}
		return nil
	},
}
