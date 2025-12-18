package main

import (
	"asashishi-agent/conf"
	"asashishi-agent/start"
)

func main() {
	if conf.Env.WebMode {

	} else {
		start.WithCliMode()
	}
}
