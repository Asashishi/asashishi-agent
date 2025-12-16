package tools

import (
	"asashishi-agent/global"
	"asashishi-agent/ui"
)

func GetStyledError(errString string) string {
	return ui.WidthStyle(
		errString,
		global.SystemErrorStyle,
	)
}
