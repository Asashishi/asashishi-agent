package main

import (
	"asashishi-agent/conf"
	"asashishi-agent/entry"
)

func init() {
	conf.InitConfig()
}

func main() {
	if conf.Env.WebMode {
		entry.WithWebMode()
	} else {
		entry.WithCliMode()
	}
}
