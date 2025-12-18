package main

import (
	"asashishi-agent/conf"
	"asashishi-agent/start"
)

func main() {
	if conf.Env.WebMode {
		start.WithWebMode()
	} else {
		start.WithCliMode()
	}
}
