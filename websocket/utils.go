package websocket

import (
	"asashishi-agent/conf"
	"asashishi-agent/global"
	"context"
	"fmt"
	"net/http"
	"time"

	ws "github.com/coder/websocket"
)

func webSocketPing(ctx context.Context, conn *ws.Conn) {
	var err error
	for {
		time.Sleep(time.Second * time.Duration(WebSocketPingDelay))
		err = conn.Ping(ctx)
		if err != nil {
			fmt.Println(global.GetStyledWarn(err.Error()))
			break
		}
	}
}

func GetWebsocketConn(ctx context.Context, conn *ws.Conn, reader *http.Request, writer http.ResponseWriter) *ws.Conn {
	var err error
	if conn != nil {
		conn.Close(ws.StatusNormalClosure, ClientExit)
	}
	if conn, err = ws.Accept(writer, reader, &ws.AcceptOptions{OriginPatterns: conf.Env.AllowOrigins}); err != nil {
		fmt.Println(global.GetStyledWarn(err.Error()))
		return nil
	}
	go webSocketPing(ctx, conn)
	return conn
}
