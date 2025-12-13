package global

import "time"

func Wait(tick float64) {
	time.Sleep(
		(time.Duration((FloatK / tick) * FloatK)) * time.Microsecond,
	)
}
