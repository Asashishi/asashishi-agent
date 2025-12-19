package websocket

import (
	"asashishi-agent/global"
	"fmt"
	"net/http"

	ws "github.com/coder/websocket"
)

func GetWebsocketConn(conn *ws.Conn, writer http.ResponseWriter, reader *http.Request) *ws.Conn {
	var err error
	if conn != nil {
		conn.Close(ws.StatusInternalError, ClientExit)
	}
	if conn, err = ws.Accept(writer, reader, nil); err != nil {
		fmt.Println(global.GetStyledWarn(err.Error()))
	}
	return conn
}
