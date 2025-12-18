package websocket

import (
	"asashishi-agent/conf"
	"context"
	"fmt"

	ws "github.com/coder/websocket"
)

func WebSocketServerInit() (context.Context, *ws.Conn) {
	var (
		err  error
		ctx  context.Context
		conn *ws.Conn
	)
	ctx = context.Background()
	if conn, _, err = ws.Dial(ctx, fmt.Sprintf(WebsocketURL, conf.Env.WebsocketPort), nil); err != nil {
		panic(err)
	}
	return ctx, conn
}
