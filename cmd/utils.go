package cmd

import (
	"asashishi-agent/conf"
	"asashishi-agent/tools"
	"os"

	"github.com/alecthomas/chroma/v2/lexers"
	"github.com/alecthomas/chroma/v2/quick"
)

func RenderFileToTerminal(path string) error {
	var (
		err      error
		content  string
		fileType string
	)
	if fileType = lexers.Match(path).Config().Name; fileType == "" {
		fileType = "plaintext"
	}
	content = tools.ReadFileContent(path)
	if err = quick.Highlight(
		os.Stdout,
		content,
		fileType,
		"terminal16m",
		conf.Env.TerminalCodeStyle,
	); err != nil {
		return err
	}
	return nil
}
