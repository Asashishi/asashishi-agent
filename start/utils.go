package start

import (
	"asashishi-agent/backup"
	"asashishi-agent/conf"
	"asashishi-agent/global"
	"net/http"

	ws "github.com/coder/websocket"
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

func GetWebsocketConn(writer http.ResponseWriter, reader *http.Request) *ws.Conn {
	var (
		err  error
		conn *ws.Conn
	)
	if conn, err = ws.Accept(writer, reader, nil); err != nil {
		conn = nil
		return GetWebsocketConn(writer, reader)
	}
	return conn
}
