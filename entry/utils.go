package entry

import (
	"asashishi-agent/backup"
	"asashishi-agent/conf"
	"asashishi-agent/global"
)

func InitCli() {
	global.SetTerminalTitle()
	if conf.Env.BackUp {
		backup.BackupFiles()
	}
	global.PrintAppBanner(conf.Env.Version)
	go global.InitGlobalCliUInput()
}

func InitWeb() {
	global.SetTerminalTitle()
	if conf.Env.BackUp {
		backup.BackupFiles()
	}
	global.PrintAppBanner(conf.Env.Version)
	go global.InitGlobalWebUInput()
}
