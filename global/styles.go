package global

import (
	"asashishi-agent/ui"
)

var LoadingStyle ui.Style = ui.Style{
	Italic:    true,
	FastBlink: true,
	Fg:        White,
}

var SystemCommentStyle ui.Style = ui.Style{
	Italic: true,
	Fg:     Grey,
}

var BannerStyle ui.Style = ui.Style{
	Italic: true,
	Fg:     Teal,
}

var VersionStyle ui.Style = ui.Style{
	Italic:    true,
	Bold:      true,
	Underline: true,
	Fg:        Violet,
}

var InputStyle ui.Style = ui.Style{
	Italic:    true,
	Underline: true,
	Fg:        White,
}

var AIOutputStyle ui.Style = ui.Style{
	Fg: LiteBlue,
}

var SystemSuccesStyle ui.Style = ui.Style{
	Fg:     Grreen,
	Italic: true,
}
