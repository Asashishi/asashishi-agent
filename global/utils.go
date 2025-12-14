package global

import (
	"fmt"
	"time"
)

func SetTerminalTitle() {
	fmt.Printf(TilteEscape, AppTitle)
}

func PrintAppBanner(version string) {
	fmt.Printf(AppBanner, version)
}

func WaitNextFrame(tick float64) {
	time.Sleep(
		(time.Duration((FloatK / tick) * FloatK)) * time.Microsecond,
	)
}
