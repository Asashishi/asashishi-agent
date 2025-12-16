package cmd

import (
	"asashishi-agent/global"
	"asashishi-agent/ui"
)

var FileContentMark string = global.LineBreakString + ui.WidthStyle(
	"-- File Content:",
	global.SystemCommentStyle,
)

var EOFMark string = global.LineBreakString + ui.WidthStyle(
	"-- EOF",
	global.SystemCommentStyle,
)

var ExceptionAtReadFile string = ui.WidthStyle(
	"Unexpect command, eg: cmd -rfile consts.py",
	global.SystemErrorStyle,
)
