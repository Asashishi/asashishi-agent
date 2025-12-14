package global

import (
	"fmt"
	"time"
)

func WaitNextFrame(tick float64) {
	time.Sleep(
		(time.Duration((FloatK / tick) * FloatK)) * time.Microsecond,
	)
}

func SetTerminalTitle() {
	fmt.Printf(TilteEscape, AppTitle)
}

func PrintAppBanner(version string) {
	fmt.Printf(AppBanner, version)
}
