package entry

import (
	"asashishi-agent/backup"
	"asashishi-agent/conf"
	"asashishi-agent/global"
	"asashishi-agent/websocket"
	"fmt"
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

func GetWebsocketConn(conn *ws.Conn, writer http.ResponseWriter, reader *http.Request) *ws.Conn {
	var err error
	if conn != nil {
		conn.Close(ws.StatusInternalError, websocket.ClientExit)
	}
	if conn, err = ws.Accept(writer, reader, nil); err != nil {
		fmt.Println(global.GetStyledWarn(err.Error()))
	}
	return conn
}
