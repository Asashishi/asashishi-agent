package tools

import "time"

func GetFormatedTime() string {
	return time.Now().Format(TimeFormatString)
}
