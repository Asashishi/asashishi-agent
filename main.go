package main

import (
	"asashishi-agent/conf"
	"asashishi-agent/start"
)

func init() {
	conf.InitConfig()
}

func main() {
	if conf.Env.WebPageMode {
		start.WithWebMode()
	} else {
		start.WithCliMode()
	}
}
