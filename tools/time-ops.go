package tools

import "time"

func GetFormatedTime() string {
	return time.Now().Format("20060102150405")
}
