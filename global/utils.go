package global

import (
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
