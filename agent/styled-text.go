package agent

import (
	"asashishi-agent/global"
	"asashishi-agent/ui"
)

var ProcessingDataClean string = ui.WidthStyle(
	"Processing ",
	global.LoadingStyle,
) + "Data Clean 🔄"

var CallToolWithoutArgs string = ui.WidthStyle(
	"Calling",
	global.LoadingStyle,
) + " %s 🔄\n"

var CallToolWithArgs string = ui.WidthStyle(
	"Calling",
	global.LoadingStyle,
) + " %s 🔄 args: %s\n"
