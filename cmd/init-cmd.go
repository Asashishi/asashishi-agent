package cmd

import (
	"asashishi-agent/tools"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/alecthomas/chroma/v2/lexers"
	"github.com/alecthomas/chroma/v2/quick"
)

var CmdTools = CmdMap[any, any]{
	"exit": func(cmd ...any) any {
		os.Exit(0)
		return nil
	},
	"cls": func(cmd ...any) any {
		exec.Command("powershell", "-Command", Clear).Run()
		return nil
	},
	"rfile": func(cmd ...any) any {
		var (
			ok       bool
			err      error
			path     string
			param    string
			content  string
			fileType string
			dParam   []string
		)
		if param, ok = cmd[0].(string); !ok {
			fmt.Println(ExceptionAtReadFile)
			return nil
		} else if dParam = strings.Split(param, " "); len(dParam) < 2 {
			fmt.Println(ExceptionAtReadFile)
			return nil
		}
		path = dParam[1]
		if fileType = lexers.Match(path).Config().Name; fileType == "" {
			fileType = "plaintext"
		}
		content = tools.ReadFileContent(path)
		if err = quick.Highlight(
			os.Stdout,
			content,
			fileType,
			"terminal16m",
			"vim",
		); err != nil {
			fmt.Println(ExceptionAtReadFile)
		}
		return nil
	},
}
