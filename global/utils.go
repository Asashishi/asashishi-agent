package global

import (
	"asashishi-agent/ui"
	"fmt"
	"time"
)

func SetTerminalTitle() {
	fmt.Printf(TilteEscape, AppTitle)
}

func PrintAppBanner(version string) {
	var styledVersion string = fmt.Sprintf(Version, version)
	fmt.Printf(AppBanner, styledVersion)
}

func WaitNextFrame(tick float64) {
	time.Sleep(
		(time.Duration((FloatK / tick) * FloatK)) * time.Microsecond,
	)
}

func GetStyledError(errString string) string {
	return ui.WidthStyle(
		errString,
		SystemErrorStyle,
	)
}

func GetStyledSuccess(sucStirng string) string {
	return ui.WidthStyle(
		sucStirng,
		SystemSuccesStyle,
	)
}

func GetStyledSystemComent(comentString string) string {
	return ui.WidthStyle(
		comentString,
		SystemCommentStyle,
	)
}
