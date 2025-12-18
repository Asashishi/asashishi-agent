package global

import (
	"asashishi-agent/ui"
)

var LoadingStyle ui.Style = ui.Style{
	Italic:    true,
	FastBlink: true,
	Fg:        White,
}

var BannerStyle ui.Style = ui.Style{
	Fg: Teal,
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

var SystemCommentStyle ui.Style = ui.Style{
	Bold:   true,
	Italic: true,
	Fg:     Grey,
}

var SystemSuccesStyle ui.Style = ui.Style{
	Fg:     Grreen,
	Bold:   true,
	Italic: true,
}

var SystemErrorStyle ui.Style = ui.Style{
	Fg:     Red,
	Bold:   true,
	Italic: true,
}

var SystemWarnStyle ui.Style = ui.Style{
	Fg:     Yelleow,
	Bold:   true,
	Italic: true,
}
