package global

import (
	"asashishi-agent/ui"
)

var Input string = ui.WidthStyle(
	"Input:",
	InputStyle,
) + SpaceString

var InputWidthLineBreakFirst string = LineBreakString + Input

var Loading string = ui.WidthStyle(
	"Loading",
	LoadingStyle,
)

var Version string = ui.WidthStyle(
	"-- v%s --",
	VersionStyle,
)

var AppBanner string = ui.WidthStyle(
	`
+--------------------------------------------------------------------------------------+
|                                                                                      |
|     /██████                                /██       /██           /██       /██     |
|    /██__  ██                              | ██      |__/          | ██      |__/     |
|   | ██  \ ██  /███████  /██████   /███████| ███████  /██  /███████| ███████  /██     |
|   | ████████ /██_____/ |____  ██ /██_____/| ██__  ██| ██ /██_____/| ██__  ██| ██     |
|   | ██__  ██|  ██████   /███████|  ██████ | ██  \ ██| ██|  ██████ | ██  \ ██| ██     |
|   | ██  | ██ \____  ██ /██__  ██ \____  ██| ██  | ██| ██ \____  ██| ██  | ██| ██     |
|   | ██  | ██ /███████/|  ███████ /███████/| ██  | ██| ██ /███████/| ██  | ██| ██     |
|   |__/  |__/|_______/  \_______/|_______/ |__/  |__/|__/|_______/ |__/  |__/|__/     |
|                                                                                      |
+--------------------------------------------------------------------------------------+
				     %s
`,
	BannerStyle,
)

var AIOutput string = ui.WidthStyle("%s", AIOutputStyle)
